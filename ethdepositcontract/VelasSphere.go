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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthdepositcontractABI is the input ABI used to generate the binding from.
const EthdepositcontractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"banNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"openInvoice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"createInvoice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"staking_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mining_addr\",\"type\":\"address\"}],\"name\":\"registerNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"old_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"new_addr\",\"type\":\"address\"}],\"name\":\"changeMiningAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"directory\",\"type\":\"bytes32\"}],\"name\":\"registerDirectory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"directory\",\"type\":\"bytes32\"}],\"name\":\"increaseDirectoryNonce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"directory\",\"type\":\"bytes32\"}],\"name\":\"getDirectoryNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EthdepositcontractBin is the compiled bytecode used for deploying new contracts.
var EthdepositcontractBin = "0x608060405264174876e80060005560326001556005805534801561002257600080fd5b50603e600381905550610c628061003a6000396000f3fe6080604052600436106100915760003560e01c80638a99c70a116100595780638a99c70a14610228578063906f014e1461028357806398a8004b146102be578063b0958532146102f9578063e708013f1461035457610091565b806306a3316a1461009b57806339bf397e146100ea5780633faba59e1461011557806351cff8d91461016657806388e66a54146101b7575b6100996103b8565b005b3480156100a757600080fd5b506100d4600480360360208110156100be57600080fd5b810190808035906020019092919050505061046a565b6040518082815260200191505060405180910390f35b3480156100f657600080fd5b506100ff6104c4565b6040518082815260200191505060405180910390f35b34801561012157600080fd5b506101646004803603602081101561013857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104ce565b005b34801561017257600080fd5b506101b56004803603602081101561018957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105f4565b005b3480156101c357600080fd5b50610226600480360360408110156101da57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610703565b005b34801561023457600080fd5b506102816004803603604081101561024b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061078e565b005b34801561028f57600080fd5b506102bc600480360360208110156102a657600080fd5b8101908080359060200190929190505050610834565b005b3480156102ca57600080fd5b506102f7600480360360208110156102e157600080fd5b810190808035906020019092919050505061088d565b005b34801561030557600080fd5b506103526004803603604081101561031c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610943565b005b6103b66004803603604081101561036a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b44565b005b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000341161040857600080fd5b348160000160008282540192505081905550600015158160010160009054906101000a900460ff161515141561044a5760016004600082825401925050819055505b60018160010160006101000a81548160ff02191690831515021790555050565b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020549050919050565b6000600254905090565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060018160020160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506001600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008282540192505081905550600554600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002050505050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600201541161064857600080fd5b8060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201549081150290604051600060405180830381858888f193505050501580156106f4573d6000803e3d6000fd5b50600081600201819055505050565b6000600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050818160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146107c657600080fd5b6000600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905081816000018190555060018160010160006101000a81548160ff021916908315150217905550505050565b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000208190555050565b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020549050808060010191505080600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000848152602001908152602001600020819055505050565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060030160009054906101000a900460ff166109a157600080fd5b600554600760008360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015410610a1557600080fd5b6000600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160020160008460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015414610acd57600080fd5b6000600960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060010160009054906101000a900460ff16610b2b57600080fd5b8381600201600082825401925050819055505050505050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600015158160030160009054906101000a900460ff16151514610ba957600080fd5b60018160030160006101000a81548160ff0219169083151502179055506000543414610bd457600080fd5b828160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160026000828254019250508190555050505056fea265627a7a72315820694a7536e4c51b72badf9919658677a077fb2c5070c69a6c73abdfec5c81624964736f6c63430005100032"

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

// GetDirectoryNonce is a free data retrieval call binding the contract method 0x06a3316a.
//
// Solidity: function getDirectoryNonce(bytes32 directory) view returns(uint256)
func (_Ethdepositcontract *EthdepositcontractCaller) GetDirectoryNonce(opts *bind.CallOpts, directory [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethdepositcontract.contract.Call(opts, out, "getDirectoryNonce", directory)
	return *ret0, err
}

// GetDirectoryNonce is a free data retrieval call binding the contract method 0x06a3316a.
//
// Solidity: function getDirectoryNonce(bytes32 directory) view returns(uint256)
func (_Ethdepositcontract *EthdepositcontractSession) GetDirectoryNonce(directory [32]byte) (*big.Int, error) {
	return _Ethdepositcontract.Contract.GetDirectoryNonce(&_Ethdepositcontract.CallOpts, directory)
}

// GetDirectoryNonce is a free data retrieval call binding the contract method 0x06a3316a.
//
// Solidity: function getDirectoryNonce(bytes32 directory) view returns(uint256)
func (_Ethdepositcontract *EthdepositcontractCallerSession) GetDirectoryNonce(directory [32]byte) (*big.Int, error) {
	return _Ethdepositcontract.Contract.GetDirectoryNonce(&_Ethdepositcontract.CallOpts, directory)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_Ethdepositcontract *EthdepositcontractCaller) GetNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethdepositcontract.contract.Call(opts, out, "getNodeCount")
	return *ret0, err
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_Ethdepositcontract *EthdepositcontractSession) GetNodeCount() (*big.Int, error) {
	return _Ethdepositcontract.Contract.GetNodeCount(&_Ethdepositcontract.CallOpts)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_Ethdepositcontract *EthdepositcontractCallerSession) GetNodeCount() (*big.Int, error) {
	return _Ethdepositcontract.Contract.GetNodeCount(&_Ethdepositcontract.CallOpts)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address _node) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) BanNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "banNode", _node)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address _node) returns()
func (_Ethdepositcontract *EthdepositcontractSession) BanNode(_node common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.BanNode(&_Ethdepositcontract.TransactOpts, _node)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address _node) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) BanNode(_node common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.BanNode(&_Ethdepositcontract.TransactOpts, _node)
}

