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

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/xpaymentsorg/go-xpayments/internal/cli/server/proto"
)

// StatusCommand is the command to output the status of the client
type StatusCommand struct {
	*Meta2
}

// Help implements the cli.Command interface
func (p *StatusCommand) Help() string {
	return `Usage: bor status

  Output the status of the client`
}

// Synopsis implements the cli.Command interface
func (c *StatusCommand) Synopsis() string {
	return "Output the status of the client"
}

// Run implements the cli.Command interface
func (c *StatusCommand) Run(args []string) int {
	flags := c.NewFlagSet("status")
	if err := flags.Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	borClt, err := c.BorConn()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	status, err := borClt.Status(context.Background(), &empty.Empty{})
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output(printStatus(status))
	return 0
}

func printStatus(status *proto.StatusResponse) string {
	printHeader := func(h *proto.Header) string {
		return formatKV([]string{
			fmt.Sprintf("Hash|%s", h.Hash),
			fmt.Sprintf("Number|%d", h.Number),
		})
	}

	forks := make([]string, len(status.Forks)+1)
	forks[0] = "Name|Block|Enabled"
	for i, d := range status.Forks {
		forks[i+1] = fmt.Sprintf("%s|%d|%v", d.Name, d.Block, !d.Disabled)
	}

	full := []string{
		"General",
		formatKV([]string{
			fmt.Sprintf("Num peers|%d", status.NumPeers),
			fmt.Sprintf("Sync mode|%s", status.SyncMode),
		}),
		"\nCurrent Header",
		printHeader(status.CurrentHeader),
		"\nCurrent Block",
		printHeader(status.CurrentBlock),
		"\nSyncing",
		formatKV([]string{
			fmt.Sprintf("Current block|%d", status.Syncing.CurrentBlock),
			fmt.Sprintf("Highest block|%d", status.Syncing.HighestBlock),
			fmt.Sprintf("Starting block|%d", status.Syncing.StartingBlock),
		}),
		"\nForks",
		formatList(forks),
	}
	return strings.Join(full, "\n")
}
