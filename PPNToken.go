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

// PPNTokenABI is the input ABI used to generate the binding from.
const PPNTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDeccimal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FreezeTokenHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// PPNToken is an auto generated Go binding around an Ethereum contract.
type PPNToken struct {
	PPNTokenCaller     // Read-only binding to the contract
	PPNTokenTransactor // Write-only binding to the contract
	PPNTokenFilterer   // Log filterer for contract events
}

// PPNTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PPNTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PPNTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PPNTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PPNTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PPNTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PPNTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PPNTokenSession struct {
	Contract     *PPNToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PPNTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PPNTokenCallerSession struct {
	Contract *PPNTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PPNTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PPNTokenTransactorSession struct {
	Contract     *PPNTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PPNTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PPNTokenRaw struct {
	Contract *PPNToken // Generic contract binding to access the raw methods on
}

// PPNTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PPNTokenCallerRaw struct {
	Contract *PPNTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PPNTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PPNTokenTransactorRaw struct {
	Contract *PPNTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPPNToken creates a new instance of PPNToken, bound to a specific deployed contract.
func NewPPNToken(address common.Address, backend bind.ContractBackend) (*PPNToken, error) {
	contract, err := bindPPNToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PPNToken{PPNTokenCaller: PPNTokenCaller{contract: contract}, PPNTokenTransactor: PPNTokenTransactor{contract: contract}, PPNTokenFilterer: PPNTokenFilterer{contract: contract}}, nil
}

// NewPPNTokenCaller creates a new read-only instance of PPNToken, bound to a specific deployed contract.
func NewPPNTokenCaller(address common.Address, caller bind.ContractCaller) (*PPNTokenCaller, error) {
	contract, err := bindPPNToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PPNTokenCaller{contract: contract}, nil
}

// NewPPNTokenTransactor creates a new write-only instance of PPNToken, bound to a specific deployed contract.
func NewPPNTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PPNTokenTransactor, error) {
	contract, err := bindPPNToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PPNTokenTransactor{contract: contract}, nil
}

// NewPPNTokenFilterer creates a new log filterer instance of PPNToken, bound to a specific deployed contract.
func NewPPNTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PPNTokenFilterer, error) {
	contract, err := bindPPNToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PPNTokenFilterer{contract: contract}, nil
}

// bindPPNToken binds a generic wrapper to an already deployed contract.
func bindPPNToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PPNTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PPNToken *PPNTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PPNToken.Contract.PPNTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PPNToken *PPNTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PPNToken.Contract.PPNTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PPNToken *PPNTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PPNToken.Contract.PPNTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PPNToken *PPNTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PPNToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PPNToken *PPNTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PPNToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PPNToken *PPNTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PPNToken.Contract.contract.Transact(opts, method, params...)
}

// FreezeTokenHolder is a free data retrieval call binding the contract method 0xedde18e0.
//
// Solidity: function FreezeTokenHolder() constant returns(address)
func (_PPNToken *PPNTokenCaller) FreezeTokenHolder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PPNToken.contract.Call(opts, out, "FreezeTokenHolder")
	return *ret0, err
}

// FreezeTokenHolder is a free data retrieval call binding the contract method 0xedde18e0.
//
// Solidity: function FreezeTokenHolder() constant returns(address)
func (_PPNToken *PPNTokenSession) FreezeTokenHolder() (common.Address, error) {
	return _PPNToken.Contract.FreezeTokenHolder(&_PPNToken.CallOpts)
}

