// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

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

// ProtonManagerABI is the input ABI used to generate the binding from.
const ProtonManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"protonAddress\",\"type\":\"bytes32\"}],\"name\":\"bindProtonAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"changeServicePrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"protonAddresses\",\"type\":\"bytes32[]\"}],\"name\":\"multiBind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"protonAddresses\",\"type\":\"bytes32[]\"}],\"name\":\"multiUnbind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"protonAddress\",\"type\":\"bytes32\"}],\"name\":\"unbindProtonAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"protonTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"checkBinder\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"protonAddress\",\"type\":\"bytes32\"}],\"name\":\"checkProtonAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"EtherCounter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ProtonCoin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ProtonUserMap\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TokenNoPerProtonAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProtonManager is an auto generated Go binding around an Ethereum contract.
type ProtonManager struct {
	ProtonManagerCaller     // Read-only binding to the contract
	ProtonManagerTransactor // Write-only binding to the contract
	ProtonManagerFilterer   // Log filterer for contract events
}

// ProtonManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtonManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtonManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtonManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtonManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtonManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtonManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtonManagerSession struct {
	Contract     *ProtonManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtonManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtonManagerCallerSession struct {
	Contract *ProtonManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProtonManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtonManagerTransactorSession struct {
	Contract     *ProtonManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProtonManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtonManagerRaw struct {
	Contract *ProtonManager // Generic contract binding to access the raw methods on
}

// ProtonManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtonManagerCallerRaw struct {
	Contract *ProtonManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ProtonManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtonManagerTransactorRaw struct {
	Contract *ProtonManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtonManager creates a new instance of ProtonManager, bound to a specific deployed contract.
func NewProtonManager(address common.Address, backend bind.ContractBackend) (*ProtonManager, error) {
	contract, err := bindProtonManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtonManager{ProtonManagerCaller: ProtonManagerCaller{contract: contract}, ProtonManagerTransactor: ProtonManagerTransactor{contract: contract}, ProtonManagerFilterer: ProtonManagerFilterer{contract: contract}}, nil
}

// NewProtonManagerCaller creates a new read-only instance of ProtonManager, bound to a specific deployed contract.
func NewProtonManagerCaller(address common.Address, caller bind.ContractCaller) (*ProtonManagerCaller, error) {
	contract, err := bindProtonManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtonManagerCaller{contract: contract}, nil
}

// NewProtonManagerTransactor creates a new write-only instance of ProtonManager, bound to a specific deployed contract.
func NewProtonManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtonManagerTransactor, error) {
	contract, err := bindProtonManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtonManagerTransactor{contract: contract}, nil
}

// NewProtonManagerFilterer creates a new log filterer instance of ProtonManager, bound to a specific deployed contract.
func NewProtonManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtonManagerFilterer, error) {
	contract, err := bindProtonManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtonManagerFilterer{contract: contract}, nil
}

// bindProtonManager binds a generic wrapper to an already deployed contract.
func bindProtonManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtonManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtonManager *ProtonManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProtonManager.Contract.ProtonManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtonManager *ProtonManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtonManager.Contract.ProtonManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtonManager *ProtonManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtonManager.Contract.ProtonManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtonManager *ProtonManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProtonManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtonManager *ProtonManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtonManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtonManager *ProtonManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtonManager.Contract.contract.Transact(opts, method, params...)
}

// EtherCounter is a free data retrieval call binding the contract method 0x8f048ca1.
//
// Solidity: function EtherCounter(address ) constant returns(uint256)
func (_ProtonManager *ProtonManagerCaller) EtherCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProtonManager.contract.Call(opts, out, "EtherCounter", arg0)
	return *ret0, err
}

// EtherCounter is a free data retrieval call binding the contract method 0x8f048ca1.
//
// Solidity: function EtherCounter(address ) constant returns(uint256)
func (_ProtonManager *ProtonManagerSession) EtherCounter(arg0 common.Address) (*big.Int, error) {
	return _ProtonManager.Contract.EtherCounter(&_ProtonManager.CallOpts, arg0)
}

// EtherCounter is a free data retrieval call binding the contract method 0x8f048ca1.
//
// Solidity: function EtherCounter(address ) constant returns(uint256)
func (_ProtonManager *ProtonManagerCallerSession) EtherCounter(arg0 common.Address) (*big.Int, error) {
	return _ProtonManager.Contract.EtherCounter(&_ProtonManager.CallOpts, arg0)
}

