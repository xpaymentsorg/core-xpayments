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

	"github.com/xpaymentsorg/go-xpayments/internal/cli/flagset"
	"github.com/xpaymentsorg/go-xpayments/internal/cli/server/proto"
)

// PeersRemoveCommand is the command to group the peers commands
type PeersRemoveCommand struct {
	*Meta2

	trusted bool
}

// Help implements the cli.Command interface
func (p *PeersRemoveCommand) Help() string {
	return `Usage: bor peers remove <enode>

  Disconnects the local client from a connected peer if exists.

  ` + p.Flags().Help()
}

func (p *PeersRemoveCommand) Flags() *flagset.Flagset {
	flags := p.NewFlagSet("peers remove")

	flags.BoolFlag(&flagset.BoolFlag{
		Name:  "trusted",
		Usage: "Add the peer as a trusted",
		Value: &p.trusted,
	})

	return flags
}

// Synopsis implements the cli.Command interface
func (c *PeersRemoveCommand) Synopsis() string {
	return "Disconnects a peer from the client"
}

// Run implements the cli.Command interface
func (c *PeersRemoveCommand) Run(args []string) int {
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

	req := &proto.PeersRemoveRequest{
		Enode:   args[0],
		Trusted: c.trusted,
	}
	if _, err := borClt.PeersRemove(context.Background(), req); err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	return 0
}
