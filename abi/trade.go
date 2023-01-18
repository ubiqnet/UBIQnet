// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TradeNodeInfo is an auto generated low-level Go binding around an user-defined struct.
type TradeNodeInfo struct {
	StartTime   *big.Int
	CreateTime  *big.Int
	Duration    uint64
	TotalFee    *big.Int
	AppAddress  common.Address
	ExpiresTime *big.Int
	IsBurned    bool
}

// TradeMetaData contains all meta data concerning the Trade contract.
var TradeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"appAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyAddress\",\"type\":\"address\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnExpireToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fac\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cpu\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mem\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bandwidth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gpu\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_storage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_appFactory\",\"type\":\"address\"}],\"name\":\"setERC721Address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"appAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_nodeAddress\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][][]\",\"name\":\"_tokenIds\",\"type\":\"uint256[][][]\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"minerAddress\",\"type\":\"address\"}],\"name\":\"setTokenIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"appAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"stopApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getExpireToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestOrderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"getOrderBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"getOrderInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"appAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiresTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBurned\",\"type\":\"bool\"}],\"internalType\":\"structTrade.NodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getOrderList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TradeABI is the input ABI used to generate the binding from.
// Deprecated: Use TradeMetaData.ABI instead.
var TradeABI = TradeMetaData.ABI

// Trade is an auto generated Go binding around an Ethereum contract.
type Trade struct {
	TradeCaller     // Read-only binding to the contract
	TradeTransactor // Write-only binding to the contract
	TradeFilterer   // Log filterer for contract events
}

// TradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradeSession struct {
	Contract     *Trade            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradeCallerSession struct {
	Contract *TradeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradeTransactorSession struct {
	Contract     *TradeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradeRaw struct {
	Contract *Trade // Generic contract binding to access the raw methods on
}

// TradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradeCallerRaw struct {
	Contract *TradeCaller // Generic read-only contract binding to access the raw methods on
}

// TradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradeTransactorRaw struct {
	Contract *TradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrade creates a new instance of Trade, bound to a specific deployed contract.
func NewTrade(address common.Address, backend bind.ContractBackend) (*Trade, error) {
	contract, err := bindTrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trade{TradeCaller: TradeCaller{contract: contract}, TradeTransactor: TradeTransactor{contract: contract}, TradeFilterer: TradeFilterer{contract: contract}}, nil
}

// NewTradeCaller creates a new read-only instance of Trade, bound to a specific deployed contract.
func NewTradeCaller(address common.Address, caller bind.ContractCaller) (*TradeCaller, error) {
	contract, err := bindTrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradeCaller{contract: contract}, nil
}

// NewTradeTransactor creates a new write-only instance of Trade, bound to a specific deployed contract.
func NewTradeTransactor(address common.Address, transactor bind.ContractTransactor) (*TradeTransactor, error) {
	contract, err := bindTrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradeTransactor{contract: contract}, nil
}

// NewTradeFilterer creates a new log filterer instance of Trade, bound to a specific deployed contract.
func NewTradeFilterer(address common.Address, filterer bind.ContractFilterer) (*TradeFilterer, error) {
	contract, err := bindTrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradeFilterer{contract: contract}, nil
}

// bindTrade binds a generic wrapper to an already deployed contract.
func bindTrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade *TradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trade.Contract.TradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade *TradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.Contract.TradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade *TradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade.Contract.TradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade *TradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade *TradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade *TradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade.Contract.contract.Transact(opts, method, params...)
}

