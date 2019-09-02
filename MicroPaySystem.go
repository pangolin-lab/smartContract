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
const MicroPaySystemABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"TokenDecimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"mainAddr\",\"type\":\"address\"},{\"name\":\"typ\",\"type\":\"uint8\"}],\"name\":\"SetPoolType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PacketPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"daysAfter\",\"type\":\"uint256\"}],\"name\":\"ChangeDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"MinerPools\",\"outputs\":[{\"name\":\"ID\",\"type\":\"uint32\"},{\"name\":\"poolType\",\"type\":\"uint8\"},{\"name\":\"mainAddr\",\"type\":\"address\"},{\"name\":\"payer\",\"type\":\"address\"},{\"name\":\"subAddr\",\"type\":\"bytes32\"},{\"name\":\"guaranteedNo\",\"type\":\"uint256\"},{\"name\":\"shortName\",\"type\":\"string\"},{\"name\":\"detailInfos\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"AllMySubPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetPoolSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetPoolAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"ChangeBandWithPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MinMinerCostInToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"tokenNo\",\"type\":\"uint256\"},{\"name\":\"poolAddr\",\"type\":\"address\"}],\"name\":\"BuyPacket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"MicroPaymentChannels\",\"outputs\":[{\"name\":\"mainAddr\",\"type\":\"address\"},{\"name\":\"payerAddr\",\"type\":\"address\"},{\"name\":\"remindTokens\",\"type\":\"uint256\"},{\"name\":\"remindPackets\",\"type\":\"uint256\"},{\"name\":\"expiration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allSubPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"mainAddr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"desc\",\"type\":\"string\"}],\"name\":\"ChangePoolSettings\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MinerPoolsAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCost\",\"type\":\"uint256\"}],\"name\":\"ChangeMinPoolCost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"TokenBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"gno\",\"type\":\"uint256\"},{\"name\":\"mainAddr\",\"type\":\"address\"},{\"name\":\"subAddr\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"desc\",\"type\":\"string\"}],\"name\":\"RegAsMinerPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Duration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MinPoolCostInToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCost\",\"type\":\"uint256\"}],\"name\":\"ChangeMinMinerCost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"ta\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

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

// AllMySubPools is a free data retrieval call binding the contract method 0x2aec99fd.
//
// Solidity: function AllMySubPools(address userAddress) constant returns(address[])
func (_MicroPaySystem *MicroPaySystemCaller) AllMySubPools(opts *bind.CallOpts, userAddress common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "AllMySubPools", userAddress)
	return *ret0, err
}

// AllMySubPools is a free data retrieval call binding the contract method 0x2aec99fd.
//
// Solidity: function AllMySubPools(address userAddress) constant returns(address[])
func (_MicroPaySystem *MicroPaySystemSession) AllMySubPools(userAddress common.Address) ([]common.Address, error) {
	return _MicroPaySystem.Contract.AllMySubPools(&_MicroPaySystem.CallOpts, userAddress)
}

// AllMySubPools is a free data retrieval call binding the contract method 0x2aec99fd.
//
// Solidity: function AllMySubPools(address userAddress) constant returns(address[])
func (_MicroPaySystem *MicroPaySystemCallerSession) AllMySubPools(userAddress common.Address) ([]common.Address, error) {
	return _MicroPaySystem.Contract.AllMySubPools(&_MicroPaySystem.CallOpts, userAddress)
}

// Duration is a free data retrieval call binding the contract method 0xccb30941.
//
// Solidity: function Duration() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "Duration")
	return *ret0, err
}

// Duration is a free data retrieval call binding the contract method 0xccb30941.
//
// Solidity: function Duration() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemSession) Duration() (*big.Int, error) {
	return _MicroPaySystem.Contract.Duration(&_MicroPaySystem.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0xccb30941.
//
// Solidity: function Duration() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCallerSession) Duration() (*big.Int, error) {
	return _MicroPaySystem.Contract.Duration(&_MicroPaySystem.CallOpts)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0x3e6cc5b2.
//
// Solidity: function GetPoolAddress() constant returns(address[])
func (_MicroPaySystem *MicroPaySystemCaller) GetPoolAddress(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "GetPoolAddress")
	return *ret0, err
}

