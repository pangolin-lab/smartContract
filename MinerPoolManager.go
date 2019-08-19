// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

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

// MinerPoolManagerABI is the input ABI used to generate the binding from.
const MinerPoolManagerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"TokenDecimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guaranteeAmount\",\"type\":\"uint256\"},{\"name\":\"desc\",\"type\":\"string\"}],\"name\":\"RegAsMinerPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newDesc\",\"type\":\"string\"}],\"name\":\"ChangeDesc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TopPoolList\",\"outputs\":[{\"name\":\"poolEthAddr\",\"type\":\"address\"},{\"name\":\"guaranteedNo\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenNo\",\"type\":\"uint256\"}],\"name\":\"RefundGuarantToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenNo\",\"type\":\"uint256\"}],\"name\":\"AddGuarantToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MinPoolCostInToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"ta\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MinerPoolManager is an auto generated Go binding around an Ethereum contract.
type MinerPoolManager struct {
	MinerPoolManagerCaller     // Read-only binding to the contract
	MinerPoolManagerTransactor // Write-only binding to the contract
	MinerPoolManagerFilterer   // Log filterer for contract events
}

// MinerPoolManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinerPoolManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerPoolManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinerPoolManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerPoolManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinerPoolManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerPoolManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinerPoolManagerSession struct {
	Contract     *MinerPoolManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinerPoolManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinerPoolManagerCallerSession struct {
	Contract *MinerPoolManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MinerPoolManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinerPoolManagerTransactorSession struct {
	Contract     *MinerPoolManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MinerPoolManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinerPoolManagerRaw struct {
	Contract *MinerPoolManager // Generic contract binding to access the raw methods on
}

// MinerPoolManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinerPoolManagerCallerRaw struct {
	Contract *MinerPoolManagerCaller // Generic read-only contract binding to access the raw methods on
}

// MinerPoolManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinerPoolManagerTransactorRaw struct {
	Contract *MinerPoolManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinerPoolManager creates a new instance of MinerPoolManager, bound to a specific deployed contract.
func NewMinerPoolManager(address common.Address, backend bind.ContractBackend) (*MinerPoolManager, error) {
	contract, err := bindMinerPoolManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinerPoolManager{MinerPoolManagerCaller: MinerPoolManagerCaller{contract: contract}, MinerPoolManagerTransactor: MinerPoolManagerTransactor{contract: contract}, MinerPoolManagerFilterer: MinerPoolManagerFilterer{contract: contract}}, nil
}

// NewMinerPoolManagerCaller creates a new read-only instance of MinerPoolManager, bound to a specific deployed contract.
func NewMinerPoolManagerCaller(address common.Address, caller bind.ContractCaller) (*MinerPoolManagerCaller, error) {
	contract, err := bindMinerPoolManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinerPoolManagerCaller{contract: contract}, nil
}

// NewMinerPoolManagerTransactor creates a new write-only instance of MinerPoolManager, bound to a specific deployed contract.
func NewMinerPoolManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MinerPoolManagerTransactor, error) {
	contract, err := bindMinerPoolManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinerPoolManagerTransactor{contract: contract}, nil
}

// NewMinerPoolManagerFilterer creates a new log filterer instance of MinerPoolManager, bound to a specific deployed contract.
func NewMinerPoolManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MinerPoolManagerFilterer, error) {
	contract, err := bindMinerPoolManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinerPoolManagerFilterer{contract: contract}, nil
}

// bindMinerPoolManager binds a generic wrapper to an already deployed contract.
func bindMinerPoolManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinerPoolManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinerPoolManager *MinerPoolManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinerPoolManager.Contract.MinerPoolManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinerPoolManager *MinerPoolManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.MinerPoolManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinerPoolManager *MinerPoolManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.MinerPoolManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinerPoolManager *MinerPoolManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinerPoolManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinerPoolManager *MinerPoolManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinerPoolManager *MinerPoolManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.contract.Transact(opts, method, params...)
}

// MinPoolCostInToken is a free data retrieval call binding the contract method 0xf6cb3dfc.
//
// Solidity: function MinPoolCostInToken() constant returns(uint256)
func (_MinerPoolManager *MinerPoolManagerCaller) MinPoolCostInToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MinerPoolManager.contract.Call(opts, out, "MinPoolCostInToken")
	return *ret0, err
}

// MinPoolCostInToken is a free data retrieval call binding the contract method 0xf6cb3dfc.
//
// Solidity: function MinPoolCostInToken() constant returns(uint256)
func (_MinerPoolManager *MinerPoolManagerSession) MinPoolCostInToken() (*big.Int, error) {
	return _MinerPoolManager.Contract.MinPoolCostInToken(&_MinerPoolManager.CallOpts)
}

