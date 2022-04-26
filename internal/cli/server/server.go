// Copyright 2022 The go-xpayments Authors
// This file is part of the go-xpayments library.
//
// Copyright 2022 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/xpaymentsorg/go-xpayments/accounts/keystore"
	"github.com/xpaymentsorg/go-xpayments/eth"
	"github.com/xpaymentsorg/go-xpayments/eth/tracers"
	"github.com/xpaymentsorg/go-xpayments/ethstats"
	"github.com/xpaymentsorg/go-xpayments/graphql"
	"github.com/xpaymentsorg/go-xpayments/internal/cli/server/proto"
	"github.com/xpaymentsorg/go-xpayments/log"
	"github.com/xpaymentsorg/go-xpayments/metrics"
	"github.com/xpaymentsorg/go-xpayments/metrics/influxdb"
	"github.com/xpaymentsorg/go-xpayments/metrics/prometheus"
	"github.com/xpaymentsorg/go-xpayments/node"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedBorServer
	node       *node.Node
	backend    *eth.Ethereum
	grpcServer *grpc.Server
	tracer     *sdktrace.TracerProvider
	config     *Config
}

func NewServer(config *Config) (*Server, error) {
	srv := &Server{
		config: config,
	}

	// start the logger
	setupLogger(config.LogLevel)

	if err := srv.setupGRPCServer(config.GRPC.Addr); err != nil {
		return nil, err
	}

	// load the chain genesis
	if err := config.loadChain(); err != nil {
		return nil, err
	}

	// create the node/stack
	nodeCfg, err := config.buildNode()
	if err != nil {
		return nil, err
	}
	stack, err := node.New(nodeCfg)
	if err != nil {
		return nil, err
	}
	srv.node = stack

	// setup account manager (only keystore)
	{
		keydir := stack.KeyStoreDir()
		n, p := keystore.StandardScryptN, keystore.StandardScryptP
		if config.Accounts.UseLightweightKDF {
			n, p = keystore.LightScryptN, keystore.LightScryptP
		}
		stack.AccountManager().AddBackend(keystore.NewKeyStore(keydir, n, p))
	}

	// register the ethereum backend
	ethCfg, err := config.buildEth(stack)
	if err != nil {
		return nil, err
	}

	backend, err := eth.New(stack, ethCfg)
	if err != nil {
		return nil, err
	}
	srv.backend = backend

	// debug tracing is enabled by default
	stack.RegisterAPIs(tracers.APIs(backend.APIBackend))

	// graphql is started from another place
	if config.JsonRPC.Graphql.Enabled {
		if err := graphql.New(stack, backend.APIBackend, config.JsonRPC.Cors, config.JsonRPC.VHost); err != nil {
			return nil, fmt.Errorf("failed to register the GraphQL service: %v", err)
		}
	}

	// register ethash service
	if config.Ethstats != "" {
		if err := ethstats.New(stack, backend.APIBackend, backend.Engine(), config.Ethstats); err != nil {
			return nil, err
		}
	}

	// sealing (if enabled) or in dev mode
	if config.Sealer.Enabled || config.Developer.Enabled {
		if err := backend.StartMining(1); err != nil {
			return nil, err
		}
	}

	if err := srv.setupMetrics(config.Telemetry, config.Name); err != nil {
		return nil, err
	}

	// start the node
	if err := srv.node.Start(); err != nil {
		return nil, err
	}
	return srv, nil
}

func (s *Server) Stop() {
	s.node.Close()

	// shutdown the tracer
	if s.tracer != nil {
		if err := s.tracer.Shutdown(context.Background()); err != nil {
			log.Error("Failed to shutdown open telemetry tracer")
		}
	}
}

