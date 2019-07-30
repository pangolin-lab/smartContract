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

// SimpleProtonManagerABI is the input ABI used to generate the binding from.
const SimpleProtonManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"bytes32\"}],\"name\":\"bind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"changeServicePrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"bytes32[]\"}],\"name\":\"multiBind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"bytes32[]\"}],\"name\":\"multiUnbind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"bytes32\"}],\"name\":\"unbind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"AddressesCounter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AddressesUnderMyAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"basicBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"BoundedEthAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"bytes32\"}],\"name\":\"check\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ethAddr\",\"type\":\"address\"}],\"name\":\"search\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"servicePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SimpleProtonManager is an auto generated Go binding around an Ethereum contract.
type SimpleProtonManager struct {
	SimpleProtonManagerCaller     // Read-only binding to the contract
	SimpleProtonManagerTransactor // Write-only binding to the contract
	SimpleProtonManagerFilterer   // Log filterer for contract events
}

// SimpleProtonManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleProtonManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleProtonManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleProtonManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleProtonManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleProtonManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleProtonManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleProtonManagerSession struct {
	Contract     *SimpleProtonManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SimpleProtonManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleProtonManagerCallerSession struct {
	Contract *SimpleProtonManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SimpleProtonManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleProtonManagerTransactorSession struct {
	Contract     *SimpleProtonManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SimpleProtonManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleProtonManagerRaw struct {
	Contract *SimpleProtonManager // Generic contract binding to access the raw methods on
}

// SimpleProtonManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleProtonManagerCallerRaw struct {
	Contract *SimpleProtonManagerCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleProtonManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleProtonManagerTransactorRaw struct {
	Contract *SimpleProtonManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleProtonManager creates a new instance of SimpleProtonManager, bound to a specific deployed contract.
func NewSimpleProtonManager(address common.Address, backend bind.ContractBackend) (*SimpleProtonManager, error) {
	contract, err := bindSimpleProtonManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleProtonManager{SimpleProtonManagerCaller: SimpleProtonManagerCaller{contract: contract}, SimpleProtonManagerTransactor: SimpleProtonManagerTransactor{contract: contract}, SimpleProtonManagerFilterer: SimpleProtonManagerFilterer{contract: contract}}, nil
}

// NewSimpleProtonManagerCaller creates a new read-only instance of SimpleProtonManager, bound to a specific deployed contract.
func NewSimpleProtonManagerCaller(address common.Address, caller bind.ContractCaller) (*SimpleProtonManagerCaller, error) {
	contract, err := bindSimpleProtonManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleProtonManagerCaller{contract: contract}, nil
}

// NewSimpleProtonManagerTransactor creates a new write-only instance of SimpleProtonManager, bound to a specific deployed contract.
func NewSimpleProtonManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleProtonManagerTransactor, error) {
	contract, err := bindSimpleProtonManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleProtonManagerTransactor{contract: contract}, nil
}

// NewSimpleProtonManagerFilterer creates a new log filterer instance of SimpleProtonManager, bound to a specific deployed contract.
func NewSimpleProtonManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleProtonManagerFilterer, error) {
	contract, err := bindSimpleProtonManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleProtonManagerFilterer{contract: contract}, nil
}

// bindSimpleProtonManager binds a generic wrapper to an already deployed contract.
func bindSimpleProtonManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleProtonManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleProtonManager *SimpleProtonManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleProtonManager.Contract.SimpleProtonManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleProtonManager *SimpleProtonManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.SimpleProtonManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleProtonManager *SimpleProtonManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.SimpleProtonManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleProtonManager *SimpleProtonManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleProtonManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleProtonManager *SimpleProtonManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleProtonManager *SimpleProtonManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.contract.Transact(opts, method, params...)
}

// AddressesCounter is a free data retrieval call binding the contract method 0x0fd98c14.
//
// Solidity: function AddressesCounter(address ) constant returns(uint256)
func (_SimpleProtonManager *SimpleProtonManagerCaller) AddressesCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleProtonManager.contract.Call(opts, out, "AddressesCounter", arg0)
	return *ret0, err
}

// AddressesCounter is a free data retrieval call binding the contract method 0x0fd98c14.
//
// Solidity: function AddressesCounter(address ) constant returns(uint256)
func (_SimpleProtonManager *SimpleProtonManagerSession) AddressesCounter(arg0 common.Address) (*big.Int, error) {
	return _SimpleProtonManager.Contract.AddressesCounter(&_SimpleProtonManager.CallOpts, arg0)
}