// MinPoolCostInToken is a free data retrieval call binding the contract method 0xf6cb3dfc.
//
// Solidity: function MinPoolCostInToken() constant returns(uint256)
func (_MinerPoolManager *MinerPoolManagerCallerSession) MinPoolCostInToken() (*big.Int, error) {
	return _MinerPoolManager.Contract.MinPoolCostInToken(&_MinerPoolManager.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x0910d1b0.
//
// Solidity: function TokenDecimals() constant returns(uint256)
func (_MinerPoolManager *MinerPoolManagerCaller) TokenDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MinerPoolManager.contract.Call(opts, out, "TokenDecimals")
	return *ret0, err
}

// TokenDecimals is a free data retrieval call binding the contract method 0x0910d1b0.
//
// Solidity: function TokenDecimals() constant returns(uint256)
func (_MinerPoolManager *MinerPoolManagerSession) TokenDecimals() (*big.Int, error) {
	return _MinerPoolManager.Contract.TokenDecimals(&_MinerPoolManager.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x0910d1b0.
//
// Solidity: function TokenDecimals() constant returns(uint256)
func (_MinerPoolManager *MinerPoolManagerCallerSession) TokenDecimals() (*big.Int, error) {
	return _MinerPoolManager.Contract.TokenDecimals(&_MinerPoolManager.CallOpts)
}

// TopPoolList is a free data retrieval call binding the contract method 0x50fc104d.
//
// Solidity: function TopPoolList(uint256 ) constant returns(address poolEthAddr, uint256 guaranteedNo)
func (_MinerPoolManager *MinerPoolManagerCaller) TopPoolList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PoolEthAddr  common.Address
	GuaranteedNo *big.Int
}, error) {
	ret := new(struct {
		PoolEthAddr  common.Address
		GuaranteedNo *big.Int
	})
	out := ret
	err := _MinerPoolManager.contract.Call(opts, out, "TopPoolList", arg0)
	return *ret, err
}

// TopPoolList is a free data retrieval call binding the contract method 0x50fc104d.
//
// Solidity: function TopPoolList(uint256 ) constant returns(address poolEthAddr, uint256 guaranteedNo)
func (_MinerPoolManager *MinerPoolManagerSession) TopPoolList(arg0 *big.Int) (struct {
	PoolEthAddr  common.Address
	GuaranteedNo *big.Int
}, error) {
	return _MinerPoolManager.Contract.TopPoolList(&_MinerPoolManager.CallOpts, arg0)
}

// TopPoolList is a free data retrieval call binding the contract method 0x50fc104d.
//
// Solidity: function TopPoolList(uint256 ) constant returns(address poolEthAddr, uint256 guaranteedNo)
func (_MinerPoolManager *MinerPoolManagerCallerSession) TopPoolList(arg0 *big.Int) (struct {
	PoolEthAddr  common.Address
	GuaranteedNo *big.Int
}, error) {
	return _MinerPoolManager.Contract.TopPoolList(&_MinerPoolManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MinerPoolManager *MinerPoolManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MinerPoolManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MinerPoolManager *MinerPoolManagerSession) Owner() (common.Address, error) {
	return _MinerPoolManager.Contract.Owner(&_MinerPoolManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MinerPoolManager *MinerPoolManagerCallerSession) Owner() (common.Address, error) {
	return _MinerPoolManager.Contract.Owner(&_MinerPoolManager.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MinerPoolManager *MinerPoolManagerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MinerPoolManager.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MinerPoolManager *MinerPoolManagerSession) Token() (common.Address, error) {
	return _MinerPoolManager.Contract.Token(&_MinerPoolManager.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MinerPoolManager *MinerPoolManagerCallerSession) Token() (common.Address, error) {
	return _MinerPoolManager.Contract.Token(&_MinerPoolManager.CallOpts)
}

// AddGuarantToken is a paid mutator transaction binding the contract method 0xf21e5ed6.
//
// Solidity: function AddGuarantToken(uint256 tokenNo) returns()
func (_MinerPoolManager *MinerPoolManagerTransactor) AddGuarantToken(opts *bind.TransactOpts, tokenNo *big.Int) (*types.Transaction, error) {
	return _MinerPoolManager.contract.Transact(opts, "AddGuarantToken", tokenNo)
}

// AddGuarantToken is a paid mutator transaction binding the contract method 0xf21e5ed6.
//
// Solidity: function AddGuarantToken(uint256 tokenNo) returns()
func (_MinerPoolManager *MinerPoolManagerSession) AddGuarantToken(tokenNo *big.Int) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.AddGuarantToken(&_MinerPoolManager.TransactOpts, tokenNo)
}

// AddGuarantToken is a paid mutator transaction binding the contract method 0xf21e5ed6.
//
// Solidity: function AddGuarantToken(uint256 tokenNo) returns()
func (_MinerPoolManager *MinerPoolManagerTransactorSession) AddGuarantToken(tokenNo *big.Int) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.AddGuarantToken(&_MinerPoolManager.TransactOpts, tokenNo)
}

// ChangeDesc is a paid mutator transaction binding the contract method 0x2f9ed7d8.
//
// Solidity: function ChangeDesc(string newDesc) returns()
func (_MinerPoolManager *MinerPoolManagerTransactor) ChangeDesc(opts *bind.TransactOpts, newDesc string) (*types.Transaction, error) {
	return _MinerPoolManager.contract.Transact(opts, "ChangeDesc", newDesc)
}

// ChangeDesc is a paid mutator transaction binding the contract method 0x2f9ed7d8.
//
// Solidity: function ChangeDesc(string newDesc) returns()
func (_MinerPoolManager *MinerPoolManagerSession) ChangeDesc(newDesc string) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.ChangeDesc(&_MinerPoolManager.TransactOpts, newDesc)
}

// ChangeDesc is a paid mutator transaction binding the contract method 0x2f9ed7d8.
//
// Solidity: function ChangeDesc(string newDesc) returns()
func (_MinerPoolManager *MinerPoolManagerTransactorSession) ChangeDesc(newDesc string) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.ChangeDesc(&_MinerPoolManager.TransactOpts, newDesc)
}

// RefundGuarantToken is a paid mutator transaction binding the contract method 0x8d55b9b5.
//
// Solidity: function RefundGuarantToken(uint256 tokenNo) returns()
func (_MinerPoolManager *MinerPoolManagerTransactor) RefundGuarantToken(opts *bind.TransactOpts, tokenNo *big.Int) (*types.Transaction, error) {
	return _MinerPoolManager.contract.Transact(opts, "RefundGuarantToken", tokenNo)
}

// RefundGuarantToken is a paid mutator transaction binding the contract method 0x8d55b9b5.
//
// Solidity: function RefundGuarantToken(uint256 tokenNo) returns()
func (_MinerPoolManager *MinerPoolManagerSession) RefundGuarantToken(tokenNo *big.Int) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.RefundGuarantToken(&_MinerPoolManager.TransactOpts, tokenNo)
}

