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

// AppMetaData contains all meta data concerning the App contract.
var AppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_templates\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"_isTemplate\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_image\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"_isExpanded\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_tradeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[][]\",\"name\":\"_tokenNeed\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getBandwidthC_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBandwidthTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBandwidth_\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCpuC_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCpuTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCpu_\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGpuC_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGpuMem_\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGpuTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImage\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIpC_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIpTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIp_\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsExpanded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsTemplate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMemC_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMemTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMem_\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMyOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeAddress\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStorageC_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStorageTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStorage_\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTemplates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenIds\",\"outputs\":[{\"internalType\":\"uint256[][][]\",\"name\":\"\",\"type\":\"uint256[][][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenNeed\",\"outputs\":[{\"internalType\":\"uint256[][]\",\"name\":\"\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isTemplate\",\"type\":\"bool\"}],\"name\":\"setIsTemplate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nodeAddress\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][][]\",\"name\":\"_tokenIds\",\"type\":\"uint256[][][]\"}],\"name\":\"setNodeAddressAndTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tradeAddress\",\"type\":\"address\"}],\"name\":\"setTradeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AppABI is the input ABI used to generate the binding from.
// Deprecated: Use AppMetaData.ABI instead.
var AppABI = AppMetaData.ABI

// App is an auto generated Go binding around an Ethereum contract.
type App struct {
	AppCaller     // Read-only binding to the contract
	AppTransactor // Write-only binding to the contract
	AppFilterer   // Log filterer for contract events
}

// AppCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppSession struct {
	Contract     *App              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppCallerSession struct {
	Contract *AppCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppTransactorSession struct {
	Contract     *AppTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppRaw struct {
	Contract *App // Generic contract binding to access the raw methods on
}

// AppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppCallerRaw struct {
	Contract *AppCaller // Generic read-only contract binding to access the raw methods on
}

// AppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppTransactorRaw struct {
	Contract *AppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApp creates a new instance of App, bound to a specific deployed contract.
func NewApp(address common.Address, backend bind.ContractBackend) (*App, error) {
	contract, err := bindApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &App{AppCaller: AppCaller{contract: contract}, AppTransactor: AppTransactor{contract: contract}, AppFilterer: AppFilterer{contract: contract}}, nil
}

// NewAppCaller creates a new read-only instance of App, bound to a specific deployed contract.
func NewAppCaller(address common.Address, caller bind.ContractCaller) (*AppCaller, error) {
	contract, err := bindApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppCaller{contract: contract}, nil
}

// NewAppTransactor creates a new write-only instance of App, bound to a specific deployed contract.
func NewAppTransactor(address common.Address, transactor bind.ContractTransactor) (*AppTransactor, error) {
	contract, err := bindApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppTransactor{contract: contract}, nil
}

// NewAppFilterer creates a new log filterer instance of App, bound to a specific deployed contract.
func NewAppFilterer(address common.Address, filterer bind.ContractFilterer) (*AppFilterer, error) {
	contract, err := bindApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppFilterer{contract: contract}, nil
}

// bindApp binds a generic wrapper to an already deployed contract.
func bindApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_App *AppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _App.Contract.AppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_App *AppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _App.Contract.AppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_App *AppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _App.Contract.AppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_App *AppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _App.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_App *AppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _App.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_App *AppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _App.Contract.contract.Transact(opts, method, params...)
}

// GetBandwidthC is a free data retrieval call binding the contract method 0x3df58f59.
//
// Solidity: function getBandwidthC_() view returns(uint256)
func (_App *AppCaller) GetBandwidthC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getBandwidthC_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBandwidthC is a free data retrieval call binding the contract method 0x3df58f59.
//
// Solidity: function getBandwidthC_() view returns(uint256)
func (_App *AppSession) GetBandwidthC() (*big.Int, error) {
	return _App.Contract.GetBandwidthC(&_App.CallOpts)
}

// GetBandwidthC is a free data retrieval call binding the contract method 0x3df58f59.
//
// Solidity: function getBandwidthC_() view returns(uint256)
func (_App *AppCallerSession) GetBandwidthC() (*big.Int, error) {
	return _App.Contract.GetBandwidthC(&_App.CallOpts)
}

