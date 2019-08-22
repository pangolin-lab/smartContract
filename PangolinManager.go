// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethInterface

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

// PangolinManagerABI is the input ABI used to generate the binding from.
const PangolinManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"changeServicePrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"bytes32\"}],\"name\":\"check\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TokenNoForOneUser\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"bytes32\"}],\"name\":\"unbind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"bytes32\"}],\"name\":\"bind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"EtherCounter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"PangolinUserRecord\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"bindingInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PangolinManager is an auto generated Go binding around an Ethereum contract.
type PangolinManager struct {
	PangolinManagerCaller     // Read-only binding to the contract
	PangolinManagerTransactor // Write-only binding to the contract
	PangolinManagerFilterer   // Log filterer for contract events
}

// PangolinManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PangolinManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PangolinManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PangolinManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PangolinManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PangolinManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PangolinManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PangolinManagerSession struct {
	Contract     *PangolinManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PangolinManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PangolinManagerCallerSession struct {
	Contract *PangolinManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PangolinManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PangolinManagerTransactorSession struct {
	Contract     *PangolinManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PangolinManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PangolinManagerRaw struct {
	Contract *PangolinManager // Generic contract binding to access the raw methods on
}

// PangolinManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PangolinManagerCallerRaw struct {
	Contract *PangolinManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PangolinManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PangolinManagerTransactorRaw struct {
	Contract *PangolinManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPangolinManager creates a new instance of PangolinManager, bound to a specific deployed contract.
func NewPangolinManager(address common.Address, backend bind.ContractBackend) (*PangolinManager, error) {
	contract, err := bindPangolinManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PangolinManager{PangolinManagerCaller: PangolinManagerCaller{contract: contract}, PangolinManagerTransactor: PangolinManagerTransactor{contract: contract}, PangolinManagerFilterer: PangolinManagerFilterer{contract: contract}}, nil
}

// NewPangolinManagerCaller creates a new read-only instance of PangolinManager, bound to a specific deployed contract.
func NewPangolinManagerCaller(address common.Address, caller bind.ContractCaller) (*PangolinManagerCaller, error) {
	contract, err := bindPangolinManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PangolinManagerCaller{contract: contract}, nil
}

// NewPangolinManagerTransactor creates a new write-only instance of PangolinManager, bound to a specific deployed contract.
func NewPangolinManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PangolinManagerTransactor, error) {
	contract, err := bindPangolinManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PangolinManagerTransactor{contract: contract}, nil
}

// NewPangolinManagerFilterer creates a new log filterer instance of PangolinManager, bound to a specific deployed contract.
func NewPangolinManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PangolinManagerFilterer, error) {
	contract, err := bindPangolinManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PangolinManagerFilterer{contract: contract}, nil
}

// bindPangolinManager binds a generic wrapper to an already deployed contract.
func bindPangolinManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PangolinManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PangolinManager *PangolinManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PangolinManager.Contract.PangolinManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PangolinManager *PangolinManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PangolinManager.Contract.PangolinManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PangolinManager *PangolinManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PangolinManager.Contract.PangolinManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PangolinManager *PangolinManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PangolinManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PangolinManager *PangolinManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PangolinManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PangolinManager *PangolinManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PangolinManager.Contract.contract.Transact(opts, method, params...)
}

// EtherCounter is a free data retrieval call binding the contract method 0x8f048ca1.
//
// Solidity: function EtherCounter(address ) constant returns(uint256)
func (_PangolinManager *PangolinManagerCaller) EtherCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PangolinManager.contract.Call(opts, out, "EtherCounter", arg0)
	return *ret0, err
}

// EtherCounter is a free data retrieval call binding the contract method 0x8f048ca1.
//
// Solidity: function EtherCounter(address ) constant returns(uint256)
func (_PangolinManager *PangolinManagerSession) EtherCounter(arg0 common.Address) (*big.Int, error) {
	return _PangolinManager.Contract.EtherCounter(&_PangolinManager.CallOpts, arg0)
}

// EtherCounter is a free data retrieval call binding the contract method 0x8f048ca1.
//
// Solidity: function EtherCounter(address ) constant returns(uint256)
func (_PangolinManager *PangolinManagerCallerSession) EtherCounter(arg0 common.Address) (*big.Int, error) {
	return _PangolinManager.Contract.EtherCounter(&_PangolinManager.CallOpts, arg0)
}

