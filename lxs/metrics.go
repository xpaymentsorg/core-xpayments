// Copyright 2022 The go-xpayments Authors
// This file is part of the go-xpayments library.
//
// The go-xpayments library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-xpayments library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-xpayments library. If not, see <http://www.gnu.org/licenses/>.

package lxs

import (
	"github.com/xpaymentsorg/go-xpayments/metrics"
	"github.com/xpaymentsorg/go-xpayments/p2p"
)

var (
	miscInPacketsMeter           = metrics.NewRegisteredMeter("lxs/misc/in/packets/total", nil)
	miscInTrafficMeter           = metrics.NewRegisteredMeter("lxs/misc/in/traffic/total", nil)
	miscInHeaderPacketsMeter     = metrics.NewRegisteredMeter("lxs/misc/in/packets/header", nil)
	miscInHeaderTrafficMeter     = metrics.NewRegisteredMeter("lxs/misc/in/traffic/header", nil)
	miscInBodyPacketsMeter       = metrics.NewRegisteredMeter("lxs/misc/in/packets/body", nil)
	miscInBodyTrafficMeter       = metrics.NewRegisteredMeter("lxs/misc/in/traffic/body", nil)
	miscInCodePacketsMeter       = metrics.NewRegisteredMeter("lxs/misc/in/packets/code", nil)
	miscInCodeTrafficMeter       = metrics.NewRegisteredMeter("lxs/misc/in/traffic/code", nil)
	miscInReceiptPacketsMeter    = metrics.NewRegisteredMeter("lxs/misc/in/packets/receipt", nil)
	miscInReceiptTrafficMeter    = metrics.NewRegisteredMeter("lxs/misc/in/traffic/receipt", nil)
	miscInTrieProofPacketsMeter  = metrics.NewRegisteredMeter("lxs/misc/in/packets/proof", nil)
	miscInTrieProofTrafficMeter  = metrics.NewRegisteredMeter("lxs/misc/in/traffic/proof", nil)
	miscInHelperTriePacketsMeter = metrics.NewRegisteredMeter("lxs/misc/in/packets/helperTrie", nil)
	miscInHelperTrieTrafficMeter = metrics.NewRegisteredMeter("lxs/misc/in/traffic/helperTrie", nil)
	miscInTxsPacketsMeter        = metrics.NewRegisteredMeter("lxs/misc/in/packets/txs", nil)
	miscInTxsTrafficMeter        = metrics.NewRegisteredMeter("lxs/misc/in/traffic/txs", nil)
	miscInTxStatusPacketsMeter   = metrics.NewRegisteredMeter("lxs/misc/in/packets/txStatus", nil)
	miscInTxStatusTrafficMeter   = metrics.NewRegisteredMeter("lxs/misc/in/traffic/txStatus", nil)

	miscOutPacketsMeter           = metrics.NewRegisteredMeter("lxs/misc/out/packets/total", nil)
	miscOutTrafficMeter           = metrics.NewRegisteredMeter("lxs/misc/out/traffic/total", nil)
	miscOutHeaderPacketsMeter     = metrics.NewRegisteredMeter("lxs/misc/out/packets/header", nil)
	miscOutHeaderTrafficMeter     = metrics.NewRegisteredMeter("lxs/misc/out/traffic/header", nil)
	miscOutBodyPacketsMeter       = metrics.NewRegisteredMeter("lxs/misc/out/packets/body", nil)
	miscOutBodyTrafficMeter       = metrics.NewRegisteredMeter("lxs/misc/out/traffic/body", nil)
	miscOutCodePacketsMeter       = metrics.NewRegisteredMeter("lxs/misc/out/packets/code", nil)
	miscOutCodeTrafficMeter       = metrics.NewRegisteredMeter("lxs/misc/out/traffic/code", nil)
	miscOutReceiptPacketsMeter    = metrics.NewRegisteredMeter("lxs/misc/out/packets/receipt", nil)
	miscOutReceiptTrafficMeter    = metrics.NewRegisteredMeter("lxs/misc/out/traffic/receipt", nil)
	miscOutTrieProofPacketsMeter  = metrics.NewRegisteredMeter("lxs/misc/out/packets/proof", nil)
	miscOutTrieProofTrafficMeter  = metrics.NewRegisteredMeter("lxs/misc/out/traffic/proof", nil)
	miscOutHelperTriePacketsMeter = metrics.NewRegisteredMeter("lxs/misc/out/packets/helperTrie", nil)
	miscOutHelperTrieTrafficMeter = metrics.NewRegisteredMeter("lxs/misc/out/traffic/helperTrie", nil)
	miscOutTxsPacketsMeter        = metrics.NewRegisteredMeter("lxs/misc/out/packets/txs", nil)
	miscOutTxsTrafficMeter        = metrics.NewRegisteredMeter("lxs/misc/out/traffic/txs", nil)
	miscOutTxStatusPacketsMeter   = metrics.NewRegisteredMeter("lxs/misc/out/packets/txStatus", nil)
	miscOutTxStatusTrafficMeter   = metrics.NewRegisteredMeter("lxs/misc/out/traffic/txStatus", nil)

	miscServingTimeHeaderTimer     = metrics.NewRegisteredTimer("lxs/misc/serve/header", nil)
	miscServingTimeBodyTimer       = metrics.NewRegisteredTimer("lxs/misc/serve/body", nil)
	miscServingTimeCodeTimer       = metrics.NewRegisteredTimer("lxs/misc/serve/code", nil)
	miscServingTimeReceiptTimer    = metrics.NewRegisteredTimer("lxs/misc/serve/receipt", nil)
	miscServingTimeTrieProofTimer  = metrics.NewRegisteredTimer("lxs/misc/serve/proof", nil)
	miscServingTimeHelperTrieTimer = metrics.NewRegisteredTimer("lxs/misc/serve/helperTrie", nil)
	miscServingTimeTxTimer         = metrics.NewRegisteredTimer("lxs/misc/serve/txs", nil)
	miscServingTimeTxStatusTimer   = metrics.NewRegisteredTimer("lxs/misc/serve/txStatus", nil)

	connectionTimer       = metrics.NewRegisteredTimer("lxs/connection/duration", nil)
	serverConnectionGauge = metrics.NewRegisteredGauge("lxs/connection/server", nil)

	totalCapacityGauge   = metrics.NewRegisteredGauge("lxs/server/totalCapacity", nil)
	totalRechargeGauge   = metrics.NewRegisteredGauge("lxs/server/totalRecharge", nil)
	blockProcessingTimer = metrics.NewRegisteredTimer("lxs/server/blockProcessingTime", nil)

	requestServedMeter               = metrics.NewRegisteredMeter("lxs/server/req/avgServedTime", nil)
	requestServedTimer               = metrics.NewRegisteredTimer("lxs/server/req/servedTime", nil)
	requestEstimatedMeter            = metrics.NewRegisteredMeter("lxs/server/req/avgEstimatedTime", nil)
	requestEstimatedTimer            = metrics.NewRegisteredTimer("lxs/server/req/estimatedTime", nil)
	relativeCostHistogram            = metrics.NewRegisteredHistogram("lxs/server/req/relative", nil, metrics.NewExpDecaySample(1028, 0.015))
	relativeCostHeaderHistogram      = metrics.NewRegisteredHistogram("lxs/server/req/relative/header", nil, metrics.NewExpDecaySample(1028, 0.015))
	relativeCostBodyHistogram        = metrics.NewRegisteredHistogram("lxs/server/req/relative/body", nil, metrics.NewExpDecaySample(1028, 0.015))
	relativeCostReceiptHistogram     = metrics.NewRegisteredHistogram("lxs/server/req/relative/receipt", nil, metrics.NewExpDecaySample(1028, 0.015))
	relativeCostCodeHistogram        = metrics.NewRegisteredHistogram("lxs/server/req/relative/code", nil, metrics.NewExpDecaySample(1028, 0.015))
	relativeCostProofHistogram       = metrics.NewRegisteredHistogram("lxs/server/req/relative/proof", nil, metrics.NewExpDecaySample(1028, 0.015))
	relativeCostHelperProofHistogram = metrics.NewRegisteredHistogram("lxs/server/req/relative/helperTrie", nil, metrics.NewExpDecaySample(1028, 0.015))
	relativeCostSendTxHistogram      = metrics.NewRegisteredHistogram("lxs/server/req/relative/txs", nil, metrics.NewExpDecaySample(1028, 0.015))
	relativeCostTxStatusHistogram    = metrics.NewRegisteredHistogram("lxs/server/req/relative/txStatus", nil, metrics.NewExpDecaySample(1028, 0.015))

	globalFactorGauge    = metrics.NewRegisteredGauge("lxs/server/globalFactor", nil)
	recentServedGauge    = metrics.NewRegisteredGauge("lxs/server/recentRequestServed", nil)
	recentEstimatedGauge = metrics.NewRegisteredGauge("lxs/server/recentRequestEstimated", nil)
	sqServedGauge        = metrics.NewRegisteredGauge("lxs/server/servingQueue/served", nil)
	sqQueuedGauge        = metrics.NewRegisteredGauge("lxs/server/servingQueue/queued", nil)

	clientFreezeMeter = metrics.NewRegisteredMeter("lxs/server/clientEvent/freeze", nil)
	clientErrorMeter  = metrics.NewRegisteredMeter("lxs/server/clientEvent/error", nil)

	requestRTT       = metrics.NewRegisteredTimer("lxs/client/req/rtt", nil)
	requestSendDelay = metrics.NewRegisteredTimer("lxs/client/req/sendDelay", nil)

	serverSelectableGauge = metrics.NewRegisteredGauge("lxs/client/serverPool/selectable", nil)
	serverDialedMeter     = metrics.NewRegisteredMeter("lxs/client/serverPool/dialed", nil)
	serverConnectedGauge  = metrics.NewRegisteredGauge("lxs/client/serverPool/connected", nil)
	sessionValueMeter     = metrics.NewRegisteredMeter("lxs/client/serverPool/sessionValue", nil)
	totalValueGauge       = metrics.NewRegisteredGauge("lxs/client/serverPool/totalValue", nil)
	suggestedTimeoutGauge = metrics.NewRegisteredGauge("lxs/client/serverPool/timeout", nil)
)

