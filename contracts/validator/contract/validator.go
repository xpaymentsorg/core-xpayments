// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
        "math/big"
        "strings"

		ethereum "github.com/xpaymentsorg/go-xpayments"
		"github.com/xpaymentsorg/go-xpayments/accounts/abi"
		"github.com/xpaymentsorg/go-xpayments/accounts/abi/bind"
		"github.com/xpaymentsorg/go-xpayments/common"
		"github.com/xpaymentsorg/go-xpayments/core/types"
		"github.com/xpaymentsorg/go-xpayments/event"
)

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146060604052600080fd00a165627a7a7230582019c23e41af0e382ccd8b18ba701e220c6627022ba89c1cb306801ea4f6ab46070029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
        parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
        if err != nil {
                return common.Address{}, nil, nil, err
        }
        address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
        if err != nil {
                return common.Address{}, nil, nil, err
        }
        return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
        SafeMathCaller     // Read-only binding to the contract
        SafeMathTransactor // Write-only binding to the contract
        SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
        contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
        contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
        contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
        Contract     *SafeMath         // Generic contract binding to set the session for
        CallOpts     bind.CallOpts     // Call options to use throughout this session
        TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
        Contract *SafeMathCaller // Generic contract caller binding to set the session for
        CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
        Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
        TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
        Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
        Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
        Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
        contract, err := bindSafeMath(address, backend, backend, backend)
        if err != nil {
                return nil, err
        }
        return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
        contract, err := bindSafeMath(address, caller, nil, nil)
        if err != nil {
                return nil, err
        }
        return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
        contract, err := bindSafeMath(address, nil, transactor, nil)
        if err != nil {
                return nil, err
        }
        return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
        contract, err := bindSafeMath(address, nil, nil, filterer)
        if err != nil {
                return nil, err
        }
        return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
        parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
        if err != nil {
                return nil, err
        }
        return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
        return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
        return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
        return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
        return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
        return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
        return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// XPSValidatorABI is the input ABI used to generate the binding from.
const XPSValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasVotedInvalid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getWithdrawCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownerToCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithdrawBlockNumbers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getLatestKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_invalidCandidate\",\"type\":\"address\"}],\"name\":\"invalidPercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"KYCString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"invalidKYCCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"candidateCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voterWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"resign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getHashCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"candidateWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwnerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_invalidCandidate\",\"type\":\"address\"}],\"name\":\"voteInvalidKYC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kychash\",\"type\":\"string\"}],\"name\":\"uploadKYC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_candidates\",\"type\":\"address[]\"},{\"name\":\"_caps\",\"type\":\"uint256[]\"},{\"name\":\"_firstOwner\",\"type\":\"address\"},{\"name\":\"_minCandidateCap\",\"type\":\"uint256\"},{\"name\":\"_minVoterCap\",\"type\":\"uint256\"},{\"name\":\"_maxValidatorNumber\",\"type\":\"uint256\"},{\"name\":\"_candidateWithdrawDelay\",\"type\":\"uint256\"},{\"name\":\"_voterWithdrawDelay\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Unvote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"Resign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kycHash\",\"type\":\"string\"}],\"name\":\"UploadedKYC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_masternodeOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_masternodes\",\"type\":\"address[]\"}],\"name\":\"InvalidatedNode\",\"type\":\"event\"}]"

