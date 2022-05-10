// Copyright 2016 The go-ethereum Authors
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

import xpaymentsorg "github.com/xpaymentsorg/go-xpayments"

// Verify that Client implements the ethereum interfaces.
var (
	_ = xpaymentsorg.ChainReader(&Client{})
	_ = xpaymentsorg.TransactionReader(&Client{})
	_ = xpaymentsorg.ChainStateReader(&Client{})
	_ = xpaymentsorg.ChainSyncReader(&Client{})
	_ = xpaymentsorg.ContractCaller(&Client{})
	_ = xpaymentsorg.GasEstimator(&Client{})
	_ = xpaymentsorg.GasPricer(&Client{})
	_ = xpaymentsorg.LogFilterer(&Client{})
	_ = xpaymentsorg.PendingStateReader(&Client{})
	// _ = ethereum.PendingStateEventer(&Client{})
	_ = xpaymentsorg.PendingContractCaller(&Client{})
)
