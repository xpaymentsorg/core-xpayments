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

package filters

import (
	"context"

	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/ethdb"
	"github.com/xpaymentsorg/go-xpayments/rpc"
)

// XPoSBlockLogsFilter can be used to retrieve and filter logs.
type XPoSBlockLogsFilter struct {
	backend Backend
	sprint  uint64

	db        ethdb.Database
	addresses []common.Address
	topics    [][]common.Hash

	block      common.Hash // Block hash if filtering a single block
	begin, end int64       // Range interval if filtering multiple blocks
}

// NewXPoSBlockLogsRangeFilter creates a new filter which uses a bloom filter on blocks to
// figure out whether a particular block is interesting or not.
func NewXPoSBlockLogsRangeFilter(backend Backend, sprint uint64, begin, end int64, addresses []common.Address, topics [][]common.Hash) *XPoSBlockLogsFilter {
	// Create a generic filter and convert it into a range filter
	filter := newXPoSBlockLogsFilter(backend, sprint, addresses, topics)
	filter.begin = begin
	filter.end = end

	return filter
}

// NewXPoSBlockLogsFilter creates a new filter which directly inspects the contents of
// a block to figure out whether it is interesting or not.
func NewXPoSBlockLogsFilter(backend Backend, sprint uint64, block common.Hash, addresses []common.Address, topics [][]common.Hash) *XPoSBlockLogsFilter {
	// Create a generic filter and convert it into a block filter
	filter := newXPoSBlockLogsFilter(backend, sprint, addresses, topics)
	filter.block = block
	return filter
}

// newXPoSBlockLogsFilter creates a generic filter that can either filter based on a block hash,
// or based on range queries. The search criteria needs to be explicitly set.
func newXPoSBlockLogsFilter(backend Backend, sprint uint64, addresses []common.Address, topics [][]common.Hash) *XPoSBlockLogsFilter {
	return &XPoSBlockLogsFilter{
		backend:   backend,
		sprint:    sprint,
		addresses: addresses,
		topics:    topics,
		db:        backend.ChainDb(),
	}
}

// Logs searches the blockchain for matching log entries, returning all from the
// first block that contains matches, updating the start of the filter accordingly.
func (f *XPoSBlockLogsFilter) Logs(ctx context.Context) ([]*types.Log, error) {
	// If we're doing singleton block filtering, execute and return
	if f.block != (common.Hash{}) {
		receipt, _ := f.backend.GetXPoSBlockReceipt(ctx, f.block)
		if receipt == nil {
			return nil, nil
		}
		return f.xposBlockLogs(ctx, receipt)
	}

	// Figure out the limits of the filter range
	header, _ := f.backend.HeaderByNumber(ctx, rpc.LatestBlockNumber)
	if header == nil {
		return nil, nil
	}
	head := header.Number.Uint64()

	if f.begin == -1 {
		f.begin = int64(head)
	}

	// adjust begin for sprint
	f.begin = currentSprintEnd(f.sprint, f.begin)

	end := f.end
	if f.end == -1 {
		end = int64(head)
	}

	// Gather all indexed logs, and finish with non indexed ones
	return f.unindexedLogs(ctx, uint64(end))
}

// unindexedLogs returns the logs matching the filter criteria based on raw block
// iteration and bloom matching.
func (f *XPoSBlockLogsFilter) unindexedLogs(ctx context.Context, end uint64) ([]*types.Log, error) {
	var logs []*types.Log

	for ; f.begin <= int64(end); f.begin = f.begin + 64 {
		header, err := f.backend.HeaderByNumber(ctx, rpc.BlockNumber(f.begin))
		if header == nil || err != nil {
			return logs, err
		}

		// get xpos block receipt
		receipt, err := f.backend.GetXPoSBlockReceipt(ctx, header.Hash())
		if receipt == nil || err != nil {
			continue
		}

		// filter xpos block logs
		found, err := f.xposBlockLogs(ctx, receipt)
		if err != nil {
			return logs, err
		}
		logs = append(logs, found...)
	}
	return logs, nil
}

// xposBlockLogs returns the logs matching the filter criteria within a single block.
func (f *XPoSBlockLogsFilter) xposBlockLogs(ctx context.Context, receipt *types.Receipt) (logs []*types.Log, err error) {
	if bloomFilter(receipt.Bloom, f.addresses, f.topics) {
		logs = filterLogs(receipt.Logs, nil, nil, f.addresses, f.topics)
	}
	return logs, nil
}

func currentSprintEnd(sprint uint64, n int64) int64 {
	m := n % int64(sprint)
	if m == 0 {
		return n
	}

	return n + 64 - m
}
