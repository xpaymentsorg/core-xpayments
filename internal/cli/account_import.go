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

	"github.com/xpaymentsorg/go-xpayments/cmd/utils"
	"github.com/xpaymentsorg/go-xpayments/crypto"
	"github.com/xpaymentsorg/go-xpayments/internal/cli/flagset"
)

type AccountImportCommand struct {
	*Meta
}

// Help implements the cli.Command interface
func (a *AccountImportCommand) Help() string {
	return `Usage: bor account import

  Import a private key into a new account.

  Import an account:

    $ bor account import key.json

  ` + a.Flags().Help()
}

func (a *AccountImportCommand) Flags() *flagset.Flagset {
	return a.NewFlagSet("account import")
}

// Synopsis implements the cli.Command interface
func (a *AccountImportCommand) Synopsis() string {
	return "Import a private key into a new account"
}

// Run implements the cli.Command interface
func (a *AccountImportCommand) Run(args []string) int {
	flags := a.Flags()
	if err := flags.Parse(args); err != nil {
		a.UI.Error(err.Error())
		return 1
	}

	args = flags.Args()
	if len(args) != 1 {
		a.UI.Error("Expected one argument")
		return 1
	}
	key, err := crypto.LoadECDSA(args[0])
	if err != nil {
		a.UI.Error(fmt.Sprintf("Failed to load the private key '%s': %v", args[0], err))
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

	acct, err := keystore.ImportECDSA(key, password)
	if err != nil {
		utils.Fatalf("Could not create the account: %v", err)
	}
	a.UI.Output(fmt.Sprintf("Account created: %s", acct.Address.String()))
	return 0
}
