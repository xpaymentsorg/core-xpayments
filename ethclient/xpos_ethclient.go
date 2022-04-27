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

package ethclient

import (
	"context"

	ethereum "github.com/xpaymentsorg/go-xpayments"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/core/types"
)

// GetRootHash returns the merkle root of the block headers
func (ec *Client) GetRootHash(ctx context.Context, startBlockNumber uint64, endBlockNumber uint64) (string, error) {
	var rootHash string
	if err := ec.c.CallContext(ctx, &rootHash, "eth_getRootHash", startBlockNumber, endBlockNumber); err != nil {
		return "", err
	}
	return rootHash, nil
}

// GetXPoSBlockReceipt returns bor block receipt
func (ec *Client) GetXPoSBlockReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	var r *types.Receipt
	err := ec.c.CallContext(ctx, &r, "eth_getXPoSBlockReceipt", hash)
	if err == nil && r == nil {
		return nil, ethereum.NotFound
	}
	return r, err
}