// AddressesCounter is a free data retrieval call binding the contract method 0x0fd98c14.
//
// Solidity: function AddressesCounter(address ) constant returns(uint256)
func (_SimpleProtonManager *SimpleProtonManagerCallerSession) AddressesCounter(arg0 common.Address) (*big.Int, error) {
	return _SimpleProtonManager.Contract.AddressesCounter(&_SimpleProtonManager.CallOpts, arg0)
}

// AddressesUnderMyAccount is a free data retrieval call binding the contract method 0x4f7dc6e7.
//
// Solidity: function AddressesUnderMyAccount(address , uint256 ) constant returns(bytes32)
func (_SimpleProtonManager *SimpleProtonManagerCaller) AddressesUnderMyAccount(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SimpleProtonManager.contract.Call(opts, out, "AddressesUnderMyAccount", arg0, arg1)
	return *ret0, err
}

// AddressesUnderMyAccount is a free data retrieval call binding the contract method 0x4f7dc6e7.
//
// Solidity: function AddressesUnderMyAccount(address , uint256 ) constant returns(bytes32)
func (_SimpleProtonManager *SimpleProtonManagerSession) AddressesUnderMyAccount(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _SimpleProtonManager.Contract.AddressesUnderMyAccount(&_SimpleProtonManager.CallOpts, arg0, arg1)
}

// AddressesUnderMyAccount is a free data retrieval call binding the contract method 0x4f7dc6e7.
//
// Solidity: function AddressesUnderMyAccount(address , uint256 ) constant returns(bytes32)
func (_SimpleProtonManager *SimpleProtonManagerCallerSession) AddressesUnderMyAccount(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _SimpleProtonManager.Contract.AddressesUnderMyAccount(&_SimpleProtonManager.CallOpts, arg0, arg1)
}

// BoundedEthAddress is a free data retrieval call binding the contract method 0x6fb12a5e.
//
// Solidity: function BoundedEthAddress(bytes32 ) constant returns(address)
func (_SimpleProtonManager *SimpleProtonManagerCaller) BoundedEthAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SimpleProtonManager.contract.Call(opts, out, "BoundedEthAddress", arg0)
	return *ret0, err
}

// BoundedEthAddress is a free data retrieval call binding the contract method 0x6fb12a5e.
//
// Solidity: function BoundedEthAddress(bytes32 ) constant returns(address)
func (_SimpleProtonManager *SimpleProtonManagerSession) BoundedEthAddress(arg0 [32]byte) (common.Address, error) {
	return _SimpleProtonManager.Contract.BoundedEthAddress(&_SimpleProtonManager.CallOpts, arg0)
}

// BoundedEthAddress is a free data retrieval call binding the contract method 0x6fb12a5e.
//
// Solidity: function BoundedEthAddress(bytes32 ) constant returns(address)
func (_SimpleProtonManager *SimpleProtonManagerCallerSession) BoundedEthAddress(arg0 [32]byte) (common.Address, error) {
	return _SimpleProtonManager.Contract.BoundedEthAddress(&_SimpleProtonManager.CallOpts, arg0)
}

// BasicBalance is a free data retrieval call binding the contract method 0x903e13ba.
//
// Solidity: function basicBalance(address target) constant returns(uint256, uint256)
func (_SimpleProtonManager *SimpleProtonManagerCaller) BasicBalance(opts *bind.CallOpts, target common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _SimpleProtonManager.contract.Call(opts, out, "basicBalance", target)
	return *ret0, *ret1, err
}

// BasicBalance is a free data retrieval call binding the contract method 0x903e13ba.
//
// Solidity: function basicBalance(address target) constant returns(uint256, uint256)
func (_SimpleProtonManager *SimpleProtonManagerSession) BasicBalance(target common.Address) (*big.Int, *big.Int, error) {
	return _SimpleProtonManager.Contract.BasicBalance(&_SimpleProtonManager.CallOpts, target)
}

