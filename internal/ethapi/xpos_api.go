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

package ethapi

import (
	"context"

	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/core/types"
)

// GetRootHash returns root hash for given start and end block
func (s *PublicBlockChainAPI) GetXPoSRootHash(ctx context.Context, starBlockNr uint64, endBlockNr uint64) (string, error) {
	root, err := s.b.GetRootHash(ctx, starBlockNr, endBlockNr)
	if err != nil {
		return "", err
	}
	return root, nil
}

func (s *PublicBlockChainAPI) GetXPoSBlockReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	return s.b.GetXPoSBlockReceipt(ctx, hash)
}

//
// XPoS transaction utils
//

func (s *PublicBlockChainAPI) appendRPCMarshalXPoSTransaction(ctx context.Context, block *types.Block, fields map[string]interface{}, fullTx bool) map[string]interface{} {
	if block != nil {
		txHash := types.GetDerivedXPoSTxHash(types.XPoSReceiptKey(block.Number().Uint64(), block.Hash()))
		xposTx, blockHash, blockNumber, txIndex, _ := s.b.GetXPoSBlockTransactionWithBlockHash(ctx, txHash, block.Hash())
		if xposTx != nil {
			formattedTxs := fields["transactions"].([]interface{})
			if fullTx {
				marshalledTx := newRPCTransaction(xposTx, blockHash, blockNumber, txIndex, block.BaseFee(), s.b.ChainConfig())
				// newRPCTransaction calculates hash based on RLP of the transaction data.
				// In case of xpos block tx, we need simple derived tx hash (same as function argument) instead of RLP hash
				marshalledTx.Hash = txHash
				fields["transactions"] = append(formattedTxs, marshalledTx)
			} else {
				fields["transactions"] = append(formattedTxs, txHash)
			}
		}
	}
	return fields
}