// PangolinUserRecord is a free data retrieval call binding the contract method 0x9379ac5b.
//
// Solidity: function PangolinUserRecord(bytes32 ) constant returns(address)
func (_PangolinManager *PangolinManagerCaller) PangolinUserRecord(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PangolinManager.contract.Call(opts, out, "PangolinUserRecord", arg0)
	return *ret0, err
}

// PangolinUserRecord is a free data retrieval call binding the contract method 0x9379ac5b.
//
// Solidity: function PangolinUserRecord(bytes32 ) constant returns(address)
func (_PangolinManager *PangolinManagerSession) PangolinUserRecord(arg0 [32]byte) (common.Address, error) {
	return _PangolinManager.Contract.PangolinUserRecord(&_PangolinManager.CallOpts, arg0)
}

// PangolinUserRecord is a free data retrieval call binding the contract method 0x9379ac5b.
//
// Solidity: function PangolinUserRecord(bytes32 ) constant returns(address)
func (_PangolinManager *PangolinManagerCallerSession) PangolinUserRecord(arg0 [32]byte) (common.Address, error) {
	return _PangolinManager.Contract.PangolinUserRecord(&_PangolinManager.CallOpts, arg0)
}

// TokenNoForOneUser is a free data retrieval call binding the contract method 0x63692e12.
//
// Solidity: function TokenNoForOneUser() constant returns(uint256)
func (_PangolinManager *PangolinManagerCaller) TokenNoForOneUser(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PangolinManager.contract.Call(opts, out, "TokenNoForOneUser")
	return *ret0, err
}

// TokenNoForOneUser is a free data retrieval call binding the contract method 0x63692e12.
//
// Solidity: function TokenNoForOneUser() constant returns(uint256)
func (_PangolinManager *PangolinManagerSession) TokenNoForOneUser() (*big.Int, error) {
	return _PangolinManager.Contract.TokenNoForOneUser(&_PangolinManager.CallOpts)
}

// TokenNoForOneUser is a free data retrieval call binding the contract method 0x63692e12.
//
// Solidity: function TokenNoForOneUser() constant returns(uint256)
func (_PangolinManager *PangolinManagerCallerSession) TokenNoForOneUser() (*big.Int, error) {
	return _PangolinManager.Contract.TokenNoForOneUser(&_PangolinManager.CallOpts)
}

// BindingInfo is a free data retrieval call binding the contract method 0x965be58d.
//
// Solidity: function bindingInfo(address userAddr) constant returns(uint256, uint256, uint256)
func (_PangolinManager *PangolinManagerCaller) BindingInfo(opts *bind.CallOpts, userAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
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
	err := _PangolinManager.contract.Call(opts, out, "bindingInfo", userAddr)
	return *ret0, *ret1, *ret2, err
}

// BindingInfo is a free data retrieval call binding the contract method 0x965be58d.
//
// Solidity: function bindingInfo(address userAddr) constant returns(uint256, uint256, uint256)
func (_PangolinManager *PangolinManagerSession) BindingInfo(userAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _PangolinManager.Contract.BindingInfo(&_PangolinManager.CallOpts, userAddr)
}

