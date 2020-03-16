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
const EthdepositcontractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"proposePricing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nodes\",\"type\":\"address[]\"}],\"name\":\"depositWithNodes\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"height_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height_end\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"createInvoice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"registerNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthdepositcontractBin is the compiled bytecode used for deploying new contracts.
var EthdepositcontractBin = "0x608060405234801561001057600080fd5b506001600080018190555060016000600101819055506001600060020181905550600160006003018190555060008060040160006101000a81548160ff0219169083151502179055506105f8806100686000396000f3fe60806040526004361061004a5760003560e01c80633296972f1461010057806351cff8d914610159578063672d7a0d146101aa57806376c58850146101fb57806396ea8630146102b3575b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050348160060160008282540192505081905550600081600001600082015481600001556001820154816001015560028201548160020155600382015481600301556004820160009054906101000a900460ff168160040160006101000a81548160ff02191690831515021790555090505050005b34801561010c57600080fd5b506101576004803603608081101561012357600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610340565b005b34801561016557600080fd5b506101a86004803603602081101561017c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103d8565b005b3480156101b657600080fd5b506101f9600480360360208110156101cd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610485565b005b6102b16004803603602081101561021157600080fd5b810190808035906020019064010000000081111561022e57600080fd5b82018360208201111561024057600080fd5b8035906020019184602083028401116401000000008311171561026257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506104e3565b005b3480156102bf57600080fd5b5061033e600480360360e08110156102d657600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291905050506105ba565b005b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508481600001600001819055508381600001600101819055508281600001600201819055508181600001600301819055506001600060040160006101000a81548160ff0219169083151502179055505050505050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600101541161042c57600080fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc82600101549081150290604051600060405180830381858888f19350505050158015610476573d6000803e3d6000fd5b50600081600101819055505050565b6001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548160ff02191690831515021790555050565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905034816006016000828254019250508190555060001515600060040160009054906101000a900460ff16151514156105b657600081600001600082015481600001556001820154816001015560028201548160020155600382015481600301556004820160009054906101000a900460ff168160040160006101000a81548160ff0219169083151502179055509050505b5050565b5050505050505056fea265627a7a7231582013a2774eeae4215a6d8898621b795e8f9ee68ce1f8008ff41ca38d7db43d63e064736f6c63430005100032"

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

// DepositWithNodes is a paid mutator transaction binding the contract method 0x76c58850.
//
// Solidity: function depositWithNodes(address[] _nodes) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) DepositWithNodes(opts *bind.TransactOpts, _nodes []common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "depositWithNodes", _nodes)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0x76c58850.
//
// Solidity: function depositWithNodes(address[] _nodes) returns()
func (_Ethdepositcontract *EthdepositcontractSession) DepositWithNodes(_nodes []common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.DepositWithNodes(&_Ethdepositcontract.TransactOpts, _nodes)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0x76c58850.
//
// Solidity: function depositWithNodes(address[] _nodes) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) DepositWithNodes(_nodes []common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.DepositWithNodes(&_Ethdepositcontract.TransactOpts, _nodes)
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
