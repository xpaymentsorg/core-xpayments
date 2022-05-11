package common

import (
	"math/big"
)

const (
	RewardMasterPercent        = 90
	RewardVoterPercent         = 0
	RewardFoundationPercent    = 10
	HexSignMethod              = "e341eaa4"
	HexSetSecret               = "34d38600"
	HexSetOpening              = "e11f5ba2"
	EpocBlockSecret            = 800
	EpocBlockOpening           = 850
	EpocBlockRandomize         = 900
	MaxMasternodes             = 18
	MaxMasternodesV2           = 108
	LimitPenaltyEpoch          = 4
	BlocksPerYearTest          = uint64(200000)
	BlocksPerYear              = uint64(15768000)
	LimitThresholdNonceInQueue = 10
	DefaultMinGasPrice         = 250000000
	MergeSignRange             = 15
	RangeReturnSigner          = 150
	MinimunMinerBlockPerEpoch  = 1

	OneYear                    = uint64(365 * 86400)
	LiquidateLendingTradeBlock = uint64(100)
)

var Rewound = uint64(0)

var TIP2019Block = big.NewInt(1)
var TIPSigning = big.NewInt(3000000)
var TIPRandomize = big.NewInt(3464000)

var TIPIncreaseMasternodes = big.NewInt(5000000) // Upgrade MN Count at Block.
var TIPNoHalvingMNReward = big.NewInt(38383838)  // hardfork no halving masternodes reward
var BlackListHFNumber = uint64(38383838)
var TIPXPSX = big.NewInt(38383838)
var TIPXPSXLending = big.NewInt(38383838)
var TIPXPSXCancellationFee = big.NewInt(38383838)
var TIPXPSXCancellationFeeTestnet = big.NewInt(38383838)

var TIPXPSXTestnet = big.NewInt(38383838)
var IsTestnet bool = false
var StoreRewardFolder string
var RollbackHash Hash
var BasePrice = big.NewInt(1000000000000000000)                       // 1
var RelayerLockedFund = big.NewInt(20000)                             // 20000 XPS
var RelayerFee = big.NewInt(1000000000000000)                         // 0.001
var XPSXBaseFee = big.NewInt(10000)                                   // 1 / XPSXBaseFee
var RelayerCancelFee = big.NewInt(100000000000000)                    // 0.0001
var XPSXBaseCancelFee = new(big.Int).Mul(XPSXBaseFee, big.NewInt(10)) // 1/ (XPSXBaseFee *10)
var RelayerLendingFee = big.NewInt(10000000000000000)                 // 0.01
var RelayerLendingCancelFee = big.NewInt(1000000000000000)            // 0.001
var BaseLendingInterest = big.NewInt(100000000)                       // 1e8

var MinGasPrice = big.NewInt(DefaultMinGasPrice)
var RelayerRegistrationSMC = "0x16c63b79f9C8784168103C0b74E6A59EC2de4a02"
var RelayerRegistrationSMCTestnet = "0xA1996F69f47ba14Cb7f661010A7C31974277958c"
var LendingRegistrationSMC = "0x7d761afd7ff65a79e4173897594a194e3c506e57"
var LendingRegistrationSMCTestnet = "0x28d7fC2Cf5c18203aaCD7459EFC6Af0643C97bE8"
var TRC21IssuerSMCTestNet = HexToAddress("0x0E2C88753131CE01c7551B726b28BFD04e44003F")
var TRC21IssuerSMC = HexToAddress("0x8c0faeb5C6bEd2129b8674F262Fd45c4e9468bee")
var XPSXListingSMC = HexToAddress("0xDE34dD0f536170993E8CFF639DdFfCF1A85D3E53")
var XPSXListingSMCTestNet = HexToAddress("0x14B2Bf043b9c31827A472CE4F94294fE9a6277e0")
var TRC21GasPriceBefore = big.NewInt(2500)
var TRC21GasPrice = big.NewInt(250000000)
var RateTopUp = big.NewInt(90) // 90%
var BaseTopUp = big.NewInt(100)
var BaseRecall = big.NewInt(100)
var TIPTRC21Fee = big.NewInt(38383838)
var TIPTRC21FeeTestnet = big.NewInt(38383838)
var LimitTimeFinality = uint64(30) // limit in 30 block

var IgnoreSignerCheckBlockArray = map[uint64]bool{
	uint64(1032300):  true,
	uint64(1033200):  true,
	uint64(27307800): true,
	uint64(28270800): true,
}
var Blacklist = map[Address]bool{
	HexToAddress("0x5248bfb72fd4f234e062d3e9bb76f08643004fcd"): true,
}
