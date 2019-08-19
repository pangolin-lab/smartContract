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

// MicroPaySystemABI is the input ABI used to generate the binding from.
const MicroPaySystemABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"TokenDecimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"typ\",\"type\":\"uint8\"},{\"name\":\"targetAddress\",\"type\":\"bytes32\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"name\":\"poolEthAddr\",\"type\":\"address\"}],\"name\":\"BuyPacket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SettlePriceForBuyer\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MinUserCostInToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MinMinerCostInToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"ChangePriceForBuyer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"ta\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MicroPaySystem is an auto generated Go binding around an Ethereum contract.
type MicroPaySystem struct {
	MicroPaySystemCaller     // Read-only binding to the contract
	MicroPaySystemTransactor // Write-only binding to the contract
	MicroPaySystemFilterer   // Log filterer for contract events
}

// MicroPaySystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type MicroPaySystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MicroPaySystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MicroPaySystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MicroPaySystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MicroPaySystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MicroPaySystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MicroPaySystemSession struct {
	Contract     *MicroPaySystem   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MicroPaySystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MicroPaySystemCallerSession struct {
	Contract *MicroPaySystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MicroPaySystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MicroPaySystemTransactorSession struct {
	Contract     *MicroPaySystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MicroPaySystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type MicroPaySystemRaw struct {
	Contract *MicroPaySystem // Generic contract binding to access the raw methods on
}

// MicroPaySystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MicroPaySystemCallerRaw struct {
	Contract *MicroPaySystemCaller // Generic read-only contract binding to access the raw methods on
}

// MicroPaySystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MicroPaySystemTransactorRaw struct {
	Contract *MicroPaySystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMicroPaySystem creates a new instance of MicroPaySystem, bound to a specific deployed contract.
func NewMicroPaySystem(address common.Address, backend bind.ContractBackend) (*MicroPaySystem, error) {
	contract, err := bindMicroPaySystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystem{MicroPaySystemCaller: MicroPaySystemCaller{contract: contract}, MicroPaySystemTransactor: MicroPaySystemTransactor{contract: contract}, MicroPaySystemFilterer: MicroPaySystemFilterer{contract: contract}}, nil
}

// NewMicroPaySystemCaller creates a new read-only instance of MicroPaySystem, bound to a specific deployed contract.
func NewMicroPaySystemCaller(address common.Address, caller bind.ContractCaller) (*MicroPaySystemCaller, error) {
	contract, err := bindMicroPaySystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystemCaller{contract: contract}, nil
}

// NewMicroPaySystemTransactor creates a new write-only instance of MicroPaySystem, bound to a specific deployed contract.
func NewMicroPaySystemTransactor(address common.Address, transactor bind.ContractTransactor) (*MicroPaySystemTransactor, error) {
	contract, err := bindMicroPaySystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystemTransactor{contract: contract}, nil
}

// NewMicroPaySystemFilterer creates a new log filterer instance of MicroPaySystem, bound to a specific deployed contract.
func NewMicroPaySystemFilterer(address common.Address, filterer bind.ContractFilterer) (*MicroPaySystemFilterer, error) {
	contract, err := bindMicroPaySystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MicroPaySystemFilterer{contract: contract}, nil
}

// bindMicroPaySystem binds a generic wrapper to an already deployed contract.
func bindMicroPaySystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MicroPaySystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MicroPaySystem *MicroPaySystemRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MicroPaySystem.Contract.MicroPaySystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MicroPaySystem *MicroPaySystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.MicroPaySystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MicroPaySystem *MicroPaySystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.MicroPaySystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MicroPaySystem *MicroPaySystemCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MicroPaySystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MicroPaySystem *MicroPaySystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MicroPaySystem *MicroPaySystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.contract.Transact(opts, method, params...)
}

// MinMinerCostInToken is a free data retrieval call binding the contract method 0x4f5fa00b.
//
// Solidity: function MinMinerCostInToken() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCaller) MinMinerCostInToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "MinMinerCostInToken")
	return *ret0, err
}

// MinMinerCostInToken is a free data retrieval call binding the contract method 0x4f5fa00b.
//
// Solidity: function MinMinerCostInToken() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemSession) MinMinerCostInToken() (*big.Int, error) {
	return _MicroPaySystem.Contract.MinMinerCostInToken(&_MicroPaySystem.CallOpts)
}

// MinMinerCostInToken is a free data retrieval call binding the contract method 0x4f5fa00b.
//
// Solidity: function MinMinerCostInToken() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCallerSession) MinMinerCostInToken() (*big.Int, error) {
	return _MicroPaySystem.Contract.MinMinerCostInToken(&_MicroPaySystem.CallOpts)
}

// MinUserCostInToken is a free data retrieval call binding the contract method 0x3a335b15.
//
// Solidity: function MinUserCostInToken() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCaller) MinUserCostInToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "MinUserCostInToken")
	return *ret0, err
}

// MinUserCostInToken is a free data retrieval call binding the contract method 0x3a335b15.
//
// Solidity: function MinUserCostInToken() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemSession) MinUserCostInToken() (*big.Int, error) {
	return _MicroPaySystem.Contract.MinUserCostInToken(&_MicroPaySystem.CallOpts)
}

