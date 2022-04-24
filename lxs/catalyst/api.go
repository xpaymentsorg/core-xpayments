// Copyright 2022 The go-xpayments Authors
// This file is part of the go-xpayments library.
//
// The go-xpayments library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-xpayments library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-xpayments library. If not, see <http://www.gnu.org/licenses/>.

// Package catalyst implements the temporary xps1/xps2 RPC integration.
package catalyst

import (
	"errors"
	"fmt"

	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/core/beacon"
	"github.com/xpaymentsorg/go-xpayments/log"
	"github.com/xpaymentsorg/go-xpayments/lxs"
	"github.com/xpaymentsorg/go-xpayments/node"
	"github.com/xpaymentsorg/go-xpayments/rpc"
)

// Register adds catalyst APIs to the light client.
func Register(stack *node.Node, backend *lxs.LightxPayments) error {
	log.Warn("Catalyst mode enabled", "protocol", "lxs")
	stack.RegisterAPIs([]rpc.API{
		{
			Namespace:     "engine",
			Version:       "1.0",
			Service:       NewConsensusAPI(backend),
			Public:        true,
			Authenticated: true,
		},
	})
	return nil
}

type ConsensusAPI struct {
	lxs *lxs.LightxPayments
}

// NewConsensusAPI creates a new consensus api for the given backend.
// The underlying blockchain needs to have a valid terminal total difficulty set.
func NewConsensusAPI(lxs *lxs.LightxPayments) *ConsensusAPI {
	if lxs.BlockChain().Config().TerminalTotalDifficulty == nil {
		panic("Catalyst started without valid total difficulty")
	}
	return &ConsensusAPI{lxs: lxs}
}

// ForkchoiceUpdatedV1 has several responsibilities:
// If the method is called with an empty head block:
// 		we return success, which can be used to check if the catalyst mode is enabled
// If the total difficulty was not reached:
// 		we return INVALID
// If the finalizedBlockHash is set:
// 		we check if we have the finalizedBlockHash in our db, if not we start a sync
// We try to set our blockchain to the headBlock
// If there are payloadAttributes:
//      we return an error since block creation is not supported in lxs mode
func (api *ConsensusAPI) ForkchoiceUpdatedV1(heads beacon.ForkchoiceStateV1, payloadAttributes *beacon.PayloadAttributesV1) (beacon.ForkChoiceResponse, error) {
	if heads.HeadBlockHash == (common.Hash{}) {
		log.Warn("Forkchoice requested update to zero hash")
		return beacon.STATUS_INVALID, nil // TODO(karalabe): Why does someone send us this?
	}
	if err := api.checkTerminalTotalDifficulty(heads.HeadBlockHash); err != nil {
		if header := api.lxs.BlockChain().GetHeaderByHash(heads.HeadBlockHash); header == nil {
			// TODO (MariusVanDerWijden) trigger sync
			return beacon.STATUS_SYNCING, nil
		}
		return beacon.STATUS_INVALID, err
	}
	// If the finalized block is set, check if it is in our blockchain
	if heads.FinalizedBlockHash != (common.Hash{}) {
		if header := api.lxs.BlockChain().GetHeaderByHash(heads.FinalizedBlockHash); header == nil {
			// TODO (MariusVanDerWijden) trigger sync
			return beacon.STATUS_SYNCING, nil
		}
	}
	// SetHead
	if err := api.setHead(heads.HeadBlockHash); err != nil {
		return beacon.STATUS_INVALID, err
	}
	if payloadAttributes != nil {
		return beacon.STATUS_INVALID, errors.New("not supported")
	}
	return api.validForkChoiceResponse(), nil
}

// GetPayloadV1 returns a cached payload by id. It's not supported in lxs mode.
func (api *ConsensusAPI) GetPayloadV1(payloadID beacon.PayloadID) (*beacon.ExecutableDataV1, error) {
	return nil, &beacon.GenericServerError
}

