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
	"github.com/mitchellh/cli"
)

// PeersCommand is the command to group the peers commands
type PeersCommand struct {
	UI cli.Ui
}

// Help implements the cli.Command interface
func (c *PeersCommand) Help() string {
	return `Usage: bor peers <subcommand>

  This command groups actions to interact with peers.
	
  List the connected peers:
  
    $ bor peers list
	
  Add a new peer by enode:
  
    $ bor peers add <enode>

  Remove a connected peer by enode:

    $ bor peers remove <enode>

  Display information about a peer:

    $ bor peers status <peer id>`
}

// Synopsis implements the cli.Command interface
func (c *PeersCommand) Synopsis() string {
	return "Interact with peers"
}

// Run implements the cli.Command interface
func (c *PeersCommand) Run(args []string) int {
	return cli.RunResultHelp
}
