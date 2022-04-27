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

package eth

import (
	"context"
	"errors"

	ethereum "github.com/xpaymentsorg/go-xpayments"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/consensus/xpos"
	"github.com/xpaymentsorg/go-xpayments/core"
	"github.com/xpaymentsorg/go-xpayments/core/rawdb"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/event"
)

// GetRootHash returns root hash for given start and end block
func (b *EthAPIBackend) GetRootHash(ctx context.Context, starBlockNr uint64, endBlockNr uint64) (string, error) {
	var api *xpos.API
	for _, _api := range b.eth.Engine().APIs(b.eth.BlockChain()) {
		if _api.Namespace == "xpos" {
			api = _api.Service.(*xpos.API)
		}
	}

	if api == nil {
		return "", errors.New("Only available in XPoS engine")
	}

	root, err := api.GetRootHash(starBlockNr, endBlockNr)
	if err != nil {
		return "", err
	}
	return root, nil
}

// GetXPoSBlockReceipt returns xpos block receipt
func (b *EthAPIBackend) GetXPoSBlockReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	receipt := b.eth.blockchain.GetXPoSReceiptByHash(hash)
	if receipt == nil {
		return nil, ethereum.NotFound
	}

	return receipt, nil
}

// GetXPoSBlockLogs returns xpos block logs
func (b *EthAPIBackend) GetXPoSBlockLogs(ctx context.Context, hash common.Hash) ([]*types.Log, error) {
	receipt := b.eth.blockchain.GetXPoSReceiptByHash(hash)
	if receipt == nil {
		return nil, nil
	}
	return receipt.Logs, nil
}

// GetXPoSBlockTransaction returns xpos block tx
func (b *EthAPIBackend) GetXPoSBlockTransaction(ctx context.Context, hash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	tx, blockHash, blockNumber, index := rawdb.ReadXPoSTransaction(b.eth.ChainDb(), hash)
	return tx, blockHash, blockNumber, index, nil
}

func (b *EthAPIBackend) GetXPoSBlockTransactionWithBlockHash(ctx context.Context, txHash common.Hash, blockHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	tx, blockHash, blockNumber, index := rawdb.ReadXPoSTransactionWithBlockHash(b.eth.ChainDb(), txHash, blockHash)
	return tx, blockHash, blockNumber, index, nil
}

// SubscribeStateSyncEvent subscribes to state sync event
func (b *EthAPIBackend) SubscribeStateSyncEvent(ch chan<- core.StateSyncEvent) event.Subscription {
	return b.eth.BlockChain().SubscribeStateSyncEvent(ch)
}

// SubscribeChain2HeadEvent subscribes to reorg/head/fork event
func (b *EthAPIBackend) SubscribeChain2HeadEvent(ch chan<- core.Chain2HeadEvent) event.Subscription {
	return b.eth.BlockChain().SubscribeChain2HeadEvent(ch)
}