// XPSValidatorBin is the compiled bytecode used for deploying new contracts.
const XPSValidatorBin = `0x606060405260006009556000600a5534156200001a57600080fd5b604051620020ea380380620020ea833981016040528080518201919060200180518201919060200180519190602001805191906020018051919060200180519190602001805191906020018051600b879055600c869055600d859055600e849055600f819055915060009050885160095560078054600181016200009f83826200033e565b50600091825260208220018054600160a060020a031916600160a060020a038a16179055600a8054600101905590505b88518110156200032f576008805460018101620000ed83826200033e565b916000526020600020900160008b84815181106200010757fe5b90602001906020020151909190916101000a815481600160a060020a030219169083600160a060020a031602179055505060606040519081016040908152600160a060020a03891682526001602083015281018983815181106200016757fe5b906020019060200201519052600160008b84815181106200018457fe5b90602001906020020151600160a060020a03168152602081019190915260400160002081518154600160a060020a031916600160a060020a039190911617815560208201518154901515740100000000000000000000000000000000000000000260a060020a60ff0219909116178155604082015160019091015550600260008a83815181106200021157fe5b90602001906020020151600160a060020a0316815260208101919091526040016000208054600181016200024683826200033e565b5060009182526020808320919091018054600160a060020a031916600160a060020a038b16908117909155825260069052604090208054600181016200028d83826200033e565b916000526020600020900160008b8481518110620002a757fe5b90602001906020020151909190916101000a815481600160a060020a030219169083600160a060020a0316021790555050600b54600160008b8481518110620002ec57fe5b90602001906020020151600160a060020a03908116825260208083019390935260409182016000908120918c1681526002909101909252902055600101620000cf565b5050505050505050506200038e565b8154818355818115116200036557600083815260209020620003659181019083016200036a565b505050565b6200038b91905b8082111562000387576000815560010162000371565b5090565b90565b611d4c806200039e6000396000f3006060604052600436106101955763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301267951811461019a578063025e7c27146101b057806302aa9be2146101e257806306a49fce146102045780630db026221461026a5780630e3e4fb81461028f57806315febd68146102c85780632a3640b1146102de5780632d15cc04146103005780632f9c4bba1461031f578063302b68721461033257806332658652146103575780633477ee2e146103ed578063441a3e701461040357806358e7525f1461041c5780635b860d271461043b5780635b9cd8cc1461045a5780636dd7d8ea1461047c57806372e44a3814610490578063a9a981a3146104af578063a9ff959e146104c2578063ae6e43f5146104d5578063b642facd146104f4578063c45607df14610513578063d09f1ab414610532578063d161c76714610545578063d51b9e9314610558578063d55b7dff14610577578063ef18374a1461058a578063f2ee3c7d1461059d578063f5c95125146105bc578063f8ac9dd5146105da575b600080fd5b6101ae600160a060020a03600435166105ed565b005b34156101bb57600080fd5b6101c6600435610934565b604051600160a060020a03909116815260200160405180910390f35b34156101ed57600080fd5b6101ae600160a060020a036004351660243561095c565b341561020f57600080fd5b610217610b8f565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561025657808201518382015260200161023e565b505050509050019250505060405180910390f35b341561027557600080fd5b61027d610bf8565b60405190815260200160405180910390f35b341561029a57600080fd5b6102b4600160a060020a0360043581169060243516610bfe565b604051901515815260200160405180910390f35b34156102d357600080fd5b61027d600435610c1e565b34156102e957600080fd5b6101c6600160a060020a0360043516602435610c46565b341561030b57600080fd5b610217600160a060020a0360043516610c7d565b341561032a57600080fd5b610217610d0a565b341561033d57600080fd5b61027d600160a060020a0360043581169060243516610d8c565b341561036257600080fd5b610376600160a060020a0360043516610dbb565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156103b257808201518382015260200161039a565b50505050905090810190601f1680156103df5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103f857600080fd5b6101c6600435610f08565b341561040e57600080fd5b6101ae600435602435610f16565b341561042757600080fd5b61027d600160a060020a036004351661107d565b341561044657600080fd5b61027d600160a060020a036004351661109c565b341561046557600080fd5b610376600160a060020a036004351660243561110f565b6101ae600160a060020a03600435166111d8565b341561049b57600080fd5b61027d600160a060020a0360043516611388565b34156104ba57600080fd5b61027d61139a565b34156104cd57600080fd5b61027d6113a0565b34156104e057600080fd5b6101ae600160a060020a03600435166113a6565b34156104ff57600080fd5b6101c6600160a060020a0360043516611630565b341561051e57600080fd5b61027d600160a060020a036004351661164e565b341561053d57600080fd5b61027d611669565b341561055057600080fd5b61027d61166f565b341561056357600080fd5b6102b4600160a060020a0360043516611675565b341561058257600080fd5b61027d61169a565b341561059557600080fd5b61027d6116a0565b34156105a857600080fd5b6101ae600160a060020a03600435166116a6565b34156105c757600080fd5b6101ae6004803560248101910135611ab2565b34156105e557600080fd5b61027d611b55565b600b546000903410156105ff57600080fd5b600160a060020a03331660009081526003602052604090205415158061063b5750600160a060020a033316600090815260066020526040812054115b151561064657600080fd5b600160a060020a038216600090815260016020526040902054829060a060020a900460ff161561067557600080fd5b600160a060020a038316600090815260016020819052604090912001546106a2903463ffffffff611b5b16565b9150600880548060010182816106b89190611b83565b5060009182526020909120018054600160a060020a031916600160a060020a03851617905560606040519081016040908152600160a060020a03338116835260016020808501829052838501879052918716600090815291522081518154600160a060020a031916600160a060020a03919091161781556020820151815490151560a060020a0274ff0000000000000000000000000000000000000000199091161781556040820151600191820155600160a060020a0380861660009081526020928352604080822033909316825260029092019092529020546107a391503463ffffffff611b5b16565b600160a060020a03808516600090815260016020818152604080842033909516845260029094019052919020919091556009546107e59163ffffffff611b5b16565b600955600160a060020a03331660009081526006602052604090205415156108485760078054600181016108198382611b83565b5060009182526020909120018054600160a060020a03191633600160a060020a0316179055600a805460010190555b600160a060020a03331660009081526006602052604090208054600181016108708382611b83565b5060009182526020808320919091018054600160a060020a031916600160a060020a038716908117909155825260029052604090208054600181016108b58382611b83565b5060009182526020909120018054600160a060020a03191633600160a060020a038116919091179091557f7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1908434604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b600780548290811061094257fe5b600091825260209091200154600160a060020a0316905081565b600160a060020a0380831660009081526001602090815260408083203390941683526002909301905290812054839083908190101561099a57600080fd5b600160a060020a0382811660009081526001602052604090205433821691161415610a0857600b54600160a060020a0380841660009081526001602090815260408083203390941683526002909301905220546109fd908363ffffffff611b7116565b1015610a0857600080fd5b600160a060020a03851660009081526001602081905260409091200154610a35908563ffffffff611b7116565b600160a060020a038087166000908152600160208181526040808420928301959095553390931682526002019091522054610a76908563ffffffff611b7116565b600160a060020a038087166000908152600160209081526040808320339094168352600290930190522055600f54610ab4904363ffffffff611b5b16565b600160a060020a033316600090815260208181526040808320848452909152902054909350610ae9908563ffffffff611b5b16565b600160a060020a0333166000818152602081815260408083208884528083529083209490945591815290526001908101805490918101610b298382611b83565b5060009182526020909120018390557faa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2338686604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050505050565b610b97611bac565b6008805480602002602001604051908101604052809291908181526020018280548015610bed57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610bcf575b505050505090505b90565b600a5481565b600560209081526000928352604080842090915290825290205460ff1681565b600160a060020a0333166000908152602081815260408083208484529091529020545b919050565b600660205281600052604060002081815481101515610c6157fe5b600091825260209091200154600160a060020a03169150829050565b610c85611bac565b6002600083600160a060020a0316600160a060020a03168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610cfe57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610ce0575b50505050509050919050565b610d12611bac565b60008033600160a060020a0316600160a060020a03168152602001908152602001600020600101805480602002602001604051908101604052809291908181526020018280548015610bed57602002820191906000526020600020905b815481526020019060010190808311610d6f575050505050905090565b600160a060020a0391821660009081526001602090815260408083209390941682526002909201909152205490565b610dc3611bac565b610dcc82611675565b15610ee05760036000610dde84611630565b600160a060020a0316600160a060020a03168152602001908152602001600020600160036000610e0d86611630565b600160a060020a031681526020810191909152604001600020548254919003908110610e3557fe5b90600052602060002090018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ed45780601f10610ea957610100808354040283529160200191610ed4565b820191906000526020600020905b815481529060010190602001808311610eb757829003601f168201915b50505050509050610c41565b600160a060020a038216600090815260036020526040902080546000198101908110610e3557fe5b600880548290811061094257fe5b60008282828211610f2657600080fd5b4382901015610f3457600080fd5b600160a060020a03331660009081526020818152604080832085845290915281205411610f6057600080fd5b600160a060020a0333166000908152602081905260409020600101805483919083908110610f8a57fe5b60009182526020909120015414610fa057600080fd5b600160a060020a03331660008181526020818152604080832089845280835290832080549084905593835291905260010180549194509085908110610fe157fe5b6000918252602082200155600160a060020a03331683156108fc0284604051600060405180830381858888f19350505050151561101d57600080fd5b7ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b5683386856040518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390a15050505050565b600160a060020a03166000908152600160208190526040909120015490565b600160a060020a0381166000908152600160205260408120548190839060a060020a900460ff1615156110ce57600080fd5b6110d784611630565b91506110e16116a0565b600160a060020a03831660009081526004602052604090205460640281151561110657fe5b04949350505050565b60036020528160005260406000208181548110151561112a57fe5b9060005260206000209001600091509150508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111d05780601f106111a5576101008083540402835291602001916111d0565b820191906000526020600020905b8154815290600101906020018083116111b357829003601f168201915b505050505081565b600c543410156111e757600080fd5b600160a060020a038116600090815260016020526040902054819060a060020a900460ff16151561121757600080fd5b600160a060020a03821660009081526001602081905260409091200154611244903463ffffffff611b5b16565b600160a060020a03808416600090815260016020818152604080842092830195909555339093168252600201909152205415156112c957600160a060020a03821660009081526002602052604090208054600181016112a38382611b83565b5060009182526020909120018054600160a060020a03191633600160a060020a03161790555b600160a060020a038083166000908152600160209081526040808320339094168352600290930190522054611304903463ffffffff611b5b16565b600160a060020a03808416600090815260016020908152604080832033948516845260020190915290819020929092557f66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc918490349051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050565b60046020526000908152604090205481565b60095481565b600f5481565b600160a060020a0381811660009081526001602052604081205490918291829185913382169116146113d757600080fd5b600160a060020a038516600090815260016020526040902054859060a060020a900460ff16151561140757600080fd5b600160a060020a0386166000908152600160208190526040909120805474ff0000000000000000000000000000000000000000191690556009546114509163ffffffff611b7116565b600955600094505b6008548510156114cd5785600160a060020a031660088681548110151561147b57fe5b600091825260209091200154600160a060020a031614156114c25760088054869081106114a457fe5b60009182526020909120018054600160a060020a03191690556114cd565b600190940193611458565b600160a060020a03808716600081815260016020818152604080842033909616845260028601825283205493909252908190529190910154909450611518908563ffffffff611b7116565b600160a060020a0380881660009081526001602081815260408084209283019590955533909316825260020190915290812055600e5461155e904363ffffffff611b5b16565b600160a060020a033316600090815260208181526040808320848452909152902054909350611593908563ffffffff611b5b16565b600160a060020a03331660008181526020818152604080832088845280835290832094909455918152905260019081018054909181016115d38382611b83565b5060009182526020909120018390557f4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d33387604051600160a060020a039283168152911660208201526040908101905180910390a1505050505050565b600160a060020a039081166000908152600160205260409020541690565b600160a060020a031660009081526003602052604090205490565b600d5481565b600e5481565b600160a060020a031660009081526001602052604090205460a060020a900460ff1690565b600b5481565b600a5490565b6000806116b1611bac565b33600160a060020a0381166000908152600160205260408120549091829182919060a060020a900460ff1615156116e757600080fd5b600160a060020a038816600090815260016020526040902054889060a060020a900460ff16151561171757600080fd5b61172033611630565b975061172b89611630565b600160a060020a03808a1660009081526005602090815260408083209385168352929052205490975060ff161561176157600080fd5b600160a060020a038089166000908152600560209081526040808320938b168352928152828220805460ff19166001908117909155600490915291902080549091019055604b6117af6116a0565b600160a060020a0389166000908152600460205260409020546064028115156117d457fe5b0410611aa757600854600019016040518059106117ee5750595b9080825280602002602001820160405250955060009450600093505b60085484101561199d5786600160a060020a031661184a60088681548110151561183057fe5b600091825260209091200154600160a060020a0316611630565b600160a060020a031614156119925760095461186d90600163ffffffff611b7116565b600955600880548590811061187e57fe5b6000918252602090912001546001860195600160a060020a03909116908790815181106118a757fe5b600160a060020a0390921660209283029091019091015260088054859081106118cc57fe5b600091825260208220018054600160a060020a0319169055600880546001929190879081106118f757fe5b6000918252602080832090910154600160a060020a0390811684528382019490945260409283018220805474ffffffffffffffffffffffffffffffffffffffffff19168155600101829055928a1681526003909252812061195791611bbe565b600160a060020a038716600090815260066020526040812061197891611bdf565b600160a060020a0387166000908152600460205260408120555b60019093019261180a565b600092505b600754831015611a215786600160a060020a03166007848154811015156119c557fe5b600091825260209091200154600160a060020a03161415611a165760078054849081106119ee57fe5b60009182526020909120018054600160a060020a0319169055600a8054600019019055611a21565b6001909201916119a2565b7fe18d61a5bf4aa2ab40afc88aa9039d27ae17ff4ec1c65f5f414df6f02ce8b35e8787604051600160a060020a038316815260406020820181815290820183818151815260200191508051906020019060200280838360005b83811015611a92578082015183820152602001611a7a565b50505050905001935050505060405180910390a15b505050505050505050565b600160a060020a0333166000908152600360205260409020805460018101611ada8382611bfd565b60009283526020909220611af091018484611c21565b50507f949360d814b28a3b393a68909efe1fee120ee09cac30f360a0f80ab5415a611a338383604051600160a060020a038416815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a15050565b600c5481565b600082820183811015611b6a57fe5b9392505050565b600082821115611b7d57fe5b50900390565b815481835581811511611ba757600083815260209020611ba7918101908301611c9f565b505050565b60206040519081016040526000815290565b5080546000825590600052602060002090810190611bdc9190611cb9565b50565b5080546000825590600052602060002090810190611bdc9190611c9f565b815481835581811511611ba757600083815260209020611ba7918101908301611cb9565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611c625782800160ff19823516178555611c8f565b82800160010185558215611c8f579182015b82811115611c8f578235825591602001919060010190611c74565b50611c9b929150611c9f565b5090565b610bf591905b80821115611c9b5760008155600101611ca5565b610bf591905b80821115611c9b576000611cd38282611cdc565b50600101611cbf565b50805460018160011615610100020316600290046000825580601f10611d025750611bdc565b601f016020900490600052602060002090810190611bdc9190611c9f5600a165627a7a7230582030232256991da7e7e43ad1763531adcc407c6b390d388506d61590c9b2295f7e0029`

