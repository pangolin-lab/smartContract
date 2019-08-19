// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// PPNAdminABI is the input ABI used to generate the binding from.
const PPNAdminABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"MinerPoolManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"mpm\",\"type\":\"address\"}],\"name\":\"ChangeMinerPoolManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MinerManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"mm\",\"type\":\"address\"}],\"name\":\"ChangeMinerManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PayChannel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pc\",\"type\":\"address\"}],\"name\":\"ChangeMicroPay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initToken\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PPNAdmin is an auto generated Go binding around an Ethereum contract.
type PPNAdmin struct {
	PPNAdminCaller     // Read-only binding to the contract
	PPNAdminTransactor // Write-only binding to the contract
	PPNAdminFilterer   // Log filterer for contract events
}

// PPNAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type PPNAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PPNAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PPNAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PPNAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PPNAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PPNAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PPNAdminSession struct {
	Contract     *PPNAdmin         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PPNAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PPNAdminCallerSession struct {
	Contract *PPNAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PPNAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PPNAdminTransactorSession struct {
	Contract     *PPNAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PPNAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type PPNAdminRaw struct {
	Contract *PPNAdmin // Generic contract binding to access the raw methods on
}

// PPNAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PPNAdminCallerRaw struct {
	Contract *PPNAdminCaller // Generic read-only contract binding to access the raw methods on
}

// PPNAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PPNAdminTransactorRaw struct {
	Contract *PPNAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPPNAdmin creates a new instance of PPNAdmin, bound to a specific deployed contract.
func NewPPNAdmin(address common.Address, backend bind.ContractBackend) (*PPNAdmin, error) {
	contract, err := bindPPNAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PPNAdmin{PPNAdminCaller: PPNAdminCaller{contract: contract}, PPNAdminTransactor: PPNAdminTransactor{contract: contract}, PPNAdminFilterer: PPNAdminFilterer{contract: contract}}, nil
}

// NewPPNAdminCaller creates a new read-only instance of PPNAdmin, bound to a specific deployed contract.
func NewPPNAdminCaller(address common.Address, caller bind.ContractCaller) (*PPNAdminCaller, error) {
	contract, err := bindPPNAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PPNAdminCaller{contract: contract}, nil
}

// NewPPNAdminTransactor creates a new write-only instance of PPNAdmin, bound to a specific deployed contract.
func NewPPNAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*PPNAdminTransactor, error) {
	contract, err := bindPPNAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PPNAdminTransactor{contract: contract}, nil
}

// NewPPNAdminFilterer creates a new log filterer instance of PPNAdmin, bound to a specific deployed contract.
func NewPPNAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*PPNAdminFilterer, error) {
	contract, err := bindPPNAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PPNAdminFilterer{contract: contract}, nil
}

// bindPPNAdmin binds a generic wrapper to an already deployed contract.
func bindPPNAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PPNAdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PPNAdmin *PPNAdminRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PPNAdmin.Contract.PPNAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PPNAdmin *PPNAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PPNAdmin.Contract.PPNAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PPNAdmin *PPNAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PPNAdmin.Contract.PPNAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PPNAdmin *PPNAdminCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PPNAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PPNAdmin *PPNAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PPNAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PPNAdmin *PPNAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PPNAdmin.Contract.contract.Transact(opts, method, params...)
}

// MinerManager is a free data retrieval call binding the contract method 0x6f270d68.
//
// Solidity: function MinerManager() constant returns(address)
func (_PPNAdmin *PPNAdminCaller) MinerManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PPNAdmin.contract.Call(opts, out, "MinerManager")
	return *ret0, err
}

// MinerManager is a free data retrieval call binding the contract method 0x6f270d68.
//
// Solidity: function MinerManager() constant returns(address)
func (_PPNAdmin *PPNAdminSession) MinerManager() (common.Address, error) {
	return _PPNAdmin.Contract.MinerManager(&_PPNAdmin.CallOpts)
}

// MinerManager is a free data retrieval call binding the contract method 0x6f270d68.
//
// Solidity: function MinerManager() constant returns(address)
func (_PPNAdmin *PPNAdminCallerSession) MinerManager() (common.Address, error) {
	return _PPNAdmin.Contract.MinerManager(&_PPNAdmin.CallOpts)
}

// MinerPoolManager is a free data retrieval call binding the contract method 0x69de0ca3.
//
// Solidity: function MinerPoolManager() constant returns(address)
func (_PPNAdmin *PPNAdminCaller) MinerPoolManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PPNAdmin.contract.Call(opts, out, "MinerPoolManager")
	return *ret0, err
}

// MinerPoolManager is a free data retrieval call binding the contract method 0x69de0ca3.
//
// Solidity: function MinerPoolManager() constant returns(address)
func (_PPNAdmin *PPNAdminSession) MinerPoolManager() (common.Address, error) {
	return _PPNAdmin.Contract.MinerPoolManager(&_PPNAdmin.CallOpts)
}

// MinerPoolManager is a free data retrieval call binding the contract method 0x69de0ca3.
//
// Solidity: function MinerPoolManager() constant returns(address)
func (_PPNAdmin *PPNAdminCallerSession) MinerPoolManager() (common.Address, error) {
	return _PPNAdmin.Contract.MinerPoolManager(&_PPNAdmin.CallOpts)
}

// PayChannel is a free data retrieval call binding the contract method 0xd6bccf26.
//
// Solidity: function PayChannel() constant returns(address)
func (_PPNAdmin *PPNAdminCaller) PayChannel(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PPNAdmin.contract.Call(opts, out, "PayChannel")
	return *ret0, err
}

