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
	"math/big"

	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/core/vm"
)

var transferLogSig = common.HexToHash("0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4")
var transferFeeLogSig = common.HexToHash("0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63")
var feeAddress = common.HexToAddress("0x0000000000000000000000000000000000001010")
var bigZero = big.NewInt(0)

// AddTransferLog adds transfer log into state
func AddTransferLog(
	state vm.StateDB,

	sender,
	recipient common.Address,

	amount,
	input1,
	input2,
	output1,
	output2 *big.Int,
) {
	addTransferLog(
		state,
		transferLogSig,

		sender,
		recipient,

		amount,
		input1,
		input2,
		output1,
		output2,
	)
}

// AddFeeTransferLog adds transfer log into state
// Deprecating transfer log and will be removed in future fork. PLEASE DO NOT USE this transfer log going forward. Parameters won't get updated as expected going forward with EIP1559
func AddFeeTransferLog(
	state vm.StateDB,

	sender,
	recipient common.Address,

	amount,
	input1,
	input2,
	output1,
	output2 *big.Int,
) {
	addTransferLog(
		state,
		transferFeeLogSig,

		sender,
		recipient,

		amount,
		input1,
		input2,
		output1,
		output2,
	)
}

// addTransferLog adds transfer log into state
func addTransferLog(
	state vm.StateDB,
	eventSig common.Hash,

	sender,
	recipient common.Address,

	amount,
	input1,
	input2,
	output1,
	output2 *big.Int,
) {
	// ignore if amount is 0
	if amount.Cmp(bigZero) <= 0 {
		return
	}

	dataInputs := []*big.Int{
		amount,
		input1,
		input2,
		output1,
		output2,
	}

	var data []byte
	for _, v := range dataInputs {
		data = append(data, common.LeftPadBytes(v.Bytes(), 32)...)
	}

	// add transfer log
	state.AddLog(&types.Log{
		Address: feeAddress,
		Topics: []common.Hash{
			eventSig,
			feeAddress.Hash(),
			sender.Hash(),
			recipient.Hash(),
		},
		Data: data,
	})
}
