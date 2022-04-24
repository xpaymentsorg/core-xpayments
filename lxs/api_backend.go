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

package lxs

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/xpaymentsorg/go-xpayments"
	"github.com/xpaymentsorg/go-xpayments/accounts"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/consensus"
	"github.com/xpaymentsorg/go-xpayments/core"
	"github.com/xpaymentsorg/go-xpayments/core/bloombits"
	"github.com/xpaymentsorg/go-xpayments/core/rawdb"
	"github.com/xpaymentsorg/go-xpayments/core/state"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/core/vm"
	"github.com/xpaymentsorg/go-xpayments/event"
	"github.com/xpaymentsorg/go-xpayments/light"
	"github.com/xpaymentsorg/go-xpayments/params"
	"github.com/xpaymentsorg/go-xpayments/rpc"
	"github.com/xpaymentsorg/go-xpayments/xps/gasprice"
	"github.com/xpaymentsorg/go-xpayments/xpsdb"
)

type LxsApiBackend struct {
	extRPCEnabled       bool
	allowUnprotectedTxs bool
	xps                 *LightxPayments
	gpo                 *gasprice.Oracle
}

func (b *LxsApiBackend) ChainConfig() *params.ChainConfig {
	return b.xps.chainConfig
}

func (b *LxsApiBackend) CurrentBlock() *types.Block {
	return types.NewBlockWithHeader(b.xps.BlockChain().CurrentHeader())
}

func (b *LxsApiBackend) SetHead(number uint64) {
	b.xps.handler.downloader.Cancel()
	b.xps.blockchain.SetHead(number)
}

func (b *LxsApiBackend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	// Return the latest current as the pending one since there
	// is no pending notion in the light client. TODO(rjl493456442)
	// unify the behavior of `HeaderByNumber` and `PendingBlockAndReceipts`.
	if number == rpc.PendingBlockNumber {
		return b.xps.blockchain.CurrentHeader(), nil
	}
	if number == rpc.LatestBlockNumber {
		return b.xps.blockchain.CurrentHeader(), nil
	}
	return b.xps.blockchain.GetHeaderByNumberOdr(ctx, uint64(number))
}

func (b *LxsApiBackend) HeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Header, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.HeaderByNumber(ctx, blockNr)
	}
	if hash, ok := blockNrOrHash.Hash(); ok {
		header, err := b.HeaderByHash(ctx, hash)
		if err != nil {
			return nil, err
		}
		if header == nil {
			return nil, errors.New("header for hash not found")
		}
		if blockNrOrHash.RequireCanonical && b.xps.blockchain.GetCanonicalHash(header.Number.Uint64()) != hash {
			return nil, errors.New("hash is not currently canonical")
		}
		return header, nil
	}
	return nil, errors.New("invalid arguments; neither block nor hash specified")
}

func (b *LxsApiBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return b.xps.blockchain.GetHeaderByHash(hash), nil
}

func (b *LxsApiBackend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	header, err := b.HeaderByNumber(ctx, number)
	if header == nil || err != nil {
		return nil, err
	}
	return b.BlockByHash(ctx, header.Hash())
}

func (b *LxsApiBackend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return b.xps.blockchain.GetBlockByHash(ctx, hash)
}

func (b *LxsApiBackend) BlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Block, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.BlockByNumber(ctx, blockNr)
	}
	if hash, ok := blockNrOrHash.Hash(); ok {
		block, err := b.BlockByHash(ctx, hash)
		if err != nil {
			return nil, err
		}
		if block == nil {
			return nil, errors.New("header found, but block body is missing")
		}
		if blockNrOrHash.RequireCanonical && b.xps.blockchain.GetCanonicalHash(block.NumberU64()) != hash {
			return nil, errors.New("hash is not currently canonical")
		}
		return block, nil
	}
	return nil, errors.New("invalid arguments; neither block nor hash specified")
}

func (b *LxsApiBackend) PendingBlockAndReceipts() (*types.Block, types.Receipts) {
	return nil, nil
}

func (b *LxsApiBackend) StateAndHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	header, err := b.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, nil, err
	}
	if header == nil {
		return nil, nil, errors.New("header not found")
	}
	return light.NewState(ctx, header, b.xps.odr), header, nil
}

func (b *LxsApiBackend) StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *types.Header, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.StateAndHeaderByNumber(ctx, blockNr)
	}
	if hash, ok := blockNrOrHash.Hash(); ok {
		header := b.xps.blockchain.GetHeaderByHash(hash)
		if header == nil {
			return nil, nil, errors.New("header for hash not found")
		}
		if blockNrOrHash.RequireCanonical && b.xps.blockchain.GetCanonicalHash(header.Number.Uint64()) != hash {
			return nil, nil, errors.New("hash is not currently canonical")
		}
		return light.NewState(ctx, header, b.xps.odr), header, nil
	}
	return nil, nil, errors.New("invalid arguments; neither block nor hash specified")
}

func (b *LxsApiBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	if number := rawdb.ReadHeaderNumber(b.xps.chainDb, hash); number != nil {
		return light.GetBlockReceipts(ctx, b.xps.odr, hash, *number)
	}
	return nil, nil
}

