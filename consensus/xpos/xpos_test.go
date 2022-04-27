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

package xpos

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/common/hexutil"
	"github.com/xpaymentsorg/go-xpayments/core"
	"github.com/xpaymentsorg/go-xpayments/core/rawdb"
	"github.com/xpaymentsorg/go-xpayments/core/state"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/core/vm"
	"github.com/xpaymentsorg/go-xpayments/params"
)

func TestGenesisContractChange(t *testing.T) {
	addr0 := common.Address{0x1}

	b := &XPoS{
		config: &params.XPoSConfig{
			Sprint: 10, // skip sprint transactions in sprint
			BlockAlloc: map[string]interface{}{
				// write as interface since that is how it is decoded in genesis
				"2": map[string]interface{}{
					addr0.Hex(): map[string]interface{}{
						"code":    hexutil.Bytes{0x1, 0x2},
						"balance": "0",
					},
				},
				"4": map[string]interface{}{
					addr0.Hex(): map[string]interface{}{
						"code":    hexutil.Bytes{0x1, 0x3},
						"balance": "0x1000",
					},
				},
			},
		},
	}

	genspec := &core.Genesis{
		Alloc: map[common.Address]core.GenesisAccount{
			addr0: {
				Balance: big.NewInt(0),
				Code:    []byte{0x1, 0x1},
			},
		},
	}

	db := rawdb.NewMemoryDatabase()
	genesis := genspec.MustCommit(db)

	statedb, err := state.New(genesis.Root(), state.NewDatabase(db), nil)
	assert.NoError(t, err)

	config := params.ChainConfig{}
	chain, err := core.NewBlockChain(db, nil, &config, b, vm.Config{}, nil, nil)
	assert.NoError(t, err)

	addBlock := func(root common.Hash, num int64) (common.Hash, *state.StateDB) {
		h := &types.Header{
			ParentHash: root,
			Number:     big.NewInt(num),
		}
		b.Finalize(chain, h, statedb, nil, nil)

		// write state to database
		root, err := statedb.Commit(false)
		assert.NoError(t, err)
		assert.NoError(t, statedb.Database().TrieDB().Commit(root, true, nil))

		statedb, err := state.New(h.Root, state.NewDatabase(db), nil)
		assert.NoError(t, err)

		return root, statedb
	}

	assert.Equal(t, statedb.GetCode(addr0), []byte{0x1, 0x1})

	root := genesis.Root()

	// code does not change
	root, statedb = addBlock(root, 1)
	assert.Equal(t, statedb.GetCode(addr0), []byte{0x1, 0x1})

	// code changes 1st time
	root, statedb = addBlock(root, 2)
	assert.Equal(t, statedb.GetCode(addr0), []byte{0x1, 0x2})

	// code same as 1st change
	root, statedb = addBlock(root, 3)
	assert.Equal(t, statedb.GetCode(addr0), []byte{0x1, 0x2})

	// code changes 2nd time
	_, statedb = addBlock(root, 4)
	assert.Equal(t, statedb.GetCode(addr0), []byte{0x1, 0x3})

	// make sure balance change DOES NOT take effect
	assert.Equal(t, statedb.GetBalance(addr0), big.NewInt(0))
}

func TestEncodeSigHeaderJaipur(t *testing.T) {
	// As part of the EIP-1559 fork in mumbai, an incorrect seal hash
	// was used for XPoS that did not included the BaseFee. The Jaipur
	// block is a hard fork to fix that.
	h := &types.Header{
		Difficulty: new(big.Int),
		Number:     big.NewInt(1),
		Extra:      make([]byte, 32+65),
	}

	var (
		// hash for the block without the BaseFee
		hashWithoutBaseFee = common.HexToHash("0x1be13e83939b3c4701ee57a34e10c9290ce07b0e53af0fe90b812c6881826e36")
		// hash for the block with the baseFee
		hashWithBaseFee = common.HexToHash("0xc55b0cac99161f71bde1423a091426b1b5b4d7598e5981ad802cce712771965b")
	)

	// Jaipur NOT enabled and BaseFee not set
	hash := SealHash(h, &params.XPoSConfig{JaipurBlock: 10})
	assert.Equal(t, hash, hashWithoutBaseFee)

	// Jaipur enabled (Jaipur=0) and BaseFee not set
	hash = SealHash(h, &params.XPoSConfig{JaipurBlock: 0})
	assert.Equal(t, hash, hashWithoutBaseFee)

	h.BaseFee = big.NewInt(2)

	// Jaipur enabled (Jaipur=Header block) and BaseFee set
	hash = SealHash(h, &params.XPoSConfig{JaipurBlock: 1})
	assert.Equal(t, hash, hashWithBaseFee)

	// Jaipur NOT enabled and BaseFee set
	hash = SealHash(h, &params.XPoSConfig{JaipurBlock: 10})
	assert.Equal(t, hash, hashWithoutBaseFee)
}