// ProtonCoin is a free data retrieval call binding the contract method 0xf26eec0f.
//
// Solidity: function ProtonCoin() constant returns(address)
func (_ProtonManager *ProtonManagerCaller) ProtonCoin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProtonManager.contract.Call(opts, out, "ProtonCoin")
	return *ret0, err
}

// ProtonCoin is a free data retrieval call binding the contract method 0xf26eec0f.
//
// Solidity: function ProtonCoin() constant returns(address)
func (_ProtonManager *ProtonManagerSession) ProtonCoin() (common.Address, error) {
	return _ProtonManager.Contract.ProtonCoin(&_ProtonManager.CallOpts)
}

// ProtonCoin is a free data retrieval call binding the contract method 0xf26eec0f.
//
// Solidity: function ProtonCoin() constant returns(address)
func (_ProtonManager *ProtonManagerCallerSession) ProtonCoin() (common.Address, error) {
	return _ProtonManager.Contract.ProtonCoin(&_ProtonManager.CallOpts)
}

// ProtonUserMap is a free data retrieval call binding the contract method 0x8c7064de.
//
// Solidity: function ProtonUserMap(bytes32 ) constant returns(address)
func (_ProtonManager *ProtonManagerCaller) ProtonUserMap(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProtonManager.contract.Call(opts, out, "ProtonUserMap", arg0)
	return *ret0, err
}

// ProtonUserMap is a free data retrieval call binding the contract method 0x8c7064de.
//
// Solidity: function ProtonUserMap(bytes32 ) constant returns(address)
func (_ProtonManager *ProtonManagerSession) ProtonUserMap(arg0 [32]byte) (common.Address, error) {
	return _ProtonManager.Contract.ProtonUserMap(&_ProtonManager.CallOpts, arg0)
}

// ProtonUserMap is a free data retrieval call binding the contract method 0x8c7064de.
//
// Solidity: function ProtonUserMap(bytes32 ) constant returns(address)
func (_ProtonManager *ProtonManagerCallerSession) ProtonUserMap(arg0 [32]byte) (common.Address, error) {
	return _ProtonManager.Contract.ProtonUserMap(&_ProtonManager.CallOpts, arg0)
}

// TokenNoPerProtonAccount is a free data retrieval call binding the contract method 0x2e27b20d.
//
// Solidity: function TokenNoPerProtonAccount() constant returns(uint256)
func (_ProtonManager *ProtonManagerCaller) TokenNoPerProtonAccount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProtonManager.contract.Call(opts, out, "TokenNoPerProtonAccount")
	return *ret0, err
}

// TokenNoPerProtonAccount is a free data retrieval call binding the contract method 0x2e27b20d.
//
// Solidity: function TokenNoPerProtonAccount() constant returns(uint256)
func (_ProtonManager *ProtonManagerSession) TokenNoPerProtonAccount() (*big.Int, error) {
	return _ProtonManager.Contract.TokenNoPerProtonAccount(&_ProtonManager.CallOpts)
}

// TokenNoPerProtonAccount is a free data retrieval call binding the contract method 0x2e27b20d.
//
// Solidity: function TokenNoPerProtonAccount() constant returns(uint256)
func (_ProtonManager *ProtonManagerCallerSession) TokenNoPerProtonAccount() (*big.Int, error) {
	return _ProtonManager.Contract.TokenNoPerProtonAccount(&_ProtonManager.CallOpts)
}

// CheckBinder is a free data retrieval call binding the contract method 0x2d5827d8.
//
// Solidity: function checkBinder(address userAddr) constant returns(uint256, uint256, uint256)
func (_ProtonManager *ProtonManagerCaller) CheckBinder(opts *bind.CallOpts, userAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ProtonManager.contract.Call(opts, out, "checkBinder", userAddr)
	return *ret0, *ret1, *ret2, err
}

// CheckBinder is a free data retrieval call binding the contract method 0x2d5827d8.
//
// Solidity: function checkBinder(address userAddr) constant returns(uint256, uint256, uint256)
func (_ProtonManager *ProtonManagerSession) CheckBinder(userAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _ProtonManager.Contract.CheckBinder(&_ProtonManager.CallOpts, userAddr)
}