// BasicBalance is a free data retrieval call binding the contract method 0x903e13ba.
//
// Solidity: function basicBalance(address target) constant returns(uint256, uint256)
func (_SimpleProtonManager *SimpleProtonManagerCallerSession) BasicBalance(target common.Address) (*big.Int, *big.Int, error) {
	return _SimpleProtonManager.Contract.BasicBalance(&_SimpleProtonManager.CallOpts, target)
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(bytes32 addr) constant returns(bool)
func (_SimpleProtonManager *SimpleProtonManagerCaller) Check(opts *bind.CallOpts, addr [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SimpleProtonManager.contract.Call(opts, out, "check", addr)
	return *ret0, err
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(bytes32 addr) constant returns(bool)
func (_SimpleProtonManager *SimpleProtonManagerSession) Check(addr [32]byte) (bool, error) {
	return _SimpleProtonManager.Contract.Check(&_SimpleProtonManager.CallOpts, addr)
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(bytes32 addr) constant returns(bool)
func (_SimpleProtonManager *SimpleProtonManagerCallerSession) Check(addr [32]byte) (bool, error) {
	return _SimpleProtonManager.Contract.Check(&_SimpleProtonManager.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SimpleProtonManager *SimpleProtonManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SimpleProtonManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SimpleProtonManager *SimpleProtonManagerSession) Owner() (common.Address, error) {
	return _SimpleProtonManager.Contract.Owner(&_SimpleProtonManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SimpleProtonManager *SimpleProtonManagerCallerSession) Owner() (common.Address, error) {
	return _SimpleProtonManager.Contract.Owner(&_SimpleProtonManager.CallOpts)
}

// Search is a free data retrieval call binding the contract method 0x657a2512.
//
// Solidity: function search(address ethAddr) constant returns(bytes32[])
func (_SimpleProtonManager *SimpleProtonManagerCaller) Search(opts *bind.CallOpts, ethAddr common.Address) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _SimpleProtonManager.contract.Call(opts, out, "search", ethAddr)
	return *ret0, err
}

// Search is a free data retrieval call binding the contract method 0x657a2512.
//
// Solidity: function search(address ethAddr) constant returns(bytes32[])
func (_SimpleProtonManager *SimpleProtonManagerSession) Search(ethAddr common.Address) ([][32]byte, error) {
	return _SimpleProtonManager.Contract.Search(&_SimpleProtonManager.CallOpts, ethAddr)
}

// Search is a free data retrieval call binding the contract method 0x657a2512.
//
// Solidity: function search(address ethAddr) constant returns(bytes32[])
func (_SimpleProtonManager *SimpleProtonManagerCallerSession) Search(ethAddr common.Address) ([][32]byte, error) {
	return _SimpleProtonManager.Contract.Search(&_SimpleProtonManager.CallOpts, ethAddr)
}

// ServicePrice is a free data retrieval call binding the contract method 0xec67d2d3.
//
// Solidity: function servicePrice() constant returns(uint256)
func (_SimpleProtonManager *SimpleProtonManagerCaller) ServicePrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleProtonManager.contract.Call(opts, out, "servicePrice")
	return *ret0, err
}

// ServicePrice is a free data retrieval call binding the contract method 0xec67d2d3.
//
// Solidity: function servicePrice() constant returns(uint256)
func (_SimpleProtonManager *SimpleProtonManagerSession) ServicePrice() (*big.Int, error) {
	return _SimpleProtonManager.Contract.ServicePrice(&_SimpleProtonManager.CallOpts)
}

// ServicePrice is a free data retrieval call binding the contract method 0xec67d2d3.
//
// Solidity: function servicePrice() constant returns(uint256)
func (_SimpleProtonManager *SimpleProtonManagerCallerSession) ServicePrice() (*big.Int, error) {
	return _SimpleProtonManager.Contract.ServicePrice(&_SimpleProtonManager.CallOpts)
}

// Bind is a paid mutator transaction binding the contract method 0x709152a1.
//
// Solidity: function bind(bytes32 addr) returns()
func (_SimpleProtonManager *SimpleProtonManagerTransactor) Bind(opts *bind.TransactOpts, addr [32]byte) (*types.Transaction, error) {
	return _SimpleProtonManager.contract.Transact(opts, "bind", addr)
}

// Bind is a paid mutator transaction binding the contract method 0x709152a1.
//
// Solidity: function bind(bytes32 addr) returns()
func (_SimpleProtonManager *SimpleProtonManagerSession) Bind(addr [32]byte) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.Bind(&_SimpleProtonManager.TransactOpts, addr)
}

// Bind is a paid mutator transaction binding the contract method 0x709152a1.
//
// Solidity: function bind(bytes32 addr) returns()
func (_SimpleProtonManager *SimpleProtonManagerTransactorSession) Bind(addr [32]byte) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.Bind(&_SimpleProtonManager.TransactOpts, addr)
}

// ChangeServicePrice is a paid mutator transaction binding the contract method 0x0987b62c.
//
// Solidity: function changeServicePrice(uint256 price) returns()
func (_SimpleProtonManager *SimpleProtonManagerTransactor) ChangeServicePrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _SimpleProtonManager.contract.Transact(opts, "changeServicePrice", price)
}

// ChangeServicePrice is a paid mutator transaction binding the contract method 0x0987b62c.
//
// Solidity: function changeServicePrice(uint256 price) returns()
func (_SimpleProtonManager *SimpleProtonManagerSession) ChangeServicePrice(price *big.Int) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.ChangeServicePrice(&_SimpleProtonManager.TransactOpts, price)
}

