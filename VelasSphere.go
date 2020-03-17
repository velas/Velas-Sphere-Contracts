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
const EthdepositcontractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"proposePricing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pull\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_places\",\"type\":\"uint256\"}],\"name\":\"depositWithNodes\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"height_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height_end\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"createInvoice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"registerNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthdepositcontractBin is the compiled bytecode used for deploying new contracts.
var EthdepositcontractBin = "0x608060405264174876e800600055603260015534801561001e57600080fd5b5060016006600001819055506001600660010181905550600160066002018190555060016006600301819055506000600660040160006101000a81548160ff021916908315150217905550603e600381905550610c87806100806000396000f3fe60806040526004361061004a5760003560e01c80633296972f1461005457806351cff8d9146100ad578063672d7a0d146100fe5780638e8b45c61461014257806396ea86301461017a575b610052610207565b005b34801561006057600080fd5b506100ab6004803603608081101561007757600080fd5b81019080803590602001909291908035906020019092919080359060200190929190803590602001909291905050506103cb565b005b3480156100b957600080fd5b506100fc600480360360208110156100d057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610465565b005b6101406004803603602081101561011457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610550565b005b6101786004803603604081101561015857600080fd5b810190808035906020019092919080359060200190929190505050610616565b005b34801561018657600080fd5b50610205600480360360e081101561019d57600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919050505061069b565b005b6000341161021457600080fd5b34600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206008016000828254019250508190555060001515600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060090160009054906101000a900460ff161515141561036e576006600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001600082015481600001556001820154816001015560028201548160020155600382015481600301556004820160009054906101000a900460ff168160040160006101000a81548160ff02191690831515021790555090505060016004600082825401925050819055505b6001600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060090160006101000a81548160ff021916908315150217905550565b6000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905084816000016000018190555083816000016001018190555082816000016002018190555081816000016003018190555060018160000160040160006101000a81548160ff0219169083151502179055505050505050565b6000600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160010154116104b957600080fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc600c60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549081150290604051600060405180830381858888f19350505050158015610541573d6000803e3d6000fd5b50600081600101819055505050565b6000600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600015158160020160009054906101000a900460ff161515146105b557600080fd5b60018160020160006101000a81548160ff02191690831515021790555060005434146105e057600080fd5b6105e8610a28565b8160030160010181905550600554816003016000018190555060016002600082825401925050819055505050565b61061e610207565b6000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905081816006016001018190555082816006016000018190555060018160050160006101000a81548160ff021916908315150217905550505050565b6000600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060020160009054906101000a900460ff166106f957600080fd5b6000600d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600115158160050160009054906101000a900460ff161515141561077557816003016000015481600601600001541461077457600080fd5b5b6000600e60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600d0154141561084f57898160000181905550888160010181905550878160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508681600801600001819055508581600801600101819055508481600801600201819055508381600801600301819055506108bc565b8981600001541461085f57600080fd5b8881600101541461086f57600080fd5b8681600801600001541461088257600080fd5b8581600801600101541461089557600080fd5b848160080160020154146108a857600080fd5b838160080160030154146108bb57600080fd5b5b600181600d016000828254019250508190555060006108e18360000189898989610a88565b90506001548a014310610a025760035482600d01541061090e576109058982610ac4565b50505050610a1f565b600e60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600382016000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff02191690555050600882016000808201600090556001820160009055600282016000905560038201600090555050600c820160009055600d820160009055505050505050610a1f565b605e82600d01541415610a1a57610a198982610ac4565b5b505050505b50505050505050565b600080600b600060055481526020019081526020016000209050605e816001015410610a605760016005600082825401925050819055505b600081600101546001901b905060018260010160008282540192505081905550809250505090565b60008085876000015402810190508487600101540281019050828760030154028101905083876002015402810190508091505095945050505050565b600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060080154811115610b1357600080fd5b80600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060080160008282540392505081905550600e60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600382016000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff02191690555050600882016000808201600090556001820160009055600282016000905560038201600090555050600c820160009055600d8201600090555050505056fea265627a7a7231582068d7d7fd972ed23483f46f1bef288a2a097d3d80bf775ef0fee9842e43c4554564736f6c63430005100032"

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

// DepositWithNodes is a paid mutator transaction binding the contract method 0x8e8b45c6.
//
// Solidity: function depositWithNodes(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) DepositWithNodes(opts *bind.TransactOpts, _pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "depositWithNodes", _pull, _places)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0x8e8b45c6.
//
// Solidity: function depositWithNodes(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractSession) DepositWithNodes(_pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.DepositWithNodes(&_Ethdepositcontract.TransactOpts, _pull, _places)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0x8e8b45c6.
//
// Solidity: function depositWithNodes(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) DepositWithNodes(_pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.DepositWithNodes(&_Ethdepositcontract.TransactOpts, _pull, _places)
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
