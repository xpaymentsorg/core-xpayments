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

package types

import (
	"encoding/binary"
	"math/big"
	"sort"

	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/crypto"
)

// TenToTheFive - To be used while sorting xpos logs
//
// Sorted using ( blockNumber * (10 ** 5) + logIndex )
const TenToTheFiveXPoS uint64 = 100000

var (
	xposReceiptPrefix = []byte("xpayments-xpos-receipt-") // xposReceiptPrefix + number + block hash -> xpos block receipt

	// SystemAddress address for system sender
	SystemAddressXPoS = common.HexToAddress("0xffffFFFfFFffffffffffffffFfFFFfffFFFfFFfE")
)

// XPoSReceiptKey = xposReceiptPrefix + num (uint64 big endian) + hash
func XPoSReceiptKey(number uint64, hash common.Hash) []byte {
	enc := make([]byte, 8)
	binary.BigEndian.PutUint64(enc, number)
	return append(append(xposReceiptPrefix, enc...), hash.Bytes()...)
}

// GetDerivedXPoSTxHash get derived tx hash from receipt key
func GetDerivedXPoSTxHash(receiptKey []byte) common.Hash {
	return common.BytesToHash(crypto.Keccak256(receiptKey))
}

// NewXPoSTransaction create new xpos transaction for xpos receipt
func NewXPoSTransaction() *Transaction {
	return NewTransaction(0, common.Address{}, big.NewInt(0), 0, big.NewInt(0), make([]byte, 0))
}

// DeriveFieldsForXPoSReceipt fills the receipts with their computed fields based on consensus
// data and contextual infos like containing block and transactions.
func DeriveFieldsForXPoSReceipt(receipt *Receipt, hash common.Hash, number uint64, receipts Receipts) error {
	// get derived tx hash
	txHash := GetDerivedXPoSTxHash(XPoSReceiptKey(number, hash))
	txIndex := uint(len(receipts))

	// set tx hash and tx index
	receipt.TxHash = txHash
	receipt.TransactionIndex = txIndex
	receipt.BlockHash = hash
	receipt.BlockNumber = big.NewInt(0).SetUint64(number)

	logIndex := 0
	for i := 0; i < len(receipts); i++ {
		logIndex += len(receipts[i].Logs)
	}

	// The derived log fields can simply be set from the block and transaction
	for j := 0; j < len(receipt.Logs); j++ {
		receipt.Logs[j].BlockNumber = number
		receipt.Logs[j].BlockHash = hash
		receipt.Logs[j].TxHash = txHash
		receipt.Logs[j].TxIndex = txIndex
		receipt.Logs[j].Index = uint(logIndex)
		logIndex++
	}
	return nil
}

// DeriveFieldsForXPoSLogs fills the receipts with their computed fields based on consensus
// data and contextual infos like containing block and transactions.
func DeriveFieldsForXPoSLogs(logs []*Log, hash common.Hash, number uint64, txIndex uint, logIndex uint) {
	// get derived tx hash
	txHash := GetDerivedXPoSTxHash(XPoSReceiptKey(number, hash))

	// the derived log fields can simply be set from the block and transaction
	for j := 0; j < len(logs); j++ {
		logs[j].BlockNumber = number
		logs[j].BlockHash = hash
		logs[j].TxHash = txHash
		logs[j].TxIndex = txIndex
		logs[j].Index = logIndex
		logIndex++
	}
}

// MergeXPoSLogs merges receipt logs and block receipt logs
func MergeXPoSLogs(logs []*Log, xposLogs []*Log) []*Log {
	result := append(logs, xposLogs...)

	sort.SliceStable(result, func(i int, j int) bool {
		return (result[i].BlockNumber*TenToTheFiveXPoS + uint64(result[i].Index)) < (result[j].BlockNumber*TenToTheFiveXPoS + uint64(result[j].Index))
	})

	return result
}