func (s *Server) setupMetrics(config *TelemetryConfig, serviceName string) error {
	metrics.Enabled = config.Enabled
	metrics.EnabledExpensive = config.Expensive

	if !metrics.Enabled {
		// metrics are disabled, do not set up any sink
		return nil
	}

	log.Info("Enabling metrics collection")

	// influxdb
	if v1Enabled, v2Enabled := config.InfluxDB.V1Enabled, config.InfluxDB.V2Enabled; v1Enabled || v2Enabled {
		if v1Enabled && v2Enabled {
			return fmt.Errorf("both influx v1 and influx v2 cannot be enabled")
		}

		cfg := config.InfluxDB
		tags := cfg.Tags
		endpoint := cfg.Endpoint

		if v1Enabled {
			log.Info("Enabling metrics export to InfluxDB (v1)")
			go influxdb.InfluxDBWithTags(metrics.DefaultRegistry, 10*time.Second, endpoint, cfg.Database, cfg.Username, cfg.Password, "geth.", tags)
		}
		if v2Enabled {
			log.Info("Enabling metrics export to InfluxDB (v2)")
			go influxdb.InfluxDBV2WithTags(metrics.DefaultRegistry, 10*time.Second, endpoint, cfg.Token, cfg.Bucket, cfg.Organization, "geth.", tags)
		}
	}

	// Start system runtime metrics collection
	go metrics.CollectProcessMetrics(3 * time.Second)

	if config.PrometheusAddr != "" {

		prometheusMux := http.NewServeMux()

		prometheusMux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
			prometheus.Handler(metrics.DefaultRegistry)
		})

		promServer := &http.Server{
			Addr:    config.PrometheusAddr,
			Handler: prometheusMux,
		}

		go func() {
			if err := promServer.ListenAndServe(); err != nil {
				log.Error("Failure in running Prometheus server", "err", err)
			}
		}()

	}

	if config.OpenCollectorEndpoint != "" {
		// setup open collector tracer
		ctx := context.Background()

		res, err := resource.New(ctx,
			resource.WithAttributes(
				// the service name used to display traces in backends
				semconv.ServiceNameKey.String(serviceName),
			),
		)
		if err != nil {
			return fmt.Errorf("failed to create open telemetry resource for service: %v", err)
		}

		// Set up a trace exporter
		traceExporter, err := otlptracegrpc.New(
			ctx,
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(config.OpenCollectorEndpoint),
		)
		if err != nil {
			return fmt.Errorf("failed to create open telemetry tracer exporter for service: %v", err)
		}

		// Register the trace exporter with a TracerProvider, using a batch
		// span processor to aggregate spans before export.
		bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
		tracerProvider := sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(res),
			sdktrace.WithSpanProcessor(bsp),
		)
		otel.SetTracerProvider(tracerProvider)

		// set global propagator to tracecontext (the default is no-op).
		otel.SetTextMapPropagator(propagation.TraceContext{})

		// set the tracer
		s.tracer = tracerProvider
	}

	return nil
}

func (s *Server) setupGRPCServer(addr string) error {
	s.grpcServer = grpc.NewServer(s.withLoggingUnaryInterceptor())
	proto.RegisterBorServer(s.grpcServer, s)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			log.Error("failed to serve grpc server", "err", err)
		}
	}()

	log.Info("GRPC Server started", "addr", addr)
	return nil
}

func (s *Server) withLoggingUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(s.loggingServerInterceptor)
}

func (s *Server) loggingServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)
	log.Trace("Request", "method", info.FullMethod, "duration", time.Since(start), "error", err)
	return h, err
}

func setupLogger(logLevel string) {
	output := io.Writer(os.Stderr)
	usecolor := (isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())) && os.Getenv("TERM") != "dumb"
	if usecolor {
		output = colorable.NewColorableStderr()
	}
	ostream := log.StreamHandler(output, log.TerminalFormat(usecolor))
	glogger := log.NewGlogHandler(ostream)

	// logging
	lvl, err := log.LvlFromString(strings.ToLower(logLevel))
	if err == nil {
		glogger.Verbosity(lvl)
	} else {
		glogger.Verbosity(log.LvlInfo)
	}
	log.Root().SetHandler(glogger)
}
