// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethdepositcontract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthdepositcontractABI is the input ABI used to generate the binding from.
const EthdepositcontractABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resource\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voices\",\"type\":\"uint256\"}],\"name\":\"createInvoice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthdepositcontractBin is the compiled bytecode used for deploying new contracts.
var EthdepositcontractBin = "0x608060405234801561001057600080fd5b50610147806100206000396000f3fe6080604052600436106100345760003560e01c80630e7ccabc1461003957806351cff8d914610088578063b6b55f25146100d9575b600080fd5b34801561004557600080fd5b506100866004803603606081101561005c57600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050610107565b005b34801561009457600080fd5b506100d7600480360360208110156100ab57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061010c565b005b610105600480360360208110156100ef57600080fd5b810190808035906020019092919050505061010f565b005b505050565b50565b5056fea265627a7a72315820e132f1da32c5c74b147af6ab96a02d830b4e28c701b25c62035a3dde4ad86c2a64736f6c63430005100032"

// DeployEthdepositcontract deploys a new Ethereum contract, binding an instance of Ethdepositcontract to it.
func DeployEthdepositcontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ethdepositcontract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthdepositcontractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthdepositcontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ethdepositcontract{EthdepositcontractCaller: EthdepositcontractCaller{contract: contract}, EthdepositcontractTransactor: EthdepositcontractTransactor{contract: contract}, EthdepositcontractFilterer: EthdepositcontractFilterer{contract: contract}}, nil
}

// Ethdepositcontract is an auto generated Go binding around an Ethereum contract.
type Ethdepositcontract struct {
	EthdepositcontractCaller     // Read-only binding to the contract
	EthdepositcontractTransactor // Write-only binding to the contract
	EthdepositcontractFilterer   // Log filterer for contract events
}

// EthdepositcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthdepositcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthdepositcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthdepositcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthdepositcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthdepositcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthdepositcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthdepositcontractSession struct {
	Contract     *Ethdepositcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthdepositcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthdepositcontractCallerSession struct {
	Contract *EthdepositcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// EthdepositcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthdepositcontractTransactorSession struct {
	Contract     *EthdepositcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// EthdepositcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthdepositcontractRaw struct {
	Contract *Ethdepositcontract // Generic contract binding to access the raw methods on
}

// EthdepositcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthdepositcontractCallerRaw struct {
	Contract *EthdepositcontractCaller // Generic read-only contract binding to access the raw methods on
}

// EthdepositcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthdepositcontractTransactorRaw struct {
	Contract *EthdepositcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthdepositcontract creates a new instance of Ethdepositcontract, bound to a specific deployed contract.
func NewEthdepositcontract(address common.Address, backend bind.ContractBackend) (*Ethdepositcontract, error) {
	contract, err := bindEthdepositcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethdepositcontract{EthdepositcontractCaller: EthdepositcontractCaller{contract: contract}, EthdepositcontractTransactor: EthdepositcontractTransactor{contract: contract}, EthdepositcontractFilterer: EthdepositcontractFilterer{contract: contract}}, nil
}

// NewEthdepositcontractCaller creates a new read-only instance of Ethdepositcontract, bound to a specific deployed contract.
func NewEthdepositcontractCaller(address common.Address, caller bind.ContractCaller) (*EthdepositcontractCaller, error) {
	contract, err := bindEthdepositcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthdepositcontractCaller{contract: contract}, nil
}

// NewEthdepositcontractTransactor creates a new write-only instance of Ethdepositcontract, bound to a specific deployed contract.
func NewEthdepositcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*EthdepositcontractTransactor, error) {
	contract, err := bindEthdepositcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthdepositcontractTransactor{contract: contract}, nil
}

// NewEthdepositcontractFilterer creates a new log filterer instance of Ethdepositcontract, bound to a specific deployed contract.
func NewEthdepositcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*EthdepositcontractFilterer, error) {
	contract, err := bindEthdepositcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthdepositcontractFilterer{contract: contract}, nil
}

// bindEthdepositcontract binds a generic wrapper to an already deployed contract.
func bindEthdepositcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthdepositcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethdepositcontract *EthdepositcontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethdepositcontract.Contract.EthdepositcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethdepositcontract *EthdepositcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.EthdepositcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethdepositcontract *EthdepositcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.EthdepositcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethdepositcontract *EthdepositcontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethdepositcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethdepositcontract *EthdepositcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethdepositcontract *EthdepositcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.contract.Transact(opts, method, params...)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x0e7ccabc.
//
// Solidity: function createInvoice(bytes32 hash, uint256 sum, uint256 voices) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) CreateInvoice(opts *bind.TransactOpts, hash [32]byte, sum *big.Int, voices *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "createInvoice", hash, sum, voices)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x0e7ccabc.
//
// Solidity: function createInvoice(bytes32 hash, uint256 sum, uint256 voices) returns()
func (_Ethdepositcontract *EthdepositcontractSession) CreateInvoice(hash [32]byte, sum *big.Int, voices *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.CreateInvoice(&_Ethdepositcontract.TransactOpts, hash, sum, voices)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x0e7ccabc.
//
// Solidity: function createInvoice(bytes32 hash, uint256 sum, uint256 voices) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) CreateInvoice(hash [32]byte, sum *big.Int, voices *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.CreateInvoice(&_Ethdepositcontract.TransactOpts, hash, sum, voices)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 resource) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) Deposit(opts *bind.TransactOpts, resource *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "deposit", resource)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 resource) returns()
func (_Ethdepositcontract *EthdepositcontractSession) Deposit(resource *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.Deposit(&_Ethdepositcontract.TransactOpts, resource)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 resource) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) Deposit(resource *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.Deposit(&_Ethdepositcontract.TransactOpts, resource)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address user) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) Withdraw(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "withdraw", user)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address user) returns()
func (_Ethdepositcontract *EthdepositcontractSession) Withdraw(user common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.Withdraw(&_Ethdepositcontract.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address user) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) Withdraw(user common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.Withdraw(&_Ethdepositcontract.TransactOpts, user)
}