// ExecutePayloadV1 creates an Xps1 block, inserts it in the chain, and returns the status of the chain.
func (api *ConsensusAPI) ExecutePayloadV1(params beacon.ExecutableDataV1) (beacon.PayloadStatusV1, error) {
	block, err := beacon.ExecutableDataToBlock(params)
	if err != nil {
		return api.invalid(), err
	}
	if !api.lxs.BlockChain().HasHeader(block.ParentHash(), block.NumberU64()-1) {
		/*
			TODO (MariusVanDerWijden) reenable once sync is merged
			if err := api.xps.Downloader().BeaconSync(api.xps.SyncMode(), block.Header()); err != nil {
				return SYNCING, err
			}
		*/
		// TODO (MariusVanDerWijden) we should return nil here not empty hash
		return beacon.PayloadStatusV1{Status: beacon.SYNCING, LatestValidHash: nil}, nil
	}
	parent := api.lxs.BlockChain().GetHeaderByHash(params.ParentHash)
	if parent == nil {
		return api.invalid(), fmt.Errorf("could not find parent %x", params.ParentHash)
	}
	td := api.lxs.BlockChain().GetTd(parent.Hash(), block.NumberU64()-1)
	ttd := api.lxs.BlockChain().Config().TerminalTotalDifficulty
	if td.Cmp(ttd) < 0 {
		return api.invalid(), fmt.Errorf("can not execute payload on top of block with low td got: %v threshold %v", td, ttd)
	}
	if err = api.lxs.BlockChain().InsertHeader(block.Header()); err != nil {
		return api.invalid(), err
	}
	if merger := api.lxs.Merger(); !merger.TDDReached() {
		merger.ReachTTD()
	}
	hash := block.Hash()
	return beacon.PayloadStatusV1{Status: beacon.VALID, LatestValidHash: &hash}, nil
}

func (api *ConsensusAPI) validForkChoiceResponse() beacon.ForkChoiceResponse {
	currentHash := api.lxs.BlockChain().CurrentHeader().Hash()
	return beacon.ForkChoiceResponse{
		PayloadStatus: beacon.PayloadStatusV1{Status: beacon.VALID, LatestValidHash: &currentHash},
	}
}

// invalid returns a response "INVALID" with the latest valid hash set to the current head.
func (api *ConsensusAPI) invalid() beacon.PayloadStatusV1 {
	currentHash := api.lxs.BlockChain().CurrentHeader().Hash()
	return beacon.PayloadStatusV1{Status: beacon.INVALID, LatestValidHash: &currentHash}
}

func (api *ConsensusAPI) checkTerminalTotalDifficulty(head common.Hash) error {
	// shortcut if we entered PoS already
	if api.lxs.Merger().PoSFinalized() {
		return nil
	}
	// make sure the parent has enough terminal total difficulty
	header := api.lxs.BlockChain().GetHeaderByHash(head)
	if header == nil {
		return &beacon.GenericServerError
	}
	td := api.lxs.BlockChain().GetTd(header.Hash(), header.Number.Uint64())
	if td != nil && td.Cmp(api.lxs.BlockChain().Config().TerminalTotalDifficulty) < 0 {
		return &beacon.InvalidTB
	}
	return nil
}

// setHead is called to perform a force choice.
func (api *ConsensusAPI) setHead(newHead common.Hash) error {
	log.Info("Setting head", "head", newHead)

	headHeader := api.lxs.BlockChain().CurrentHeader()
	if headHeader.Hash() == newHead {
		return nil
	}
	newHeadHeader := api.lxs.BlockChain().GetHeaderByHash(newHead)
	if newHeadHeader == nil {
		return &beacon.GenericServerError
	}
	if err := api.lxs.BlockChain().SetChainHead(newHeadHeader); err != nil {
		return err
	}
	// Trigger the transition if it's the first `NewHead` event.
	if merger := api.lxs.Merger(); !merger.PoSFinalized() {
		merger.FinalizePoS()
	}
	return nil
}