// CheckBinder is a free data retrieval call binding the contract method 0x2d5827d8.
//
// Solidity: function checkBinder(address userAddr) constant returns(uint256, uint256, uint256)
func (_ProtonManager *ProtonManagerCallerSession) CheckBinder(userAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _ProtonManager.Contract.CheckBinder(&_ProtonManager.CallOpts, userAddr)
}

// CheckProtonAddress is a free data retrieval call binding the contract method 0xe46cfff3.
//
// Solidity: function checkProtonAddress(bytes32 protonAddress) constant returns(address, uint256, uint256)
func (_ProtonManager *ProtonManagerCaller) CheckProtonAddress(opts *bind.CallOpts, protonAddress [32]byte) (common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ProtonManager.contract.Call(opts, out, "checkProtonAddress", protonAddress)
	return *ret0, *ret1, *ret2, err
}

// CheckProtonAddress is a free data retrieval call binding the contract method 0xe46cfff3.
//
// Solidity: function checkProtonAddress(bytes32 protonAddress) constant returns(address, uint256, uint256)
func (_ProtonManager *ProtonManagerSession) CheckProtonAddress(protonAddress [32]byte) (common.Address, *big.Int, *big.Int, error) {
	return _ProtonManager.Contract.CheckProtonAddress(&_ProtonManager.CallOpts, protonAddress)
}