// BindingInfo is a free data retrieval call binding the contract method 0x965be58d.
//
// Solidity: function bindingInfo(address userAddr) constant returns(uint256, uint256, uint256)
func (_PangolinManager *PangolinManagerCallerSession) BindingInfo(userAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _PangolinManager.Contract.BindingInfo(&_PangolinManager.CallOpts, userAddr)
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(bytes32 addr) constant returns(address, uint256, uint256)
func (_PangolinManager *PangolinManagerCaller) Check(opts *bind.CallOpts, addr [32]byte) (common.Address, *big.Int, *big.Int, error) {
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
	err := _PangolinManager.contract.Call(opts, out, "check", addr)
	return *ret0, *ret1, *ret2, err
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(bytes32 addr) constant returns(address, uint256, uint256)
func (_PangolinManager *PangolinManagerSession) Check(addr [32]byte) (common.Address, *big.Int, *big.Int, error) {
	return _PangolinManager.Contract.Check(&_PangolinManager.CallOpts, addr)
}

// Check is a free data retrieval call binding the contract method 0x399e0792.
//
// Solidity: function check(bytes32 addr) constant returns(address, uint256, uint256)
func (_PangolinManager *PangolinManagerCallerSession) Check(addr [32]byte) (common.Address, *big.Int, *big.Int, error) {
	return _PangolinManager.Contract.Check(&_PangolinManager.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PangolinManager *PangolinManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PangolinManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PangolinManager *PangolinManagerSession) Owner() (common.Address, error) {
	return _PangolinManager.Contract.Owner(&_PangolinManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PangolinManager *PangolinManagerCallerSession) Owner() (common.Address, error) {
	return _PangolinManager.Contract.Owner(&_PangolinManager.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PangolinManager *PangolinManagerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PangolinManager.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PangolinManager *PangolinManagerSession) Token() (common.Address, error) {
	return _PangolinManager.Contract.Token(&_PangolinManager.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PangolinManager *PangolinManagerCallerSession) Token() (common.Address, error) {
	return _PangolinManager.Contract.Token(&_PangolinManager.CallOpts)
}

// Bind is a paid mutator transaction binding the contract method 0x709152a1.
//
// Solidity: function bind(bytes32 addr) returns()
func (_PangolinManager *PangolinManagerTransactor) Bind(opts *bind.TransactOpts, addr [32]byte) (*types.Transaction, error) {
	return _PangolinManager.contract.Transact(opts, "bind", addr)
}

// Bind is a paid mutator transaction binding the contract method 0x709152a1.
//
// Solidity: function bind(bytes32 addr) returns()
func (_PangolinManager *PangolinManagerSession) Bind(addr [32]byte) (*types.Transaction, error) {
	return _PangolinManager.Contract.Bind(&_PangolinManager.TransactOpts, addr)
}

// Bind is a paid mutator transaction binding the contract method 0x709152a1.
//
// Solidity: function bind(bytes32 addr) returns()
func (_PangolinManager *PangolinManagerTransactorSession) Bind(addr [32]byte) (*types.Transaction, error) {
	return _PangolinManager.Contract.Bind(&_PangolinManager.TransactOpts, addr)
}

// ChangeServicePrice is a paid mutator transaction binding the contract method 0x0987b62c.
//
// Solidity: function changeServicePrice(uint256 price) returns()
func (_PangolinManager *PangolinManagerTransactor) ChangeServicePrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _PangolinManager.contract.Transact(opts, "changeServicePrice", price)
}

// ChangeServicePrice is a paid mutator transaction binding the contract method 0x0987b62c.
//
// Solidity: function changeServicePrice(uint256 price) returns()
func (_PangolinManager *PangolinManagerSession) ChangeServicePrice(price *big.Int) (*types.Transaction, error) {
	return _PangolinManager.Contract.ChangeServicePrice(&_PangolinManager.TransactOpts, price)
}

// ChangeServicePrice is a paid mutator transaction binding the contract method 0x0987b62c.
//
// Solidity: function changeServicePrice(uint256 price) returns()
func (_PangolinManager *PangolinManagerTransactorSession) ChangeServicePrice(price *big.Int) (*types.Transaction, error) {
	return _PangolinManager.Contract.ChangeServicePrice(&_PangolinManager.TransactOpts, price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PangolinManager *PangolinManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PangolinManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PangolinManager *PangolinManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PangolinManager.Contract.TransferOwnership(&_PangolinManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PangolinManager *PangolinManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PangolinManager.Contract.TransferOwnership(&_PangolinManager.TransactOpts, newOwner)
}

// Unbind is a paid mutator transaction binding the contract method 0x6e08f245.
//
// Solidity: function unbind(bytes32 addr) returns()
func (_PangolinManager *PangolinManagerTransactor) Unbind(opts *bind.TransactOpts, addr [32]byte) (*types.Transaction, error) {
	return _PangolinManager.contract.Transact(opts, "unbind", addr)
}

// Unbind is a paid mutator transaction binding the contract method 0x6e08f245.
//
// Solidity: function unbind(bytes32 addr) returns()
func (_PangolinManager *PangolinManagerSession) Unbind(addr [32]byte) (*types.Transaction, error) {
	return _PangolinManager.Contract.Unbind(&_PangolinManager.TransactOpts, addr)
}

// Unbind is a paid mutator transaction binding the contract method 0x6e08f245.
//
// Solidity: function unbind(bytes32 addr) returns()
func (_PangolinManager *PangolinManagerTransactorSession) Unbind(addr [32]byte) (*types.Transaction, error) {
	return _PangolinManager.Contract.Unbind(&_PangolinManager.TransactOpts, addr)
}