// DeployXPSValidator deploys a new Ethereum contract, binding an instance of XPSValidator to it.
func DeployXPSValidator(auth *bind.TransactOpts, backend bind.ContractBackend, _candidates []common.Address, _caps []*big.Int, _firstOwner common.Address, _minCandidateCap *big.Int, _minVoterCap *big.Int, _maxValidatorNumber *big.Int, _candidateWithdrawDelay *big.Int, _voterWithdrawDelay *big.Int) (common.Address, *types.Transaction, *XPSValidator, error) {
        parsed, err := abi.JSON(strings.NewReader(XPSValidatorABI))
        if err != nil {
                return common.Address{}, nil, nil, err
        }
        address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XPSValidatorBin), backend, _candidates, _caps, _firstOwner, _minCandidateCap, _minVoterCap, _maxValidatorNumber, _candidateWithdrawDelay, _voterWithdrawDelay)
        if err != nil {
                return common.Address{}, nil, nil, err
        }
        return address, tx, &XPSValidator{XPSValidatorCaller: XPSValidatorCaller{contract: contract}, XPSValidatorTransactor: XPSValidatorTransactor{contract: contract}, XPSValidatorFilterer: XPSValidatorFilterer{contract: contract}}, nil
}

// XPSValidator is an auto generated Go binding around an Ethereum contract.
type XPSValidator struct {
        XPSValidatorCaller     // Read-only binding to the contract
        XPSValidatorTransactor // Write-only binding to the contract
        XPSValidatorFilterer   // Log filterer for contract events
}

// XPSValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type XPSValidatorCaller struct {
        contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XPSValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XPSValidatorTransactor struct {
        contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XPSValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XPSValidatorFilterer struct {
        contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XPSValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XPSValidatorSession struct {
        Contract     *XPSValidator     // Generic contract binding to set the session for
        CallOpts     bind.CallOpts     // Call options to use throughout this session
        TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XPSValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XPSValidatorCallerSession struct {
        Contract *XPSValidatorCaller // Generic contract caller binding to set the session for
        CallOpts bind.CallOpts       // Call options to use throughout this session
}

// XPSValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XPSValidatorTransactorSession struct {
        Contract     *XPSValidatorTransactor // Generic contract transactor binding to set the session for
        TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// XPSValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type XPSValidatorRaw struct {
        Contract *XPSValidator // Generic contract binding to access the raw methods on
}

// XPSValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XPSValidatorCallerRaw struct {
        Contract *XPSValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// XPSValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XPSValidatorTransactorRaw struct {
        Contract *XPSValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXPSValidator creates a new instance of XPSValidator, bound to a specific deployed contract.
func NewXPSValidator(address common.Address, backend bind.ContractBackend) (*XPSValidator, error) {
        contract, err := bindXPSValidator(address, backend, backend, backend)
        if err != nil {
                return nil, err
        }
        return &XPSValidator{XPSValidatorCaller: XPSValidatorCaller{contract: contract}, XPSValidatorTransactor: XPSValidatorTransactor{contract: contract}, XPSValidatorFilterer: XPSValidatorFilterer{contract: contract}}, nil
}

// NewXPSValidatorCaller creates a new read-only instance of XPSValidator, bound to a specific deployed contract.
func NewXPSValidatorCaller(address common.Address, caller bind.ContractCaller) (*XPSValidatorCaller, error) {
        contract, err := bindXPSValidator(address, caller, nil, nil)
        if err != nil {
                return nil, err
        }
        return &XPSValidatorCaller{contract: contract}, nil
}

// NewXPSValidatorTransactor creates a new write-only instance of XPSValidator, bound to a specific deployed contract.
func NewXPSValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*XPSValidatorTransactor, error) {
        contract, err := bindXPSValidator(address, nil, transactor, nil)
        if err != nil {
                return nil, err
        }
        return &XPSValidatorTransactor{contract: contract}, nil
}

// NewXPSValidatorFilterer creates a new log filterer instance of XPSValidator, bound to a specific deployed contract.
func NewXPSValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*XPSValidatorFilterer, error) {
        contract, err := bindXPSValidator(address, nil, nil, filterer)
        if err != nil {
                return nil, err
        }
        return &XPSValidatorFilterer{contract: contract}, nil
}

// bindXPSValidator binds a generic wrapper to an already deployed contract.
func bindXPSValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
        parsed, err := abi.JSON(strings.NewReader(XPSValidatorABI))
        if err != nil {
                return nil, err
        }
        return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XPSValidator *XPSValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
        return _XPSValidator.Contract.XPSValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XPSValidator *XPSValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
        return _XPSValidator.Contract.XPSValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XPSValidator *XPSValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
        return _XPSValidator.Contract.XPSValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XPSValidator *XPSValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
        return _XPSValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XPSValidator *XPSValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
        return _XPSValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XPSValidator *XPSValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
        return _XPSValidator.Contract.contract.Transact(opts, method, params...)
}

// KYCString is a free data retrieval call binding the contract method 0x5b9cd8cc.
//
// Solidity: function KYCString( address,  uint256) constant returns(string)
func (_XPSValidator *XPSValidatorCaller) KYCString(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
        var (
                ret0 = new(string)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "KYCString", arg0, arg1)
        return *ret0, err
}

// KYCString is a free data retrieval call binding the contract method 0x5b9cd8cc.
//
// Solidity: function KYCString( address,  uint256) constant returns(string)
func (_XPSValidator *XPSValidatorSession) KYCString(arg0 common.Address, arg1 *big.Int) (string, error) {
        return _XPSValidator.Contract.KYCString(&_XPSValidator.CallOpts, arg0, arg1)
}

