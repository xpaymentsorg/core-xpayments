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

package filters

import (
	"context"

	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/core/rawdb"
	"github.com/xpaymentsorg/go-xpayments/core/types"
)

func (b *testBackend) GetXPoSBlockReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	number := rawdb.ReadHeaderNumber(b.db, hash)
	if number == nil {
		return nil, nil
	}

	receipt := rawdb.ReadXPoSReceipt(b.db, hash, *number)
	if receipt == nil {
		return nil, nil
	}
	return receipt, nil
}

func (b *testBackend) GetXPoSBlockLogs(ctx context.Context, hash common.Hash) ([]*types.Log, error) {
	receipt, err := b.GetXPoSBlockReceipt(ctx, hash)
	if receipt == nil || err != nil {
		return nil, nil
	}
	return receipt.Logs, nil
}
