// Copyright 2022 The go-xpayments Authors
// This file is part of the go-xpayments library.
//
// The go-xpayments library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-xpayments library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-xpayments library. If not, see <http://www.gnu.org/licenses/>.

package lxs

import (
	"sync"
	"time"

	"github.com/xpaymentsorg/go-xpayments/common/math"
	"github.com/xpaymentsorg/go-xpayments/core"
	"github.com/xpaymentsorg/go-xpayments/log"
	"github.com/xpaymentsorg/go-xpayments/xpsdb"
)

// pruner is responsible for pruning historical light chain data.
type pruner struct {
	db       xpsdb.Database
	indexers []*core.ChainIndexer
	closeCh  chan struct{}
	wg       sync.WaitGroup
}

// newPruner returns a light chain pruner instance.
func newPruner(db xpsdb.Database, indexers ...*core.ChainIndexer) *pruner {
	pruner := &pruner{
		db:       db,
		indexers: indexers,
		closeCh:  make(chan struct{}),
	}
	pruner.wg.Add(1)
	go pruner.loop()
	return pruner
}

// close notifies all background goroutines belonging to pruner to exit.
func (p *pruner) close() {
	close(p.closeCh)
	p.wg.Wait()
}

// loop periodically queries the status of chain indexers and prunes useless
// historical chain data. Notably, whenever Gpay restarts, it will iterate
// all historical sections even they don't exist at all(below checkpoint) so
// that light client can prune cached chain data that was ODRed after pruning
// that section.
func (p *pruner) loop() {
	defer p.wg.Done()

	// cleanTicker is the ticker used to trigger a history clean 2 times a day.
	var cleanTicker = time.NewTicker(12 * time.Hour)
	defer cleanTicker.Stop()

	// pruning finds the sections that have been processed by all indexers
	// and deletes all historical chain data.
	// Note, if some indexers don't support pruning(e.g. xps.BloomIndexer),
	// pruning operations can be silently ignored.
	pruning := func() {
		min := uint64(math.MaxUint64)
		for _, indexer := range p.indexers {
			sections, _, _ := indexer.Sections()
			if sections < min {
				min = sections
			}
		}
		// Always keep the latest section data in database.
		if min < 2 || len(p.indexers) == 0 {
			return
		}
		for _, indexer := range p.indexers {
			if err := indexer.Prune(min - 2); err != nil {
				log.Debug("Failed to prune historical data", "err", err)
				return
			}
		}
		p.db.Compact(nil, nil) // Compact entire database, ensure all removed data are deleted.
	}
	for {
		pruning()
		select {
		case <-cleanTicker.C:
		case <-p.closeCh:
			return
		}
	}
}