// GetPoolAddress is a free data retrieval call binding the contract method 0x3e6cc5b2.
//
// Solidity: function GetPoolAddress() constant returns(address[])
func (_MicroPaySystem *MicroPaySystemSession) GetPoolAddress() ([]common.Address, error) {
	return _MicroPaySystem.Contract.GetPoolAddress(&_MicroPaySystem.CallOpts)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0x3e6cc5b2.
//
// Solidity: function GetPoolAddress() constant returns(address[])
func (_MicroPaySystem *MicroPaySystemCallerSession) GetPoolAddress() ([]common.Address, error) {
	return _MicroPaySystem.Contract.GetPoolAddress(&_MicroPaySystem.CallOpts)
}

// GetPoolSize is a free data retrieval call binding the contract method 0x38339ea1.
//
// Solidity: function GetPoolSize() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCaller) GetPoolSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "GetPoolSize")
	return *ret0, err
}

// GetPoolSize is a free data retrieval call binding the contract method 0x38339ea1.
//
// Solidity: function GetPoolSize() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemSession) GetPoolSize() (*big.Int, error) {
	return _MicroPaySystem.Contract.GetPoolSize(&_MicroPaySystem.CallOpts)
}

// GetPoolSize is a free data retrieval call binding the contract method 0x38339ea1.
//
// Solidity: function GetPoolSize() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCallerSession) GetPoolSize() (*big.Int, error) {
	return _MicroPaySystem.Contract.GetPoolSize(&_MicroPaySystem.CallOpts)
}

// MicroPaymentChannels is a free data retrieval call binding the contract method 0x53488ab7.
//
// Solidity: function MicroPaymentChannels(address , address ) constant returns(address mainAddr, address payerAddr, uint256 remindTokens, uint256 remindPackets, uint256 expiration)
func (_MicroPaySystem *MicroPaySystemCaller) MicroPaymentChannels(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	MainAddr      common.Address
	PayerAddr     common.Address
	RemindTokens  *big.Int
	RemindPackets *big.Int
	Expiration    *big.Int
}, error) {
	ret := new(struct {
		MainAddr      common.Address
		PayerAddr     common.Address
		RemindTokens  *big.Int
		RemindPackets *big.Int
		Expiration    *big.Int
	})
	out := ret
	err := _MicroPaySystem.contract.Call(opts, out, "MicroPaymentChannels", arg0, arg1)
	return *ret, err
}

// MicroPaymentChannels is a free data retrieval call binding the contract method 0x53488ab7.
//
// Solidity: function MicroPaymentChannels(address , address ) constant returns(address mainAddr, address payerAddr, uint256 remindTokens, uint256 remindPackets, uint256 expiration)
func (_MicroPaySystem *MicroPaySystemSession) MicroPaymentChannels(arg0 common.Address, arg1 common.Address) (struct {
	MainAddr      common.Address
	PayerAddr     common.Address
	RemindTokens  *big.Int
	RemindPackets *big.Int
	Expiration    *big.Int
}, error) {
	return _MicroPaySystem.Contract.MicroPaymentChannels(&_MicroPaySystem.CallOpts, arg0, arg1)
}

// MicroPaymentChannels is a free data retrieval call binding the contract method 0x53488ab7.
//
// Solidity: function MicroPaymentChannels(address , address ) constant returns(address mainAddr, address payerAddr, uint256 remindTokens, uint256 remindPackets, uint256 expiration)
func (_MicroPaySystem *MicroPaySystemCallerSession) MicroPaymentChannels(arg0 common.Address, arg1 common.Address) (struct {
	MainAddr      common.Address
	PayerAddr     common.Address
	RemindTokens  *big.Int
	RemindPackets *big.Int
	Expiration    *big.Int
}, error) {
	return _MicroPaySystem.Contract.MicroPaymentChannels(&_MicroPaySystem.CallOpts, arg0, arg1)
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

// MinPoolCostInToken is a free data retrieval call binding the contract method 0xf6cb3dfc.
//
// Solidity: function MinPoolCostInToken() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCaller) MinPoolCostInToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "MinPoolCostInToken")
	return *ret0, err
}

