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

package bor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"time"

	"github.com/xpaymentsorg/go-xpayments/log"
)

var (
	stateFetchLimit = 50
)

// ResponseWithHeight defines a response object type that wraps an original
// response with a height.
type ResponseWithHeight struct {
	Height string          `json:"height"`
	Result json.RawMessage `json:"result"`
}

type IHeimdallClient interface {
	Fetch(path string, query string) (*ResponseWithHeight, error)
	FetchWithRetry(path string, query string) (*ResponseWithHeight, error)
	FetchStateSyncEvents(fromID uint64, to int64) ([]*EventRecordWithTime, error)
	Close()
}

type HeimdallClient struct {
	urlString string
	client    http.Client
	closeCh   chan struct{}
}

func NewHeimdallClient(urlString string) (*HeimdallClient, error) {
	h := &HeimdallClient{
		urlString: urlString,
		client: http.Client{
			Timeout: time.Duration(5 * time.Second),
		},
		closeCh: make(chan struct{}),
	}
	return h, nil
}

func (h *HeimdallClient) FetchStateSyncEvents(fromID uint64, to int64) ([]*EventRecordWithTime, error) {
	eventRecords := make([]*EventRecordWithTime, 0)
	for {
		queryParams := fmt.Sprintf("from-id=%d&to-time=%d&limit=%d", fromID, to, stateFetchLimit)
		log.Info("Fetching state sync events", "queryParams", queryParams)
		response, err := h.FetchWithRetry("clerk/event-record/list", queryParams)
		if err != nil {
			return nil, err
		}
		var _eventRecords []*EventRecordWithTime
		if response.Result == nil { // status 204
			break
		}
		if err := json.Unmarshal(response.Result, &_eventRecords); err != nil {
			return nil, err
		}
		eventRecords = append(eventRecords, _eventRecords...)
		if len(_eventRecords) < stateFetchLimit {
			break
		}
		fromID += uint64(stateFetchLimit)
	}

	sort.SliceStable(eventRecords, func(i, j int) bool {
		return eventRecords[i].ID < eventRecords[j].ID
	})
	return eventRecords, nil
}

// Fetch fetches response from heimdall
func (h *HeimdallClient) Fetch(rawPath string, rawQuery string) (*ResponseWithHeight, error) {
	u, err := url.Parse(h.urlString)
	if err != nil {
		return nil, err
	}

	u.Path = rawPath
	u.RawQuery = rawQuery

	return h.internalFetch(u)
}

// FetchWithRetry returns data from heimdall with retry
func (h *HeimdallClient) FetchWithRetry(rawPath string, rawQuery string) (*ResponseWithHeight, error) {
	u, err := url.Parse(h.urlString)
	if err != nil {
		return nil, err
	}

	u.Path = rawPath
	u.RawQuery = rawQuery

	// create a new ticker for retrying the request
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-h.closeCh:
			log.Debug("Shutdown detected, terminating request")
			return nil, errShutdownDetected
		case <-ticker.C:
			res, err := h.internalFetch(u)
			if err == nil && res != nil {
				return res, nil
			}
			log.Info("Retrying again in 5 seconds to fetch data from Heimdall", "path", u.Path)
		}
	}
}

// internal fetch method
func (h *HeimdallClient) internalFetch(u *url.URL) (*ResponseWithHeight, error) {
	res, err := h.client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// check status code
	if res.StatusCode != 200 && res.StatusCode != 204 {
		return nil, fmt.Errorf("Error while fetching data from Heimdall")
	}

	// unmarshall data from buffer
	var response ResponseWithHeight
	if res.StatusCode == 204 {
		return &response, nil
	}

	// get response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Close sends a signal to stop the running process
func (h *HeimdallClient) Close() {
	close(h.closeCh)
	h.client.CloseIdleConnections()
}
