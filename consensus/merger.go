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

package consensus

import (
	"fmt"
	"sync"

	"github.com/xpaymentsorg/go-xpayments/core/rawdb"
	"github.com/xpaymentsorg/go-xpayments/log"
	"github.com/xpaymentsorg/go-xpayments/rlp"
	"github.com/xpaymentsorg/go-xpayments/xpsdb"
)

// transitionStatus describes the status of xps1/2 transition. This switch
// between modes is a one-way action which is triggered by corresponding
// consensus-layer message.
type transitionStatus struct {
	LeftPoW    bool // The flag is set when the first NewHead message received
	EnteredPoS bool // The flag is set when the first FinalisedBlock message received
}

// Merger is an internal help structure used to track the xps1/2 transition status.
// It's a common structure can be used in both full node and light client.
type Merger struct {
	db     xpsdb.KeyValueStore
	status transitionStatus
	mu     sync.RWMutex
}

// NewMerger creates a new Merger which stores its transition status in the provided db.
func NewMerger(db xpsdb.KeyValueStore) *Merger {
	var status transitionStatus
	blob := rawdb.ReadTransitionStatus(db)
	if len(blob) != 0 {
		if err := rlp.DecodeBytes(blob, &status); err != nil {
			log.Crit("Failed to decode the transition status", "err", err)
		}
	}
	return &Merger{
		db:     db,
		status: status,
	}
}

// ReachTTD is called whenever the first NewHead message received
// from the consensus-layer.
func (m *Merger) ReachTTD() {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.status.LeftPoW {
		return
	}
	m.status = transitionStatus{LeftPoW: true}
	blob, err := rlp.EncodeToBytes(m.status)
	if err != nil {
		panic(fmt.Sprintf("Failed to encode the transition status: %v", err))
	}
	rawdb.WriteTransitionStatus(m.db, blob)
	log.Info("Left PoW stage")
}

// FinalizePoS is called whenever the first FinalisedBlock message received
// from the consensus-layer.
func (m *Merger) FinalizePoS() {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.status.EnteredPoS {
		return
	}
	m.status = transitionStatus{LeftPoW: true, EnteredPoS: true}
	blob, err := rlp.EncodeToBytes(m.status)
	if err != nil {
		panic(fmt.Sprintf("Failed to encode the transition status: %v", err))
	}
	rawdb.WriteTransitionStatus(m.db, blob)
	log.Info("Entered PoS stage")
}

// TDDReached reports whether the chain has left the PoW stage.
func (m *Merger) TDDReached() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.status.LeftPoW
}

// PoSFinalized reports whether the chain has entered the PoS stage.
func (m *Merger) PoSFinalized() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.status.EnteredPoS
}