// MinPoolCostInToken is a free data retrieval call binding the contract method 0xf6cb3dfc.
//
// Solidity: function MinPoolCostInToken() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemSession) MinPoolCostInToken() (*big.Int, error) {
	return _MicroPaySystem.Contract.MinPoolCostInToken(&_MicroPaySystem.CallOpts)
}

// MinPoolCostInToken is a free data retrieval call binding the contract method 0xf6cb3dfc.
//
// Solidity: function MinPoolCostInToken() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCallerSession) MinPoolCostInToken() (*big.Int, error) {
	return _MicroPaySystem.Contract.MinPoolCostInToken(&_MicroPaySystem.CallOpts)
}

// MinerPools is a free data retrieval call binding the contract method 0x291c3a5c.
//
// Solidity: function MinerPools(address ) constant returns(uint32 ID, uint8 poolType, address mainAddr, address payer, bytes32 subAddr, uint256 guaranteedNo, string shortName, string detailInfos)
func (_MicroPaySystem *MicroPaySystemCaller) MinerPools(opts *bind.CallOpts, arg0 common.Address) (struct {
	ID           uint32
	PoolType     uint8
	MainAddr     common.Address
	Payer        common.Address
	SubAddr      [32]byte
	GuaranteedNo *big.Int
	ShortName    string
	DetailInfos  string
}, error) {
	ret := new(struct {
		ID           uint32
		PoolType     uint8
		MainAddr     common.Address
		Payer        common.Address
		SubAddr      [32]byte
		GuaranteedNo *big.Int
		ShortName    string
		DetailInfos  string
	})
	out := ret
	err := _MicroPaySystem.contract.Call(opts, out, "MinerPools", arg0)
	return *ret, err
}

// MinerPools is a free data retrieval call binding the contract method 0x291c3a5c.
//
// Solidity: function MinerPools(address ) constant returns(uint32 ID, uint8 poolType, address mainAddr, address payer, bytes32 subAddr, uint256 guaranteedNo, string shortName, string detailInfos)
func (_MicroPaySystem *MicroPaySystemSession) MinerPools(arg0 common.Address) (struct {
	ID           uint32
	PoolType     uint8
	MainAddr     common.Address
	Payer        common.Address
	SubAddr      [32]byte
	GuaranteedNo *big.Int
	ShortName    string
	DetailInfos  string
}, error) {
	return _MicroPaySystem.Contract.MinerPools(&_MicroPaySystem.CallOpts, arg0)
}

// MinerPools is a free data retrieval call binding the contract method 0x291c3a5c.
//
// Solidity: function MinerPools(address ) constant returns(uint32 ID, uint8 poolType, address mainAddr, address payer, bytes32 subAddr, uint256 guaranteedNo, string shortName, string detailInfos)
func (_MicroPaySystem *MicroPaySystemCallerSession) MinerPools(arg0 common.Address) (struct {
	ID           uint32
	PoolType     uint8
	MainAddr     common.Address
	Payer        common.Address
	SubAddr      [32]byte
	GuaranteedNo *big.Int
	ShortName    string
	DetailInfos  string
}, error) {
	return _MicroPaySystem.Contract.MinerPools(&_MicroPaySystem.CallOpts, arg0)
}

// MinerPoolsAddresses is a free data retrieval call binding the contract method 0xb1d9c348.
//
// Solidity: function MinerPoolsAddresses(uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemCaller) MinerPoolsAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "MinerPoolsAddresses", arg0)
	return *ret0, err
}

// MinerPoolsAddresses is a free data retrieval call binding the contract method 0xb1d9c348.
//
// Solidity: function MinerPoolsAddresses(uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemSession) MinerPoolsAddresses(arg0 *big.Int) (common.Address, error) {
	return _MicroPaySystem.Contract.MinerPoolsAddresses(&_MicroPaySystem.CallOpts, arg0)
}

