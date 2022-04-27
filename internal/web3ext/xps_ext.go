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

package web3ext

// XPSJs xps related apis
const XPSJs = `
web3._extend({
	property: 'xps',
	methods: [
		new web3._extend.Method({
			name: 'getSnapshot',
			call: 'xps_getSnapshot',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getAuthor',
			call: 'xps_getAuthor',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getSnapshotAtHash',
			call: 'xps_getSnapshotAtHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getSigners',
			call: 'xps_getSigners',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getSignersAtHash',
			call: 'xps_getSignersAtHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getCurrentProposer',
			call: 'xps_getCurrentProposer',
			params: 0
		}),
		new web3._extend.Method({
			name: 'getCurrentValidators',
			call: 'xps_getCurrentValidators',
			params: 0
		}),
		new web3._extend.Method({
			name: 'getRootHash',
			call: 'xps_getRootHash',
			params: 2,
		}),
	]
});
`