// KYCString is a free data retrieval call binding the contract method 0x5b9cd8cc.
//
// Solidity: function KYCString( address,  uint256) constant returns(string)
func (_XPSValidator *XPSValidatorCallerSession) KYCString(arg0 common.Address, arg1 *big.Int) (string, error) {
        return _XPSValidator.Contract.KYCString(&_XPSValidator.CallOpts, arg0, arg1)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) CandidateCount(opts *bind.CallOpts) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "candidateCount")
        return *ret0, err
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) CandidateCount() (*big.Int, error) {
        return _XPSValidator.Contract.CandidateCount(&_XPSValidator.CallOpts)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) CandidateCount() (*big.Int, error) {
        return _XPSValidator.Contract.CandidateCount(&_XPSValidator.CallOpts)
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) CandidateWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "candidateWithdrawDelay")
        return *ret0, err
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) CandidateWithdrawDelay() (*big.Int, error) {
        return _XPSValidator.Contract.CandidateWithdrawDelay(&_XPSValidator.CallOpts)
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) CandidateWithdrawDelay() (*big.Int, error) {
        return _XPSValidator.Contract.CandidateWithdrawDelay(&_XPSValidator.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_XPSValidator *XPSValidatorCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
        var (
                ret0 = new(common.Address)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "candidates", arg0)
        return *ret0, err
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_XPSValidator *XPSValidatorSession) Candidates(arg0 *big.Int) (common.Address, error) {
        return _XPSValidator.Contract.Candidates(&_XPSValidator.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_XPSValidator *XPSValidatorCallerSession) Candidates(arg0 *big.Int) (common.Address, error) {
        return _XPSValidator.Contract.Candidates(&_XPSValidator.CallOpts, arg0)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) GetCandidateCap(opts *bind.CallOpts, _candidate common.Address) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "getCandidateCap", _candidate)
        return *ret0, err
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
        return _XPSValidator.Contract.GetCandidateCap(&_XPSValidator.CallOpts, _candidate)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
        return _XPSValidator.Contract.GetCandidateCap(&_XPSValidator.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_XPSValidator *XPSValidatorCaller) GetCandidateOwner(opts *bind.CallOpts, _candidate common.Address) (common.Address, error) {
        var (
                ret0 = new(common.Address)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "getCandidateOwner", _candidate)
        return *ret0, err
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_XPSValidator *XPSValidatorSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
        return _XPSValidator.Contract.GetCandidateOwner(&_XPSValidator.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_XPSValidator *XPSValidatorCallerSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
        return _XPSValidator.Contract.GetCandidateOwner(&_XPSValidator.CallOpts, _candidate)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_XPSValidator *XPSValidatorCaller) GetCandidates(opts *bind.CallOpts) ([]common.Address, error) {
        var (
                ret0 = new([]common.Address)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "getCandidates")
        return *ret0, err
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_XPSValidator *XPSValidatorSession) GetCandidates() ([]common.Address, error) {
        return _XPSValidator.Contract.GetCandidates(&_XPSValidator.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_XPSValidator *XPSValidatorCallerSession) GetCandidates() ([]common.Address, error) {
        return _XPSValidator.Contract.GetCandidates(&_XPSValidator.CallOpts)
}

// GetHashCount is a free data retrieval call binding the contract method 0xc45607df.
//
// Solidity: function getHashCount(_address address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) GetHashCount(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "getHashCount", _address)
        return *ret0, err
}

// GetHashCount is a free data retrieval call binding the contract method 0xc45607df.
//
// Solidity: function getHashCount(_address address) constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) GetHashCount(_address common.Address) (*big.Int, error) {
        return _XPSValidator.Contract.GetHashCount(&_XPSValidator.CallOpts, _address)
}

// GetHashCount is a free data retrieval call binding the contract method 0xc45607df.
//
// Solidity: function getHashCount(_address address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) GetHashCount(_address common.Address) (*big.Int, error) {
        return _XPSValidator.Contract.GetHashCount(&_XPSValidator.CallOpts, _address)
}

// GetLatestKYC is a free data retrieval call binding the contract method 0x32658652.
//
// Solidity: function getLatestKYC(_address address) constant returns(string)
func (_XPSValidator *XPSValidatorCaller) GetLatestKYC(opts *bind.CallOpts, _address common.Address) (string, error) {
        var (
                ret0 = new(string)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "getLatestKYC", _address)
        return *ret0, err
}

// GetLatestKYC is a free data retrieval call binding the contract method 0x32658652.
//
// Solidity: function getLatestKYC(_address address) constant returns(string)
func (_XPSValidator *XPSValidatorSession) GetLatestKYC(_address common.Address) (string, error) {
        return _XPSValidator.Contract.GetLatestKYC(&_XPSValidator.CallOpts, _address)
}

// GetLatestKYC is a free data retrieval call binding the contract method 0x32658652.
//
// Solidity: function getLatestKYC(_address address) constant returns(string)
func (_XPSValidator *XPSValidatorCallerSession) GetLatestKYC(_address common.Address) (string, error) {
        return _XPSValidator.Contract.GetLatestKYC(&_XPSValidator.CallOpts, _address)
}

// GetOwnerCount is a free data retrieval call binding the contract method 0xef18374a.
//
// Solidity: function getOwnerCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) GetOwnerCount(opts *bind.CallOpts) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "getOwnerCount")
        return *ret0, err
}

// GetOwnerCount is a free data retrieval call binding the contract method 0xef18374a.
//
// Solidity: function getOwnerCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) GetOwnerCount() (*big.Int, error) {
        return _XPSValidator.Contract.GetOwnerCount(&_XPSValidator.CallOpts)
}

// GetOwnerCount is a free data retrieval call binding the contract method 0xef18374a.
//
// Solidity: function getOwnerCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) GetOwnerCount() (*big.Int, error) {
        return _XPSValidator.Contract.GetOwnerCount(&_XPSValidator.CallOpts)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) GetVoterCap(opts *bind.CallOpts, _candidate common.Address, _voter common.Address) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "getVoterCap", _candidate, _voter)
        return *ret0, err
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
        return _XPSValidator.Contract.GetVoterCap(&_XPSValidator.CallOpts, _candidate, _voter)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
        return _XPSValidator.Contract.GetVoterCap(&_XPSValidator.CallOpts, _candidate, _voter)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_XPSValidator *XPSValidatorCaller) GetVoters(opts *bind.CallOpts, _candidate common.Address) ([]common.Address, error) {
        var (
                ret0 = new([]common.Address)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "getVoters", _candidate)
        return *ret0, err
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_XPSValidator *XPSValidatorSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
        return _XPSValidator.Contract.GetVoters(&_XPSValidator.CallOpts, _candidate)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_XPSValidator *XPSValidatorCallerSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
        return _XPSValidator.Contract.GetVoters(&_XPSValidator.CallOpts, _candidate)
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_XPSValidator *XPSValidatorCaller) GetWithdrawBlockNumbers(opts *bind.CallOpts) ([]*big.Int, error) {
        var (
                ret0 = new([]*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "getWithdrawBlockNumbers")
        return *ret0, err
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_XPSValidator *XPSValidatorSession) GetWithdrawBlockNumbers() ([]*big.Int, error) {
        return _XPSValidator.Contract.GetWithdrawBlockNumbers(&_XPSValidator.CallOpts)
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_XPSValidator *XPSValidatorCallerSession) GetWithdrawBlockNumbers() ([]*big.Int, error) {
        return _XPSValidator.Contract.GetWithdrawBlockNumbers(&_XPSValidator.CallOpts)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) GetWithdrawCap(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "getWithdrawCap", _blockNumber)
        return *ret0, err
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) GetWithdrawCap(_blockNumber *big.Int) (*big.Int, error) {
        return _XPSValidator.Contract.GetWithdrawCap(&_XPSValidator.CallOpts, _blockNumber)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) GetWithdrawCap(_blockNumber *big.Int) (*big.Int, error) {
        return _XPSValidator.Contract.GetWithdrawCap(&_XPSValidator.CallOpts, _blockNumber)
}

// HasVotedInvalid is a free data retrieval call binding the contract method 0x0e3e4fb8.
//
// Solidity: function hasVotedInvalid( address,  address) constant returns(bool)
func (_XPSValidator *XPSValidatorCaller) HasVotedInvalid(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
        var (
                ret0 = new(bool)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "hasVotedInvalid", arg0, arg1)
        return *ret0, err
}

// HasVotedInvalid is a free data retrieval call binding the contract method 0x0e3e4fb8.
//
// Solidity: function hasVotedInvalid( address,  address) constant returns(bool)
func (_XPSValidator *XPSValidatorSession) HasVotedInvalid(arg0 common.Address, arg1 common.Address) (bool, error) {
        return _XPSValidator.Contract.HasVotedInvalid(&_XPSValidator.CallOpts, arg0, arg1)
}

