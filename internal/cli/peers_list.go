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

package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/xpaymentsorg/go-xpayments/internal/cli/flagset"
	"github.com/xpaymentsorg/go-xpayments/internal/cli/server/proto"
)

// PeersListCommand is the command to group the peers commands
type PeersListCommand struct {
	*Meta2
}

// Help implements the cli.Command interface
func (p *PeersListCommand) Help() string {
	return `Usage: bor peers list

  Lists the connected peers

  ` + p.Flags().Help()
}

func (p *PeersListCommand) Flags() *flagset.Flagset {
	flags := p.NewFlagSet("peers list")

	return flags
}

// Synopsis implements the cli.Command interface
func (c *PeersListCommand) Synopsis() string {
	return ""
}

// Run implements the cli.Command interface
func (c *PeersListCommand) Run(args []string) int {
	flags := c.Flags()
	if err := flags.Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	borClt, err := c.BorConn()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	req := &proto.PeersListRequest{}
	resp, err := borClt.PeersList(context.Background(), req)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output(formatPeers(resp.Peers))
	return 0
}

func formatPeers(peers []*proto.Peer) string {
	if len(peers) == 0 {
		return "No peers found"
	}

	rows := make([]string, len(peers)+1)
	rows[0] = "ID|Enode|Name|Caps|Static|Trusted"
	for i, d := range peers {
		enode := strings.TrimPrefix(d.Enode, "enode://")

		rows[i+1] = fmt.Sprintf("%s|%s|%s|%s|%v|%v",
			d.Id,
			enode[:10],
			d.Name,
			strings.Join(d.Caps, ","),
			d.Static,
			d.Trusted)
	}
	return formatList(rows)
}