// MinerPoolsAddresses is a free data retrieval call binding the contract method 0xb1d9c348.
//
// Solidity: function MinerPoolsAddresses(uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemCallerSession) MinerPoolsAddresses(arg0 *big.Int) (common.Address, error) {
	return _MicroPaySystem.Contract.MinerPoolsAddresses(&_MicroPaySystem.CallOpts, arg0)
}

// PacketPrice is a free data retrieval call binding the contract method 0x19f2170f.
//
// Solidity: function PacketPrice() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCaller) PacketPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "PacketPrice")
	return *ret0, err
}

// PacketPrice is a free data retrieval call binding the contract method 0x19f2170f.
//
// Solidity: function PacketPrice() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemSession) PacketPrice() (*big.Int, error) {
	return _MicroPaySystem.Contract.PacketPrice(&_MicroPaySystem.CallOpts)
}

// PacketPrice is a free data retrieval call binding the contract method 0x19f2170f.
//
// Solidity: function PacketPrice() constant returns(uint256)
func (_MicroPaySystem *MicroPaySystemCallerSession) PacketPrice() (*big.Int, error) {
	return _MicroPaySystem.Contract.PacketPrice(&_MicroPaySystem.CallOpts)
}

// TokenBalance is a free data retrieval call binding the contract method 0xb6e2b395.
//
// Solidity: function TokenBalance(address userAddress) constant returns(uint256, uint256)
func (_MicroPaySystem *MicroPaySystemCaller) TokenBalance(opts *bind.CallOpts, userAddress common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MicroPaySystem.contract.Call(opts, out, "TokenBalance", userAddress)
	return *ret0, *ret1, err
}

// TokenBalance is a free data retrieval call binding the contract method 0xb6e2b395.
//
// Solidity: function TokenBalance(address userAddress) constant returns(uint256, uint256)
func (_MicroPaySystem *MicroPaySystemSession) TokenBalance(userAddress common.Address) (*big.Int, *big.Int, error) {
	return _MicroPaySystem.Contract.TokenBalance(&_MicroPaySystem.CallOpts, userAddress)
}

// TokenBalance is a free data retrieval call binding the contract method 0xb6e2b395.
//
// Solidity: function TokenBalance(address userAddress) constant returns(uint256, uint256)
func (_MicroPaySystem *MicroPaySystemCallerSession) TokenBalance(userAddress common.Address) (*big.Int, *big.Int, error) {
	return _MicroPaySystem.Contract.TokenBalance(&_MicroPaySystem.CallOpts, userAddress)
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

// AllSubPools is a free data retrieval call binding the contract method 0x5bbddc0a.
//
// Solidity: function allSubPools(address , uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemCaller) AllSubPools(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MicroPaySystem.contract.Call(opts, out, "allSubPools", arg0, arg1)
	return *ret0, err
}

// AllSubPools is a free data retrieval call binding the contract method 0x5bbddc0a.
//
// Solidity: function allSubPools(address , uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemSession) AllSubPools(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _MicroPaySystem.Contract.AllSubPools(&_MicroPaySystem.CallOpts, arg0, arg1)
}

// AllSubPools is a free data retrieval call binding the contract method 0x5bbddc0a.
//
// Solidity: function allSubPools(address , uint256 ) constant returns(address)
func (_MicroPaySystem *MicroPaySystemCallerSession) AllSubPools(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _MicroPaySystem.Contract.AllSubPools(&_MicroPaySystem.CallOpts, arg0, arg1)
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

// BuyPacket is a paid mutator transaction binding the contract method 0x51f922ac.
//
// Solidity: function BuyPacket(address user, uint256 tokenNo, address poolAddr) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) BuyPacket(opts *bind.TransactOpts, user common.Address, tokenNo *big.Int, poolAddr common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "BuyPacket", user, tokenNo, poolAddr)
}

// BuyPacket is a paid mutator transaction binding the contract method 0x51f922ac.
//
// Solidity: function BuyPacket(address user, uint256 tokenNo, address poolAddr) returns()
func (_MicroPaySystem *MicroPaySystemSession) BuyPacket(user common.Address, tokenNo *big.Int, poolAddr common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.BuyPacket(&_MicroPaySystem.TransactOpts, user, tokenNo, poolAddr)
}