func (b *LxsApiBackend) GetLogs(ctx context.Context, hash common.Hash) ([][]*types.Log, error) {
	if number := rawdb.ReadHeaderNumber(b.xps.chainDb, hash); number != nil {
		return light.GetBlockLogs(ctx, b.xps.odr, hash, *number)
	}
	return nil, nil
}

func (b *LxsApiBackend) GetTd(ctx context.Context, hash common.Hash) *big.Int {
	if number := rawdb.ReadHeaderNumber(b.xps.chainDb, hash); number != nil {
		return b.xps.blockchain.GetTdOdr(ctx, hash, *number)
	}
	return nil
}

func (b *LxsApiBackend) GetXVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmConfig *vm.Config) (*vm.XVM, func() error, error) {
	if vmConfig == nil {
		vmConfig = new(vm.Config)
	}
	txContext := core.NewXVMTxContext(msg)
	context := core.NewXVMBlockContext(header, b.xps.blockchain, nil)
	return vm.NewXVM(context, txContext, state, b.xps.chainConfig, *vmConfig), state.Error, nil
}

func (b *LxsApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.xps.txPool.Add(ctx, signedTx)
}

func (b *LxsApiBackend) RemoveTx(txHash common.Hash) {
	b.xps.txPool.RemoveTx(txHash)
}

func (b *LxsApiBackend) GetPoolTransactions() (types.Transactions, error) {
	return b.xps.txPool.GetTransactions()
}

func (b *LxsApiBackend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	return b.xps.txPool.GetTransaction(txHash)
}

func (b *LxsApiBackend) GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	return light.GetTransaction(ctx, b.xps.odr, txHash)
}

func (b *LxsApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.xps.txPool.GetNonce(ctx, addr)
}

func (b *LxsApiBackend) Stats() (pending int, queued int) {
	return b.xps.txPool.Stats(), 0
}

func (b *LxsApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.xps.txPool.Content()
}

func (b *LxsApiBackend) TxPoolContentFrom(addr common.Address) (types.Transactions, types.Transactions) {
	return b.xps.txPool.ContentFrom(addr)
}

func (b *LxsApiBackend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return b.xps.txPool.SubscribeNewTxsEvent(ch)
}

func (b *LxsApiBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.xps.blockchain.SubscribeChainEvent(ch)
}

func (b *LxsApiBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.xps.blockchain.SubscribeChainHeadEvent(ch)
}

func (b *LxsApiBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return b.xps.blockchain.SubscribeChainSideEvent(ch)
}

func (b *LxsApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.xps.blockchain.SubscribeLogsEvent(ch)
}

func (b *LxsApiBackend) SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		<-quit
		return nil
	})
}

func (b *LxsApiBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.xps.blockchain.SubscribeRemovedLogsEvent(ch)
}

func (b *LxsApiBackend) SyncProgress() xpayments.SyncProgress {
	return b.xps.Downloader().Progress()
}

func (b *LxsApiBackend) ProtocolVersion() int {
	return b.xps.LxsVersion() + 10000
}

func (b *LxsApiBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestTipCap(ctx)
}

func (b *LxsApiBackend) FeeHistory(ctx context.Context, blockCount int, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (firstBlock *big.Int, reward [][]*big.Int, baseFee []*big.Int, gasUsedRatio []float64, err error) {
	return b.gpo.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
}

func (b *LxsApiBackend) ChainDb() xpsdb.Database {
	return b.xps.chainDb
}

func (b *LxsApiBackend) AccountManager() *accounts.Manager {
	return b.xps.accountManager
}

func (b *LxsApiBackend) ExtRPCEnabled() bool {
	return b.extRPCEnabled
}

func (b *LxsApiBackend) UnprotectedAllowed() bool {
	return b.allowUnprotectedTxs
}

func (b *LxsApiBackend) RPCGasCap() uint64 {
	return b.xps.config.RPCGasCap
}

func (b *LxsApiBackend) RPCXVMTimeout() time.Duration {
	return b.xps.config.RPCXVMTimeout
}

func (b *LxsApiBackend) RPCTxFeeCap() float64 {
	return b.xps.config.RPCTxFeeCap
}

func (b *LxsApiBackend) BloomStatus() (uint64, uint64) {
	if b.xps.bloomIndexer == nil {
		return 0, 0
	}
	sections, _, _ := b.xps.bloomIndexer.Sections()
	return params.BloomBitsBlocksClient, sections
}

func (b *LxsApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.xps.bloomRequests)
	}
}

func (b *LxsApiBackend) Engine() consensus.Engine {
	return b.xps.engine
}

func (b *LxsApiBackend) CurrentHeader() *types.Header {
	return b.xps.blockchain.CurrentHeader()
}

func (b *LxsApiBackend) StateAtBlock(ctx context.Context, block *types.Block, reexec uint64, base *state.StateDB, checkLive bool, preferDisk bool) (*state.StateDB, error) {
	return b.xps.stateAtBlock(ctx, block, reexec)
}

func (b *LxsApiBackend) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (core.Message, vm.BlockContext, *state.StateDB, error) {
	return b.xps.stateAtTransaction(ctx, block, txIndex, reexec)
}
