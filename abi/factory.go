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

// MinerFactorytokenInfo is an auto generated low-level Go binding around an user-defined struct.
type MinerFactorytokenInfo struct {
	IpTokenIds        []*big.Int
	BandwidthTokenIds []*big.Int
	GpuTokenIds       []*big.Int
	StorageTokenIds   []*big.Int
	CpuTokenIds       []*big.Int
	MemTokenIds       []*big.Int
}

// MinerFactoryMetaData contains all meta data concerning the MinerFactory contract.
var MinerFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"MinerDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bandwidthAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burnNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cpuAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"deployMiner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"getIdsByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"ipTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"bandwidthTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gpuTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"storageTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"cpuTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"memTokenIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structMinerFactory.tokenInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"}],\"name\":\"getNodeById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gpuAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bandwidth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gpu\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_storage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cpu\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mem\",\"type\":\"address\"}],\"name\":\"initAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ipAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"memAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_info\",\"type\":\"string\"}],\"name\":\"mintNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setTokenPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MinerFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MinerFactoryMetaData.ABI instead.
var MinerFactoryABI = MinerFactoryMetaData.ABI

// MinerFactory is an auto generated Go binding around an Ethereum contract.
type MinerFactory struct {
	MinerFactoryCaller     // Read-only binding to the contract
	MinerFactoryTransactor // Write-only binding to the contract
	MinerFactoryFilterer   // Log filterer for contract events
}

// MinerFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinerFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinerFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinerFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinerFactorySession struct {
	Contract     *MinerFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinerFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinerFactoryCallerSession struct {
	Contract *MinerFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MinerFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinerFactoryTransactorSession struct {
	Contract     *MinerFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MinerFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinerFactoryRaw struct {
	Contract *MinerFactory // Generic contract binding to access the raw methods on
}

// MinerFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinerFactoryCallerRaw struct {
	Contract *MinerFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MinerFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinerFactoryTransactorRaw struct {
	Contract *MinerFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinerFactory creates a new instance of MinerFactory, bound to a specific deployed contract.
func NewMinerFactory(address common.Address, backend bind.ContractBackend) (*MinerFactory, error) {
	contract, err := bindMinerFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinerFactory{MinerFactoryCaller: MinerFactoryCaller{contract: contract}, MinerFactoryTransactor: MinerFactoryTransactor{contract: contract}, MinerFactoryFilterer: MinerFactoryFilterer{contract: contract}}, nil
}

// NewMinerFactoryCaller creates a new read-only instance of MinerFactory, bound to a specific deployed contract.
func NewMinerFactoryCaller(address common.Address, caller bind.ContractCaller) (*MinerFactoryCaller, error) {
	contract, err := bindMinerFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinerFactoryCaller{contract: contract}, nil
}

// NewMinerFactoryTransactor creates a new write-only instance of MinerFactory, bound to a specific deployed contract.
func NewMinerFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MinerFactoryTransactor, error) {
	contract, err := bindMinerFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinerFactoryTransactor{contract: contract}, nil
}

// NewMinerFactoryFilterer creates a new log filterer instance of MinerFactory, bound to a specific deployed contract.
func NewMinerFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MinerFactoryFilterer, error) {
	contract, err := bindMinerFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinerFactoryFilterer{contract: contract}, nil
}

// bindMinerFactory binds a generic wrapper to an already deployed contract.
func bindMinerFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinerFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinerFactory *MinerFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinerFactory.Contract.MinerFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinerFactory *MinerFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinerFactory.Contract.MinerFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinerFactory *MinerFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinerFactory.Contract.MinerFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinerFactory *MinerFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinerFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinerFactory *MinerFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinerFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinerFactory *MinerFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinerFactory.Contract.contract.Transact(opts, method, params...)
}

// BandwidthAddress is a free data retrieval call binding the contract method 0xbe1e3951.
//
// Solidity: function bandwidthAddress() view returns(address)
func (_MinerFactory *MinerFactoryCaller) BandwidthAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "bandwidthAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BandwidthAddress is a free data retrieval call binding the contract method 0xbe1e3951.
//
// Solidity: function bandwidthAddress() view returns(address)
func (_MinerFactory *MinerFactorySession) BandwidthAddress() (common.Address, error) {
	return _MinerFactory.Contract.BandwidthAddress(&_MinerFactory.CallOpts)
}