// BuyPacket is a paid mutator transaction binding the contract method 0x51f922ac.
//
// Solidity: function BuyPacket(address user, uint256 tokenNo, address poolAddr) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) BuyPacket(user common.Address, tokenNo *big.Int, poolAddr common.Address) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.BuyPacket(&_MicroPaySystem.TransactOpts, user, tokenNo, poolAddr)
}

// ChangeBandWithPrice is a paid mutator transaction binding the contract method 0x4435341d.
//
// Solidity: function ChangeBandWithPrice(uint256 newPrice) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) ChangeBandWithPrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "ChangeBandWithPrice", newPrice)
}

// ChangeBandWithPrice is a paid mutator transaction binding the contract method 0x4435341d.
//
// Solidity: function ChangeBandWithPrice(uint256 newPrice) returns()
func (_MicroPaySystem *MicroPaySystemSession) ChangeBandWithPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangeBandWithPrice(&_MicroPaySystem.TransactOpts, newPrice)
}

// ChangeBandWithPrice is a paid mutator transaction binding the contract method 0x4435341d.
//
// Solidity: function ChangeBandWithPrice(uint256 newPrice) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) ChangeBandWithPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangeBandWithPrice(&_MicroPaySystem.TransactOpts, newPrice)
}

// ChangeDuration is a paid mutator transaction binding the contract method 0x1deacdd1.
//
// Solidity: function ChangeDuration(uint256 daysAfter) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) ChangeDuration(opts *bind.TransactOpts, daysAfter *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "ChangeDuration", daysAfter)
}

// ChangeDuration is a paid mutator transaction binding the contract method 0x1deacdd1.
//
// Solidity: function ChangeDuration(uint256 daysAfter) returns()
func (_MicroPaySystem *MicroPaySystemSession) ChangeDuration(daysAfter *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangeDuration(&_MicroPaySystem.TransactOpts, daysAfter)
}

// ChangeDuration is a paid mutator transaction binding the contract method 0x1deacdd1.
//
// Solidity: function ChangeDuration(uint256 daysAfter) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) ChangeDuration(daysAfter *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangeDuration(&_MicroPaySystem.TransactOpts, daysAfter)
}

// ChangeMinMinerCost is a paid mutator transaction binding the contract method 0xff13c103.
//
// Solidity: function ChangeMinMinerCost(uint256 newCost) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) ChangeMinMinerCost(opts *bind.TransactOpts, newCost *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "ChangeMinMinerCost", newCost)
}

// ChangeMinMinerCost is a paid mutator transaction binding the contract method 0xff13c103.
//
// Solidity: function ChangeMinMinerCost(uint256 newCost) returns()
func (_MicroPaySystem *MicroPaySystemSession) ChangeMinMinerCost(newCost *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangeMinMinerCost(&_MicroPaySystem.TransactOpts, newCost)
}

// ChangeMinMinerCost is a paid mutator transaction binding the contract method 0xff13c103.
//
// Solidity: function ChangeMinMinerCost(uint256 newCost) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) ChangeMinMinerCost(newCost *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangeMinMinerCost(&_MicroPaySystem.TransactOpts, newCost)
}

// ChangeMinPoolCost is a paid mutator transaction binding the contract method 0xb583e623.
//
// Solidity: function ChangeMinPoolCost(uint256 newCost) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) ChangeMinPoolCost(opts *bind.TransactOpts, newCost *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "ChangeMinPoolCost", newCost)
}

// ChangeMinPoolCost is a paid mutator transaction binding the contract method 0xb583e623.
//
// Solidity: function ChangeMinPoolCost(uint256 newCost) returns()
func (_MicroPaySystem *MicroPaySystemSession) ChangeMinPoolCost(newCost *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangeMinPoolCost(&_MicroPaySystem.TransactOpts, newCost)
}

