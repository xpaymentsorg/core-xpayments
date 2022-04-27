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

package filters

import (
	"bytes"
	"context"
	"errors"

	ethereum "github.com/xpaymentsorg/go-xpayments"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/params"
	"github.com/xpaymentsorg/go-xpayments/rpc"
)

// SetChainConfig sets chain config
func (api *PublicFilterAPI) SetChainConfig(chainConfig *params.ChainConfig) {
	api.chainConfig = chainConfig
}

func (api *PublicFilterAPI) GetXPoSBlockLogs(ctx context.Context, crit FilterCriteria) ([]*types.Log, error) {
	if api.chainConfig == nil {
		return nil, errors.New("No chain config found. Proper PublicFilterAPI initialization required")
	}

	// get sprint from xpos config
	sprint := api.chainConfig.XPoS.Sprint

	var filter *XPoSBlockLogsFilter
	if crit.BlockHash != nil {
		// Block filter requested, construct a single-shot filter
		filter = NewXPoSBlockLogsFilter(api.backend, sprint, *crit.BlockHash, crit.Addresses, crit.Topics)
	} else {
		// Convert the RPC block numbers into internal representations
		begin := rpc.LatestBlockNumber.Int64()
		if crit.FromBlock != nil {
			begin = crit.FromBlock.Int64()
		}
		end := rpc.LatestBlockNumber.Int64()
		if crit.ToBlock != nil {
			end = crit.ToBlock.Int64()
		}
		// Construct the range filter
		filter = NewXPoSBlockLogsRangeFilter(api.backend, sprint, begin, end, crit.Addresses, crit.Topics)
	}

	// Run the filter and return all the logs
	logs, err := filter.Logs(ctx)
	if err != nil {
		return nil, err
	}
	return returnLogs(logs), err
}

// NewDeposits send a notification each time a new deposit received from bridge.
func (api *PublicFilterAPI) NewDeposits(ctx context.Context, crit ethereum.StateSyncFilter) (*rpc.Subscription, error) {
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}

	rpcSub := notifier.CreateSubscription()
	go func() {
		stateSyncData := make(chan *types.StateSyncData)
		stateSyncSub := api.events.SubscribeNewDepositsXPoS(stateSyncData)

		for {
			select {
			case h := <-stateSyncData:
				if crit.ID == h.ID || bytes.Compare(crit.Contract.Bytes(), h.Contract.Bytes()) == 0 ||
					(crit.ID == 0 && crit.Contract == common.Address{}) {
					notifier.Notify(rpcSub.ID, h)
				}
			case <-rpcSub.Err():
				stateSyncSub.Unsubscribe()
				return
			case <-notifier.Closed():
				stateSyncSub.Unsubscribe()
				return
			}
		}
	}()

	return rpcSub, nil
}
