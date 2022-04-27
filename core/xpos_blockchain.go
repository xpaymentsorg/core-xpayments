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

package core

import (
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/core/rawdb"
	"github.com/xpaymentsorg/go-xpayments/core/types"
)

// GetXPoSReceiptByHash retrieves the xpos block receipt in a given block.
func (bc *BlockChain) GetXPoSReceiptByHash(hash common.Hash) *types.Receipt {
	if receipt, ok := bc.xposReceiptsCache.Get(hash); ok {
		return receipt.(*types.Receipt)
	}

	// read header from hash
	number := rawdb.ReadHeaderNumber(bc.db, hash)
	if number == nil {
		return nil
	}

	// read xpos reciept by hash and number
	receipt := rawdb.ReadXPoSReceipt(bc.db, hash, *number)
	if receipt == nil {
		return nil
	}

	// add into xpos receipt cache
	bc.xposReceiptsCache.Add(hash, receipt)
	return receipt
}