// PayChannel is a free data retrieval call binding the contract method 0xd6bccf26.
//
// Solidity: function PayChannel() constant returns(address)
func (_PPNAdmin *PPNAdminSession) PayChannel() (common.Address, error) {
	return _PPNAdmin.Contract.PayChannel(&_PPNAdmin.CallOpts)
}

// PayChannel is a free data retrieval call binding the contract method 0xd6bccf26.
//
// Solidity: function PayChannel() constant returns(address)
func (_PPNAdmin *PPNAdminCallerSession) PayChannel() (common.Address, error) {
	return _PPNAdmin.Contract.PayChannel(&_PPNAdmin.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0xc2cba306.
//
// Solidity: function TokenAddress() constant returns(address)
func (_PPNAdmin *PPNAdminCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PPNAdmin.contract.Call(opts, out, "TokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0xc2cba306.
//
// Solidity: function TokenAddress() constant returns(address)
func (_PPNAdmin *PPNAdminSession) TokenAddress() (common.Address, error) {
	return _PPNAdmin.Contract.TokenAddress(&_PPNAdmin.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0xc2cba306.
//
// Solidity: function TokenAddress() constant returns(address)
func (_PPNAdmin *PPNAdminCallerSession) TokenAddress() (common.Address, error) {
	return _PPNAdmin.Contract.TokenAddress(&_PPNAdmin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PPNAdmin *PPNAdminCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PPNAdmin.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PPNAdmin *PPNAdminSession) Owner() (common.Address, error) {
	return _PPNAdmin.Contract.Owner(&_PPNAdmin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PPNAdmin *PPNAdminCallerSession) Owner() (common.Address, error) {
	return _PPNAdmin.Contract.Owner(&_PPNAdmin.CallOpts)
}

// ChangeMicroPay is a paid mutator transaction binding the contract method 0xe035d60a.
//
// Solidity: function ChangeMicroPay(address pc) returns()
func (_PPNAdmin *PPNAdminTransactor) ChangeMicroPay(opts *bind.TransactOpts, pc common.Address) (*types.Transaction, error) {
	return _PPNAdmin.contract.Transact(opts, "ChangeMicroPay", pc)
}

// ChangeMicroPay is a paid mutator transaction binding the contract method 0xe035d60a.
//
// Solidity: function ChangeMicroPay(address pc) returns()
func (_PPNAdmin *PPNAdminSession) ChangeMicroPay(pc common.Address) (*types.Transaction, error) {
	return _PPNAdmin.Contract.ChangeMicroPay(&_PPNAdmin.TransactOpts, pc)
}

// ChangeMicroPay is a paid mutator transaction binding the contract method 0xe035d60a.
//
// Solidity: function ChangeMicroPay(address pc) returns()
func (_PPNAdmin *PPNAdminTransactorSession) ChangeMicroPay(pc common.Address) (*types.Transaction, error) {
	return _PPNAdmin.Contract.ChangeMicroPay(&_PPNAdmin.TransactOpts, pc)
}

// ChangeMinerManager is a paid mutator transaction binding the contract method 0xc5a09d3b.
//
// Solidity: function ChangeMinerManager(address mm) returns()
func (_PPNAdmin *PPNAdminTransactor) ChangeMinerManager(opts *bind.TransactOpts, mm common.Address) (*types.Transaction, error) {
	return _PPNAdmin.contract.Transact(opts, "ChangeMinerManager", mm)
}

// ChangeMinerManager is a paid mutator transaction binding the contract method 0xc5a09d3b.
//
// Solidity: function ChangeMinerManager(address mm) returns()
func (_PPNAdmin *PPNAdminSession) ChangeMinerManager(mm common.Address) (*types.Transaction, error) {
	return _PPNAdmin.Contract.ChangeMinerManager(&_PPNAdmin.TransactOpts, mm)
}

// ChangeMinerManager is a paid mutator transaction binding the contract method 0xc5a09d3b.
//
// Solidity: function ChangeMinerManager(address mm) returns()
func (_PPNAdmin *PPNAdminTransactorSession) ChangeMinerManager(mm common.Address) (*types.Transaction, error) {
	return _PPNAdmin.Contract.ChangeMinerManager(&_PPNAdmin.TransactOpts, mm)
}

// ChangeMinerPoolManager is a paid mutator transaction binding the contract method 0x6afe7424.
//
// Solidity: function ChangeMinerPoolManager(address mpm) returns()
func (_PPNAdmin *PPNAdminTransactor) ChangeMinerPoolManager(opts *bind.TransactOpts, mpm common.Address) (*types.Transaction, error) {
	return _PPNAdmin.contract.Transact(opts, "ChangeMinerPoolManager", mpm)
}

// ChangeMinerPoolManager is a paid mutator transaction binding the contract method 0x6afe7424.
//
// Solidity: function ChangeMinerPoolManager(address mpm) returns()
func (_PPNAdmin *PPNAdminSession) ChangeMinerPoolManager(mpm common.Address) (*types.Transaction, error) {
	return _PPNAdmin.Contract.ChangeMinerPoolManager(&_PPNAdmin.TransactOpts, mpm)
}

// ChangeMinerPoolManager is a paid mutator transaction binding the contract method 0x6afe7424.
//
// Solidity: function ChangeMinerPoolManager(address mpm) returns()
func (_PPNAdmin *PPNAdminTransactorSession) ChangeMinerPoolManager(mpm common.Address) (*types.Transaction, error) {
	return _PPNAdmin.Contract.ChangeMinerPoolManager(&_PPNAdmin.TransactOpts, mpm)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PPNAdmin *PPNAdminTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PPNAdmin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PPNAdmin *PPNAdminSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PPNAdmin.Contract.TransferOwnership(&_PPNAdmin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PPNAdmin *PPNAdminTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PPNAdmin.Contract.TransferOwnership(&_PPNAdmin.TransactOpts, newOwner)
}
