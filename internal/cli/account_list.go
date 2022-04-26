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

	"github.com/xpaymentsorg/go-xpayments/accounts"
	"github.com/xpaymentsorg/go-xpayments/internal/cli/flagset"
)

type AccountListCommand struct {
	*Meta
}

// Help implements the cli.Command interface
func (a *AccountListCommand) Help() string {
	return `Usage: bor account list

  List the local accounts.

  ` + a.Flags().Help()
}

func (a *AccountListCommand) Flags() *flagset.Flagset {
	return a.NewFlagSet("account list")
}

// Synopsis implements the cli.Command interface
func (a *AccountListCommand) Synopsis() string {
	return "List the local accounts"
}

// Run implements the cli.Command interface
func (a *AccountListCommand) Run(args []string) int {
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
	a.UI.Output(formatAccounts(keystore.Accounts()))
	return 0
}

func formatAccounts(accts []accounts.Account) string {
	if len(accts) == 0 {
		return "No accounts found"
	}

	rows := make([]string, len(accts)+1)
	rows[0] = "Index|Address"
	for i, d := range accts {
		rows[i+1] = fmt.Sprintf("%d|%s",
			i,
			d.Address.String())
	}
	return formatList(rows)
}
