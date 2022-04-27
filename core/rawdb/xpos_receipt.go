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

package rawdb

import (
	"math/big"

	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/ethdb"
	"github.com/xpaymentsorg/go-xpayments/log"
	"github.com/xpaymentsorg/go-xpayments/rlp"
)

var (
	// xpos receipt key
	xposReceiptKey = types.XPoSReceiptKey

	// xpos derived tx hash
	getDerivedXPoSTxHash = types.GetDerivedXPoSTxHash

	// xposTxLookupPrefix + hash -> transaction/receipt lookup metadata
	xposTxLookupPrefix = []byte("xpayments-xpos-tx-lookup-")

	// freezerXPoSReceiptTable indicates the name of the freezer xpos receipts table.
	freezerXPoSReceiptTable = "xpayments-xpos-receipts"
)

// xposTxLookupKey = xposTxLookupPrefix + xpos tx hash
func xposTxLookupKey(hash common.Hash) []byte {
	return append(xposTxLookupPrefix, hash.Bytes()...)
}

// HasXPoSReceipt verifies the existence of all block receipt belonging
// to a block.
func HasXPoSReceipt(db ethdb.Reader, hash common.Hash, number uint64) bool {
	if has, err := db.Ancient(freezerHashTable, number); err == nil && common.BytesToHash(has) == hash {
		return true
	}

	if has, err := db.Has(xposReceiptKey(number, hash)); !has || err != nil {
		return false
	}

	return true
}

// ReadXPoSReceiptRLP retrieves the block receipt belonging to a block in RLP encoding.
func ReadXPoSReceiptRLP(db ethdb.Reader, hash common.Hash, number uint64) rlp.RawValue {
	// First try to look up the data in ancient database. Extra hash
	// comparison is necessary since ancient database only maintains
	// the canonical data.
	data, _ := db.Ancient(freezerXPoSReceiptTable, number)
	if len(data) > 0 {
		h, _ := db.Ancient(freezerHashTable, number)
		if common.BytesToHash(h) == hash {
			return data
		}
	}
	// Then try to look up the data in leveldb.
	data, _ = db.Get(xposReceiptKey(number, hash))
	if len(data) > 0 {
		return data
	}
	// In the background freezer is moving data from leveldb to flatten files.
	// So during the first check for ancient db, the data is not yet in there,
	// but when we reach into leveldb, the data was already moved. That would
	// result in a not found error.
	data, _ = db.Ancient(freezerXPoSReceiptTable, number)
	if len(data) > 0 {
		h, _ := db.Ancient(freezerHashTable, number)
		if common.BytesToHash(h) == hash {
			return data
		}
	}
	return nil // Can't find the data anywhere.
}

// ReadRawXPoSReceipt retrieves the block receipt belonging to a block.
// The receipt metadata fields are not guaranteed to be populated, so they
// should not be used. Use ReadXPoSReceipt instead if the metadata is needed.
func ReadRawXPoSReceipt(db ethdb.Reader, hash common.Hash, number uint64) *types.Receipt {
	// Retrieve the flattened receipt slice
	data := ReadXPoSReceiptRLP(db, hash, number)
	if data == nil || len(data) == 0 {
		return nil
	}

	// Convert the receipts from their storage form to their internal representation
	var storageReceipt types.ReceiptForStorage
	if err := rlp.DecodeBytes(data, &storageReceipt); err != nil {
		log.Error("Invalid receipt array RLP", "hash", hash, "err", err)
		return nil
	}

	return (*types.Receipt)(&storageReceipt)
}

// ReadXPoSReceipt retrieves all the xpos block receipts belonging to a block, including
// its correspoinding metadata fields. If it is unable to populate these metadata
// fields then nil is returned.
func ReadXPoSReceipt(db ethdb.Reader, hash common.Hash, number uint64) *types.Receipt {
	// We're deriving many fields from the block body, retrieve beside the receipt
	xposReceipt := ReadRawXPoSReceipt(db, hash, number)
	if xposReceipt == nil {
		return nil
	}

	// We're deriving many fields from the block body, retrieve beside the receipt
	receipts := ReadRawReceipts(db, hash, number)
	if receipts == nil {
		return nil
	}

	body := ReadBody(db, hash, number)
	if body == nil {
		log.Error("Missing body but have xpos receipt", "hash", hash, "number", number)
		return nil
	}

	if err := types.DeriveFieldsForXPoSReceipt(xposReceipt, hash, number, receipts); err != nil {
		log.Error("Failed to derive xpos receipt fields", "hash", hash, "number", number, "err", err)
		return nil
	}
	return xposReceipt
}