// BandwidthAddress is a free data retrieval call binding the contract method 0xbe1e3951.
//
// Solidity: function bandwidthAddress() view returns(address)
func (_MinerFactory *MinerFactoryCallerSession) BandwidthAddress() (common.Address, error) {
	return _MinerFactory.Contract.BandwidthAddress(&_MinerFactory.CallOpts)
}

// CpuAddress is a free data retrieval call binding the contract method 0xc5e034f2.
//
// Solidity: function cpuAddress() view returns(address)
func (_MinerFactory *MinerFactoryCaller) CpuAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "cpuAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CpuAddress is a free data retrieval call binding the contract method 0xc5e034f2.
//
// Solidity: function cpuAddress() view returns(address)
func (_MinerFactory *MinerFactorySession) CpuAddress() (common.Address, error) {
	return _MinerFactory.Contract.CpuAddress(&_MinerFactory.CallOpts)
}

// CpuAddress is a free data retrieval call binding the contract method 0xc5e034f2.
//
// Solidity: function cpuAddress() view returns(address)
func (_MinerFactory *MinerFactoryCallerSession) CpuAddress() (common.Address, error) {
	return _MinerFactory.Contract.CpuAddress(&_MinerFactory.CallOpts)
}

// GetIdsByMiner is a free data retrieval call binding the contract method 0x3266f0aa.
//
// Solidity: function getIdsByMiner(address miner) view returns((uint256[],uint256[],uint256[],uint256[],uint256[],uint256[]))
func (_MinerFactory *MinerFactoryCaller) GetIdsByMiner(opts *bind.CallOpts, miner common.Address) (MinerFactorytokenInfo, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "getIdsByMiner", miner)

	if err != nil {
		return *new(MinerFactorytokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(MinerFactorytokenInfo)).(*MinerFactorytokenInfo)

	return out0, err

}

// GetIdsByMiner is a free data retrieval call binding the contract method 0x3266f0aa.
//
// Solidity: function getIdsByMiner(address miner) view returns((uint256[],uint256[],uint256[],uint256[],uint256[],uint256[]))
func (_MinerFactory *MinerFactorySession) GetIdsByMiner(miner common.Address) (MinerFactorytokenInfo, error) {
	return _MinerFactory.Contract.GetIdsByMiner(&_MinerFactory.CallOpts, miner)
}

// GetIdsByMiner is a free data retrieval call binding the contract method 0x3266f0aa.
//
// Solidity: function getIdsByMiner(address miner) view returns((uint256[],uint256[],uint256[],uint256[],uint256[],uint256[]))
func (_MinerFactory *MinerFactoryCallerSession) GetIdsByMiner(miner common.Address) (MinerFactorytokenInfo, error) {
	return _MinerFactory.Contract.GetIdsByMiner(&_MinerFactory.CallOpts, miner)
}