// CheckProtonAddress is a free data retrieval call binding the contract method 0xe46cfff3.
//
// Solidity: function checkProtonAddress(bytes32 protonAddress) constant returns(address, uint256, uint256)
func (_ProtonManager *ProtonManagerCallerSession) CheckProtonAddress(protonAddress [32]byte) (common.Address, *big.Int, *big.Int, error) {
	return _ProtonManager.Contract.CheckProtonAddress(&_ProtonManager.CallOpts, protonAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProtonManager *ProtonManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProtonManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProtonManager *ProtonManagerSession) Owner() (common.Address, error) {
	return _ProtonManager.Contract.Owner(&_ProtonManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProtonManager *ProtonManagerCallerSession) Owner() (common.Address, error) {
	return _ProtonManager.Contract.Owner(&_ProtonManager.CallOpts)
}

// BindProtonAddress is a paid mutator transaction binding the contract method 0x0d4711c2.
//
// Solidity: function bindProtonAddress(bytes32 protonAddress) returns()
func (_ProtonManager *ProtonManagerTransactor) BindProtonAddress(opts *bind.TransactOpts, protonAddress [32]byte) (*types.Transaction, error) {
	return _ProtonManager.contract.Transact(opts, "bindProtonAddress", protonAddress)
}

// BindProtonAddress is a paid mutator transaction binding the contract method 0x0d4711c2.
//
// Solidity: function bindProtonAddress(bytes32 protonAddress) returns()
func (_ProtonManager *ProtonManagerSession) BindProtonAddress(protonAddress [32]byte) (*types.Transaction, error) {
	return _ProtonManager.Contract.BindProtonAddress(&_ProtonManager.TransactOpts, protonAddress)
}

// BindProtonAddress is a paid mutator transaction binding the contract method 0x0d4711c2.
//
// Solidity: function bindProtonAddress(bytes32 protonAddress) returns()
func (_ProtonManager *ProtonManagerTransactorSession) BindProtonAddress(protonAddress [32]byte) (*types.Transaction, error) {
	return _ProtonManager.Contract.BindProtonAddress(&_ProtonManager.TransactOpts, protonAddress)
}

// ChangeServicePrice is a paid mutator transaction binding the contract method 0x0987b62c.
//
// Solidity: function changeServicePrice(uint256 price) returns()
func (_ProtonManager *ProtonManagerTransactor) ChangeServicePrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _ProtonManager.contract.Transact(opts, "changeServicePrice", price)
}

// ChangeServicePrice is a paid mutator transaction binding the contract method 0x0987b62c.
//
// Solidity: function changeServicePrice(uint256 price) returns()
func (_ProtonManager *ProtonManagerSession) ChangeServicePrice(price *big.Int) (*types.Transaction, error) {
	return _ProtonManager.Contract.ChangeServicePrice(&_ProtonManager.TransactOpts, price)
}

// ChangeServicePrice is a paid mutator transaction binding the contract method 0x0987b62c.
//
// Solidity: function changeServicePrice(uint256 price) returns()
func (_ProtonManager *ProtonManagerTransactorSession) ChangeServicePrice(price *big.Int) (*types.Transaction, error) {
	return _ProtonManager.Contract.ChangeServicePrice(&_ProtonManager.TransactOpts, price)
}

// MultiBind is a paid mutator transaction binding the contract method 0x55a3083d.
//
// Solidity: function multiBind(bytes32[] protonAddresses) returns()
func (_ProtonManager *ProtonManagerTransactor) MultiBind(opts *bind.TransactOpts, protonAddresses [][32]byte) (*types.Transaction, error) {
	return _ProtonManager.contract.Transact(opts, "multiBind", protonAddresses)
}

// MultiBind is a paid mutator transaction binding the contract method 0x55a3083d.
//
// Solidity: function multiBind(bytes32[] protonAddresses) returns()
func (_ProtonManager *ProtonManagerSession) MultiBind(protonAddresses [][32]byte) (*types.Transaction, error) {
	return _ProtonManager.Contract.MultiBind(&_ProtonManager.TransactOpts, protonAddresses)
}

// MultiBind is a paid mutator transaction binding the contract method 0x55a3083d.
//
// Solidity: function multiBind(bytes32[] protonAddresses) returns()
func (_ProtonManager *ProtonManagerTransactorSession) MultiBind(protonAddresses [][32]byte) (*types.Transaction, error) {
	return _ProtonManager.Contract.MultiBind(&_ProtonManager.TransactOpts, protonAddresses)
}

// MultiUnbind is a paid mutator transaction binding the contract method 0x7a264485.
//
// Solidity: function multiUnbind(bytes32[] protonAddresses) returns()
func (_ProtonManager *ProtonManagerTransactor) MultiUnbind(opts *bind.TransactOpts, protonAddresses [][32]byte) (*types.Transaction, error) {
	return _ProtonManager.contract.Transact(opts, "multiUnbind", protonAddresses)
}

// MultiUnbind is a paid mutator transaction binding the contract method 0x7a264485.
//
// Solidity: function multiUnbind(bytes32[] protonAddresses) returns()
func (_ProtonManager *ProtonManagerSession) MultiUnbind(protonAddresses [][32]byte) (*types.Transaction, error) {
	return _ProtonManager.Contract.MultiUnbind(&_ProtonManager.TransactOpts, protonAddresses)
}

// MultiUnbind is a paid mutator transaction binding the contract method 0x7a264485.
//
// Solidity: function multiUnbind(bytes32[] protonAddresses) returns()
func (_ProtonManager *ProtonManagerTransactorSession) MultiUnbind(protonAddresses [][32]byte) (*types.Transaction, error) {
	return _ProtonManager.Contract.MultiUnbind(&_ProtonManager.TransactOpts, protonAddresses)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProtonManager *ProtonManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProtonManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProtonManager *ProtonManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProtonManager.Contract.TransferOwnership(&_ProtonManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProtonManager *ProtonManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProtonManager.Contract.TransferOwnership(&_ProtonManager.TransactOpts, newOwner)
}

// UnbindProtonAddress is a paid mutator transaction binding the contract method 0xfe47f2e6.
//
// Solidity: function unbindProtonAddress(bytes32 protonAddress) returns()
func (_ProtonManager *ProtonManagerTransactor) UnbindProtonAddress(opts *bind.TransactOpts, protonAddress [32]byte) (*types.Transaction, error) {
	return _ProtonManager.contract.Transact(opts, "unbindProtonAddress", protonAddress)
}

// UnbindProtonAddress is a paid mutator transaction binding the contract method 0xfe47f2e6.
//
// Solidity: function unbindProtonAddress(bytes32 protonAddress) returns()
func (_ProtonManager *ProtonManagerSession) UnbindProtonAddress(protonAddress [32]byte) (*types.Transaction, error) {
	return _ProtonManager.Contract.UnbindProtonAddress(&_ProtonManager.TransactOpts, protonAddress)
}

// UnbindProtonAddress is a paid mutator transaction binding the contract method 0xfe47f2e6.
//
// Solidity: function unbindProtonAddress(bytes32 protonAddress) returns()
func (_ProtonManager *ProtonManagerTransactorSession) UnbindProtonAddress(protonAddress [32]byte) (*types.Transaction, error) {
	return _ProtonManager.Contract.UnbindProtonAddress(&_ProtonManager.TransactOpts, protonAddress)
}