// ChangeMinPoolCost is a paid mutator transaction binding the contract method 0xb583e623.
//
// Solidity: function ChangeMinPoolCost(uint256 newCost) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) ChangeMinPoolCost(newCost *big.Int) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangeMinPoolCost(&_MicroPaySystem.TransactOpts, newCost)
}

// ChangePoolSettings is a paid mutator transaction binding the contract method 0xa61bd1ed.
//
// Solidity: function ChangePoolSettings(address mainAddr, string name, string desc) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) ChangePoolSettings(opts *bind.TransactOpts, mainAddr common.Address, name string, desc string) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "ChangePoolSettings", mainAddr, name, desc)
}

// ChangePoolSettings is a paid mutator transaction binding the contract method 0xa61bd1ed.
//
// Solidity: function ChangePoolSettings(address mainAddr, string name, string desc) returns()
func (_MicroPaySystem *MicroPaySystemSession) ChangePoolSettings(mainAddr common.Address, name string, desc string) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangePoolSettings(&_MicroPaySystem.TransactOpts, mainAddr, name, desc)
}

// ChangePoolSettings is a paid mutator transaction binding the contract method 0xa61bd1ed.
//
// Solidity: function ChangePoolSettings(address mainAddr, string name, string desc) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) ChangePoolSettings(mainAddr common.Address, name string, desc string) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.ChangePoolSettings(&_MicroPaySystem.TransactOpts, mainAddr, name, desc)
}

// RegAsMinerPool is a paid mutator transaction binding the contract method 0xb8d8f9ef.
//
// Solidity: function RegAsMinerPool(uint256 gno, address mainAddr, bytes32 subAddr, string name, string desc) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) RegAsMinerPool(opts *bind.TransactOpts, gno *big.Int, mainAddr common.Address, subAddr [32]byte, name string, desc string) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "RegAsMinerPool", gno, mainAddr, subAddr, name, desc)
}

// RegAsMinerPool is a paid mutator transaction binding the contract method 0xb8d8f9ef.
//
// Solidity: function RegAsMinerPool(uint256 gno, address mainAddr, bytes32 subAddr, string name, string desc) returns()
func (_MicroPaySystem *MicroPaySystemSession) RegAsMinerPool(gno *big.Int, mainAddr common.Address, subAddr [32]byte, name string, desc string) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.RegAsMinerPool(&_MicroPaySystem.TransactOpts, gno, mainAddr, subAddr, name, desc)
}

// RegAsMinerPool is a paid mutator transaction binding the contract method 0xb8d8f9ef.
//
// Solidity: function RegAsMinerPool(uint256 gno, address mainAddr, bytes32 subAddr, string name, string desc) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) RegAsMinerPool(gno *big.Int, mainAddr common.Address, subAddr [32]byte, name string, desc string) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.RegAsMinerPool(&_MicroPaySystem.TransactOpts, gno, mainAddr, subAddr, name, desc)
}

// SetPoolType is a paid mutator transaction binding the contract method 0x0ff45f38.
//
// Solidity: function SetPoolType(address mainAddr, uint8 typ) returns()
func (_MicroPaySystem *MicroPaySystemTransactor) SetPoolType(opts *bind.TransactOpts, mainAddr common.Address, typ uint8) (*types.Transaction, error) {
	return _MicroPaySystem.contract.Transact(opts, "SetPoolType", mainAddr, typ)
}

// SetPoolType is a paid mutator transaction binding the contract method 0x0ff45f38.
//
// Solidity: function SetPoolType(address mainAddr, uint8 typ) returns()
func (_MicroPaySystem *MicroPaySystemSession) SetPoolType(mainAddr common.Address, typ uint8) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.SetPoolType(&_MicroPaySystem.TransactOpts, mainAddr, typ)
}

// SetPoolType is a paid mutator transaction binding the contract method 0x0ff45f38.
//
// Solidity: function SetPoolType(address mainAddr, uint8 typ) returns()
func (_MicroPaySystem *MicroPaySystemTransactorSession) SetPoolType(mainAddr common.Address, typ uint8) (*types.Transaction, error) {
	return _MicroPaySystem.Contract.SetPoolType(&_MicroPaySystem.TransactOpts, mainAddr, typ)
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
