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

package chains

import (
	"math/big"

	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/core"
	"github.com/xpaymentsorg/go-xpayments/crypto"
	"github.com/xpaymentsorg/go-xpayments/params"
)

// GetDeveloperChain returns the developer mode configs.
func GetDeveloperChain(period uint64, faucet common.Address) *Chain {
	// Override the default period to the user requested one
	config := *params.AllCliqueProtocolChanges
	config.Clique = &params.CliqueConfig{
		Period: period,
		Epoch:  config.Clique.Epoch,
	}

	// Assemble and return the chain having genesis with the
	// precompiles and faucet pre-funded
	return &Chain{
		Hash:      common.Hash{},
		NetworkId: 1337,
		Genesis: &core.Genesis{
			Config:     &config,
			ExtraData:  append(append(make([]byte, 32), faucet[:]...), make([]byte, crypto.SignatureLength)...),
			GasLimit:   11500000,
			BaseFee:    big.NewInt(params.InitialBaseFee),
			Difficulty: big.NewInt(1),
			Alloc: map[common.Address]core.GenesisAccount{
				common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
				common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
				common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
				common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
				common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
				common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
				common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
				common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
				common.BytesToAddress([]byte{9}): {Balance: big.NewInt(1)}, // BLAKE2b
				faucet:                           {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
			},
		},
		Bootnodes: []string{},
	}
}
