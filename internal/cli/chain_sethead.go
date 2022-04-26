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
	"strconv"

	"github.com/xpaymentsorg/go-xpayments/internal/cli/flagset"
	"github.com/xpaymentsorg/go-xpayments/internal/cli/server/proto"
)

// ChainSetHeadCommand is the command to group the peers commands
type ChainSetHeadCommand struct {
	*Meta2

	yes bool
}

// Help implements the cli.Command interface
func (c *ChainSetHeadCommand) Help() string {
	return `Usage: bor chain sethead <number> [--yes]

  This command sets the current chain to a certain block`
}

func (c *ChainSetHeadCommand) Flags() *flagset.Flagset {
	flags := c.NewFlagSet("chain sethead")

	flags.BoolFlag(&flagset.BoolFlag{
		Name:    "yes",
		Usage:   "Force set head",
		Default: false,
		Value:   &c.yes,
	})
	return flags
}

// Synopsis implements the cli.Command interface
func (c *ChainSetHeadCommand) Synopsis() string {
	return "Set the new head of the chain"
}

// Run implements the cli.Command interface
func (c *ChainSetHeadCommand) Run(args []string) int {
	flags := c.Flags()
	if err := flags.Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	args = flags.Args()
	if len(args) != 1 {
		c.UI.Error("No number provided")
		return 1
	}

	borClt, err := c.BorConn()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	arg := args[0]
	fmt.Println(arg)

	number, err := strconv.Atoi(arg)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	if !c.yes {
		response, err := c.UI.Ask("Are you sure you want to reset the database? (y/n)")
		if err != nil {
			c.UI.Error(err.Error())
			return 1
		}
		if response != "y" {
			c.UI.Output("set head aborted")
			return 0
		}
	}

	if _, err := borClt.ChainSetHead(context.Background(), &proto.ChainSetHeadRequest{Number: uint64(number)}); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output("Done!")
	return 0
}