// HasVotedInvalid is a free data retrieval call binding the contract method 0x0e3e4fb8.
//
// Solidity: function hasVotedInvalid( address,  address) constant returns(bool)
func (_XPSValidator *XPSValidatorCallerSession) HasVotedInvalid(arg0 common.Address, arg1 common.Address) (bool, error) {
        return _XPSValidator.Contract.HasVotedInvalid(&_XPSValidator.CallOpts, arg0, arg1)
}

// InvalidKYCCount is a free data retrieval call binding the contract method 0x72e44a38.
//
// Solidity: function invalidKYCCount( address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) InvalidKYCCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "invalidKYCCount", arg0)
        return *ret0, err
}

// InvalidKYCCount is a free data retrieval call binding the contract method 0x72e44a38.
//
// Solidity: function invalidKYCCount( address) constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) InvalidKYCCount(arg0 common.Address) (*big.Int, error) {
        return _XPSValidator.Contract.InvalidKYCCount(&_XPSValidator.CallOpts, arg0)
}

// InvalidKYCCount is a free data retrieval call binding the contract method 0x72e44a38.
//
// Solidity: function invalidKYCCount( address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) InvalidKYCCount(arg0 common.Address) (*big.Int, error) {
        return _XPSValidator.Contract.InvalidKYCCount(&_XPSValidator.CallOpts, arg0)
}

// InvalidPercent is a free data retrieval call binding the contract method 0x5b860d27.
//
// Solidity: function invalidPercent(_invalidCandidate address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) InvalidPercent(opts *bind.CallOpts, _invalidCandidate common.Address) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "invalidPercent", _invalidCandidate)
        return *ret0, err
}

// InvalidPercent is a free data retrieval call binding the contract method 0x5b860d27.
//
// Solidity: function invalidPercent(_invalidCandidate address) constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) InvalidPercent(_invalidCandidate common.Address) (*big.Int, error) {
        return _XPSValidator.Contract.InvalidPercent(&_XPSValidator.CallOpts, _invalidCandidate)
}

// InvalidPercent is a free data retrieval call binding the contract method 0x5b860d27.
//
// Solidity: function invalidPercent(_invalidCandidate address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) InvalidPercent(_invalidCandidate common.Address) (*big.Int, error) {
        return _XPSValidator.Contract.InvalidPercent(&_XPSValidator.CallOpts, _invalidCandidate)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_XPSValidator *XPSValidatorCaller) IsCandidate(opts *bind.CallOpts, _candidate common.Address) (bool, error) {
        var (
                ret0 = new(bool)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "isCandidate", _candidate)
        return *ret0, err
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_XPSValidator *XPSValidatorSession) IsCandidate(_candidate common.Address) (bool, error) {
        return _XPSValidator.Contract.IsCandidate(&_XPSValidator.CallOpts, _candidate)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_XPSValidator *XPSValidatorCallerSession) IsCandidate(_candidate common.Address) (bool, error) {
        return _XPSValidator.Contract.IsCandidate(&_XPSValidator.CallOpts, _candidate)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) MaxValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "maxValidatorNumber")
        return *ret0, err
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) MaxValidatorNumber() (*big.Int, error) {
        return _XPSValidator.Contract.MaxValidatorNumber(&_XPSValidator.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) MaxValidatorNumber() (*big.Int, error) {
        return _XPSValidator.Contract.MaxValidatorNumber(&_XPSValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) MinCandidateCap(opts *bind.CallOpts) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "minCandidateCap")
        return *ret0, err
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) MinCandidateCap() (*big.Int, error) {
        return _XPSValidator.Contract.MinCandidateCap(&_XPSValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) MinCandidateCap() (*big.Int, error) {
        return _XPSValidator.Contract.MinCandidateCap(&_XPSValidator.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) MinVoterCap(opts *bind.CallOpts) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "minVoterCap")
        return *ret0, err
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) MinVoterCap() (*big.Int, error) {
        return _XPSValidator.Contract.MinVoterCap(&_XPSValidator.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) MinVoterCap() (*big.Int, error) {
        return _XPSValidator.Contract.MinVoterCap(&_XPSValidator.CallOpts)
}

// OwnerCount is a free data retrieval call binding the contract method 0x0db02622.
//
// Solidity: function ownerCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) OwnerCount(opts *bind.CallOpts) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "ownerCount")
        return *ret0, err
}

// OwnerCount is a free data retrieval call binding the contract method 0x0db02622.
//
// Solidity: function ownerCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) OwnerCount() (*big.Int, error) {
        return _XPSValidator.Contract.OwnerCount(&_XPSValidator.CallOpts)
}

// OwnerCount is a free data retrieval call binding the contract method 0x0db02622.
//
// Solidity: function ownerCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) OwnerCount() (*big.Int, error) {
        return _XPSValidator.Contract.OwnerCount(&_XPSValidator.CallOpts)
}

// OwnerToCandidate is a free data retrieval call binding the contract method 0x2a3640b1.
//
// Solidity: function ownerToCandidate( address,  uint256) constant returns(address)
func (_XPSValidator *XPSValidatorCaller) OwnerToCandidate(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
        var (
                ret0 = new(common.Address)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "ownerToCandidate", arg0, arg1)
        return *ret0, err
}

// OwnerToCandidate is a free data retrieval call binding the contract method 0x2a3640b1.
//
// Solidity: function ownerToCandidate( address,  uint256) constant returns(address)
func (_XPSValidator *XPSValidatorSession) OwnerToCandidate(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
        return _XPSValidator.Contract.OwnerToCandidate(&_XPSValidator.CallOpts, arg0, arg1)
}

