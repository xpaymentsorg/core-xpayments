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
	"github.com/xpaymentsorg/go-xpayments/core/forkid"
	"github.com/xpaymentsorg/go-xpayments/p2p/dnsdisc"
	"github.com/xpaymentsorg/go-xpayments/p2p/enode"
	"github.com/xpaymentsorg/go-xpayments/rlp"
)

// lxsEntry is the "lxs" ENR entry. This is set for LXS servers only.
type lxsEntry struct {
	// Ignore additional fields (for forward compatibility).
	VfxVersion uint
	Rest       []rlp.RawValue `rlp:"tail"`
}

func (lxsEntry) ENRKey() string { return "lxs" }

// xpsEntry is the "xps" ENR entry. This is redeclared here to avoid depending on package xps.
type xpsEntry struct {
	ForkID forkid.ID
	Tail   []rlp.RawValue `rlp:"tail"`
}

func (xpsEntry) ENRKey() string { return "xps" }

// setupDiscovery creates the node discovery source for the xps protocol.
func (xps *LightxPayments) setupDiscovery() (enode.Iterator, error) {
	it := enode.NewFairMix(0)

	// Enable DNS discovery.
	if len(xps.config.XpsDiscoveryURLs) != 0 {
		client := dnsdisc.NewClient(dnsdisc.Config{})
		dns, err := client.NewIterator(xps.config.XpsDiscoveryURLs...)
		if err != nil {
			return nil, err
		}
		it.AddSource(dns)
	}

	// Enable DHT.
	if xps.udpEnabled {
		it.AddSource(xps.p2pServer.DiscV5.RandomNodes())
	}

	forkFilter := forkid.NewFilter(xps.blockchain)
	iterator := enode.Filter(it, func(n *enode.Node) bool { return nodeIsServer(forkFilter, n) })
	return iterator, nil
}

// nodeIsServer checks whether n is an LXS server node.
func nodeIsServer(forkFilter forkid.Filter, n *enode.Node) bool {
	var lxs lxsEntry
	var xps xpsEntry
	return n.Load(&lxs) == nil && n.Load(&xps) == nil && forkFilter(xps.ForkID) == nil
}