// ChangeMiningAddr is a paid mutator transaction binding the contract method 0x88e66a54.
//
// Solidity: function changeMiningAddr(address old_addr, address new_addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) ChangeMiningAddr(opts *bind.TransactOpts, old_addr common.Address, new_addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "changeMiningAddr", old_addr, new_addr)
}

// ChangeMiningAddr is a paid mutator transaction binding the contract method 0x88e66a54.
//
// Solidity: function changeMiningAddr(address old_addr, address new_addr) returns()
func (_Ethdepositcontract *EthdepositcontractSession) ChangeMiningAddr(old_addr common.Address, new_addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ChangeMiningAddr(&_Ethdepositcontract.TransactOpts, old_addr, new_addr)
}

// ChangeMiningAddr is a paid mutator transaction binding the contract method 0x88e66a54.
//
// Solidity: function changeMiningAddr(address old_addr, address new_addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) ChangeMiningAddr(old_addr common.Address, new_addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ChangeMiningAddr(&_Ethdepositcontract.TransactOpts, old_addr, new_addr)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0xb0958532.
//
// Solidity: function createInvoice(address user, uint256 price) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) CreateInvoice(opts *bind.TransactOpts, user common.Address, price *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "createInvoice", user, price)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0xb0958532.
//
// Solidity: function createInvoice(address user, uint256 price) returns()
func (_Ethdepositcontract *EthdepositcontractSession) CreateInvoice(user common.Address, price *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.CreateInvoice(&_Ethdepositcontract.TransactOpts, user, price)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0xb0958532.
//
// Solidity: function createInvoice(address user, uint256 price) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) CreateInvoice(user common.Address, price *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.CreateInvoice(&_Ethdepositcontract.TransactOpts, user, price)
}

// IncreaseDirectoryNonce is a paid mutator transaction binding the contract method 0x98a8004b.
//
// Solidity: function increaseDirectoryNonce(bytes32 directory) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) IncreaseDirectoryNonce(opts *bind.TransactOpts, directory [32]byte) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "increaseDirectoryNonce", directory)
}

// IncreaseDirectoryNonce is a paid mutator transaction binding the contract method 0x98a8004b.
//
// Solidity: function increaseDirectoryNonce(bytes32 directory) returns()
func (_Ethdepositcontract *EthdepositcontractSession) IncreaseDirectoryNonce(directory [32]byte) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.IncreaseDirectoryNonce(&_Ethdepositcontract.TransactOpts, directory)
}

// IncreaseDirectoryNonce is a paid mutator transaction binding the contract method 0x98a8004b.
//
// Solidity: function increaseDirectoryNonce(bytes32 directory) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) IncreaseDirectoryNonce(directory [32]byte) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.IncreaseDirectoryNonce(&_Ethdepositcontract.TransactOpts, directory)
}

// OpenInvoice is a paid mutator transaction binding the contract method 0x8a99c70a.
//
// Solidity: function openInvoice(address addr, uint256 deadline) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) OpenInvoice(opts *bind.TransactOpts, addr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "openInvoice", addr, deadline)
}

// OpenInvoice is a paid mutator transaction binding the contract method 0x8a99c70a.
//
// Solidity: function openInvoice(address addr, uint256 deadline) returns()
func (_Ethdepositcontract *EthdepositcontractSession) OpenInvoice(addr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.OpenInvoice(&_Ethdepositcontract.TransactOpts, addr, deadline)
}

// OpenInvoice is a paid mutator transaction binding the contract method 0x8a99c70a.
//
// Solidity: function openInvoice(address addr, uint256 deadline) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) OpenInvoice(addr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.OpenInvoice(&_Ethdepositcontract.TransactOpts, addr, deadline)
}

// RegisterDirectory is a paid mutator transaction binding the contract method 0x906f014e.
//
// Solidity: function registerDirectory(bytes32 directory) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) RegisterDirectory(opts *bind.TransactOpts, directory [32]byte) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "registerDirectory", directory)
}

// RegisterDirectory is a paid mutator transaction binding the contract method 0x906f014e.
//
// Solidity: function registerDirectory(bytes32 directory) returns()
func (_Ethdepositcontract *EthdepositcontractSession) RegisterDirectory(directory [32]byte) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.RegisterDirectory(&_Ethdepositcontract.TransactOpts, directory)
}

// RegisterDirectory is a paid mutator transaction binding the contract method 0x906f014e.
//
// Solidity: function registerDirectory(bytes32 directory) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) RegisterDirectory(directory [32]byte) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.RegisterDirectory(&_Ethdepositcontract.TransactOpts, directory)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xe708013f.
//
// Solidity: function registerNode(address staking_addr, address mining_addr) payable returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) RegisterNode(opts *bind.TransactOpts, staking_addr common.Address, mining_addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "registerNode", staking_addr, mining_addr)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xe708013f.
//
// Solidity: function registerNode(address staking_addr, address mining_addr) payable returns()
func (_Ethdepositcontract *EthdepositcontractSession) RegisterNode(staking_addr common.Address, mining_addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.RegisterNode(&_Ethdepositcontract.TransactOpts, staking_addr, mining_addr)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xe708013f.
//
// Solidity: function registerNode(address staking_addr, address mining_addr) payable returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) RegisterNode(staking_addr common.Address, mining_addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.RegisterNode(&_Ethdepositcontract.TransactOpts, staking_addr, mining_addr)
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

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ethdepositcontract *EthdepositcontractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.Fallback(&_Ethdepositcontract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.Fallback(&_Ethdepositcontract.TransactOpts, calldata)
}
