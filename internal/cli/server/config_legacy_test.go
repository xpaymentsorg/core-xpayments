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

	"github.com/stretchr/testify/assert"
)

func TestConfigLegacy(t *testing.T) {
	toml := `[Node.P2P]
StaticNodes = ["node1"]
TrustedNodes = ["node2"]`

	config, err := readLegacyConfig([]byte(toml))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, config.P2P.Discovery.StaticNodes, []string{"node1"})
	assert.Equal(t, config.P2P.Discovery.TrustedNodes, []string{"node2"})
}
