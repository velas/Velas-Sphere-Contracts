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
const EthdepositcontractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"proposePricing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_positions\",\"type\":\"uint256\"}],\"name\":\"depositWithNodes\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"height_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height_end\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"createInvoice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"registerNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthdepositcontractBin is the compiled bytecode used for deploying new contracts.
var EthdepositcontractBin = "0x608060405264174876e80060005534801561001957600080fd5b50600160036000018190555060016003600101819055506001600360020181905550600160038001819055506000600360040160006101000a81548160ff021916908315150217905550610906806100726000396000f3fe60806040526004361061004a5760003560e01c80633296972f1461005457806351cff8d9146100ad578063672d7a0d146100fe57806396ea863014610142578063b2492b17146101cf575b6100526101fd565b005b34801561006057600080fd5b506100ab6004803603608081101561007757600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610421565b005b3480156100b957600080fd5b506100fc600480360360208110156100d057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105ad565b005b6101406004803603602081101561011457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106d0565b005b34801561014e57600080fd5b506101cd600480360360e081101561016557600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291905050506107fa565b005b6101fb600480360360208110156101e557600080fd5b810190808035906020019092919050505061085c565b005b6000341161020a57600080fd5b34600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206006016000828254019250508190555060001515600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060070160009054906101000a900460ff16151514156102c85760016002600082825401925050819055505b6001600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060070160006101000a81548160ff02191690831515021790555060001515600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160040160009054906101000a900460ff161515141561041f576003600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001600082015481600001556001820154816001015560028201548160020155600382015481600301556004820160009054906101000a900460ff168160040160006101000a81548160ff0219169083151502179055509050505b565b83600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000018190555082600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016001018190555081600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016002018190555080600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001600301819055506001600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160040160006101000a81548160ff02191690831515021790555050505050565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154116105fc57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549081150290604051600060405180830381858888f19350505050158015610684573d6000803e3d6000fd5b506000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555050565b60001515600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff1615151461073057600080fd5b6001600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548160ff021916908315150217905550600054341461079957600080fd5b6107a16108ae565b600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301819055506001806000828254019250508190555050565b600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff1661085357600080fd5b50505050505050565b6108646101fd565b80600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206005018190555050565b6000605e60015411156108c457600090506108ce565b6001546001901b90505b9056fea265627a7a72315820766e83dbd9f0ca176f933823d5856c1b9e5240b3559862bfaa7260cbbb1a06ed64736f6c63430005100032"

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

// CreateInvoice is a paid mutator transaction binding the contract method 0x96ea8630.
//
// Solidity: function createInvoice(uint256 height_start, uint256 height_end, address user, uint256 keepPerByte, uint256 writePerByte, uint256 GPUTPerCycle, uint256 CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) CreateInvoice(opts *bind.TransactOpts, height_start *big.Int, height_end *big.Int, user common.Address, keepPerByte *big.Int, writePerByte *big.Int, GPUTPerCycle *big.Int, CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "createInvoice", height_start, height_end, user, keepPerByte, writePerByte, GPUTPerCycle, CPUTtPerCycle)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x96ea8630.
//
// Solidity: function createInvoice(uint256 height_start, uint256 height_end, address user, uint256 keepPerByte, uint256 writePerByte, uint256 GPUTPerCycle, uint256 CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractSession) CreateInvoice(height_start *big.Int, height_end *big.Int, user common.Address, keepPerByte *big.Int, writePerByte *big.Int, GPUTPerCycle *big.Int, CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.CreateInvoice(&_Ethdepositcontract.TransactOpts, height_start, height_end, user, keepPerByte, writePerByte, GPUTPerCycle, CPUTtPerCycle)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x96ea8630.
//
// Solidity: function createInvoice(uint256 height_start, uint256 height_end, address user, uint256 keepPerByte, uint256 writePerByte, uint256 GPUTPerCycle, uint256 CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) CreateInvoice(height_start *big.Int, height_end *big.Int, user common.Address, keepPerByte *big.Int, writePerByte *big.Int, GPUTPerCycle *big.Int, CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.CreateInvoice(&_Ethdepositcontract.TransactOpts, height_start, height_end, user, keepPerByte, writePerByte, GPUTPerCycle, CPUTtPerCycle)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0xb2492b17.
//
// Solidity: function depositWithNodes(uint256 _positions) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) DepositWithNodes(opts *bind.TransactOpts, _positions *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "depositWithNodes", _positions)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0xb2492b17.
//
// Solidity: function depositWithNodes(uint256 _positions) returns()
func (_Ethdepositcontract *EthdepositcontractSession) DepositWithNodes(_positions *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.DepositWithNodes(&_Ethdepositcontract.TransactOpts, _positions)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0xb2492b17.
//
// Solidity: function depositWithNodes(uint256 _positions) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) DepositWithNodes(_positions *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.DepositWithNodes(&_Ethdepositcontract.TransactOpts, _positions)
}

// ProposePricing is a paid mutator transaction binding the contract method 0x3296972f.
//
// Solidity: function proposePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) ProposePricing(opts *bind.TransactOpts, _keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "proposePricing", _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// ProposePricing is a paid mutator transaction binding the contract method 0x3296972f.
//
// Solidity: function proposePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractSession) ProposePricing(_keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ProposePricing(&_Ethdepositcontract.TransactOpts, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// ProposePricing is a paid mutator transaction binding the contract method 0x3296972f.
//
// Solidity: function proposePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) ProposePricing(_keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ProposePricing(&_Ethdepositcontract.TransactOpts, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x672d7a0d.
//
// Solidity: function registerNode(address addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) RegisterNode(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "registerNode", addr)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x672d7a0d.
//
// Solidity: function registerNode(address addr) returns()
func (_Ethdepositcontract *EthdepositcontractSession) RegisterNode(addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.RegisterNode(&_Ethdepositcontract.TransactOpts, addr)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x672d7a0d.
//
// Solidity: function registerNode(address addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) RegisterNode(addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.RegisterNode(&_Ethdepositcontract.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) Withdraw(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "withdraw", addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address addr) returns()
func (_Ethdepositcontract *EthdepositcontractSession) Withdraw(addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.Withdraw(&_Ethdepositcontract.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) Withdraw(addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.Withdraw(&_Ethdepositcontract.TransactOpts, addr)
}
