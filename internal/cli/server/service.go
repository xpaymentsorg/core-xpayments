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
	"math/big"
	"reflect"
	"strings"

	gproto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	grpc_net_conn "github.com/mitchellh/go-grpc-net-conn"
	"github.com/xpaymentsorg/go-xpayments/core"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/internal/cli/server/pprof"
	"github.com/xpaymentsorg/go-xpayments/internal/cli/server/proto"
	"github.com/xpaymentsorg/go-xpayments/p2p"
	"github.com/xpaymentsorg/go-xpayments/p2p/enode"
)

func (s *Server) Pprof(req *proto.PprofRequest, stream proto.Bor_PprofServer) error {
	var payload []byte
	var headers map[string]string
	var err error

	ctx := context.Background()
	switch req.Type {
	case proto.PprofRequest_CPU:
		payload, headers, err = pprof.CPUProfile(ctx, int(req.Seconds))
	case proto.PprofRequest_TRACE:
		payload, headers, err = pprof.Trace(ctx, int(req.Seconds))
	case proto.PprofRequest_LOOKUP:
		payload, headers, err = pprof.Profile(req.Profile, 0, 0)
	}
	if err != nil {
		return err
	}

	// open the stream and send the headers
	err = stream.Send(&proto.PprofResponse{
		Event: &proto.PprofResponse_Open_{
			Open: &proto.PprofResponse_Open{
				Headers: headers,
				Size:    int64(len(payload)),
			},
		},
	})
	if err != nil {
		return err
	}

	// Wrap our conn around the response.
	conn := &grpc_net_conn.Conn{
		Stream:  stream,
		Request: &proto.PprofResponse_Input{},
		Encode: grpc_net_conn.SimpleEncoder(func(msg gproto.Message) *[]byte {
			return &msg.(*proto.PprofResponse_Input).Data
		}),
	}
	if _, err := conn.Write(payload); err != nil {
		return err
	}

	// send the eof
	err = stream.Send(&proto.PprofResponse{
		Event: &proto.PprofResponse_Eof{},
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) PeersAdd(ctx context.Context, req *proto.PeersAddRequest) (*proto.PeersAddResponse, error) {
	node, err := enode.Parse(enode.ValidSchemes, req.Enode)
	if err != nil {
		return nil, fmt.Errorf("invalid enode: %v", err)
	}
	srv := s.node.Server()
	if req.Trusted {
		srv.AddTrustedPeer(node)
	} else {
		srv.AddPeer(node)
	}
	return &proto.PeersAddResponse{}, nil
}

func (s *Server) PeersRemove(ctx context.Context, req *proto.PeersRemoveRequest) (*proto.PeersRemoveResponse, error) {
	node, err := enode.Parse(enode.ValidSchemes, req.Enode)
	if err != nil {
		return nil, fmt.Errorf("invalid enode: %v", err)
	}
	srv := s.node.Server()
	if req.Trusted {
		srv.RemoveTrustedPeer(node)
	} else {
		srv.RemovePeer(node)
	}
	return &proto.PeersRemoveResponse{}, nil
}

func (s *Server) PeersList(ctx context.Context, req *proto.PeersListRequest) (*proto.PeersListResponse, error) {
	resp := &proto.PeersListResponse{}

	peers := s.node.Server().PeersInfo()
	for _, p := range peers {
		resp.Peers = append(resp.Peers, peerInfoToPeer(p))
	}
	return resp, nil
}

func (s *Server) PeersStatus(ctx context.Context, req *proto.PeersStatusRequest) (*proto.PeersStatusResponse, error) {
	var peerInfo *p2p.PeerInfo
	for _, p := range s.node.Server().PeersInfo() {
		if strings.HasPrefix(p.ID, req.Enode) {
			if peerInfo != nil {
				return nil, fmt.Errorf("more than one peer with the same prefix")
			}
			peerInfo = p
		}
	}
	resp := &proto.PeersStatusResponse{}
	if peerInfo != nil {
		resp.Peer = peerInfoToPeer(peerInfo)
	}
	return resp, nil
}

func peerInfoToPeer(info *p2p.PeerInfo) *proto.Peer {
	return &proto.Peer{
		Id:      info.ID,
		Enode:   info.Enode,
		Enr:     info.ENR,
		Caps:    info.Caps,
		Name:    info.Name,
		Trusted: info.Network.Trusted,
		Static:  info.Network.Static,
	}
}

func (s *Server) ChainSetHead(ctx context.Context, req *proto.ChainSetHeadRequest) (*proto.ChainSetHeadResponse, error) {
	s.backend.APIBackend.SetHead(req.Number)
	return &proto.ChainSetHeadResponse{}, nil
}

func (s *Server) Status(ctx context.Context, _ *empty.Empty) (*proto.StatusResponse, error) {
	apiBackend := s.backend.APIBackend
	syncProgress := apiBackend.SyncProgress()

	resp := &proto.StatusResponse{
		CurrentHeader: headerToProtoHeader(apiBackend.CurrentHeader()),
		CurrentBlock:  headerToProtoHeader(apiBackend.CurrentBlock().Header()),
		NumPeers:      int64(len(s.node.Server().PeersInfo())),
		SyncMode:      s.config.SyncMode,
		Syncing: &proto.StatusResponse_Syncing{
			StartingBlock: int64(syncProgress.StartingBlock),
			HighestBlock:  int64(syncProgress.HighestBlock),
			CurrentBlock:  int64(syncProgress.CurrentBlock),
		},
		Forks: gatherForks(s.config.chain.Genesis.Config, s.config.chain.Genesis.Config.Bor),
	}
	return resp, nil
}

func headerToProtoHeader(h *types.Header) *proto.Header {
	return &proto.Header{
		Hash:   h.Hash().String(),
		Number: h.Number.Uint64(),
	}
}

var bigIntT = reflect.TypeOf(new(big.Int)).Kind()

// gatherForks gathers all the fork numbers via reflection
func gatherForks(configList ...interface{}) []*proto.StatusResponse_Fork {
	var forks []*proto.StatusResponse_Fork

	for _, config := range configList {
		kind := reflect.TypeOf(config)
		for kind.Kind() == reflect.Ptr {
			kind = kind.Elem()
		}

		skip := "DAOForkBlock"

		conf := reflect.ValueOf(config).Elem()
		for i := 0; i < kind.NumField(); i++ {
			// Fetch the next field and skip non-fork rules
			field := kind.Field(i)
			if strings.Contains(field.Name, skip) {
				continue
			}
			if !strings.HasSuffix(field.Name, "Block") {
				continue
			}

			fork := &proto.StatusResponse_Fork{
				Name: strings.TrimSuffix(field.Name, "Block"),
			}

			val := conf.Field(i)
			switch field.Type.Kind() {
			case bigIntT:
				rule := val.Interface().(*big.Int)
				if rule != nil {
					fork.Block = rule.Int64()
				} else {
					fork.Disabled = true
				}
			case reflect.Uint64:
				fork.Block = int64(val.Uint())

			default:
				continue
			}

			forks = append(forks, fork)
		}
	}
	return forks
}

func convertBlockToBlockStub(blocks []*types.Block) []*proto.BlockStub {

	var blockStubs []*proto.BlockStub

	for _, block := range blocks {
		blockStub := &proto.BlockStub{
			Hash:   block.Hash().String(),
			Number: block.NumberU64(),
		}
		blockStubs = append(blockStubs, blockStub)
	}

	return blockStubs
}

func (s *Server) ChainWatch(req *proto.ChainWatchRequest, reply proto.Bor_ChainWatchServer) error {

	chain2HeadChanSize := 10

	chain2HeadCh := make(chan core.Chain2HeadEvent, chain2HeadChanSize)
	headSub := s.backend.APIBackend.SubscribeChain2HeadEvent(chain2HeadCh)
	defer headSub.Unsubscribe()

	for {
		msg := <-chain2HeadCh

		err := reply.Send(&proto.ChainWatchResponse{Type: msg.Type,
			Newchain: convertBlockToBlockStub(msg.NewChain),
			Oldchain: convertBlockToBlockStub(msg.OldChain),
		})
		if err != nil {
			return err
		}
	}
}
