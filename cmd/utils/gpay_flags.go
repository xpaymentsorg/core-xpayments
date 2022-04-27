// Copyright 2022 The go-xpayments Authors
// This file is part of the go-xpayments library.
//
// Copyright 2022 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/xpaymentsorg/go-xpayments/core"
	"github.com/xpaymentsorg/go-xpayments/eth"
	"github.com/xpaymentsorg/go-xpayments/log"
	"github.com/xpaymentsorg/go-xpayments/node"
	"gopkg.in/urfave/cli.v1"
)

var (
	//
	// XPoS Specific flags
	//

	// GenisysURLFlag flag for genisys url
	GenisysURLFlag = cli.StringFlag{
		Name:  "gpay.genisys",
		Usage: "URL of Genisys service",
		Value: "http://localhost:1818",
	}

	// WithoutGenisysFlag no genisys (for testing purpose)
	WithoutGenisysFlag = cli.BoolFlag{
		Name:  "gpay.withoutgenisys",
		Usage: "Run without Genisys service (for testing purpose)",
	}

	// XPoSFlags all gpay related flags
	XPoSFlags = []cli.Flag{
		GenisysURLFlag,
		WithoutGenisysFlag,
	}
)

func getGenesis(genesisPath string) (*core.Genesis, error) {
	log.Info("Reading genesis at ", "file", genesisPath)
	file, err := os.Open(genesisPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	genesis := new(core.Genesis)
	if err := json.NewDecoder(file).Decode(genesis); err != nil {
		return nil, err
	}
	return genesis, nil
}

// SetXPoSConfig sets gpay config
func SetXPoSConfig(ctx *cli.Context, cfg *eth.Config) {
	cfg.GenisysURL = ctx.GlobalString(GenisysURLFlag.Name)
	cfg.WithoutGenisys = ctx.GlobalBool(WithoutGenisysFlag.Name)
}

// CreateXPoSEthereum Creates gpay ethereum object from eth.Config
func CreateXPoSEthereum(cfg *eth.Config) *eth.Ethereum {
	workspace, err := ioutil.TempDir("", "gpay-command-node-")
	if err != nil {
		Fatalf("Failed to create temporary keystore: %v", err)
	}

	// Create a networkless protocol stack and start an Ethereum service within
	stack, err := node.New(&node.Config{DataDir: workspace, UseLightweightKDF: true, Name: "gpay-command-node"})
	if err != nil {
		Fatalf("Failed to create node: %v", err)
	}
	ethereum, err := eth.New(stack, cfg)
	if err != nil {
		Fatalf("Failed to register Ethereum protocol: %v", err)
	}

	// Start the node and assemble the JavaScript console around it
	if err = stack.Start(); err != nil {
		Fatalf("Failed to start stack: %v", err)
	}
	_, err = stack.Attach()
	if err != nil {
		Fatalf("Failed to attach to node: %v", err)
	}

	return ethereum
}
