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
	"bytes"

	"github.com/naoina/toml"
)

type legacyConfig struct {
	Node struct {
		P2P struct {
			StaticNodes  []string
			TrustedNodes []string
		}
	}
}

func (l *legacyConfig) Config() *Config {
	c := DefaultConfig()
	c.P2P.Discovery.StaticNodes = l.Node.P2P.StaticNodes
	c.P2P.Discovery.TrustedNodes = l.Node.P2P.TrustedNodes
	return c
}

func readLegacyConfig(data []byte) (*Config, error) {
	var legacy legacyConfig

	r := toml.NewDecoder(bytes.NewReader(data))
	if err := r.Decode(&legacy); err != nil {
		return nil, err
	}
	return legacy.Config(), nil
}