// GetExpireToken is a free data retrieval call binding the contract method 0x250ff46b.
//
// Solidity: function getExpireToken() view returns(uint256)
func (_Trade *TradeCaller) GetExpireToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trade.contract.Call(opts, &out, "getExpireToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpireToken is a free data retrieval call binding the contract method 0x250ff46b.
//
// Solidity: function getExpireToken() view returns(uint256)
func (_Trade *TradeSession) GetExpireToken() (*big.Int, error) {
	return _Trade.Contract.GetExpireToken(&_Trade.CallOpts)
}

// GetExpireToken is a free data retrieval call binding the contract method 0x250ff46b.
//
// Solidity: function getExpireToken() view returns(uint256)
func (_Trade *TradeCallerSession) GetExpireToken() (*big.Int, error) {
	return _Trade.Contract.GetExpireToken(&_Trade.CallOpts)
}

// GetLatestOrderId is a free data retrieval call binding the contract method 0x3d7fe973.
//
// Solidity: function getLatestOrderId() view returns(uint256)
func (_Trade *TradeCaller) GetLatestOrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trade.contract.Call(opts, &out, "getLatestOrderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestOrderId is a free data retrieval call binding the contract method 0x3d7fe973.
//
// Solidity: function getLatestOrderId() view returns(uint256)
func (_Trade *TradeSession) GetLatestOrderId() (*big.Int, error) {
	return _Trade.Contract.GetLatestOrderId(&_Trade.CallOpts)
}

// GetLatestOrderId is a free data retrieval call binding the contract method 0x3d7fe973.
//
// Solidity: function getLatestOrderId() view returns(uint256)
func (_Trade *TradeCallerSession) GetLatestOrderId() (*big.Int, error) {
	return _Trade.Contract.GetLatestOrderId(&_Trade.CallOpts)
}

// GetOrderBalance is a free data retrieval call binding the contract method 0x62951cec.
//
// Solidity: function getOrderBalance(uint256 orderId) view returns(uint256)
func (_Trade *TradeCaller) GetOrderBalance(opts *bind.CallOpts, orderId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Trade.contract.Call(opts, &out, "getOrderBalance", orderId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderBalance is a free data retrieval call binding the contract method 0x62951cec.
//
// Solidity: function getOrderBalance(uint256 orderId) view returns(uint256)
func (_Trade *TradeSession) GetOrderBalance(orderId *big.Int) (*big.Int, error) {
	return _Trade.Contract.GetOrderBalance(&_Trade.CallOpts, orderId)
}

// GetOrderBalance is a free data retrieval call binding the contract method 0x62951cec.
//
// Solidity: function getOrderBalance(uint256 orderId) view returns(uint256)
func (_Trade *TradeCallerSession) GetOrderBalance(orderId *big.Int) (*big.Int, error) {
	return _Trade.Contract.GetOrderBalance(&_Trade.CallOpts, orderId)
}

// GetOrderInfo is a free data retrieval call binding the contract method 0xd311636b.
//
// Solidity: function getOrderInfo(uint256 orderId) view returns((uint256,uint256,uint64,uint256,address,uint256,bool))
func (_Trade *TradeCaller) GetOrderInfo(opts *bind.CallOpts, orderId *big.Int) (TradeNodeInfo, error) {
	var out []interface{}
	err := _Trade.contract.Call(opts, &out, "getOrderInfo", orderId)

	if err != nil {
		return *new(TradeNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TradeNodeInfo)).(*TradeNodeInfo)

	return out0, err

}

// GetOrderInfo is a free data retrieval call binding the contract method 0xd311636b.
//
// Solidity: function getOrderInfo(uint256 orderId) view returns((uint256,uint256,uint64,uint256,address,uint256,bool))
func (_Trade *TradeSession) GetOrderInfo(orderId *big.Int) (TradeNodeInfo, error) {
	return _Trade.Contract.GetOrderInfo(&_Trade.CallOpts, orderId)
}

// GetOrderInfo is a free data retrieval call binding the contract method 0xd311636b.
//
// Solidity: function getOrderInfo(uint256 orderId) view returns((uint256,uint256,uint64,uint256,address,uint256,bool))
func (_Trade *TradeCallerSession) GetOrderInfo(orderId *big.Int) (TradeNodeInfo, error) {
	return _Trade.Contract.GetOrderInfo(&_Trade.CallOpts, orderId)
}

// GetOrderList is a free data retrieval call binding the contract method 0xe1ee4758.
//
// Solidity: function getOrderList(address _addr) view returns(uint256[])
func (_Trade *TradeCaller) GetOrderList(opts *bind.CallOpts, _addr common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Trade.contract.Call(opts, &out, "getOrderList", _addr)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOrderList is a free data retrieval call binding the contract method 0xe1ee4758.
//
// Solidity: function getOrderList(address _addr) view returns(uint256[])
func (_Trade *TradeSession) GetOrderList(_addr common.Address) ([]*big.Int, error) {
	return _Trade.Contract.GetOrderList(&_Trade.CallOpts, _addr)
}

// GetOrderList is a free data retrieval call binding the contract method 0xe1ee4758.
//
// Solidity: function getOrderList(address _addr) view returns(uint256[])
func (_Trade *TradeCallerSession) GetOrderList(_addr common.Address) ([]*big.Int, error) {
	return _Trade.Contract.GetOrderList(&_Trade.CallOpts, _addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trade *TradeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trade.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trade *TradeSession) Owner() (common.Address, error) {
	return _Trade.Contract.Owner(&_Trade.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trade *TradeCallerSession) Owner() (common.Address, error) {
	return _Trade.Contract.Owner(&_Trade.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Trade *TradeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Trade.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Trade *TradeSession) Paused() (bool, error) {
	return _Trade.Contract.Paused(&_Trade.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Trade *TradeCallerSession) Paused() (bool, error) {
	return _Trade.Contract.Paused(&_Trade.CallOpts)
}

// AddOrder is a paid mutator transaction binding the contract method 0x5352631f.
//
// Solidity: function addOrder(uint64 duration, uint256 totalPrice, address appAddress, address buyAddress) returns()
func (_Trade *TradeTransactor) AddOrder(opts *bind.TransactOpts, duration uint64, totalPrice *big.Int, appAddress common.Address, buyAddress common.Address) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "addOrder", duration, totalPrice, appAddress, buyAddress)
}

// AddOrder is a paid mutator transaction binding the contract method 0x5352631f.
//
// Solidity: function addOrder(uint64 duration, uint256 totalPrice, address appAddress, address buyAddress) returns()
func (_Trade *TradeSession) AddOrder(duration uint64, totalPrice *big.Int, appAddress common.Address, buyAddress common.Address) (*types.Transaction, error) {
	return _Trade.Contract.AddOrder(&_Trade.TransactOpts, duration, totalPrice, appAddress, buyAddress)
}

// AddOrder is a paid mutator transaction binding the contract method 0x5352631f.
//
// Solidity: function addOrder(uint64 duration, uint256 totalPrice, address appAddress, address buyAddress) returns()
func (_Trade *TradeTransactorSession) AddOrder(duration uint64, totalPrice *big.Int, appAddress common.Address, buyAddress common.Address) (*types.Transaction, error) {
	return _Trade.Contract.AddOrder(&_Trade.TransactOpts, duration, totalPrice, appAddress, buyAddress)
}

// BurnExpireToken is a paid mutator transaction binding the contract method 0x26ccf7c8.
//
// Solidity: function burnExpireToken() returns()
func (_Trade *TradeTransactor) BurnExpireToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "burnExpireToken")
}

// BurnExpireToken is a paid mutator transaction binding the contract method 0x26ccf7c8.
//
// Solidity: function burnExpireToken() returns()
func (_Trade *TradeSession) BurnExpireToken() (*types.Transaction, error) {
	return _Trade.Contract.BurnExpireToken(&_Trade.TransactOpts)
}

// BurnExpireToken is a paid mutator transaction binding the contract method 0x26ccf7c8.
//
// Solidity: function burnExpireToken() returns()
func (_Trade *TradeTransactorSession) BurnExpireToken() (*types.Transaction, error) {
	return _Trade.Contract.BurnExpireToken(&_Trade.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trade *TradeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trade *TradeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trade.Contract.RenounceOwnership(&_Trade.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trade *TradeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trade.Contract.RenounceOwnership(&_Trade.TransactOpts)
}

// SetERC721Address is a paid mutator transaction binding the contract method 0x2a2e9616.
//
// Solidity: function setERC721Address(address _fac, address cpu, address mem, address ip, address bandwidth, address gpu, address _storage, address _appFactory) returns()
func (_Trade *TradeTransactor) SetERC721Address(opts *bind.TransactOpts, _fac common.Address, cpu common.Address, mem common.Address, ip common.Address, bandwidth common.Address, gpu common.Address, _storage common.Address, _appFactory common.Address) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "setERC721Address", _fac, cpu, mem, ip, bandwidth, gpu, _storage, _appFactory)
}

// SetERC721Address is a paid mutator transaction binding the contract method 0x2a2e9616.
//
// Solidity: function setERC721Address(address _fac, address cpu, address mem, address ip, address bandwidth, address gpu, address _storage, address _appFactory) returns()
func (_Trade *TradeSession) SetERC721Address(_fac common.Address, cpu common.Address, mem common.Address, ip common.Address, bandwidth common.Address, gpu common.Address, _storage common.Address, _appFactory common.Address) (*types.Transaction, error) {
	return _Trade.Contract.SetERC721Address(&_Trade.TransactOpts, _fac, cpu, mem, ip, bandwidth, gpu, _storage, _appFactory)
}

// SetERC721Address is a paid mutator transaction binding the contract method 0x2a2e9616.
//
// Solidity: function setERC721Address(address _fac, address cpu, address mem, address ip, address bandwidth, address gpu, address _storage, address _appFactory) returns()
func (_Trade *TradeTransactorSession) SetERC721Address(_fac common.Address, cpu common.Address, mem common.Address, ip common.Address, bandwidth common.Address, gpu common.Address, _storage common.Address, _appFactory common.Address) (*types.Transaction, error) {
	return _Trade.Contract.SetERC721Address(&_Trade.TransactOpts, _fac, cpu, mem, ip, bandwidth, gpu, _storage, _appFactory)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Trade *TradeTransactor) SetTokenAddress(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "setTokenAddress", _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Trade *TradeSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Trade.Contract.SetTokenAddress(&_Trade.TransactOpts, _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Trade *TradeTransactorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Trade.Contract.SetTokenAddress(&_Trade.TransactOpts, _tokenAddress)
}

// SetTokenIds is a paid mutator transaction binding the contract method 0x381554d1.
//
// Solidity: function setTokenIds(address appAddress, address[] _nodeAddress, uint256[][][] _tokenIds, uint256 orderId, address minerAddress) returns()
func (_Trade *TradeTransactor) SetTokenIds(opts *bind.TransactOpts, appAddress common.Address, _nodeAddress []common.Address, _tokenIds [][][]*big.Int, orderId *big.Int, minerAddress common.Address) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "setTokenIds", appAddress, _nodeAddress, _tokenIds, orderId, minerAddress)
}

// SetTokenIds is a paid mutator transaction binding the contract method 0x381554d1.
//
// Solidity: function setTokenIds(address appAddress, address[] _nodeAddress, uint256[][][] _tokenIds, uint256 orderId, address minerAddress) returns()
func (_Trade *TradeSession) SetTokenIds(appAddress common.Address, _nodeAddress []common.Address, _tokenIds [][][]*big.Int, orderId *big.Int, minerAddress common.Address) (*types.Transaction, error) {
	return _Trade.Contract.SetTokenIds(&_Trade.TransactOpts, appAddress, _nodeAddress, _tokenIds, orderId, minerAddress)
}

// SetTokenIds is a paid mutator transaction binding the contract method 0x381554d1.
//
// Solidity: function setTokenIds(address appAddress, address[] _nodeAddress, uint256[][][] _tokenIds, uint256 orderId, address minerAddress) returns()
func (_Trade *TradeTransactorSession) SetTokenIds(appAddress common.Address, _nodeAddress []common.Address, _tokenIds [][][]*big.Int, orderId *big.Int, minerAddress common.Address) (*types.Transaction, error) {
	return _Trade.Contract.SetTokenIds(&_Trade.TransactOpts, appAddress, _nodeAddress, _tokenIds, orderId, minerAddress)
}

// StopApp is a paid mutator transaction binding the contract method 0xd6251560.
//
// Solidity: function stopApp(address appAddress, uint256 orderId) returns()
func (_Trade *TradeTransactor) StopApp(opts *bind.TransactOpts, appAddress common.Address, orderId *big.Int) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "stopApp", appAddress, orderId)
}

// StopApp is a paid mutator transaction binding the contract method 0xd6251560.
//
// Solidity: function stopApp(address appAddress, uint256 orderId) returns()
func (_Trade *TradeSession) StopApp(appAddress common.Address, orderId *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.StopApp(&_Trade.TransactOpts, appAddress, orderId)
}

// StopApp is a paid mutator transaction binding the contract method 0xd6251560.
//
// Solidity: function stopApp(address appAddress, uint256 orderId) returns()
func (_Trade *TradeTransactorSession) StopApp(appAddress common.Address, orderId *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.StopApp(&_Trade.TransactOpts, appAddress, orderId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trade *TradeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trade *TradeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trade.Contract.TransferOwnership(&_Trade.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trade *TradeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trade.Contract.TransferOwnership(&_Trade.TransactOpts, newOwner)
}

// TradeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Trade contract.
type TradeOwnershipTransferredIterator struct {
	Event *TradeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TradeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradeOwnershipTransferred)
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
		it.Event = new(TradeOwnershipTransferred)
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
func (it *TradeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradeOwnershipTransferred represents a OwnershipTransferred event raised by the Trade contract.
type TradeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trade *TradeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TradeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trade.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TradeOwnershipTransferredIterator{contract: _Trade.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trade *TradeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TradeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trade.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradeOwnershipTransferred)
				if err := _Trade.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trade *TradeFilterer) ParseOwnershipTransferred(log types.Log) (*TradeOwnershipTransferred, error) {
	event := new(TradeOwnershipTransferred)
	if err := _Trade.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Trade contract.
type TradePausedIterator struct {
	Event *TradePaused // Event containing the contract specifics and raw log

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
func (it *TradePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradePaused)
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
		it.Event = new(TradePaused)
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
func (it *TradePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradePaused represents a Paused event raised by the Trade contract.
type TradePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Trade *TradeFilterer) FilterPaused(opts *bind.FilterOpts) (*TradePausedIterator, error) {

	logs, sub, err := _Trade.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TradePausedIterator{contract: _Trade.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Trade *TradeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TradePaused) (event.Subscription, error) {

	logs, sub, err := _Trade.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradePaused)
				if err := _Trade.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Trade *TradeFilterer) ParsePaused(log types.Log) (*TradePaused, error) {
	event := new(TradePaused)
	if err := _Trade.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Trade contract.
type TradeUnpausedIterator struct {
	Event *TradeUnpaused // Event containing the contract specifics and raw log

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
func (it *TradeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradeUnpaused)
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
		it.Event = new(TradeUnpaused)
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
func (it *TradeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradeUnpaused represents a Unpaused event raised by the Trade contract.
type TradeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Trade *TradeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TradeUnpausedIterator, error) {

	logs, sub, err := _Trade.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TradeUnpausedIterator{contract: _Trade.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Trade *TradeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TradeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Trade.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradeUnpaused)
				if err := _Trade.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Trade *TradeFilterer) ParseUnpaused(log types.Log) (*TradeUnpaused, error) {
	event := new(TradeUnpaused)
	if err := _Trade.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