// OwnerToCandidate is a free data retrieval call binding the contract method 0x2a3640b1.
//
// Solidity: function ownerToCandidate( address,  uint256) constant returns(address)
func (_XPSValidator *XPSValidatorCallerSession) OwnerToCandidate(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
        return _XPSValidator.Contract.OwnerToCandidate(&_XPSValidator.CallOpts, arg0, arg1)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_XPSValidator *XPSValidatorCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
        var (
                ret0 = new(common.Address)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "owners", arg0)
        return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_XPSValidator *XPSValidatorSession) Owners(arg0 *big.Int) (common.Address, error) {
        return _XPSValidator.Contract.Owners(&_XPSValidator.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_XPSValidator *XPSValidatorCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
        return _XPSValidator.Contract.Owners(&_XPSValidator.CallOpts, arg0)
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) VoterWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
        var (
                ret0 = new(*big.Int)
        )
        out := ret0
        err := _XPSValidator.contract.Call(opts, out, "voterWithdrawDelay")
        return *ret0, err
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) VoterWithdrawDelay() (*big.Int, error) {
        return _XPSValidator.Contract.VoterWithdrawDelay(&_XPSValidator.CallOpts)
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) VoterWithdrawDelay() (*big.Int, error) {
        return _XPSValidator.Contract.VoterWithdrawDelay(&_XPSValidator.CallOpts)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_XPSValidator *XPSValidatorTransactor) Propose(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
        return _XPSValidator.contract.Transact(opts, "propose", _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_XPSValidator *XPSValidatorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
        return _XPSValidator.Contract.Propose(&_XPSValidator.TransactOpts, _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_XPSValidator *XPSValidatorTransactorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
        return _XPSValidator.Contract.Propose(&_XPSValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_XPSValidator *XPSValidatorTransactor) Resign(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
        return _XPSValidator.contract.Transact(opts, "resign", _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_XPSValidator *XPSValidatorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
        return _XPSValidator.Contract.Resign(&_XPSValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_XPSValidator *XPSValidatorTransactorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
        return _XPSValidator.Contract.Resign(&_XPSValidator.TransactOpts, _candidate)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_XPSValidator *XPSValidatorTransactor) Unvote(opts *bind.TransactOpts, _candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
        return _XPSValidator.contract.Transact(opts, "unvote", _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_XPSValidator *XPSValidatorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
        return _XPSValidator.Contract.Unvote(&_XPSValidator.TransactOpts, _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_XPSValidator *XPSValidatorTransactorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
        return _XPSValidator.Contract.Unvote(&_XPSValidator.TransactOpts, _candidate, _cap)
}

// UploadKYC is a paid mutator transaction binding the contract method 0xf5c95125.
//
// Solidity: function uploadKYC(kychash string) returns()
func (_XPSValidator *XPSValidatorTransactor) UploadKYC(opts *bind.TransactOpts, kychash string) (*types.Transaction, error) {
        return _XPSValidator.contract.Transact(opts, "uploadKYC", kychash)
}

// UploadKYC is a paid mutator transaction binding the contract method 0xf5c95125.
//
// Solidity: function uploadKYC(kychash string) returns()
func (_XPSValidator *XPSValidatorSession) UploadKYC(kychash string) (*types.Transaction, error) {
        return _XPSValidator.Contract.UploadKYC(&_XPSValidator.TransactOpts, kychash)
}

// UploadKYC is a paid mutator transaction binding the contract method 0xf5c95125.
//
// Solidity: function uploadKYC(kychash string) returns()
func (_XPSValidator *XPSValidatorTransactorSession) UploadKYC(kychash string) (*types.Transaction, error) {
        return _XPSValidator.Contract.UploadKYC(&_XPSValidator.TransactOpts, kychash)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_XPSValidator *XPSValidatorTransactor) Vote(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
        return _XPSValidator.contract.Transact(opts, "vote", _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_XPSValidator *XPSValidatorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
        return _XPSValidator.Contract.Vote(&_XPSValidator.TransactOpts, _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_XPSValidator *XPSValidatorTransactorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
        return _XPSValidator.Contract.Vote(&_XPSValidator.TransactOpts, _candidate)
}

// VoteInvalidKYC is a paid mutator transaction binding the contract method 0xf2ee3c7d.
//
// Solidity: function voteInvalidKYC(_invalidCandidate address) returns()
func (_XPSValidator *XPSValidatorTransactor) VoteInvalidKYC(opts *bind.TransactOpts, _invalidCandidate common.Address) (*types.Transaction, error) {
        return _XPSValidator.contract.Transact(opts, "voteInvalidKYC", _invalidCandidate)
}

// VoteInvalidKYC is a paid mutator transaction binding the contract method 0xf2ee3c7d.
//
// Solidity: function voteInvalidKYC(_invalidCandidate address) returns()
func (_XPSValidator *XPSValidatorSession) VoteInvalidKYC(_invalidCandidate common.Address) (*types.Transaction, error) {
        return _XPSValidator.Contract.VoteInvalidKYC(&_XPSValidator.TransactOpts, _invalidCandidate)
}

// VoteInvalidKYC is a paid mutator transaction binding the contract method 0xf2ee3c7d.
//
// Solidity: function voteInvalidKYC(_invalidCandidate address) returns()
func (_XPSValidator *XPSValidatorTransactorSession) VoteInvalidKYC(_invalidCandidate common.Address) (*types.Transaction, error) {
        return _XPSValidator.Contract.VoteInvalidKYC(&_XPSValidator.TransactOpts, _invalidCandidate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_XPSValidator *XPSValidatorTransactor) Withdraw(opts *bind.TransactOpts, _blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
        return _XPSValidator.contract.Transact(opts, "withdraw", _blockNumber, _index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_XPSValidator *XPSValidatorSession) Withdraw(_blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
        return _XPSValidator.Contract.Withdraw(&_XPSValidator.TransactOpts, _blockNumber, _index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_XPSValidator *XPSValidatorTransactorSession) Withdraw(_blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
        return _XPSValidator.Contract.Withdraw(&_XPSValidator.TransactOpts, _blockNumber, _index)
}

// XPSValidatorInvalidatedNodeIterator is returned from FilterInvalidatedNode and is used to iterate over the raw logs and unpacked data for InvalidatedNode events raised by the XPSValidator contract.
type XPSValidatorInvalidatedNodeIterator struct {
        Event *XPSValidatorInvalidatedNode // Event containing the contract specifics and raw log

        contract *bind.BoundContract // Generic contract to use for unpacking event data
        event    string              // Event name to use for unpacking event data

        logs chan types.Log        // Log channel receiving the found contract events
        sub  ethereum.Subscription // Subscription for errors, completion and termination
        done bool                  // Whether the subscription completed delivering logs
        fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XPSValidatorInvalidatedNodeIterator) Next() bool {
        // If the iterator failed, stop iterating
        if it.fail != nil {
                return false
        }
        // If the iterator completed, deliver directly whatever's available
        if it.done {
                select {
                case log := <-it.logs:
                        it.Event = new(XPSValidatorInvalidatedNode)
                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                                it.fail = err
                                return false
                        }
                        it.Event.Raw = log
                        return true

                default:
                        return false
                }
        }
        // Iterator still in progress, wait for either a data or an error event
        select {
        case log := <-it.logs:
                it.Event = new(XPSValidatorInvalidatedNode)
                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                        it.fail = err
                        return false
                }
                it.Event.Raw = log
                return true

        case err := <-it.sub.Err():
                it.done = true
                it.fail = err
                return it.Next()
        }
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XPSValidatorInvalidatedNodeIterator) Error() error {
        return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorInvalidatedNodeIterator) Close() error {
        it.sub.Unsubscribe()
        return nil
}

// XPSValidatorInvalidatedNode represents a InvalidatedNode event raised by the XPSValidator contract.
type XPSValidatorInvalidatedNode struct {
        MasternodeOwner common.Address
        Masternodes     []common.Address
        Raw             types.Log // Blockchain specific contextual infos
}

// FilterInvalidatedNode is a free log retrieval operation binding the contract event 0xe18d61a5bf4aa2ab40afc88aa9039d27ae17ff4ec1c65f5f414df6f02ce8b35e.
//
// Solidity: event InvalidatedNode(_masternodeOwner address, _masternodes address[])
func (_XPSValidator *XPSValidatorFilterer) FilterInvalidatedNode(opts *bind.FilterOpts) (*XPSValidatorInvalidatedNodeIterator, error) {

        logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "InvalidatedNode")
        if err != nil {
                return nil, err
        }
        return &XPSValidatorInvalidatedNodeIterator{contract: _XPSValidator.contract, event: "InvalidatedNode", logs: logs, sub: sub}, nil
}

// WatchInvalidatedNode is a free log subscription operation binding the contract event 0xe18d61a5bf4aa2ab40afc88aa9039d27ae17ff4ec1c65f5f414df6f02ce8b35e.
//
// Solidity: event InvalidatedNode(_masternodeOwner address, _masternodes address[])
func (_XPSValidator *XPSValidatorFilterer) WatchInvalidatedNode(opts *bind.WatchOpts, sink chan<- *XPSValidatorInvalidatedNode) (event.Subscription, error) {

        logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "InvalidatedNode")
        if err != nil {
                return nil, err
        }
        return event.NewSubscription(func(quit <-chan struct{}) error {
                defer sub.Unsubscribe()
                for {
                        select {
                        case log := <-logs:
                                // New log arrived, parse the event and forward to the user
                                event := new(XPSValidatorInvalidatedNode)
                                if err := _XPSValidator.contract.UnpackLog(event, "InvalidatedNode", log); err != nil {
                                        return err
                                }
                                event.Raw = log

                                select {
                                case sink <- event:
                                case err := <-sub.Err():
                                        return err
                                case <-quit:
                                        return nil
                                }
                        case err := <-sub.Err():
                                return err
                        case <-quit:
                                return nil
                        }
                }
        }), nil
}

// XPSValidatorProposeIterator is returned from FilterPropose and is used to iterate over the raw logs and unpacked data for Propose events raised by the XPSValidator contract.
type XPSValidatorProposeIterator struct {
        Event *XPSValidatorPropose // Event containing the contract specifics and raw log

        contract *bind.BoundContract // Generic contract to use for unpacking event data
        event    string              // Event name to use for unpacking event data

        logs chan types.Log        // Log channel receiving the found contract events
        sub  ethereum.Subscription // Subscription for errors, completion and termination
        done bool                  // Whether the subscription completed delivering logs
        fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XPSValidatorProposeIterator) Next() bool {
        // If the iterator failed, stop iterating
        if it.fail != nil {
                return false
        }
        // If the iterator completed, deliver directly whatever's available
        if it.done {
                select {
                case log := <-it.logs:
                        it.Event = new(XPSValidatorPropose)
                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                                it.fail = err
                                return false
                        }
                        it.Event.Raw = log
                        return true

                default:
                        return false
                }
        }
        // Iterator still in progress, wait for either a data or an error event
        select {
        case log := <-it.logs:
                it.Event = new(XPSValidatorPropose)
                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                        it.fail = err
                        return false
                }
                it.Event.Raw = log
                return true

        case err := <-it.sub.Err():
                it.done = true
                it.fail = err
                return it.Next()
        }
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XPSValidatorProposeIterator) Error() error {
        return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorProposeIterator) Close() error {
        it.sub.Unsubscribe()
        return nil
}

// XPSValidatorPropose represents a Propose event raised by the XPSValidator contract.
type XPSValidatorPropose struct {
        Owner     common.Address
        Candidate common.Address
        Cap       *big.Int
        Raw       types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(_owner address, _candidate address, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) FilterPropose(opts *bind.FilterOpts) (*XPSValidatorProposeIterator, error) {

        logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "Propose")
        if err != nil {
                return nil, err
        }
        return &XPSValidatorProposeIterator{contract: _XPSValidator.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(_owner address, _candidate address, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *XPSValidatorPropose) (event.Subscription, error) {

        logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "Propose")
        if err != nil {
                return nil, err
        }
        return event.NewSubscription(func(quit <-chan struct{}) error {
                defer sub.Unsubscribe()
                for {
                        select {
                        case log := <-logs:
                                // New log arrived, parse the event and forward to the user
                                event := new(XPSValidatorPropose)
                                if err := _XPSValidator.contract.UnpackLog(event, "Propose", log); err != nil {
                                        return err
                                }
                                event.Raw = log

                                select {
                                case sink <- event:
                                case err := <-sub.Err():
                                        return err
                                case <-quit:
                                        return nil
                                }
                        case err := <-sub.Err():
                                return err
                        case <-quit:
                                return nil
                        }
                }
        }), nil
}

// XPSValidatorResignIterator is returned from FilterResign and is used to iterate over the raw logs and unpacked data for Resign events raised by the XPSValidator contract.
type XPSValidatorResignIterator struct {
        Event *XPSValidatorResign // Event containing the contract specifics and raw log

        contract *bind.BoundContract // Generic contract to use for unpacking event data
        event    string              // Event name to use for unpacking event data

        logs chan types.Log        // Log channel receiving the found contract events
        sub  ethereum.Subscription // Subscription for errors, completion and termination
        done bool                  // Whether the subscription completed delivering logs
        fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XPSValidatorResignIterator) Next() bool {
        // If the iterator failed, stop iterating
        if it.fail != nil {
                return false
        }
        // If the iterator completed, deliver directly whatever's available
        if it.done {
                select {
                case log := <-it.logs:
                        it.Event = new(XPSValidatorResign)
                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                                it.fail = err
                                return false
                        }
                        it.Event.Raw = log
                        return true

                default:
                        return false
                }
        }
        // Iterator still in progress, wait for either a data or an error event
        select {
        case log := <-it.logs:
                it.Event = new(XPSValidatorResign)
                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                        it.fail = err
                        return false
                }
                it.Event.Raw = log
                return true

        case err := <-it.sub.Err():
                it.done = true
                it.fail = err
                return it.Next()
        }
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XPSValidatorResignIterator) Error() error {
        return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorResignIterator) Close() error {
        it.sub.Unsubscribe()
        return nil
}

// XPSValidatorResign represents a Resign event raised by the XPSValidator contract.
type XPSValidatorResign struct {
        Owner     common.Address
        Candidate common.Address
        Raw       types.Log // Blockchain specific contextual infos
}

// FilterResign is a free log retrieval operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(_owner address, _candidate address)
func (_XPSValidator *XPSValidatorFilterer) FilterResign(opts *bind.FilterOpts) (*XPSValidatorResignIterator, error) {

        logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "Resign")
        if err != nil {
                return nil, err
        }
        return &XPSValidatorResignIterator{contract: _XPSValidator.contract, event: "Resign", logs: logs, sub: sub}, nil
}

// WatchResign is a free log subscription operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(_owner address, _candidate address)
func (_XPSValidator *XPSValidatorFilterer) WatchResign(opts *bind.WatchOpts, sink chan<- *XPSValidatorResign) (event.Subscription, error) {

        logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "Resign")
        if err != nil {
                return nil, err
        }
        return event.NewSubscription(func(quit <-chan struct{}) error {
                defer sub.Unsubscribe()
                for {
                        select {
                        case log := <-logs:
                                // New log arrived, parse the event and forward to the user
                                event := new(XPSValidatorResign)
                                if err := _XPSValidator.contract.UnpackLog(event, "Resign", log); err != nil {
                                        return err
                                }
                                event.Raw = log

                                select {
                                case sink <- event:
                                case err := <-sub.Err():
                                        return err
                                case <-quit:
                                        return nil
                                }
                        case err := <-sub.Err():
                                return err
                        case <-quit:
                                return nil
                        }
                }
        }), nil
}

// XPSValidatorUnvoteIterator is returned from FilterUnvote and is used to iterate over the raw logs and unpacked data for Unvote events raised by the XPSValidator contract.
type XPSValidatorUnvoteIterator struct {
        Event *XPSValidatorUnvote // Event containing the contract specifics and raw log

        contract *bind.BoundContract // Generic contract to use for unpacking event data
        event    string              // Event name to use for unpacking event data

        logs chan types.Log        // Log channel receiving the found contract events
        sub  ethereum.Subscription // Subscription for errors, completion and termination
        done bool                  // Whether the subscription completed delivering logs
        fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XPSValidatorUnvoteIterator) Next() bool {
        // If the iterator failed, stop iterating
        if it.fail != nil {
                return false
        }
        // If the iterator completed, deliver directly whatever's available
        if it.done {
                select {
                case log := <-it.logs:
                        it.Event = new(XPSValidatorUnvote)
                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                                it.fail = err
                                return false
                        }
                        it.Event.Raw = log
                        return true

                default:
                        return false
                }
        }
        // Iterator still in progress, wait for either a data or an error event
        select {
        case log := <-it.logs:
                it.Event = new(XPSValidatorUnvote)
                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                        it.fail = err
                        return false
                }
                it.Event.Raw = log
                return true

        case err := <-it.sub.Err():
                it.done = true
                it.fail = err
                return it.Next()
        }
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XPSValidatorUnvoteIterator) Error() error {
        return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorUnvoteIterator) Close() error {
        it.sub.Unsubscribe()
        return nil
}

// XPSValidatorUnvote represents a Unvote event raised by the XPSValidator contract.
type XPSValidatorUnvote struct {
        Voter     common.Address
        Candidate common.Address
        Cap       *big.Int
        Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnvote is a free log retrieval operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(_voter address, _candidate address, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) FilterUnvote(opts *bind.FilterOpts) (*XPSValidatorUnvoteIterator, error) {

        logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "Unvote")
        if err != nil {
                return nil, err
        }
        return &XPSValidatorUnvoteIterator{contract: _XPSValidator.contract, event: "Unvote", logs: logs, sub: sub}, nil
}

// WatchUnvote is a free log subscription operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(_voter address, _candidate address, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) WatchUnvote(opts *bind.WatchOpts, sink chan<- *XPSValidatorUnvote) (event.Subscription, error) {

        logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "Unvote")
        if err != nil {
                return nil, err
        }
        return event.NewSubscription(func(quit <-chan struct{}) error {
                defer sub.Unsubscribe()
                for {
                        select {
                        case log := <-logs:
                                // New log arrived, parse the event and forward to the user
                                event := new(XPSValidatorUnvote)
                                if err := _XPSValidator.contract.UnpackLog(event, "Unvote", log); err != nil {
                                        return err
                                }
                                event.Raw = log

                                select {
                                case sink <- event:
                                case err := <-sub.Err():
                                        return err
                                case <-quit:
                                        return nil
                                }
                        case err := <-sub.Err():
                                return err
                        case <-quit:
                                return nil
                        }
                }
        }), nil
}

