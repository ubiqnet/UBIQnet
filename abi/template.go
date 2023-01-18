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

// TemplateMetaData contains all meta data concerning the Template contract.
var TemplateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_templateType\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"_info\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_infoC\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_image\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bandwidthC_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bandwidth_\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cpuC_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cpu_\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsTemplate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenInfo\",\"outputs\":[{\"internalType\":\"uint256[6]\",\"name\":\"\",\"type\":\"uint256[6]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gpuC_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gpu_\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"image\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ipC_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ip_\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"memC_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mem_\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isTemplate\",\"type\":\"bool\"}],\"name\":\"setIsTemplate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageC_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storage_\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"templateType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TemplateABI is the input ABI used to generate the binding from.
// Deprecated: Use TemplateMetaData.ABI instead.
var TemplateABI = TemplateMetaData.ABI

// Template is an auto generated Go binding around an Ethereum contract.
type Template struct {
	TemplateCaller     // Read-only binding to the contract
	TemplateTransactor // Write-only binding to the contract
	TemplateFilterer   // Log filterer for contract events
}

// TemplateCaller is an auto generated read-only Go binding around an Ethereum contract.
type TemplateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TemplateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TemplateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TemplateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TemplateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TemplateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TemplateSession struct {
	Contract     *Template         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TemplateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TemplateCallerSession struct {
	Contract *TemplateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TemplateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TemplateTransactorSession struct {
	Contract     *TemplateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TemplateRaw is an auto generated low-level Go binding around an Ethereum contract.
type TemplateRaw struct {
	Contract *Template // Generic contract binding to access the raw methods on
}

// TemplateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TemplateCallerRaw struct {
	Contract *TemplateCaller // Generic read-only contract binding to access the raw methods on
}

// TemplateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TemplateTransactorRaw struct {
	Contract *TemplateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTemplate creates a new instance of Template, bound to a specific deployed contract.
func NewTemplate(address common.Address, backend bind.ContractBackend) (*Template, error) {
	contract, err := bindTemplate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Template{TemplateCaller: TemplateCaller{contract: contract}, TemplateTransactor: TemplateTransactor{contract: contract}, TemplateFilterer: TemplateFilterer{contract: contract}}, nil
}

// NewTemplateCaller creates a new read-only instance of Template, bound to a specific deployed contract.
func NewTemplateCaller(address common.Address, caller bind.ContractCaller) (*TemplateCaller, error) {
	contract, err := bindTemplate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TemplateCaller{contract: contract}, nil
}

// NewTemplateTransactor creates a new write-only instance of Template, bound to a specific deployed contract.
func NewTemplateTransactor(address common.Address, transactor bind.ContractTransactor) (*TemplateTransactor, error) {
	contract, err := bindTemplate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TemplateTransactor{contract: contract}, nil
}

// NewTemplateFilterer creates a new log filterer instance of Template, bound to a specific deployed contract.
func NewTemplateFilterer(address common.Address, filterer bind.ContractFilterer) (*TemplateFilterer, error) {
	contract, err := bindTemplate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TemplateFilterer{contract: contract}, nil
}

// bindTemplate binds a generic wrapper to an already deployed contract.
func bindTemplate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TemplateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Template *TemplateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Template.Contract.TemplateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Template *TemplateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Template.Contract.TemplateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Template *TemplateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Template.Contract.TemplateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Template *TemplateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Template.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Template *TemplateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Template.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Template *TemplateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Template.Contract.contract.Transact(opts, method, params...)
}

// BandwidthC is a free data retrieval call binding the contract method 0xdbe27f5a.
//
// Solidity: function bandwidthC_() view returns(uint256)
func (_Template *TemplateCaller) BandwidthC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "bandwidthC_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BandwidthC is a free data retrieval call binding the contract method 0xdbe27f5a.
//
// Solidity: function bandwidthC_() view returns(uint256)
func (_Template *TemplateSession) BandwidthC() (*big.Int, error) {
	return _Template.Contract.BandwidthC(&_Template.CallOpts)
}

// BandwidthC is a free data retrieval call binding the contract method 0xdbe27f5a.
//
// Solidity: function bandwidthC_() view returns(uint256)
func (_Template *TemplateCallerSession) BandwidthC() (*big.Int, error) {
	return _Template.Contract.BandwidthC(&_Template.CallOpts)
}

// Bandwidth is a free data retrieval call binding the contract method 0xce3a71bd.
//
// Solidity: function bandwidth_() view returns(string)
func (_Template *TemplateCaller) Bandwidth(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "bandwidth_")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bandwidth is a free data retrieval call binding the contract method 0xce3a71bd.
//
// Solidity: function bandwidth_() view returns(string)
func (_Template *TemplateSession) Bandwidth() (string, error) {
	return _Template.Contract.Bandwidth(&_Template.CallOpts)
}

// Bandwidth is a free data retrieval call binding the contract method 0xce3a71bd.
//
// Solidity: function bandwidth_() view returns(string)
func (_Template *TemplateCallerSession) Bandwidth() (string, error) {
	return _Template.Contract.Bandwidth(&_Template.CallOpts)
}

// CpuC is a free data retrieval call binding the contract method 0x16a2fc90.
//
// Solidity: function cpuC_() view returns(uint256)
func (_Template *TemplateCaller) CpuC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "cpuC_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CpuC is a free data retrieval call binding the contract method 0x16a2fc90.
//
// Solidity: function cpuC_() view returns(uint256)
func (_Template *TemplateSession) CpuC() (*big.Int, error) {
	return _Template.Contract.CpuC(&_Template.CallOpts)
}

// CpuC is a free data retrieval call binding the contract method 0x16a2fc90.
//
// Solidity: function cpuC_() view returns(uint256)
func (_Template *TemplateCallerSession) CpuC() (*big.Int, error) {
	return _Template.Contract.CpuC(&_Template.CallOpts)
}

// Cpu is a free data retrieval call binding the contract method 0x47995a25.
//
// Solidity: function cpu_() view returns(string)
func (_Template *TemplateCaller) Cpu(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "cpu_")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Cpu is a free data retrieval call binding the contract method 0x47995a25.
//
// Solidity: function cpu_() view returns(string)
func (_Template *TemplateSession) Cpu() (string, error) {
	return _Template.Contract.Cpu(&_Template.CallOpts)
}

// Cpu is a free data retrieval call binding the contract method 0x47995a25.
//
// Solidity: function cpu_() view returns(string)
func (_Template *TemplateCallerSession) Cpu() (string, error) {
	return _Template.Contract.Cpu(&_Template.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() view returns(uint256)
func (_Template *TemplateCaller) CreateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "createTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() view returns(uint256)
func (_Template *TemplateSession) CreateTime() (*big.Int, error) {
	return _Template.Contract.CreateTime(&_Template.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() view returns(uint256)
func (_Template *TemplateCallerSession) CreateTime() (*big.Int, error) {
	return _Template.Contract.CreateTime(&_Template.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Template *TemplateCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Template *TemplateSession) Creator() (common.Address, error) {
	return _Template.Contract.Creator(&_Template.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Template *TemplateCallerSession) Creator() (common.Address, error) {
	return _Template.Contract.Creator(&_Template.CallOpts)
}

// GetIsTemplate is a free data retrieval call binding the contract method 0xf678e6fc.
//
// Solidity: function getIsTemplate() view returns(bool)
func (_Template *TemplateCaller) GetIsTemplate(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "getIsTemplate")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsTemplate is a free data retrieval call binding the contract method 0xf678e6fc.
//
// Solidity: function getIsTemplate() view returns(bool)
func (_Template *TemplateSession) GetIsTemplate() (bool, error) {
	return _Template.Contract.GetIsTemplate(&_Template.CallOpts)
}

// GetIsTemplate is a free data retrieval call binding the contract method 0xf678e6fc.
//
// Solidity: function getIsTemplate() view returns(bool)
func (_Template *TemplateCallerSession) GetIsTemplate() (bool, error) {
	return _Template.Contract.GetIsTemplate(&_Template.CallOpts)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0xabb1dc44.
//
// Solidity: function getTokenInfo() view returns(uint256[6])
func (_Template *TemplateCaller) GetTokenInfo(opts *bind.CallOpts) ([6]*big.Int, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "getTokenInfo")

	if err != nil {
		return *new([6]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([6]*big.Int)).(*[6]*big.Int)

	return out0, err

}

// GetTokenInfo is a free data retrieval call binding the contract method 0xabb1dc44.
//
// Solidity: function getTokenInfo() view returns(uint256[6])
func (_Template *TemplateSession) GetTokenInfo() ([6]*big.Int, error) {
	return _Template.Contract.GetTokenInfo(&_Template.CallOpts)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0xabb1dc44.
//
// Solidity: function getTokenInfo() view returns(uint256[6])
func (_Template *TemplateCallerSession) GetTokenInfo() ([6]*big.Int, error) {
	return _Template.Contract.GetTokenInfo(&_Template.CallOpts)
}

// GpuC is a free data retrieval call binding the contract method 0xf54da5b0.
//
// Solidity: function gpuC_() view returns(uint256)
func (_Template *TemplateCaller) GpuC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "gpuC_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GpuC is a free data retrieval call binding the contract method 0xf54da5b0.
//
// Solidity: function gpuC_() view returns(uint256)
func (_Template *TemplateSession) GpuC() (*big.Int, error) {
	return _Template.Contract.GpuC(&_Template.CallOpts)
}

// GpuC is a free data retrieval call binding the contract method 0xf54da5b0.
//
// Solidity: function gpuC_() view returns(uint256)
func (_Template *TemplateCallerSession) GpuC() (*big.Int, error) {
	return _Template.Contract.GpuC(&_Template.CallOpts)
}

// Gpu is a free data retrieval call binding the contract method 0x33147641.
//
// Solidity: function gpu_() view returns(string)
func (_Template *TemplateCaller) Gpu(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "gpu_")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Gpu is a free data retrieval call binding the contract method 0x33147641.
//
// Solidity: function gpu_() view returns(string)
func (_Template *TemplateSession) Gpu() (string, error) {
	return _Template.Contract.Gpu(&_Template.CallOpts)
}

// Gpu is a free data retrieval call binding the contract method 0x33147641.
//
// Solidity: function gpu_() view returns(string)
func (_Template *TemplateCallerSession) Gpu() (string, error) {
	return _Template.Contract.Gpu(&_Template.CallOpts)
}

// Image is a free data retrieval call binding the contract method 0xf3ccaac0.
//
// Solidity: function image() view returns(string)
func (_Template *TemplateCaller) Image(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "image")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Image is a free data retrieval call binding the contract method 0xf3ccaac0.
//
// Solidity: function image() view returns(string)
func (_Template *TemplateSession) Image() (string, error) {
	return _Template.Contract.Image(&_Template.CallOpts)
}

// Image is a free data retrieval call binding the contract method 0xf3ccaac0.
//
// Solidity: function image() view returns(string)
func (_Template *TemplateCallerSession) Image() (string, error) {
	return _Template.Contract.Image(&_Template.CallOpts)
}

// IpC is a free data retrieval call binding the contract method 0x245032d0.
//
// Solidity: function ipC_() view returns(uint256)
func (_Template *TemplateCaller) IpC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "ipC_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IpC is a free data retrieval call binding the contract method 0x245032d0.
//
// Solidity: function ipC_() view returns(uint256)
func (_Template *TemplateSession) IpC() (*big.Int, error) {
	return _Template.Contract.IpC(&_Template.CallOpts)
}

// IpC is a free data retrieval call binding the contract method 0x245032d0.
//
// Solidity: function ipC_() view returns(uint256)
func (_Template *TemplateCallerSession) IpC() (*big.Int, error) {
	return _Template.Contract.IpC(&_Template.CallOpts)
}

// Ip is a free data retrieval call binding the contract method 0x185d4a58.
//
// Solidity: function ip_() view returns(string)
func (_Template *TemplateCaller) Ip(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "ip_")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Ip is a free data retrieval call binding the contract method 0x185d4a58.
//
// Solidity: function ip_() view returns(string)
func (_Template *TemplateSession) Ip() (string, error) {
	return _Template.Contract.Ip(&_Template.CallOpts)
}

// Ip is a free data retrieval call binding the contract method 0x185d4a58.
//
// Solidity: function ip_() view returns(string)
func (_Template *TemplateCallerSession) Ip() (string, error) {
	return _Template.Contract.Ip(&_Template.CallOpts)
}

// MemC is a free data retrieval call binding the contract method 0x13ccf0ad.
//
// Solidity: function memC_() view returns(uint256)
func (_Template *TemplateCaller) MemC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "memC_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MemC is a free data retrieval call binding the contract method 0x13ccf0ad.
//
// Solidity: function memC_() view returns(uint256)
func (_Template *TemplateSession) MemC() (*big.Int, error) {
	return _Template.Contract.MemC(&_Template.CallOpts)
}

// MemC is a free data retrieval call binding the contract method 0x13ccf0ad.
//
// Solidity: function memC_() view returns(uint256)
func (_Template *TemplateCallerSession) MemC() (*big.Int, error) {
	return _Template.Contract.MemC(&_Template.CallOpts)
}

// Mem is a free data retrieval call binding the contract method 0x9335d467.
//
// Solidity: function mem_() view returns(string)
func (_Template *TemplateCaller) Mem(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "mem_")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Mem is a free data retrieval call binding the contract method 0x9335d467.
//
// Solidity: function mem_() view returns(string)
func (_Template *TemplateSession) Mem() (string, error) {
	return _Template.Contract.Mem(&_Template.CallOpts)
}

// Mem is a free data retrieval call binding the contract method 0x9335d467.
//
// Solidity: function mem_() view returns(string)
func (_Template *TemplateCallerSession) Mem() (string, error) {
	return _Template.Contract.Mem(&_Template.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Template *TemplateCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Template *TemplateSession) Name() (string, error) {
	return _Template.Contract.Name(&_Template.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Template *TemplateCallerSession) Name() (string, error) {
	return _Template.Contract.Name(&_Template.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Template *TemplateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Template *TemplateSession) Owner() (common.Address, error) {
	return _Template.Contract.Owner(&_Template.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Template *TemplateCallerSession) Owner() (common.Address, error) {
	return _Template.Contract.Owner(&_Template.CallOpts)
}

// StorageC is a free data retrieval call binding the contract method 0xb9c61981.
//
// Solidity: function storageC_() view returns(uint256)
func (_Template *TemplateCaller) StorageC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "storageC_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StorageC is a free data retrieval call binding the contract method 0xb9c61981.
//
// Solidity: function storageC_() view returns(uint256)
func (_Template *TemplateSession) StorageC() (*big.Int, error) {
	return _Template.Contract.StorageC(&_Template.CallOpts)
}

// StorageC is a free data retrieval call binding the contract method 0xb9c61981.
//
// Solidity: function storageC_() view returns(uint256)
func (_Template *TemplateCallerSession) StorageC() (*big.Int, error) {
	return _Template.Contract.StorageC(&_Template.CallOpts)
}

// Storage is a free data retrieval call binding the contract method 0xc57b13b2.
//
// Solidity: function storage_() view returns(string)
func (_Template *TemplateCaller) Storage(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "storage_")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Storage is a free data retrieval call binding the contract method 0xc57b13b2.
//
// Solidity: function storage_() view returns(string)
func (_Template *TemplateSession) Storage() (string, error) {
	return _Template.Contract.Storage(&_Template.CallOpts)
}

// Storage is a free data retrieval call binding the contract method 0xc57b13b2.
//
// Solidity: function storage_() view returns(string)
func (_Template *TemplateCallerSession) Storage() (string, error) {
	return _Template.Contract.Storage(&_Template.CallOpts)
}

// TemplateType is a free data retrieval call binding the contract method 0x5713a9c5.
//
// Solidity: function templateType() view returns(uint256)
func (_Template *TemplateCaller) TemplateType(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Template.contract.Call(opts, &out, "templateType")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TemplateType is a free data retrieval call binding the contract method 0x5713a9c5.
//
// Solidity: function templateType() view returns(uint256)
func (_Template *TemplateSession) TemplateType() (*big.Int, error) {
	return _Template.Contract.TemplateType(&_Template.CallOpts)
}

// TemplateType is a free data retrieval call binding the contract method 0x5713a9c5.
//
// Solidity: function templateType() view returns(uint256)
func (_Template *TemplateCallerSession) TemplateType() (*big.Int, error) {
	return _Template.Contract.TemplateType(&_Template.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Template *TemplateTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Template.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Template *TemplateSession) RenounceOwnership() (*types.Transaction, error) {
	return _Template.Contract.RenounceOwnership(&_Template.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Template *TemplateTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Template.Contract.RenounceOwnership(&_Template.TransactOpts)
}

// SetIsTemplate is a paid mutator transaction binding the contract method 0x1861a7ae.
//
// Solidity: function setIsTemplate(bool _isTemplate) returns()
func (_Template *TemplateTransactor) SetIsTemplate(opts *bind.TransactOpts, _isTemplate bool) (*types.Transaction, error) {
	return _Template.contract.Transact(opts, "setIsTemplate", _isTemplate)
}

// SetIsTemplate is a paid mutator transaction binding the contract method 0x1861a7ae.
//
// Solidity: function setIsTemplate(bool _isTemplate) returns()
func (_Template *TemplateSession) SetIsTemplate(_isTemplate bool) (*types.Transaction, error) {
	return _Template.Contract.SetIsTemplate(&_Template.TransactOpts, _isTemplate)
}

// SetIsTemplate is a paid mutator transaction binding the contract method 0x1861a7ae.
//
// Solidity: function setIsTemplate(bool _isTemplate) returns()
func (_Template *TemplateTransactorSession) SetIsTemplate(_isTemplate bool) (*types.Transaction, error) {
	return _Template.Contract.SetIsTemplate(&_Template.TransactOpts, _isTemplate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Template *TemplateTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Template.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Template *TemplateSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Template.Contract.TransferOwnership(&_Template.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Template *TemplateTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Template.Contract.TransferOwnership(&_Template.TransactOpts, newOwner)
}

// TemplateOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Template contract.
type TemplateOwnershipTransferredIterator struct {
	Event *TemplateOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TemplateOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TemplateOwnershipTransferred)
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
		it.Event = new(TemplateOwnershipTransferred)
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
func (it *TemplateOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TemplateOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TemplateOwnershipTransferred represents a OwnershipTransferred event raised by the Template contract.
type TemplateOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Template *TemplateFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TemplateOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Template.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TemplateOwnershipTransferredIterator{contract: _Template.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Template *TemplateFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TemplateOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Template.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TemplateOwnershipTransferred)
				if err := _Template.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Template *TemplateFilterer) ParseOwnershipTransferred(log types.Log) (*TemplateOwnershipTransferred, error) {
	event := new(TemplateOwnershipTransferred)
	if err := _Template.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
