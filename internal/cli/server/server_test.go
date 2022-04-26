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

package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer_DeveloperMode(t *testing.T) {

	// get the default config
	config := DefaultConfig()

	// enable developer mode
	config.Developer.Enabled = true
	config.Developer.Period = 2 // block time

	// start the server
	server, err1 := NewServer(config)
	if err1 != nil {
		t.Fatalf("failed to start server: %v", err1)
	}

	// record the initial block number
	blockNumber := server.backend.BlockChain().CurrentBlock().Header().Number.Int64()

	var i int64 = 0
	for i = 0; i < 10; i++ {
		// We expect the node to mine blocks every `config.Developer.Period` time period
		time.Sleep(time.Duration(config.Developer.Period) * time.Second)
		currBlock := server.backend.BlockChain().CurrentBlock().Header().Number.Int64()
		expected := blockNumber + i + 1
		if res := assert.Equal(t, currBlock, expected); res == false {
			break
		}
	}

	// stop the server
	server.Stop()
}