// WriteXPoSReceipt stores all the xpos receipt belonging to a block.
func WriteXPoSReceipt(db ethdb.KeyValueWriter, hash common.Hash, number uint64, xposReceipt *types.ReceiptForStorage) {
	// Convert the xpos receipt into their storage form and serialize them
	bytes, err := rlp.EncodeToBytes(xposReceipt)
	if err != nil {
		log.Crit("Failed to encode xpos receipt", "err", err)
	}

	// Store the flattened receipt slice
	if err := db.Put(xposReceiptKey(number, hash), bytes); err != nil {
		log.Crit("Failed to store xpos receipt", "err", err)
	}
}

// DeleteXPoSReceipt removes receipt data associated with a block hash.
func DeleteXPoSReceipt(db ethdb.KeyValueWriter, hash common.Hash, number uint64) {
	key := xposReceiptKey(number, hash)

	if err := db.Delete(key); err != nil {
		log.Crit("Failed to delete xpos receipt", "err", err)
	}
}

// ReadXPoSTransactionWithBlockHash retrieves a specific xpos (fake) transaction by tx hash and block hash, along with
// its added positional metadata.
func ReadXPoSTransactionWithBlockHash(db ethdb.Reader, txHash common.Hash, blockHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64) {
	blockNumber := ReadXPoSTxLookupEntry(db, txHash)
	if blockNumber == nil {
		return nil, common.Hash{}, 0, 0
	}

	body := ReadBody(db, blockHash, *blockNumber)
	if body == nil {
		log.Error("Transaction referenced missing", "number", blockNumber, "hash", blockHash)
		return nil, common.Hash{}, 0, 0
	}

	// fetch receipt and return it
	return types.NewXPoSTransaction(), blockHash, *blockNumber, uint64(len(body.Transactions))
}

// ReadXPoSTransaction retrieves a specific xpos (fake) transaction by hash, along with
// its added positional metadata.
func ReadXPoSTransaction(db ethdb.Reader, hash common.Hash) (*types.Transaction, common.Hash, uint64, uint64) {
	blockNumber := ReadXPoSTxLookupEntry(db, hash)
	if blockNumber == nil {
		return nil, common.Hash{}, 0, 0
	}

	blockHash := ReadCanonicalHash(db, *blockNumber)
	if blockHash == (common.Hash{}) {
		return nil, common.Hash{}, 0, 0
	}

	body := ReadBody(db, blockHash, *blockNumber)
	if body == nil {
		log.Error("Transaction referenced missing", "number", blockNumber, "hash", blockHash)
		return nil, common.Hash{}, 0, 0
	}

	// fetch receipt and return it
	return types.NewXPoSTransaction(), blockHash, *blockNumber, uint64(len(body.Transactions))
}

//
// Indexes for reverse lookup
//

// ReadXPoSTxLookupEntry retrieves the positional metadata associated with a transaction
// hash to allow retrieving the xpos transaction or xpos receipt using tx hash.
func ReadXPoSTxLookupEntry(db ethdb.Reader, txHash common.Hash) *uint64 {
	data, _ := db.Get(xposTxLookupKey(txHash))
	if len(data) == 0 {
		return nil
	}

	number := new(big.Int).SetBytes(data).Uint64()
	return &number
}

// WriteXPoSTxLookupEntry stores a positional metadata for xpos transaction using block hash and block number
func WriteXPoSTxLookupEntry(db ethdb.KeyValueWriter, hash common.Hash, number uint64) {
	txHash := types.GetDerivedXPoSTxHash(xposReceiptKey(number, hash))
	if err := db.Put(xposTxLookupKey(txHash), big.NewInt(0).SetUint64(number).Bytes()); err != nil {
		log.Crit("Failed to store xpos transaction lookup entry", "err", err)
	}
}

// DeleteXPoSTxLookupEntry removes xpos transaction data associated with block hash and block number
func DeleteXPoSTxLookupEntry(db ethdb.KeyValueWriter, hash common.Hash, number uint64) {
	txHash := types.GetDerivedXPoSTxHash(xposReceiptKey(number, hash))
	DeleteXPoSTxLookupEntryByTxHash(db, txHash)
}

// DeleteXPoSTxLookupEntryByTxHash removes xpos transaction data associated with a xpos tx hash.
func DeleteXPoSTxLookupEntryByTxHash(db ethdb.KeyValueWriter, txHash common.Hash) {
	if err := db.Delete(xposTxLookupKey(txHash)); err != nil {
		log.Crit("Failed to delete xpos transaction lookup entry", "err", err)
	}
}