// XPSValidatorUploadedKYCIterator is returned from FilterUploadedKYC and is used to iterate over the raw logs and unpacked data for UploadedKYC events raised by the XPSValidator contract.
type XPSValidatorUploadedKYCIterator struct {
        Event *XPSValidatorUploadedKYC // Event containing the contract specifics and raw log

        contract *bind.BoundContract // Generic contract to use for unpacking event data
        event    string              // Event name to use for unpacking event data

        logs chan types.Log        // Log channel receiving the found contract events
        sub  ethereum.Subscription // Subscription for errors, completion and termination
        done bool                  // Whether the subscription completed delivering logs
        fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XPSValidatorUploadedKYCIterator) Next() bool {
        // If the iterator failed, stop iterating
        if it.fail != nil {
                return false
        }
        // If the iterator completed, deliver directly whatever's available
        if it.done {
                select {
                case log := <-it.logs:
                        it.Event = new(XPSValidatorUploadedKYC)
                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                                it.fail = err
                                return false
                        }
                        it.Event.Raw = log
                        return true

                default:
                        return false
                }
        }
        // Iterator still in progress, wait for either a data or an error event
        select {
        case log := <-it.logs:
                it.Event = new(XPSValidatorUploadedKYC)
                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                        it.fail = err
                        return false
                }
                it.Event.Raw = log
                return true

        case err := <-it.sub.Err():
                it.done = true
                it.fail = err
                return it.Next()
        }
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XPSValidatorUploadedKYCIterator) Error() error {
        return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorUploadedKYCIterator) Close() error {
        it.sub.Unsubscribe()
        return nil
}

