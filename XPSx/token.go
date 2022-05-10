package XPSx

import (
	"math/big"
	"strings"

	xpaymentsorg "github.com/xpaymentsorg/go-xpayments"
	"github.com/xpaymentsorg/go-xpayments/accounts/abi"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/consensus"
	"github.com/xpaymentsorg/go-xpayments/contracts/XPSx/contract"
	"github.com/xpaymentsorg/go-xpayments/core"
	"github.com/xpaymentsorg/go-xpayments/core/state"
	"github.com/xpaymentsorg/go-xpayments/log"
)

// GetTokenAbi return token abi
func GetTokenAbi() (*abi.ABI, error) {
	contractABI, err := abi.JSON(strings.NewReader(contract.TRC21ABI))
	if err != nil {
		return nil, err
	}
	return &contractABI, nil
}

// RunContract run smart contract
func RunContract(chain consensus.ChainContext, statedb *state.StateDB, contractAddr common.Address, abi *abi.ABI, method string, args ...interface{}) (interface{}, error) {
	input, err := abi.Pack(method)
	if err != nil {
		return nil, err
	}
	fakeCaller := common.HexToAddress("0x0000000000000000000000000000000000000001")
	msg := xpaymentsorg.CallMsg{To: &contractAddr, Data: input, From: fakeCaller}
	result, err := core.CallContractWithState(msg, chain, statedb)
	if err != nil {
		return nil, err
	}
	var unpackResult interface{}
	err = abi.Unpack(&unpackResult, method, result)
	if err != nil {
		return nil, err
	}
	return unpackResult, nil
}

func (XPSx *XPSX) GetTokenDecimal(chain consensus.ChainContext, statedb *state.StateDB, tokenAddr common.Address) (*big.Int, error) {
	if tokenDecimal, ok := XPSx.tokenDecimalCache.Get(tokenAddr); ok {
		return tokenDecimal.(*big.Int), nil
	}
	if tokenAddr.String() == common.XPSNativeAddress {
		XPSx.tokenDecimalCache.Add(tokenAddr, common.BasePrice)
		return common.BasePrice, nil
	}
	var decimals uint8
	defer func() {
		log.Debug("GetTokenDecimal from ", "relayerSMC", common.RelayerRegistrationSMC, "tokenAddr", tokenAddr.Hex(), "decimals", decimals)
	}()
	contractABI, err := GetTokenAbi()
	if err != nil {
		return nil, err
	}
	stateCopy := statedb.Copy()
	result, err := RunContract(chain, stateCopy, tokenAddr, contractABI, "decimals")
	if err != nil {
		return nil, err
	}
	decimals = result.(uint8)

	tokenDecimal := new(big.Int).SetUint64(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	XPSx.tokenDecimalCache.Add(tokenAddr, tokenDecimal)
	return tokenDecimal, nil
}

// FIXME: using in unit tests only
func (XPSx *XPSX) SetTokenDecimal(token common.Address, decimal *big.Int) {
	XPSx.tokenDecimalCache.Add(token, decimal)
}