// meteredMsgReadWriter is a wrapper around a p2p.MsgReadWriter, capable of
// accumulating the above defined metrics based on the data stream contents.
type meteredMsgReadWriter struct {
	p2p.MsgReadWriter     // Wrapped message stream to meter
	version           int // Protocol version to select correct meters
}

// newMeteredMsgWriter wraps a p2p MsgReadWriter with metering support. If the
// metrics system is disabled, this function returns the original object.
func newMeteredMsgWriter(rw p2p.MsgReadWriter, version int) p2p.MsgReadWriter {
	if !metrics.Enabled {
		return rw
	}
	return &meteredMsgReadWriter{MsgReadWriter: rw, version: version}
}

func (rw *meteredMsgReadWriter) ReadMsg() (p2p.Msg, error) {
	// Read the message and short circuit in case of an error
	msg, err := rw.MsgReadWriter.ReadMsg()
	if err != nil {
		return msg, err
	}
	// Account for the data traffic
	packets, traffic := miscInPacketsMeter, miscInTrafficMeter
	packets.Mark(1)
	traffic.Mark(int64(msg.Size))

	return msg, err
}

func (rw *meteredMsgReadWriter) WriteMsg(msg p2p.Msg) error {
	// Account for the data traffic
	packets, traffic := miscOutPacketsMeter, miscOutTrafficMeter
	packets.Mark(1)
	traffic.Mark(int64(msg.Size))

	// Send the packet to the p2p layer
	return rw.MsgReadWriter.WriteMsg(msg)
}
