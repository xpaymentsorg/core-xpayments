package XPSx

import (
	"math/big"

	"github.com/xpaymentsorg/go-xpayments/accounts/abi/bind"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/contracts/XPSx/contract"
)

type TRC21Issuer struct {
	*contract.TRC21IssuerSession
	contractBackend bind.ContractBackend
}

func NewTRC21Issuer(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*TRC21Issuer, error) {
	contractObject, err := contract.NewTRC21Issuer(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &TRC21Issuer{
		&contract.TRC21IssuerSession{
			Contract:     contractObject,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployTRC21Issuer(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, minApply *big.Int) (common.Address, *TRC21Issuer, error) {
	contractAddr, _, _, err := contract.DeployTRC21Issuer(transactOpts, contractBackend, minApply)
	if err != nil {
		return contractAddr, nil, err
	}
	contractObject, err := NewTRC21Issuer(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, contractObject, nil
}