// GetNodeById is a free data retrieval call binding the contract method 0x0e6a8496.
//
// Solidity: function getNodeById(uint256 nodeId) view returns(address)
func (_MinerFactory *MinerFactoryCaller) GetNodeById(opts *bind.CallOpts, nodeId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "getNodeById", nodeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeById is a free data retrieval call binding the contract method 0x0e6a8496.
//
// Solidity: function getNodeById(uint256 nodeId) view returns(address)
func (_MinerFactory *MinerFactorySession) GetNodeById(nodeId *big.Int) (common.Address, error) {
	return _MinerFactory.Contract.GetNodeById(&_MinerFactory.CallOpts, nodeId)
}

// GetNodeById is a free data retrieval call binding the contract method 0x0e6a8496.
//
// Solidity: function getNodeById(uint256 nodeId) view returns(address)
func (_MinerFactory *MinerFactoryCallerSession) GetNodeById(nodeId *big.Int) (common.Address, error) {
	return _MinerFactory.Contract.GetNodeById(&_MinerFactory.CallOpts, nodeId)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_MinerFactory *MinerFactoryCaller) GetNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "getNodeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_MinerFactory *MinerFactorySession) GetNodeCount() (*big.Int, error) {
	return _MinerFactory.Contract.GetNodeCount(&_MinerFactory.CallOpts)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_MinerFactory *MinerFactoryCallerSession) GetNodeCount() (*big.Int, error) {
	return _MinerFactory.Contract.GetNodeCount(&_MinerFactory.CallOpts)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0xf586c6d9.
//
// Solidity: function getPoolAddress() view returns(address)
func (_MinerFactory *MinerFactoryCaller) GetPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "getPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolAddress is a free data retrieval call binding the contract method 0xf586c6d9.
//
// Solidity: function getPoolAddress() view returns(address)
func (_MinerFactory *MinerFactorySession) GetPoolAddress() (common.Address, error) {
	return _MinerFactory.Contract.GetPoolAddress(&_MinerFactory.CallOpts)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0xf586c6d9.
//
// Solidity: function getPoolAddress() view returns(address)
func (_MinerFactory *MinerFactoryCallerSession) GetPoolAddress() (common.Address, error) {
	return _MinerFactory.Contract.GetPoolAddress(&_MinerFactory.CallOpts)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0xc9f7153c.
//
// Solidity: function getTokenPrice(address contractAddress, uint256 tokenId) view returns(uint256)
func (_MinerFactory *MinerFactoryCaller) GetTokenPrice(opts *bind.CallOpts, contractAddress common.Address, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "getTokenPrice", contractAddress, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenPrice is a free data retrieval call binding the contract method 0xc9f7153c.
//
// Solidity: function getTokenPrice(address contractAddress, uint256 tokenId) view returns(uint256)
func (_MinerFactory *MinerFactorySession) GetTokenPrice(contractAddress common.Address, tokenId *big.Int) (*big.Int, error) {
	return _MinerFactory.Contract.GetTokenPrice(&_MinerFactory.CallOpts, contractAddress, tokenId)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0xc9f7153c.
//
// Solidity: function getTokenPrice(address contractAddress, uint256 tokenId) view returns(uint256)
func (_MinerFactory *MinerFactoryCallerSession) GetTokenPrice(contractAddress common.Address, tokenId *big.Int) (*big.Int, error) {
	return _MinerFactory.Contract.GetTokenPrice(&_MinerFactory.CallOpts, contractAddress, tokenId)
}

// GpuAddress is a free data retrieval call binding the contract method 0xb96519dc.
//
// Solidity: function gpuAddress() view returns(address)
func (_MinerFactory *MinerFactoryCaller) GpuAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "gpuAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GpuAddress is a free data retrieval call binding the contract method 0xb96519dc.
//
// Solidity: function gpuAddress() view returns(address)
func (_MinerFactory *MinerFactorySession) GpuAddress() (common.Address, error) {
	return _MinerFactory.Contract.GpuAddress(&_MinerFactory.CallOpts)
}

// GpuAddress is a free data retrieval call binding the contract method 0xb96519dc.
//
// Solidity: function gpuAddress() view returns(address)
func (_MinerFactory *MinerFactoryCallerSession) GpuAddress() (common.Address, error) {
	return _MinerFactory.Contract.GpuAddress(&_MinerFactory.CallOpts)
}

// IpAddress is a free data retrieval call binding the contract method 0xf6725d4b.
//
// Solidity: function ipAddress() view returns(address)
func (_MinerFactory *MinerFactoryCaller) IpAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "ipAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IpAddress is a free data retrieval call binding the contract method 0xf6725d4b.
//
// Solidity: function ipAddress() view returns(address)
func (_MinerFactory *MinerFactorySession) IpAddress() (common.Address, error) {
	return _MinerFactory.Contract.IpAddress(&_MinerFactory.CallOpts)
}

// IpAddress is a free data retrieval call binding the contract method 0xf6725d4b.
//
// Solidity: function ipAddress() view returns(address)
func (_MinerFactory *MinerFactoryCallerSession) IpAddress() (common.Address, error) {
	return _MinerFactory.Contract.IpAddress(&_MinerFactory.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address contractAddress) view returns(address)
func (_MinerFactory *MinerFactoryCaller) IsActive(opts *bind.CallOpts, contractAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "isActive", contractAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address contractAddress) view returns(address)
func (_MinerFactory *MinerFactorySession) IsActive(contractAddress common.Address) (common.Address, error) {
	return _MinerFactory.Contract.IsActive(&_MinerFactory.CallOpts, contractAddress)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address contractAddress) view returns(address)
func (_MinerFactory *MinerFactoryCallerSession) IsActive(contractAddress common.Address) (common.Address, error) {
	return _MinerFactory.Contract.IsActive(&_MinerFactory.CallOpts, contractAddress)
}

// MemAddress is a free data retrieval call binding the contract method 0x4a16f471.
//
// Solidity: function memAddress() view returns(address)
func (_MinerFactory *MinerFactoryCaller) MemAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "memAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MemAddress is a free data retrieval call binding the contract method 0x4a16f471.
//
// Solidity: function memAddress() view returns(address)
func (_MinerFactory *MinerFactorySession) MemAddress() (common.Address, error) {
	return _MinerFactory.Contract.MemAddress(&_MinerFactory.CallOpts)
}

// MemAddress is a free data retrieval call binding the contract method 0x4a16f471.
//
// Solidity: function memAddress() view returns(address)
func (_MinerFactory *MinerFactoryCallerSession) MemAddress() (common.Address, error) {
	return _MinerFactory.Contract.MemAddress(&_MinerFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MinerFactory *MinerFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MinerFactory *MinerFactorySession) Owner() (common.Address, error) {
	return _MinerFactory.Contract.Owner(&_MinerFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MinerFactory *MinerFactoryCallerSession) Owner() (common.Address, error) {
	return _MinerFactory.Contract.Owner(&_MinerFactory.CallOpts)
}

// StorageAddress is a free data retrieval call binding the contract method 0x85aa92a7.
//
// Solidity: function storageAddress() view returns(address)
func (_MinerFactory *MinerFactoryCaller) StorageAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinerFactory.contract.Call(opts, &out, "storageAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageAddress is a free data retrieval call binding the contract method 0x85aa92a7.
//
// Solidity: function storageAddress() view returns(address)
func (_MinerFactory *MinerFactorySession) StorageAddress() (common.Address, error) {
	return _MinerFactory.Contract.StorageAddress(&_MinerFactory.CallOpts)
}

// StorageAddress is a free data retrieval call binding the contract method 0x85aa92a7.
//
// Solidity: function storageAddress() view returns(address)
func (_MinerFactory *MinerFactoryCallerSession) StorageAddress() (common.Address, error) {
	return _MinerFactory.Contract.StorageAddress(&_MinerFactory.CallOpts)
}

// BurnNfts is a paid mutator transaction binding the contract method 0x3e7ca1bb.
//
// Solidity: function burnNfts(address miner, address contractAddress, uint256 tokenId) returns()
func (_MinerFactory *MinerFactoryTransactor) BurnNfts(opts *bind.TransactOpts, miner common.Address, contractAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MinerFactory.contract.Transact(opts, "burnNfts", miner, contractAddress, tokenId)
}

// BurnNfts is a paid mutator transaction binding the contract method 0x3e7ca1bb.
//
// Solidity: function burnNfts(address miner, address contractAddress, uint256 tokenId) returns()
func (_MinerFactory *MinerFactorySession) BurnNfts(miner common.Address, contractAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MinerFactory.Contract.BurnNfts(&_MinerFactory.TransactOpts, miner, contractAddress, tokenId)
}

// BurnNfts is a paid mutator transaction binding the contract method 0x3e7ca1bb.
//
// Solidity: function burnNfts(address miner, address contractAddress, uint256 tokenId) returns()
func (_MinerFactory *MinerFactoryTransactorSession) BurnNfts(miner common.Address, contractAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MinerFactory.Contract.BurnNfts(&_MinerFactory.TransactOpts, miner, contractAddress, tokenId)
}

// DeployMiner is a paid mutator transaction binding the contract method 0xb9299900.
//
// Solidity: function deployMiner(address issuer) returns(address)
func (_MinerFactory *MinerFactoryTransactor) DeployMiner(opts *bind.TransactOpts, issuer common.Address) (*types.Transaction, error) {
	return _MinerFactory.contract.Transact(opts, "deployMiner", issuer)
}

// DeployMiner is a paid mutator transaction binding the contract method 0xb9299900.
//
// Solidity: function deployMiner(address issuer) returns(address)
func (_MinerFactory *MinerFactorySession) DeployMiner(issuer common.Address) (*types.Transaction, error) {
	return _MinerFactory.Contract.DeployMiner(&_MinerFactory.TransactOpts, issuer)
}

// DeployMiner is a paid mutator transaction binding the contract method 0xb9299900.
//
// Solidity: function deployMiner(address issuer) returns(address)
func (_MinerFactory *MinerFactoryTransactorSession) DeployMiner(issuer common.Address) (*types.Transaction, error) {
	return _MinerFactory.Contract.DeployMiner(&_MinerFactory.TransactOpts, issuer)
}

// InitAddress is a paid mutator transaction binding the contract method 0x4553fc8c.
//
// Solidity: function initAddress(address _poolAddress, address tokenAddr, address ip, address bandwidth, address gpu, address _storage, address cpu, address mem) returns()
func (_MinerFactory *MinerFactoryTransactor) InitAddress(opts *bind.TransactOpts, _poolAddress common.Address, tokenAddr common.Address, ip common.Address, bandwidth common.Address, gpu common.Address, _storage common.Address, cpu common.Address, mem common.Address) (*types.Transaction, error) {
	return _MinerFactory.contract.Transact(opts, "initAddress", _poolAddress, tokenAddr, ip, bandwidth, gpu, _storage, cpu, mem)
}

// InitAddress is a paid mutator transaction binding the contract method 0x4553fc8c.
//
// Solidity: function initAddress(address _poolAddress, address tokenAddr, address ip, address bandwidth, address gpu, address _storage, address cpu, address mem) returns()
func (_MinerFactory *MinerFactorySession) InitAddress(_poolAddress common.Address, tokenAddr common.Address, ip common.Address, bandwidth common.Address, gpu common.Address, _storage common.Address, cpu common.Address, mem common.Address) (*types.Transaction, error) {
	return _MinerFactory.Contract.InitAddress(&_MinerFactory.TransactOpts, _poolAddress, tokenAddr, ip, bandwidth, gpu, _storage, cpu, mem)
}

// InitAddress is a paid mutator transaction binding the contract method 0x4553fc8c.
//
// Solidity: function initAddress(address _poolAddress, address tokenAddr, address ip, address bandwidth, address gpu, address _storage, address cpu, address mem) returns()
func (_MinerFactory *MinerFactoryTransactorSession) InitAddress(_poolAddress common.Address, tokenAddr common.Address, ip common.Address, bandwidth common.Address, gpu common.Address, _storage common.Address, cpu common.Address, mem common.Address) (*types.Transaction, error) {
	return _MinerFactory.Contract.InitAddress(&_MinerFactory.TransactOpts, _poolAddress, tokenAddr, ip, bandwidth, gpu, _storage, cpu, mem)
}

// MintNfts is a paid mutator transaction binding the contract method 0x44076ece.
//
// Solidity: function mintNfts(address miner, address contractAddress, string _info) returns()
func (_MinerFactory *MinerFactoryTransactor) MintNfts(opts *bind.TransactOpts, miner common.Address, contractAddress common.Address, _info string) (*types.Transaction, error) {
	return _MinerFactory.contract.Transact(opts, "mintNfts", miner, contractAddress, _info)
}

// MintNfts is a paid mutator transaction binding the contract method 0x44076ece.
//
// Solidity: function mintNfts(address miner, address contractAddress, string _info) returns()
func (_MinerFactory *MinerFactorySession) MintNfts(miner common.Address, contractAddress common.Address, _info string) (*types.Transaction, error) {
	return _MinerFactory.Contract.MintNfts(&_MinerFactory.TransactOpts, miner, contractAddress, _info)
}

// MintNfts is a paid mutator transaction binding the contract method 0x44076ece.
//
// Solidity: function mintNfts(address miner, address contractAddress, string _info) returns()
func (_MinerFactory *MinerFactoryTransactorSession) MintNfts(miner common.Address, contractAddress common.Address, _info string) (*types.Transaction, error) {
	return _MinerFactory.Contract.MintNfts(&_MinerFactory.TransactOpts, miner, contractAddress, _info)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MinerFactory *MinerFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinerFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MinerFactory *MinerFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _MinerFactory.Contract.RenounceOwnership(&_MinerFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MinerFactory *MinerFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MinerFactory.Contract.RenounceOwnership(&_MinerFactory.TransactOpts)
}

// SetTokenPrice is a paid mutator transaction binding the contract method 0xeb1c0316.
//
// Solidity: function setTokenPrice(address contractAddress, uint256 tokenId, uint256 _price) returns()
func (_MinerFactory *MinerFactoryTransactor) SetTokenPrice(opts *bind.TransactOpts, contractAddress common.Address, tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _MinerFactory.contract.Transact(opts, "setTokenPrice", contractAddress, tokenId, _price)
}

// SetTokenPrice is a paid mutator transaction binding the contract method 0xeb1c0316.
//
// Solidity: function setTokenPrice(address contractAddress, uint256 tokenId, uint256 _price) returns()
func (_MinerFactory *MinerFactorySession) SetTokenPrice(contractAddress common.Address, tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _MinerFactory.Contract.SetTokenPrice(&_MinerFactory.TransactOpts, contractAddress, tokenId, _price)
}

// SetTokenPrice is a paid mutator transaction binding the contract method 0xeb1c0316.
//
// Solidity: function setTokenPrice(address contractAddress, uint256 tokenId, uint256 _price) returns()
func (_MinerFactory *MinerFactoryTransactorSession) SetTokenPrice(contractAddress common.Address, tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _MinerFactory.Contract.SetTokenPrice(&_MinerFactory.TransactOpts, contractAddress, tokenId, _price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MinerFactory *MinerFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MinerFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MinerFactory *MinerFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MinerFactory.Contract.TransferOwnership(&_MinerFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MinerFactory *MinerFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MinerFactory.Contract.TransferOwnership(&_MinerFactory.TransactOpts, newOwner)
}

// MinerFactoryMinerDeployedIterator is returned from FilterMinerDeployed and is used to iterate over the raw logs and unpacked data for MinerDeployed events raised by the MinerFactory contract.
type MinerFactoryMinerDeployedIterator struct {
	Event *MinerFactoryMinerDeployed // Event containing the contract specifics and raw log

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
func (it *MinerFactoryMinerDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinerFactoryMinerDeployed)
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
		it.Event = new(MinerFactoryMinerDeployed)
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
func (it *MinerFactoryMinerDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinerFactoryMinerDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinerFactoryMinerDeployed represents a MinerDeployed event raised by the MinerFactory contract.
type MinerFactoryMinerDeployed struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMinerDeployed is a free log retrieval operation binding the contract event 0xe80bab348976231d375fb6b307274bc236ff7055ba1817e8f19e8811283061f1.
//
// Solidity: event MinerDeployed(address contractAddress)
func (_MinerFactory *MinerFactoryFilterer) FilterMinerDeployed(opts *bind.FilterOpts) (*MinerFactoryMinerDeployedIterator, error) {

	logs, sub, err := _MinerFactory.contract.FilterLogs(opts, "MinerDeployed")
	if err != nil {
		return nil, err
	}
	return &MinerFactoryMinerDeployedIterator{contract: _MinerFactory.contract, event: "MinerDeployed", logs: logs, sub: sub}, nil
}

// WatchMinerDeployed is a free log subscription operation binding the contract event 0xe80bab348976231d375fb6b307274bc236ff7055ba1817e8f19e8811283061f1.
//
// Solidity: event MinerDeployed(address contractAddress)
func (_MinerFactory *MinerFactoryFilterer) WatchMinerDeployed(opts *bind.WatchOpts, sink chan<- *MinerFactoryMinerDeployed) (event.Subscription, error) {

	logs, sub, err := _MinerFactory.contract.WatchLogs(opts, "MinerDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinerFactoryMinerDeployed)
				if err := _MinerFactory.contract.UnpackLog(event, "MinerDeployed", log); err != nil {
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

// ParseMinerDeployed is a log parse operation binding the contract event 0xe80bab348976231d375fb6b307274bc236ff7055ba1817e8f19e8811283061f1.
//
// Solidity: event MinerDeployed(address contractAddress)
func (_MinerFactory *MinerFactoryFilterer) ParseMinerDeployed(log types.Log) (*MinerFactoryMinerDeployed, error) {
	event := new(MinerFactoryMinerDeployed)
	if err := _MinerFactory.contract.UnpackLog(event, "MinerDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinerFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MinerFactory contract.
type MinerFactoryOwnershipTransferredIterator struct {
	Event *MinerFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MinerFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinerFactoryOwnershipTransferred)
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
		it.Event = new(MinerFactoryOwnershipTransferred)
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
func (it *MinerFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinerFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinerFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the MinerFactory contract.
type MinerFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MinerFactory *MinerFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MinerFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MinerFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MinerFactoryOwnershipTransferredIterator{contract: _MinerFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MinerFactory *MinerFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MinerFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MinerFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinerFactoryOwnershipTransferred)
				if err := _MinerFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MinerFactory *MinerFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*MinerFactoryOwnershipTransferred, error) {
	event := new(MinerFactoryOwnershipTransferred)
	if err := _MinerFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