// XPSValidatorUploadedKYC represents a UploadedKYC event raised by the XPSValidator contract.
type XPSValidatorUploadedKYC struct {
        Owner   common.Address
        KycHash string
        Raw     types.Log // Blockchain specific contextual infos
}

// FilterUploadedKYC is a free log retrieval operation binding the contract event 0x949360d814b28a3b393a68909efe1fee120ee09cac30f360a0f80ab5415a611a.
//
// Solidity: event UploadedKYC(_owner address, kycHash string)
func (_XPSValidator *XPSValidatorFilterer) FilterUploadedKYC(opts *bind.FilterOpts) (*XPSValidatorUploadedKYCIterator, error) {

        logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "UploadedKYC")
        if err != nil {
                return nil, err
        }
        return &XPSValidatorUploadedKYCIterator{contract: _XPSValidator.contract, event: "UploadedKYC", logs: logs, sub: sub}, nil
}

// WatchUploadedKYC is a free log subscription operation binding the contract event 0x949360d814b28a3b393a68909efe1fee120ee09cac30f360a0f80ab5415a611a.
//
// Solidity: event UploadedKYC(_owner address, kycHash string)
func (_XPSValidator *XPSValidatorFilterer) WatchUploadedKYC(opts *bind.WatchOpts, sink chan<- *XPSValidatorUploadedKYC) (event.Subscription, error) {

        logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "UploadedKYC")
        if err != nil {
                return nil, err
        }
        return event.NewSubscription(func(quit <-chan struct{}) error {
                defer sub.Unsubscribe()
                for {
                        select {
                        case log := <-logs:
                                // New log arrived, parse the event and forward to the user
                                event := new(XPSValidatorUploadedKYC)
                                if err := _XPSValidator.contract.UnpackLog(event, "UploadedKYC", log); err != nil {
                                        return err
                                }
                                event.Raw = log

                                select {
                                case sink <- event:
                                case err := <-sub.Err():
                                        return err
                                case <-quit:
                                        return nil
                                }
                        case err := <-sub.Err():
                                return err
                        case <-quit:
                                return nil
                        }
                }
        }), nil
}

// XPSValidatorVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the XPSValidator contract.
type XPSValidatorVoteIterator struct {
        Event *XPSValidatorVote // Event containing the contract specifics and raw log

        contract *bind.BoundContract // Generic contract to use for unpacking event data
        event    string              // Event name to use for unpacking event data

        logs chan types.Log        // Log channel receiving the found contract events
        sub  ethereum.Subscription // Subscription for errors, completion and termination
        done bool                  // Whether the subscription completed delivering logs
        fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XPSValidatorVoteIterator) Next() bool {
        // If the iterator failed, stop iterating
        if it.fail != nil {
                return false
        }
        // If the iterator completed, deliver directly whatever's available
        if it.done {
                select {
                case log := <-it.logs:
                        it.Event = new(XPSValidatorVote)
                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                                it.fail = err
                                return false
                        }
                        it.Event.Raw = log
                        return true

                default:
                        return false
                }
        }
        // Iterator still in progress, wait for either a data or an error event
        select {
        case log := <-it.logs:
                it.Event = new(XPSValidatorVote)
                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                        it.fail = err
                        return false
                }
                it.Event.Raw = log
                return true

        case err := <-it.sub.Err():
                it.done = true
                it.fail = err
                return it.Next()
        }
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XPSValidatorVoteIterator) Error() error {
        return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorVoteIterator) Close() error {
        it.sub.Unsubscribe()
        return nil
}

// XPSValidatorVote represents a Vote event raised by the XPSValidator contract.
type XPSValidatorVote struct {
        Voter     common.Address
        Candidate common.Address
        Cap       *big.Int
        Raw       types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(_voter address, _candidate address, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) FilterVote(opts *bind.FilterOpts) (*XPSValidatorVoteIterator, error) {

        logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "Vote")
        if err != nil {
                return nil, err
        }
        return &XPSValidatorVoteIterator{contract: _XPSValidator.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(_voter address, _candidate address, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *XPSValidatorVote) (event.Subscription, error) {

        logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "Vote")
        if err != nil {
                return nil, err
        }
        return event.NewSubscription(func(quit <-chan struct{}) error {
                defer sub.Unsubscribe()
                for {
                        select {
                        case log := <-logs:
                                // New log arrived, parse the event and forward to the user
                                event := new(XPSValidatorVote)
                                if err := _XPSValidator.contract.UnpackLog(event, "Vote", log); err != nil {
                                        return err
                                }
                                event.Raw = log

                                select {
                                case sink <- event:
                                case err := <-sub.Err():
                                        return err
                                case <-quit:
                                        return nil
                                }
                        case err := <-sub.Err():
                                return err
                        case <-quit:
                                return nil
                        }
                }
        }), nil
}

// XPSValidatorWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the XPSValidator contract.
type XPSValidatorWithdrawIterator struct {
        Event *XPSValidatorWithdraw // Event containing the contract specifics and raw log

        contract *bind.BoundContract // Generic contract to use for unpacking event data
        event    string              // Event name to use for unpacking event data

        logs chan types.Log        // Log channel receiving the found contract events
        sub  ethereum.Subscription // Subscription for errors, completion and termination
        done bool                  // Whether the subscription completed delivering logs
        fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XPSValidatorWithdrawIterator) Next() bool {
        // If the iterator failed, stop iterating
        if it.fail != nil {
                return false
        }
        // If the iterator completed, deliver directly whatever's available
        if it.done {
                select {
                case log := <-it.logs:
                        it.Event = new(XPSValidatorWithdraw)
                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                                it.fail = err
                                return false
                        }
                        it.Event.Raw = log
                        return true

                default:
                        return false
                }
        }
        // Iterator still in progress, wait for either a data or an error event
        select {
        case log := <-it.logs:
                it.Event = new(XPSValidatorWithdraw)
                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
                        it.fail = err
                        return false
                }
                it.Event.Raw = log
                return true

        case err := <-it.sub.Err():
                it.done = true
                it.fail = err
                return it.Next()
        }
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XPSValidatorWithdrawIterator) Error() error {
        return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorWithdrawIterator) Close() error {
        it.sub.Unsubscribe()
        return nil
}

// XPSValidatorWithdraw represents a Withdraw event raised by the XPSValidator contract.
type XPSValidatorWithdraw struct {
        Owner       common.Address
        BlockNumber *big.Int
        Cap         *big.Int
        Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(_owner address, _blockNumber uint256, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) FilterWithdraw(opts *bind.FilterOpts) (*XPSValidatorWithdrawIterator, error) {

        logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "Withdraw")
        if err != nil {
                return nil, err
        }
        return &XPSValidatorWithdrawIterator{contract: _XPSValidator.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(_owner address, _blockNumber uint256, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *XPSValidatorWithdraw) (event.Subscription, error) {

        logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "Withdraw")
        if err != nil {
                return nil, err
        }
        return event.NewSubscription(func(quit <-chan struct{}) error {
                defer sub.Unsubscribe()
                for {
                        select {
                        case log := <-logs:
                                // New log arrived, parse the event and forward to the user
                                event := new(XPSValidatorWithdraw)
                                if err := _XPSValidator.contract.UnpackLog(event, "Withdraw", log); err != nil {
                                        return err
                                }
                                event.Raw = log

                                select {
                                case sink <- event:
                                case err := <-sub.Err():
                                        return err
                                case <-quit:
                                        return nil
                                }
                        case err := <-sub.Err():
                                return err
                        case <-quit:
                                return nil
                        }
                }
        }), nil
}