// FreezeTokenHolder is a free data retrieval call binding the contract method 0xedde18e0.
//
// Solidity: function FreezeTokenHolder() constant returns(address)
func (_PPNToken *PPNTokenCallerSession) FreezeTokenHolder() (common.Address, error) {
	return _PPNToken.Contract.FreezeTokenHolder(&_PPNToken.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_PPNToken *PPNTokenCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PPNToken.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_PPNToken *PPNTokenSession) INITIALSUPPLY() (*big.Int, error) {
	return _PPNToken.Contract.INITIALSUPPLY(&_PPNToken.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_PPNToken *PPNTokenCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _PPNToken.Contract.INITIALSUPPLY(&_PPNToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_PPNToken *PPNTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PPNToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_PPNToken *PPNTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PPNToken.Contract.Allowance(&_PPNToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_PPNToken *PPNTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PPNToken.Contract.Allowance(&_PPNToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_PPNToken *PPNTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PPNToken.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_PPNToken *PPNTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PPNToken.Contract.BalanceOf(&_PPNToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_PPNToken *PPNTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PPNToken.Contract.BalanceOf(&_PPNToken.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PPNToken *PPNTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PPNToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PPNToken *PPNTokenSession) Decimals() (uint8, error) {
	return _PPNToken.Contract.Decimals(&_PPNToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PPNToken *PPNTokenCallerSession) Decimals() (uint8, error) {
	return _PPNToken.Contract.Decimals(&_PPNToken.CallOpts)
}

// GetDeccimal is a free data retrieval call binding the contract method 0xb9a21589.
//
// Solidity: function getDeccimal() constant returns(uint256)
func (_PPNToken *PPNTokenCaller) GetDeccimal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PPNToken.contract.Call(opts, out, "getDeccimal")
	return *ret0, err
}

// GetDeccimal is a free data retrieval call binding the contract method 0xb9a21589.
//
// Solidity: function getDeccimal() constant returns(uint256)
func (_PPNToken *PPNTokenSession) GetDeccimal() (*big.Int, error) {
	return _PPNToken.Contract.GetDeccimal(&_PPNToken.CallOpts)
}

// GetDeccimal is a free data retrieval call binding the contract method 0xb9a21589.
//
// Solidity: function getDeccimal() constant returns(uint256)
func (_PPNToken *PPNTokenCallerSession) GetDeccimal() (*big.Int, error) {
	return _PPNToken.Contract.GetDeccimal(&_PPNToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PPNToken *PPNTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PPNToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PPNToken *PPNTokenSession) Name() (string, error) {
	return _PPNToken.Contract.Name(&_PPNToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PPNToken *PPNTokenCallerSession) Name() (string, error) {
	return _PPNToken.Contract.Name(&_PPNToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PPNToken *PPNTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PPNToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PPNToken *PPNTokenSession) Symbol() (string, error) {
	return _PPNToken.Contract.Symbol(&_PPNToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PPNToken *PPNTokenCallerSession) Symbol() (string, error) {
	return _PPNToken.Contract.Symbol(&_PPNToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PPNToken *PPNTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PPNToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PPNToken *PPNTokenSession) TotalSupply() (*big.Int, error) {
	return _PPNToken.Contract.TotalSupply(&_PPNToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PPNToken *PPNTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _PPNToken.Contract.TotalSupply(&_PPNToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PPNToken *PPNTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PPNToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PPNToken *PPNTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PPNToken.Contract.Approve(&_PPNToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PPNToken *PPNTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PPNToken.Contract.Approve(&_PPNToken.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PPNToken *PPNTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PPNToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PPNToken *PPNTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PPNToken.Contract.DecreaseAllowance(&_PPNToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PPNToken *PPNTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PPNToken.Contract.DecreaseAllowance(&_PPNToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PPNToken *PPNTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PPNToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PPNToken *PPNTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PPNToken.Contract.IncreaseAllowance(&_PPNToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PPNToken *PPNTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PPNToken.Contract.IncreaseAllowance(&_PPNToken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PPNToken *PPNTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PPNToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PPNToken *PPNTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PPNToken.Contract.Transfer(&_PPNToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PPNToken *PPNTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PPNToken.Contract.Transfer(&_PPNToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PPNToken *PPNTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PPNToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PPNToken *PPNTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PPNToken.Contract.TransferFrom(&_PPNToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PPNToken *PPNTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PPNToken.Contract.TransferFrom(&_PPNToken.TransactOpts, from, to, value)
}

// PPNTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PPNToken contract.
type PPNTokenApprovalIterator struct {
	Event *PPNTokenApproval // Event containing the contract specifics and raw log

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
func (it *PPNTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PPNTokenApproval)
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
		it.Event = new(PPNTokenApproval)
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
func (it *PPNTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PPNTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PPNTokenApproval represents a Approval event raised by the PPNToken contract.
type PPNTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PPNToken *PPNTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PPNTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PPNToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PPNTokenApprovalIterator{contract: _PPNToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PPNToken *PPNTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PPNTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PPNToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PPNTokenApproval)
				if err := _PPNToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PPNToken *PPNTokenFilterer) ParseApproval(log types.Log) (*PPNTokenApproval, error) {
	event := new(PPNTokenApproval)
	if err := _PPNToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PPNTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PPNToken contract.
type PPNTokenTransferIterator struct {
	Event *PPNTokenTransfer // Event containing the contract specifics and raw log

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
func (it *PPNTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PPNTokenTransfer)
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
		it.Event = new(PPNTokenTransfer)
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
func (it *PPNTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PPNTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PPNTokenTransfer represents a Transfer event raised by the PPNToken contract.
type PPNTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PPNToken *PPNTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PPNTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PPNToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PPNTokenTransferIterator{contract: _PPNToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PPNToken *PPNTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PPNTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PPNToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PPNTokenTransfer)
				if err := _PPNToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PPNToken *PPNTokenFilterer) ParseTransfer(log types.Log) (*PPNTokenTransfer, error) {
	event := new(PPNTokenTransfer)
	if err := _PPNToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
