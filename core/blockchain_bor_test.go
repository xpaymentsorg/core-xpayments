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
	"testing"
	"time"

	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/consensus/ethash"
	"github.com/xpaymentsorg/go-xpayments/core/rawdb"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/core/vm"
	"github.com/xpaymentsorg/go-xpayments/crypto"
	"github.com/xpaymentsorg/go-xpayments/params"
)

func TestChain2HeadEvent(t *testing.T) {
	var (
		db      = rawdb.NewMemoryDatabase()
		key1, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
		addr1   = crypto.PubkeyToAddress(key1.PublicKey)
		gspec   = &Genesis{
			Config: params.TestChainConfig,
			Alloc:  GenesisAlloc{addr1: {Balance: big.NewInt(10000000000000000)}},
		}
		genesis = gspec.MustCommit(db)
		signer  = types.LatestSigner(gspec.Config)
	)

	blockchain, _ := NewBlockChain(db, nil, gspec.Config, ethash.NewFaker(), vm.Config{}, nil, nil)
	defer blockchain.Stop()

	chain2HeadCh := make(chan Chain2HeadEvent, 64)
	blockchain.SubscribeChain2HeadEvent(chain2HeadCh)

	chain, _ := GenerateChain(gspec.Config, genesis, ethash.NewFaker(), db, 3, func(i int, gen *BlockGen) {})
	if _, err := blockchain.InsertChain(chain); err != nil {
		t.Fatalf("failed to insert chain: %v", err)
	}

	replacementBlocks, _ := GenerateChain(gspec.Config, genesis, ethash.NewFaker(), db, 4, func(i int, gen *BlockGen) {
		tx, err := types.SignTx(types.NewContractCreation(gen.TxNonce(addr1), new(big.Int), 1000000, gen.header.BaseFee, nil), signer, key1)
		if i == 2 {
			gen.OffsetTime(-9)
		}
		if err != nil {
			t.Fatalf("failed to create tx: %v", err)
		}
		gen.AddTx(tx)
	})

	if _, err := blockchain.InsertChain(replacementBlocks); err != nil {
		t.Fatalf("failed to insert chain: %v", err)
	}

	type eventTest struct {
		Type    string
		Added   []common.Hash
		Removed []common.Hash
	}

	readEvent := func(expect *eventTest) {
		select {
		case ev := <-chain2HeadCh:
			if ev.Type != expect.Type {
				t.Fatal("Type mismatch")
			}

			if len(ev.NewChain) != len(expect.Added) {
				t.Fatal("Newchain and Added Array Size don't match")
			}
			if len(ev.OldChain) != len(expect.Removed) {
				t.Fatal("Oldchain and Removed Array Size don't match")
			}

			for j := 0; j < len(ev.OldChain); j++ {
				if ev.OldChain[j].Hash() != expect.Removed[j] {
					t.Fatal("Oldchain hashes Do Not Match")
				}
			}
			for j := 0; j < len(ev.NewChain); j++ {
				if ev.NewChain[j].Hash() != expect.Added[j] {
					t.Fatalf("Newchain hashes Do Not Match %s %s", ev.NewChain[j].Hash(), expect.Added[j])
				}
			}
		case <-time.After(2 * time.Second):
			t.Fatal("timeout")
		}
	}

	// head event
	readEvent(&eventTest{
		Type: Chain2HeadCanonicalEvent,
		Added: []common.Hash{
			chain[0].Hash(),
			chain[1].Hash(),
			chain[2].Hash(),
		}})

	// fork event
	readEvent(&eventTest{
		Type: Chain2HeadForkEvent,
		Added: []common.Hash{
			replacementBlocks[0].Hash(),
		}})

	// fork event
	readEvent(&eventTest{
		Type: Chain2HeadForkEvent,
		Added: []common.Hash{
			replacementBlocks[1].Hash(),
		}})

	// reorg event
	//In this event the channel recieves an array of Blocks in NewChain and OldChain
	readEvent(&eventTest{
		Type: Chain2HeadReorgEvent,
		Added: []common.Hash{
			replacementBlocks[2].Hash(),
			replacementBlocks[1].Hash(),
			replacementBlocks[0].Hash(),
		},
		Removed: []common.Hash{
			chain[2].Hash(),
			chain[1].Hash(),
			chain[0].Hash(),
		},
	})

	// head event
	readEvent(&eventTest{
		Type: Chain2HeadCanonicalEvent,
		Added: []common.Hash{
			replacementBlocks[2].Hash(),
			replacementBlocks[3].Hash(),
		}})

}