// RefundGuarantToken is a paid mutator transaction binding the contract method 0x8d55b9b5.
//
// Solidity: function RefundGuarantToken(uint256 tokenNo) returns()
func (_MinerPoolManager *MinerPoolManagerTransactorSession) RefundGuarantToken(tokenNo *big.Int) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.RefundGuarantToken(&_MinerPoolManager.TransactOpts, tokenNo)
}

// RegAsMinerPool is a paid mutator transaction binding the contract method 0x1ce4e7ec.
//
// Solidity: function RegAsMinerPool(uint256 guaranteeAmount, string desc) returns()
func (_MinerPoolManager *MinerPoolManagerTransactor) RegAsMinerPool(opts *bind.TransactOpts, guaranteeAmount *big.Int, desc string) (*types.Transaction, error) {
	return _MinerPoolManager.contract.Transact(opts, "RegAsMinerPool", guaranteeAmount, desc)
}

// RegAsMinerPool is a paid mutator transaction binding the contract method 0x1ce4e7ec.
//
// Solidity: function RegAsMinerPool(uint256 guaranteeAmount, string desc) returns()
func (_MinerPoolManager *MinerPoolManagerSession) RegAsMinerPool(guaranteeAmount *big.Int, desc string) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.RegAsMinerPool(&_MinerPoolManager.TransactOpts, guaranteeAmount, desc)
}

// RegAsMinerPool is a paid mutator transaction binding the contract method 0x1ce4e7ec.
//
// Solidity: function RegAsMinerPool(uint256 guaranteeAmount, string desc) returns()
func (_MinerPoolManager *MinerPoolManagerTransactorSession) RegAsMinerPool(guaranteeAmount *big.Int, desc string) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.RegAsMinerPool(&_MinerPoolManager.TransactOpts, guaranteeAmount, desc)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MinerPoolManager *MinerPoolManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MinerPoolManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MinerPoolManager *MinerPoolManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.TransferOwnership(&_MinerPoolManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MinerPoolManager *MinerPoolManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MinerPoolManager.Contract.TransferOwnership(&_MinerPoolManager.TransactOpts, newOwner)
}
