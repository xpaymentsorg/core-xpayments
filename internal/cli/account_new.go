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
	"fmt"

	"github.com/xpaymentsorg/go-xpayments/internal/cli/flagset"
)

type AccountNewCommand struct {
	*Meta
}

// Help implements the cli.Command interface
func (a *AccountNewCommand) Help() string {
	return `Usage: bor account new

  Create a new local account.

  ` + a.Flags().Help()
}

func (a *AccountNewCommand) Flags() *flagset.Flagset {
	return a.NewFlagSet("account new")
}

// Synopsis implements the cli.Command interface
func (a *AccountNewCommand) Synopsis() string {
	return "Create a new local account"
}

// Run implements the cli.Command interface
func (a *AccountNewCommand) Run(args []string) int {
	flags := a.Flags()
	if err := flags.Parse(args); err != nil {
		a.UI.Error(err.Error())
		return 1
	}

	keystore, err := a.GetKeystore()
	if err != nil {
		a.UI.Error(fmt.Sprintf("Failed to get keystore: %v", err))
		return 1
	}

	password, err := a.AskPassword()
	if err != nil {
		a.UI.Error(err.Error())
		return 1
	}

	account, err := keystore.NewAccount(password)
	if err != nil {
		a.UI.Error(fmt.Sprintf("Failed to create new account: %v", err))
		return 1
	}

	a.UI.Output("\nYour new key was generated")
	a.UI.Output(fmt.Sprintf("Public address of the key:   %s", account.Address.Hex()))
	a.UI.Output(fmt.Sprintf("Path of the secret key file: %s", account.URL.Path))

	return 0
}