// GetBandwidthTokenIds is a free data retrieval call binding the contract method 0xa6ea3e98.
//
// Solidity: function getBandwidthTokenIds() view returns(uint256[])
func (_App *AppCaller) GetBandwidthTokenIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getBandwidthTokenIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBandwidthTokenIds is a free data retrieval call binding the contract method 0xa6ea3e98.
//
// Solidity: function getBandwidthTokenIds() view returns(uint256[])
func (_App *AppSession) GetBandwidthTokenIds() ([]*big.Int, error) {
	return _App.Contract.GetBandwidthTokenIds(&_App.CallOpts)
}

// GetBandwidthTokenIds is a free data retrieval call binding the contract method 0xa6ea3e98.
//
// Solidity: function getBandwidthTokenIds() view returns(uint256[])
func (_App *AppCallerSession) GetBandwidthTokenIds() ([]*big.Int, error) {
	return _App.Contract.GetBandwidthTokenIds(&_App.CallOpts)
}

// GetBandwidth is a free data retrieval call binding the contract method 0xe5c8eee2.
//
// Solidity: function getBandwidth_() view returns(string[])
func (_App *AppCaller) GetBandwidth(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getBandwidth_")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetBandwidth is a free data retrieval call binding the contract method 0xe5c8eee2.
//
// Solidity: function getBandwidth_() view returns(string[])
func (_App *AppSession) GetBandwidth() ([]string, error) {
	return _App.Contract.GetBandwidth(&_App.CallOpts)
}

// GetBandwidth is a free data retrieval call binding the contract method 0xe5c8eee2.
//
// Solidity: function getBandwidth_() view returns(string[])
func (_App *AppCallerSession) GetBandwidth() ([]string, error) {
	return _App.Contract.GetBandwidth(&_App.CallOpts)
}

// GetCpuC is a free data retrieval call binding the contract method 0x4a4022eb.
//
// Solidity: function getCpuC_() view returns(uint256)
func (_App *AppCaller) GetCpuC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getCpuC_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCpuC is a free data retrieval call binding the contract method 0x4a4022eb.
//
// Solidity: function getCpuC_() view returns(uint256)
func (_App *AppSession) GetCpuC() (*big.Int, error) {
	return _App.Contract.GetCpuC(&_App.CallOpts)
}

// GetCpuC is a free data retrieval call binding the contract method 0x4a4022eb.
//
// Solidity: function getCpuC_() view returns(uint256)
func (_App *AppCallerSession) GetCpuC() (*big.Int, error) {
	return _App.Contract.GetCpuC(&_App.CallOpts)
}

// GetCpuTokenIds is a free data retrieval call binding the contract method 0x00d2b30f.
//
// Solidity: function getCpuTokenIds() view returns(uint256[])
func (_App *AppCaller) GetCpuTokenIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getCpuTokenIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCpuTokenIds is a free data retrieval call binding the contract method 0x00d2b30f.
//
// Solidity: function getCpuTokenIds() view returns(uint256[])
func (_App *AppSession) GetCpuTokenIds() ([]*big.Int, error) {
	return _App.Contract.GetCpuTokenIds(&_App.CallOpts)
}

// GetCpuTokenIds is a free data retrieval call binding the contract method 0x00d2b30f.
//
// Solidity: function getCpuTokenIds() view returns(uint256[])
func (_App *AppCallerSession) GetCpuTokenIds() ([]*big.Int, error) {
	return _App.Contract.GetCpuTokenIds(&_App.CallOpts)
}

// GetCpu is a free data retrieval call binding the contract method 0x27c3554c.
//
// Solidity: function getCpu_() view returns(string[])
func (_App *AppCaller) GetCpu(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getCpu_")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetCpu is a free data retrieval call binding the contract method 0x27c3554c.
//
// Solidity: function getCpu_() view returns(string[])
func (_App *AppSession) GetCpu() ([]string, error) {
	return _App.Contract.GetCpu(&_App.CallOpts)
}

// GetCpu is a free data retrieval call binding the contract method 0x27c3554c.
//
// Solidity: function getCpu_() view returns(string[])
func (_App *AppCallerSession) GetCpu() ([]string, error) {
	return _App.Contract.GetCpu(&_App.CallOpts)
}

// GetGpuC is a free data retrieval call binding the contract method 0x22895c1d.
//
// Solidity: function getGpuC_() view returns(uint256)
func (_App *AppCaller) GetGpuC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getGpuC_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGpuC is a free data retrieval call binding the contract method 0x22895c1d.
//
// Solidity: function getGpuC_() view returns(uint256)
func (_App *AppSession) GetGpuC() (*big.Int, error) {
	return _App.Contract.GetGpuC(&_App.CallOpts)
}

// GetGpuC is a free data retrieval call binding the contract method 0x22895c1d.
//
// Solidity: function getGpuC_() view returns(uint256)
func (_App *AppCallerSession) GetGpuC() (*big.Int, error) {
	return _App.Contract.GetGpuC(&_App.CallOpts)
}

// GetGpuMem is a free data retrieval call binding the contract method 0xa87b8d06.
//
// Solidity: function getGpuMem_() view returns(string[])
func (_App *AppCaller) GetGpuMem(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getGpuMem_")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetGpuMem is a free data retrieval call binding the contract method 0xa87b8d06.
//
// Solidity: function getGpuMem_() view returns(string[])
func (_App *AppSession) GetGpuMem() ([]string, error) {
	return _App.Contract.GetGpuMem(&_App.CallOpts)
}

// GetGpuMem is a free data retrieval call binding the contract method 0xa87b8d06.
//
// Solidity: function getGpuMem_() view returns(string[])
func (_App *AppCallerSession) GetGpuMem() ([]string, error) {
	return _App.Contract.GetGpuMem(&_App.CallOpts)
}

// GetGpuTokenIds is a free data retrieval call binding the contract method 0x9b8bd3ba.
//
// Solidity: function getGpuTokenIds() view returns(uint256[])
func (_App *AppCaller) GetGpuTokenIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getGpuTokenIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetGpuTokenIds is a free data retrieval call binding the contract method 0x9b8bd3ba.
//
// Solidity: function getGpuTokenIds() view returns(uint256[])
func (_App *AppSession) GetGpuTokenIds() ([]*big.Int, error) {
	return _App.Contract.GetGpuTokenIds(&_App.CallOpts)
}

// GetGpuTokenIds is a free data retrieval call binding the contract method 0x9b8bd3ba.
//
// Solidity: function getGpuTokenIds() view returns(uint256[])
func (_App *AppCallerSession) GetGpuTokenIds() ([]*big.Int, error) {
	return _App.Contract.GetGpuTokenIds(&_App.CallOpts)
}

// GetImage is a free data retrieval call binding the contract method 0xf3651cbb.
//
// Solidity: function getImage() view returns(string[])
func (_App *AppCaller) GetImage(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getImage")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetImage is a free data retrieval call binding the contract method 0xf3651cbb.
//
// Solidity: function getImage() view returns(string[])
func (_App *AppSession) GetImage() ([]string, error) {
	return _App.Contract.GetImage(&_App.CallOpts)
}

// GetImage is a free data retrieval call binding the contract method 0xf3651cbb.
//
// Solidity: function getImage() view returns(string[])
func (_App *AppCallerSession) GetImage() ([]string, error) {
	return _App.Contract.GetImage(&_App.CallOpts)
}

// GetIpC is a free data retrieval call binding the contract method 0xa12713e9.
//
// Solidity: function getIpC_() view returns(uint256)
func (_App *AppCaller) GetIpC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getIpC_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIpC is a free data retrieval call binding the contract method 0xa12713e9.
//
// Solidity: function getIpC_() view returns(uint256)
func (_App *AppSession) GetIpC() (*big.Int, error) {
	return _App.Contract.GetIpC(&_App.CallOpts)
}

// GetIpC is a free data retrieval call binding the contract method 0xa12713e9.
//
// Solidity: function getIpC_() view returns(uint256)
func (_App *AppCallerSession) GetIpC() (*big.Int, error) {
	return _App.Contract.GetIpC(&_App.CallOpts)
}

// GetIpTokenIds is a free data retrieval call binding the contract method 0x0ad1952b.
//
// Solidity: function getIpTokenIds() view returns(uint256[])
func (_App *AppCaller) GetIpTokenIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getIpTokenIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetIpTokenIds is a free data retrieval call binding the contract method 0x0ad1952b.
//
// Solidity: function getIpTokenIds() view returns(uint256[])
func (_App *AppSession) GetIpTokenIds() ([]*big.Int, error) {
	return _App.Contract.GetIpTokenIds(&_App.CallOpts)
}

// GetIpTokenIds is a free data retrieval call binding the contract method 0x0ad1952b.
//
// Solidity: function getIpTokenIds() view returns(uint256[])
func (_App *AppCallerSession) GetIpTokenIds() ([]*big.Int, error) {
	return _App.Contract.GetIpTokenIds(&_App.CallOpts)
}

// GetIp is a free data retrieval call binding the contract method 0x0fedafa8.
//
// Solidity: function getIp_() view returns(string[])
func (_App *AppCaller) GetIp(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getIp_")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetIp is a free data retrieval call binding the contract method 0x0fedafa8.
//
// Solidity: function getIp_() view returns(string[])
func (_App *AppSession) GetIp() ([]string, error) {
	return _App.Contract.GetIp(&_App.CallOpts)
}

// GetIp is a free data retrieval call binding the contract method 0x0fedafa8.
//
// Solidity: function getIp_() view returns(string[])
func (_App *AppCallerSession) GetIp() ([]string, error) {
	return _App.Contract.GetIp(&_App.CallOpts)
}

// GetIsExpanded is a free data retrieval call binding the contract method 0x09599a5e.
//
// Solidity: function getIsExpanded() view returns(bool)
func (_App *AppCaller) GetIsExpanded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getIsExpanded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsExpanded is a free data retrieval call binding the contract method 0x09599a5e.
//
// Solidity: function getIsExpanded() view returns(bool)
func (_App *AppSession) GetIsExpanded() (bool, error) {
	return _App.Contract.GetIsExpanded(&_App.CallOpts)
}

// GetIsExpanded is a free data retrieval call binding the contract method 0x09599a5e.
//
// Solidity: function getIsExpanded() view returns(bool)
func (_App *AppCallerSession) GetIsExpanded() (bool, error) {
	return _App.Contract.GetIsExpanded(&_App.CallOpts)
}

// GetIsTemplate is a free data retrieval call binding the contract method 0xf678e6fc.
//
// Solidity: function getIsTemplate() view returns(bool)
func (_App *AppCaller) GetIsTemplate(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getIsTemplate")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsTemplate is a free data retrieval call binding the contract method 0xf678e6fc.
//
// Solidity: function getIsTemplate() view returns(bool)
func (_App *AppSession) GetIsTemplate() (bool, error) {
	return _App.Contract.GetIsTemplate(&_App.CallOpts)
}

// GetIsTemplate is a free data retrieval call binding the contract method 0xf678e6fc.
//
// Solidity: function getIsTemplate() view returns(bool)
func (_App *AppCallerSession) GetIsTemplate() (bool, error) {
	return _App.Contract.GetIsTemplate(&_App.CallOpts)
}

// GetMemC is a free data retrieval call binding the contract method 0xe7919519.
//
// Solidity: function getMemC_() view returns(uint256)
func (_App *AppCaller) GetMemC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getMemC_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMemC is a free data retrieval call binding the contract method 0xe7919519.
//
// Solidity: function getMemC_() view returns(uint256)
func (_App *AppSession) GetMemC() (*big.Int, error) {
	return _App.Contract.GetMemC(&_App.CallOpts)
}

// GetMemC is a free data retrieval call binding the contract method 0xe7919519.
//
// Solidity: function getMemC_() view returns(uint256)
func (_App *AppCallerSession) GetMemC() (*big.Int, error) {
	return _App.Contract.GetMemC(&_App.CallOpts)
}

// GetMemTokenIds is a free data retrieval call binding the contract method 0x4a98bd58.
//
// Solidity: function getMemTokenIds() view returns(uint256[])
func (_App *AppCaller) GetMemTokenIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getMemTokenIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetMemTokenIds is a free data retrieval call binding the contract method 0x4a98bd58.
//
// Solidity: function getMemTokenIds() view returns(uint256[])
func (_App *AppSession) GetMemTokenIds() ([]*big.Int, error) {
	return _App.Contract.GetMemTokenIds(&_App.CallOpts)
}

// GetMemTokenIds is a free data retrieval call binding the contract method 0x4a98bd58.
//
// Solidity: function getMemTokenIds() view returns(uint256[])
func (_App *AppCallerSession) GetMemTokenIds() ([]*big.Int, error) {
	return _App.Contract.GetMemTokenIds(&_App.CallOpts)
}

// GetMem is a free data retrieval call binding the contract method 0xefe8e34e.
//
// Solidity: function getMem_() view returns(string[])
func (_App *AppCaller) GetMem(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getMem_")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetMem is a free data retrieval call binding the contract method 0xefe8e34e.
//
// Solidity: function getMem_() view returns(string[])
func (_App *AppSession) GetMem() ([]string, error) {
	return _App.Contract.GetMem(&_App.CallOpts)
}

// GetMem is a free data retrieval call binding the contract method 0xefe8e34e.
//
// Solidity: function getMem_() view returns(string[])
func (_App *AppCallerSession) GetMem() ([]string, error) {
	return _App.Contract.GetMem(&_App.CallOpts)
}

// GetMyOwner is a free data retrieval call binding the contract method 0x208f1f99.
//
// Solidity: function getMyOwner() view returns(address)
func (_App *AppCaller) GetMyOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getMyOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMyOwner is a free data retrieval call binding the contract method 0x208f1f99.
//
// Solidity: function getMyOwner() view returns(address)
func (_App *AppSession) GetMyOwner() (common.Address, error) {
	return _App.Contract.GetMyOwner(&_App.CallOpts)
}

// GetMyOwner is a free data retrieval call binding the contract method 0x208f1f99.
//
// Solidity: function getMyOwner() view returns(address)
func (_App *AppCallerSession) GetMyOwner() (common.Address, error) {
	return _App.Contract.GetMyOwner(&_App.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_App *AppCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_App *AppSession) GetName() (string, error) {
	return _App.Contract.GetName(&_App.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_App *AppCallerSession) GetName() (string, error) {
	return _App.Contract.GetName(&_App.CallOpts)
}

// GetNodeAddress is a free data retrieval call binding the contract method 0x70dabc9e.
//
// Solidity: function getNodeAddress() view returns(address[])
func (_App *AppCaller) GetNodeAddress(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getNodeAddress")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetNodeAddress is a free data retrieval call binding the contract method 0x70dabc9e.
//
// Solidity: function getNodeAddress() view returns(address[])
func (_App *AppSession) GetNodeAddress() ([]common.Address, error) {
	return _App.Contract.GetNodeAddress(&_App.CallOpts)
}

// GetNodeAddress is a free data retrieval call binding the contract method 0x70dabc9e.
//
// Solidity: function getNodeAddress() view returns(address[])
func (_App *AppCallerSession) GetNodeAddress() ([]common.Address, error) {
	return _App.Contract.GetNodeAddress(&_App.CallOpts)
}

// GetStorageC is a free data retrieval call binding the contract method 0x318d4119.
//
// Solidity: function getStorageC_() view returns(uint256)
func (_App *AppCaller) GetStorageC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getStorageC_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStorageC is a free data retrieval call binding the contract method 0x318d4119.
//
// Solidity: function getStorageC_() view returns(uint256)
func (_App *AppSession) GetStorageC() (*big.Int, error) {
	return _App.Contract.GetStorageC(&_App.CallOpts)
}

// GetStorageC is a free data retrieval call binding the contract method 0x318d4119.
//
// Solidity: function getStorageC_() view returns(uint256)
func (_App *AppCallerSession) GetStorageC() (*big.Int, error) {
	return _App.Contract.GetStorageC(&_App.CallOpts)
}

// GetStorageTokenIds is a free data retrieval call binding the contract method 0x8b9983c6.
//
// Solidity: function getStorageTokenIds() view returns(uint256[])
func (_App *AppCaller) GetStorageTokenIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getStorageTokenIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetStorageTokenIds is a free data retrieval call binding the contract method 0x8b9983c6.
//
// Solidity: function getStorageTokenIds() view returns(uint256[])
func (_App *AppSession) GetStorageTokenIds() ([]*big.Int, error) {
	return _App.Contract.GetStorageTokenIds(&_App.CallOpts)
}

// GetStorageTokenIds is a free data retrieval call binding the contract method 0x8b9983c6.
//
// Solidity: function getStorageTokenIds() view returns(uint256[])
func (_App *AppCallerSession) GetStorageTokenIds() ([]*big.Int, error) {
	return _App.Contract.GetStorageTokenIds(&_App.CallOpts)
}

// GetStorage is a free data retrieval call binding the contract method 0x2b281e47.
//
// Solidity: function getStorage_() view returns(string[])
func (_App *AppCaller) GetStorage(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getStorage_")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetStorage is a free data retrieval call binding the contract method 0x2b281e47.
//
// Solidity: function getStorage_() view returns(string[])
func (_App *AppSession) GetStorage() ([]string, error) {
	return _App.Contract.GetStorage(&_App.CallOpts)
}

// GetStorage is a free data retrieval call binding the contract method 0x2b281e47.
//
// Solidity: function getStorage_() view returns(string[])
func (_App *AppCallerSession) GetStorage() ([]string, error) {
	return _App.Contract.GetStorage(&_App.CallOpts)
}

// GetTemplates is a free data retrieval call binding the contract method 0xe64ae119.
//
// Solidity: function getTemplates() view returns(address[])
func (_App *AppCaller) GetTemplates(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getTemplates")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTemplates is a free data retrieval call binding the contract method 0xe64ae119.
//
// Solidity: function getTemplates() view returns(address[])
func (_App *AppSession) GetTemplates() ([]common.Address, error) {
	return _App.Contract.GetTemplates(&_App.CallOpts)
}

// GetTemplates is a free data retrieval call binding the contract method 0xe64ae119.
//
// Solidity: function getTemplates() view returns(address[])
func (_App *AppCallerSession) GetTemplates() ([]common.Address, error) {
	return _App.Contract.GetTemplates(&_App.CallOpts)
}

// GetTokenIds is a free data retrieval call binding the contract method 0x67f718a9.
//
// Solidity: function getTokenIds() view returns(uint256[][][])
func (_App *AppCaller) GetTokenIds(opts *bind.CallOpts) ([][][]*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getTokenIds")

	if err != nil {
		return *new([][][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][][]*big.Int)).(*[][][]*big.Int)

	return out0, err

}

// GetTokenIds is a free data retrieval call binding the contract method 0x67f718a9.
//
// Solidity: function getTokenIds() view returns(uint256[][][])
func (_App *AppSession) GetTokenIds() ([][][]*big.Int, error) {
	return _App.Contract.GetTokenIds(&_App.CallOpts)
}

// GetTokenIds is a free data retrieval call binding the contract method 0x67f718a9.
//
// Solidity: function getTokenIds() view returns(uint256[][][])
func (_App *AppCallerSession) GetTokenIds() ([][][]*big.Int, error) {
	return _App.Contract.GetTokenIds(&_App.CallOpts)
}

// GetTokenNeed is a free data retrieval call binding the contract method 0x2add2fa1.
//
// Solidity: function getTokenNeed() view returns(uint256[][])
func (_App *AppCaller) GetTokenNeed(opts *bind.CallOpts) ([][]*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getTokenNeed")

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetTokenNeed is a free data retrieval call binding the contract method 0x2add2fa1.
//
// Solidity: function getTokenNeed() view returns(uint256[][])
func (_App *AppSession) GetTokenNeed() ([][]*big.Int, error) {
	return _App.Contract.GetTokenNeed(&_App.CallOpts)
}

// GetTokenNeed is a free data retrieval call binding the contract method 0x2add2fa1.
//
// Solidity: function getTokenNeed() view returns(uint256[][])
func (_App *AppCallerSession) GetTokenNeed() ([][]*big.Int, error) {
	return _App.Contract.GetTokenNeed(&_App.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_App *AppCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_App *AppSession) Owner() (common.Address, error) {
	return _App.Contract.Owner(&_App.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_App *AppCallerSession) Owner() (common.Address, error) {
	return _App.Contract.Owner(&_App.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_App *AppTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_App *AppSession) RenounceOwnership() (*types.Transaction, error) {
	return _App.Contract.RenounceOwnership(&_App.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_App *AppTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _App.Contract.RenounceOwnership(&_App.TransactOpts)
}

// SetIsTemplate is a paid mutator transaction binding the contract method 0x1861a7ae.
//
// Solidity: function setIsTemplate(bool _isTemplate) returns()
func (_App *AppTransactor) SetIsTemplate(opts *bind.TransactOpts, _isTemplate bool) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "setIsTemplate", _isTemplate)
}

// SetIsTemplate is a paid mutator transaction binding the contract method 0x1861a7ae.
//
// Solidity: function setIsTemplate(bool _isTemplate) returns()
func (_App *AppSession) SetIsTemplate(_isTemplate bool) (*types.Transaction, error) {
	return _App.Contract.SetIsTemplate(&_App.TransactOpts, _isTemplate)
}

// SetIsTemplate is a paid mutator transaction binding the contract method 0x1861a7ae.
//
// Solidity: function setIsTemplate(bool _isTemplate) returns()
func (_App *AppTransactorSession) SetIsTemplate(_isTemplate bool) (*types.Transaction, error) {
	return _App.Contract.SetIsTemplate(&_App.TransactOpts, _isTemplate)
}

// SetNodeAddressAndTokenInfo is a paid mutator transaction binding the contract method 0x2d839d41.
//
// Solidity: function setNodeAddressAndTokenInfo(address[] _nodeAddress, uint256[][][] _tokenIds) returns()
func (_App *AppTransactor) SetNodeAddressAndTokenInfo(opts *bind.TransactOpts, _nodeAddress []common.Address, _tokenIds [][][]*big.Int) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "setNodeAddressAndTokenInfo", _nodeAddress, _tokenIds)
}

// SetNodeAddressAndTokenInfo is a paid mutator transaction binding the contract method 0x2d839d41.
//
// Solidity: function setNodeAddressAndTokenInfo(address[] _nodeAddress, uint256[][][] _tokenIds) returns()
func (_App *AppSession) SetNodeAddressAndTokenInfo(_nodeAddress []common.Address, _tokenIds [][][]*big.Int) (*types.Transaction, error) {
	return _App.Contract.SetNodeAddressAndTokenInfo(&_App.TransactOpts, _nodeAddress, _tokenIds)
}

// SetNodeAddressAndTokenInfo is a paid mutator transaction binding the contract method 0x2d839d41.
//
// Solidity: function setNodeAddressAndTokenInfo(address[] _nodeAddress, uint256[][][] _tokenIds) returns()
func (_App *AppTransactorSession) SetNodeAddressAndTokenInfo(_nodeAddress []common.Address, _tokenIds [][][]*big.Int) (*types.Transaction, error) {
	return _App.Contract.SetNodeAddressAndTokenInfo(&_App.TransactOpts, _nodeAddress, _tokenIds)
}

// SetTradeAddress is a paid mutator transaction binding the contract method 0x21a9cf34.
//
// Solidity: function setTradeAddress(address _tradeAddress) returns()
func (_App *AppTransactor) SetTradeAddress(opts *bind.TransactOpts, _tradeAddress common.Address) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "setTradeAddress", _tradeAddress)
}

// SetTradeAddress is a paid mutator transaction binding the contract method 0x21a9cf34.
//
// Solidity: function setTradeAddress(address _tradeAddress) returns()
func (_App *AppSession) SetTradeAddress(_tradeAddress common.Address) (*types.Transaction, error) {
	return _App.Contract.SetTradeAddress(&_App.TransactOpts, _tradeAddress)
}

// SetTradeAddress is a paid mutator transaction binding the contract method 0x21a9cf34.
//
// Solidity: function setTradeAddress(address _tradeAddress) returns()
func (_App *AppTransactorSession) SetTradeAddress(_tradeAddress common.Address) (*types.Transaction, error) {
	return _App.Contract.SetTradeAddress(&_App.TransactOpts, _tradeAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_App *AppTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_App *AppSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _App.Contract.TransferOwnership(&_App.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_App *AppTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _App.Contract.TransferOwnership(&_App.TransactOpts, newOwner)
}

// AppOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the App contract.
type AppOwnershipTransferredIterator struct {
	Event *AppOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AppOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppOwnershipTransferred)
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
		it.Event = new(AppOwnershipTransferred)
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
func (it *AppOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppOwnershipTransferred represents a OwnershipTransferred event raised by the App contract.
type AppOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_App *AppFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AppOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _App.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AppOwnershipTransferredIterator{contract: _App.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_App *AppFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AppOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _App.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppOwnershipTransferred)
				if err := _App.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_App *AppFilterer) ParseOwnershipTransferred(log types.Log) (*AppOwnershipTransferred, error) {
	event := new(AppOwnershipTransferred)
	if err := _App.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
