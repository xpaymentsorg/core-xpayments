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

// PeersStatusCommand is the command to group the peers commands
type PeersStatusCommand struct {
	*Meta2
}

// Help implements the cli.Command interface
func (p *PeersStatusCommand) Help() string {
	return `Usage: bor peers status <peer id>

  Display the status of a peer by its id.

  ` + p.Flags().Help()
}

func (p *PeersStatusCommand) Flags() *flagset.Flagset {
	flags := p.NewFlagSet("peers status")

	return flags
}

// Synopsis implements the cli.Command interface
func (c *PeersStatusCommand) Synopsis() string {
	return "Display the status of a peer"
}

// Run implements the cli.Command interface
func (c *PeersStatusCommand) Run(args []string) int {
	flags := c.Flags()
	if err := flags.Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	args = flags.Args()
	if len(args) != 1 {
		c.UI.Error("No enode address provided")
		return 1
	}

	borClt, err := c.BorConn()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	req := &proto.PeersStatusRequest{
		Enode: args[0],
	}
	resp, err := borClt.PeersStatus(context.Background(), req)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output(formatPeer(resp.Peer))
	return 0
}

func formatPeer(peer *proto.Peer) string {
	base := formatKV([]string{
		fmt.Sprintf("Name|%s", peer.Name),
		fmt.Sprintf("ID|%s", peer.Id),
		fmt.Sprintf("ENR|%s", peer.Enr),
		fmt.Sprintf("Capabilities|%s", strings.Join(peer.Caps, ",")),
		fmt.Sprintf("Enode|%s", peer.Enode),
		fmt.Sprintf("Static|%v", peer.Static),
		fmt.Sprintf("Trusted|%v", peer.Trusted),
	})
	return base
}