// ChangeServicePrice is a paid mutator transaction binding the contract method 0x0987b62c.
//
// Solidity: function changeServicePrice(uint256 price) returns()
func (_SimpleProtonManager *SimpleProtonManagerTransactorSession) ChangeServicePrice(price *big.Int) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.ChangeServicePrice(&_SimpleProtonManager.TransactOpts, price)
}

// MultiBind is a paid mutator transaction binding the contract method 0x55a3083d.
//
// Solidity: function multiBind(bytes32[] addresses) returns()
func (_SimpleProtonManager *SimpleProtonManagerTransactor) MultiBind(opts *bind.TransactOpts, addresses [][32]byte) (*types.Transaction, error) {
	return _SimpleProtonManager.contract.Transact(opts, "multiBind", addresses)
}

// MultiBind is a paid mutator transaction binding the contract method 0x55a3083d.
//
// Solidity: function multiBind(bytes32[] addresses) returns()
func (_SimpleProtonManager *SimpleProtonManagerSession) MultiBind(addresses [][32]byte) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.MultiBind(&_SimpleProtonManager.TransactOpts, addresses)
}

// MultiBind is a paid mutator transaction binding the contract method 0x55a3083d.
//
// Solidity: function multiBind(bytes32[] addresses) returns()
func (_SimpleProtonManager *SimpleProtonManagerTransactorSession) MultiBind(addresses [][32]byte) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.MultiBind(&_SimpleProtonManager.TransactOpts, addresses)
}

// MultiUnbind is a paid mutator transaction binding the contract method 0x7a264485.
//
// Solidity: function multiUnbind(bytes32[] addresses) returns()
func (_SimpleProtonManager *SimpleProtonManagerTransactor) MultiUnbind(opts *bind.TransactOpts, addresses [][32]byte) (*types.Transaction, error) {
	return _SimpleProtonManager.contract.Transact(opts, "multiUnbind", addresses)
}

// MultiUnbind is a paid mutator transaction binding the contract method 0x7a264485.
//
// Solidity: function multiUnbind(bytes32[] addresses) returns()
func (_SimpleProtonManager *SimpleProtonManagerSession) MultiUnbind(addresses [][32]byte) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.MultiUnbind(&_SimpleProtonManager.TransactOpts, addresses)
}

// MultiUnbind is a paid mutator transaction binding the contract method 0x7a264485.
//
// Solidity: function multiUnbind(bytes32[] addresses) returns()
func (_SimpleProtonManager *SimpleProtonManagerTransactorSession) MultiUnbind(addresses [][32]byte) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.MultiUnbind(&_SimpleProtonManager.TransactOpts, addresses)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SimpleProtonManager *SimpleProtonManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SimpleProtonManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SimpleProtonManager *SimpleProtonManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.TransferOwnership(&_SimpleProtonManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SimpleProtonManager *SimpleProtonManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.TransferOwnership(&_SimpleProtonManager.TransactOpts, newOwner)
}

// Unbind is a paid mutator transaction binding the contract method 0x6e08f245.
//
// Solidity: function unbind(bytes32 addr) returns()
func (_SimpleProtonManager *SimpleProtonManagerTransactor) Unbind(opts *bind.TransactOpts, addr [32]byte) (*types.Transaction, error) {
	return _SimpleProtonManager.contract.Transact(opts, "unbind", addr)
}

// Unbind is a paid mutator transaction binding the contract method 0x6e08f245.
//
// Solidity: function unbind(bytes32 addr) returns()
func (_SimpleProtonManager *SimpleProtonManagerSession) Unbind(addr [32]byte) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.Unbind(&_SimpleProtonManager.TransactOpts, addr)
}

// Unbind is a paid mutator transaction binding the contract method 0x6e08f245.
//
// Solidity: function unbind(bytes32 addr) returns()
func (_SimpleProtonManager *SimpleProtonManagerTransactorSession) Unbind(addr [32]byte) (*types.Transaction, error) {
	return _SimpleProtonManager.Contract.Unbind(&_SimpleProtonManager.TransactOpts, addr)
}