// MinUserCostInToken is a free data retrieval call binding the contract method 0x3a335b15.
//
// Solidity: function MinUserCostInToken() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCallerSession) MinUserCostInToken() (*big.Int, error) {
	return _MicroPaySystem.Contract.MinUserCostInToken(&_MicroPaySystem.CallOpts)
}

// SettlePriceForBuyer is a free data retrieval call binding the contract method 0x21bf56ed.
//
// Solidity: function SettlePriceForBuyer() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCaller) SettlePriceForBuyer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "SettlePriceForBuyer")
	return *ret0, err
}

// SettlePriceForBuyer is a free data retrieval call binding the contract method 0x21bf56ed.
//
// Solidity: function SettlePriceForBuyer() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemSession) SettlePriceForBuyer() (*big.Int, error) {
	return _MicroPaySystem.Contract.SettlePriceForBuyer(&_MicroPaySystem.CallOpts)
}

// SettlePriceForBuyer is a free data retrieval call binding the contract method 0x21bf56ed.
//
// Solidity: function SettlePriceForBuyer() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCallerSession) SettlePriceForBuyer() (*big.Int, error) {
	return _MicroPaySystem.Contract.SettlePriceForBuyer(&_MicroPaySystem.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x0910d1b0.
//
// Solidity: function TokenDecimals() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCaller) TokenDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "TokenDecimals")
	return *ret0, err
}

// TokenDecimals is a free data retrieval call binding the contract method 0x0910d1b0.
//
// Solidity: function TokenDecimals() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemSession) TokenDecimals() (*big.Int, error) {
	return _MicroPaySystem.Contract.TokenDecimals(&_MicroPaySystem.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x0910d1b0.
//
// Solidity: function TokenDecimals() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCallerSession) TokenDecimals() (*big.Int, error) {
	return _MicroPaySystem.Contract.TokenDecimals(&_MicroPaySystem.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MicroPaySystem *MicroPaySystemCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MicroPaySystem *MicroPaySystemSession) Owner() (common.Address, error) {
	return _MicroPaySystem.Contract.Owner(&_MicroPaySystem.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MicroPaySystem *MicroPaySystemCallerSession) Owner() (common.Address, error) {
	return _MicroPaySystem.Contract.Owner(&_MicroPaySystem.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MicroPaySystem *MicroPaySystemCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MicroPaySystem *MicroPaySystemSession) Token() (common.Address, error) {
	return _MicroPaySystem.Contract.Token(&_MicroPaySystem.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_MicroPaySystem *MicroPaySystemCallerSession) Token() (common.Address, error) {
	return _MicroPaySystem.Contract.Token(&_MicroPaySystem.CallOpts)
}

// BuyPacket is a paid mutator transaction binding the contract method 0x17e58f4f.
//
// Solidity: function BuyPacket(uint8 typ, bytes32 targetAddress, uint256 tokenAmount, address poolEthAddr) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) BuyPacket(opts *bind.TransactOpts, typ uint8, targetAddress [32]byte, tokenAmount *big.Int, poolEthAddr common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "BuyPacket", typ, targetAddress, tokenAmount, poolEthAddr)
}

// BuyPacket is a paid mutator transaction binding the contract method 0x17e58f4f.
//
// Solidity: function BuyPacket(uint8 typ, bytes32 targetAddress, uint256 tokenAmount, address poolEthAddr) returns()
func (_MicroPaySystem *MicroPaySystemSession) BuyPacket(typ uint8, targetAddress [32]byte, tokenAmount *big.Int, poolEthAddr common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.BuyPacket(&_MicroPaySystem.TransactOpts, typ, targetAddress, tokenAmount, poolEthAddr)
}

// BuyPacket is a paid mutator transaction binding the contract method 0x17e58f4f.
//
// Solidity: function BuyPacket(uint8 typ, bytes32 targetAddress, uint256 tokenAmount, address poolEthAddr) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) BuyPacket(typ uint8, targetAddress [32]byte, tokenAmount *big.Int, poolEthAddr common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.BuyPacket(&_MicroPaySystem.TransactOpts, typ, targetAddress, tokenAmount, poolEthAddr)
}

// ChangePriceForBuyer is a paid mutator transaction binding the contract method 0x6b8d514f.
//
// Solidity: function ChangePriceForBuyer(uint256 newPrice) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) ChangePriceForBuyer(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "ChangePriceForBuyer", newPrice)
}

// ChangePriceForBuyer is a paid mutator transaction binding the contract method 0x6b8d514f.
//
// Solidity: function ChangePriceForBuyer(uint256 newPrice) returns()
func (_MicroPaySystem *MicroPaySystemSession) ChangePriceForBuyer(newPrice *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangePriceForBuyer(&_MicroPaySystem.TransactOpts, newPrice)
}

// ChangePriceForBuyer is a paid mutator transaction binding the contract method 0x6b8d514f.
//
// Solidity: function ChangePriceForBuyer(uint256 newPrice) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) ChangePriceForBuyer(newPrice *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangePriceForBuyer(&_MicroPaySystem.TransactOpts, newPrice)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MicroPaySystem *MicroPaySystemSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.TransferOwnership(&_MicroPaySystem.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.TransferOwnership(&_MicroPaySystem.TransactOpts, newOwner)
}
