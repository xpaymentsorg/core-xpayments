package XPSx

import (
	"github.com/xpaymentsorg/go-xpayments/accounts/abi/bind"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/contracts/XPSx/contract"
)

type LendingRelayerRegistration struct {
	*contract.LendingSession
	contractBackend bind.ContractBackend
}

func NewLendingRelayerRegistration(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*LendingRelayerRegistration, error) {
	smartContract, err := contract.NewLending(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &LendingRelayerRegistration{
		&contract.LendingSession{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployLendingRelayerRegistration(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, relayerAddr common.Address, XPSxListtingAddr common.Address) (common.Address, *LendingRelayerRegistration, error) {
	contractAddr, _, _, err := contract.DeployLending(transactOpts, contractBackend, relayerAddr, XPSxListtingAddr)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewLendingRelayerRegistration(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
