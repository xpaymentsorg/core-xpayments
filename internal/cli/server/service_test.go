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
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xpaymentsorg/go-xpayments/internal/cli/server/proto"
)

func TestGatherBlocks(t *testing.T) {
	type c struct {
		ABlock *big.Int
		BBlock *big.Int
	}
	type d struct {
		DBlock uint64
	}
	val := &c{
		BBlock: new(big.Int).SetInt64(1),
	}
	val2 := &d{
		DBlock: 10,
	}

	expect := []*proto.StatusResponse_Fork{
		{
			Name:     "A",
			Disabled: true,
		},
		{
			Name:  "B",
			Block: 1,
		},
		{
			Name:  "D",
			Block: 10,
		},
	}

	res := gatherForks(val, val2)
	assert.Equal(t, res, expect)
}
