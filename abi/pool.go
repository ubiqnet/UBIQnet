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

// MyStakeClientInfo is an auto generated low-level Go binding around an user-defined struct.
type MyStakeClientInfo struct {
	Ts        *big.Int
	Ip        *big.Int
	Bandwidth *big.Int
	Gpu       *big.Int
	Storage   *big.Int
	Cpu       *big.Int
	Mem       *big.Int
}

// MyStakeRewardInfo is an auto generated low-level Go binding around an user-defined struct.
type MyStakeRewardInfo struct {
	RewardType *big.Int
	RewardTime *big.Int
	Amount     *big.Int
}

// MyStakeStake is an auto generated low-level Go binding around an user-defined struct.
type MyStakeStake struct {
	DepositAmount           *big.Int
	StakeCreationTime       *big.Int
	LastUploadTime          *big.Int
	Returned                bool
	AlreadyWithdrawedAmount *big.Int
	AllUploadCount          *big.Int
	AllRewardAmount         *big.Int
	RewardAmount            *big.Int
	RewardUpload            *big.Int
	AllReturnAmount         *big.Int
	WalletAddr              common.Address
	IsUsed                  bool
}

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"factoryInterface\",\"outputs\":[{\"internalType\":\"contractMinerFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getAllInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deposit_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake_creation_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last_upload_time\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"returned\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"alreadyWithdrawedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AllUploadCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AllRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardUpload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AllReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isUsed\",\"type\":\"bool\"}],\"internalType\":\"structMyStake.Stake\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getAllReturnAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getAllRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getAllUploadCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllUser\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getClientInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ip\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bandwidth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_storage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mem\",\"type\":\"uint256\"}],\"internalType\":\"structMyStake.ClientInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGas\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReturnPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getRewardLog\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rewardType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structMyStake.RewardInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUploadCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUploadTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifyAddress\",\"type\":\"address\"}],\"name\":\"getVerifyTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isExistEntry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTokenSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"random\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"returnToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factoryAddress\",\"type\":\"address\"}],\"name\":\"setFactoryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ga\",\"type\":\"uint256\"}],\"name\":\"setGasAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rt\",\"type\":\"uint256\"}],\"name\":\"setReturnPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rp\",\"type\":\"uint256\"}],\"name\":\"setRewardPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_uploadCount\",\"type\":\"uint256\"}],\"name\":\"setUploadCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_uploadTime\",\"type\":\"uint256\"}],\"name\":\"setUploadTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"stakeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tsStr\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"sign\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"nodeAddress\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][][]\",\"name\":\"TokenIds\",\"type\":\"uint256[][][]\"},{\"internalType\":\"uint256[][][]\",\"name\":\"TokenPrice\",\"type\":\"uint256[][][]\"}],\"name\":\"uploadAllInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// FactoryInterface is a free data retrieval call binding the contract method 0x44c25ac2.
//
// Solidity: function factoryInterface() view returns(address)
func (_Pool *PoolCaller) FactoryInterface(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "factoryInterface")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryInterface is a free data retrieval call binding the contract method 0x44c25ac2.
//
// Solidity: function factoryInterface() view returns(address)
func (_Pool *PoolSession) FactoryInterface() (common.Address, error) {
	return _Pool.Contract.FactoryInterface(&_Pool.CallOpts)
}

// FactoryInterface is a free data retrieval call binding the contract method 0x44c25ac2.
//
// Solidity: function factoryInterface() view returns(address)
func (_Pool *PoolCallerSession) FactoryInterface() (common.Address, error) {
	return _Pool.Contract.FactoryInterface(&_Pool.CallOpts)
}

// GetAllInfo is a free data retrieval call binding the contract method 0x7e16c30c.
//
// Solidity: function getAllInfo(address _addr) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,bool))
func (_Pool *PoolCaller) GetAllInfo(opts *bind.CallOpts, _addr common.Address) (MyStakeStake, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getAllInfo", _addr)

	if err != nil {
		return *new(MyStakeStake), err
	}

	out0 := *abi.ConvertType(out[0], new(MyStakeStake)).(*MyStakeStake)

	return out0, err

}

// GetAllInfo is a free data retrieval call binding the contract method 0x7e16c30c.
//
// Solidity: function getAllInfo(address _addr) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,bool))
func (_Pool *PoolSession) GetAllInfo(_addr common.Address) (MyStakeStake, error) {
	return _Pool.Contract.GetAllInfo(&_Pool.CallOpts, _addr)
}

// GetAllInfo is a free data retrieval call binding the contract method 0x7e16c30c.
//
// Solidity: function getAllInfo(address _addr) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,bool))
func (_Pool *PoolCallerSession) GetAllInfo(_addr common.Address) (MyStakeStake, error) {
	return _Pool.Contract.GetAllInfo(&_Pool.CallOpts, _addr)
}

// GetAllReturnAmount is a free data retrieval call binding the contract method 0xe73bbb19.
//
// Solidity: function getAllReturnAmount(address _addr) view returns(uint256)
func (_Pool *PoolCaller) GetAllReturnAmount(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getAllReturnAmount", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllReturnAmount is a free data retrieval call binding the contract method 0xe73bbb19.
//
// Solidity: function getAllReturnAmount(address _addr) view returns(uint256)
func (_Pool *PoolSession) GetAllReturnAmount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetAllReturnAmount(&_Pool.CallOpts, _addr)
}

// GetAllReturnAmount is a free data retrieval call binding the contract method 0xe73bbb19.
//
// Solidity: function getAllReturnAmount(address _addr) view returns(uint256)
func (_Pool *PoolCallerSession) GetAllReturnAmount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetAllReturnAmount(&_Pool.CallOpts, _addr)
}

// GetAllRewardAmount is a free data retrieval call binding the contract method 0xdd0b7ebe.
//
// Solidity: function getAllRewardAmount(address _addr) view returns(uint256)
func (_Pool *PoolCaller) GetAllRewardAmount(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getAllRewardAmount", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllRewardAmount is a free data retrieval call binding the contract method 0xdd0b7ebe.
//
// Solidity: function getAllRewardAmount(address _addr) view returns(uint256)
func (_Pool *PoolSession) GetAllRewardAmount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetAllRewardAmount(&_Pool.CallOpts, _addr)
}

// GetAllRewardAmount is a free data retrieval call binding the contract method 0xdd0b7ebe.
//
// Solidity: function getAllRewardAmount(address _addr) view returns(uint256)
func (_Pool *PoolCallerSession) GetAllRewardAmount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetAllRewardAmount(&_Pool.CallOpts, _addr)
}

// GetAllUploadCount is a free data retrieval call binding the contract method 0xb3dce268.
//
// Solidity: function getAllUploadCount(address _addr) view returns(uint256)
func (_Pool *PoolCaller) GetAllUploadCount(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getAllUploadCount", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllUploadCount is a free data retrieval call binding the contract method 0xb3dce268.
//
// Solidity: function getAllUploadCount(address _addr) view returns(uint256)
func (_Pool *PoolSession) GetAllUploadCount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetAllUploadCount(&_Pool.CallOpts, _addr)
}

// GetAllUploadCount is a free data retrieval call binding the contract method 0xb3dce268.
//
// Solidity: function getAllUploadCount(address _addr) view returns(uint256)
func (_Pool *PoolCallerSession) GetAllUploadCount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetAllUploadCount(&_Pool.CallOpts, _addr)
}

// GetAllUser is a free data retrieval call binding the contract method 0x47531df8.
//
// Solidity: function getAllUser() view returns(address[])
func (_Pool *PoolCaller) GetAllUser(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getAllUser")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllUser is a free data retrieval call binding the contract method 0x47531df8.
//
// Solidity: function getAllUser() view returns(address[])
func (_Pool *PoolSession) GetAllUser() ([]common.Address, error) {
	return _Pool.Contract.GetAllUser(&_Pool.CallOpts)
}

// GetAllUser is a free data retrieval call binding the contract method 0x47531df8.
//
// Solidity: function getAllUser() view returns(address[])
func (_Pool *PoolCallerSession) GetAllUser() ([]common.Address, error) {
	return _Pool.Contract.GetAllUser(&_Pool.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Pool *PoolCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Pool *PoolSession) GetBalance() (*big.Int, error) {
	return _Pool.Contract.GetBalance(&_Pool.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Pool *PoolCallerSession) GetBalance() (*big.Int, error) {
	return _Pool.Contract.GetBalance(&_Pool.CallOpts)
}

// GetClientInfos is a free data retrieval call binding the contract method 0xac4ec70e.
//
// Solidity: function getClientInfos(address _addr) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Pool *PoolCaller) GetClientInfos(opts *bind.CallOpts, _addr common.Address) ([]MyStakeClientInfo, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getClientInfos", _addr)

	if err != nil {
		return *new([]MyStakeClientInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]MyStakeClientInfo)).(*[]MyStakeClientInfo)

	return out0, err

}

// GetClientInfos is a free data retrieval call binding the contract method 0xac4ec70e.
//
// Solidity: function getClientInfos(address _addr) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Pool *PoolSession) GetClientInfos(_addr common.Address) ([]MyStakeClientInfo, error) {
	return _Pool.Contract.GetClientInfos(&_Pool.CallOpts, _addr)
}

// GetClientInfos is a free data retrieval call binding the contract method 0xac4ec70e.
//
// Solidity: function getClientInfos(address _addr) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Pool *PoolCallerSession) GetClientInfos(_addr common.Address) ([]MyStakeClientInfo, error) {
	return _Pool.Contract.GetClientInfos(&_Pool.CallOpts, _addr)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address _addr) view returns(uint256)
func (_Pool *PoolCaller) GetDepositAmount(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getDepositAmount", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address _addr) view returns(uint256)
func (_Pool *PoolSession) GetDepositAmount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetDepositAmount(&_Pool.CallOpts, _addr)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address _addr) view returns(uint256)
func (_Pool *PoolCallerSession) GetDepositAmount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetDepositAmount(&_Pool.CallOpts, _addr)
}

// GetGasAmount is a free data retrieval call binding the contract method 0x8d8ea2c5.
//
// Solidity: function getGasAmount() view returns(uint256)
func (_Pool *PoolCaller) GetGasAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getGasAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasAmount is a free data retrieval call binding the contract method 0x8d8ea2c5.
//
// Solidity: function getGasAmount() view returns(uint256)
func (_Pool *PoolSession) GetGasAmount() (*big.Int, error) {
	return _Pool.Contract.GetGasAmount(&_Pool.CallOpts)
}

// GetGasAmount is a free data retrieval call binding the contract method 0x8d8ea2c5.
//
// Solidity: function getGasAmount() view returns(uint256)
func (_Pool *PoolCallerSession) GetGasAmount() (*big.Int, error) {
	return _Pool.Contract.GetGasAmount(&_Pool.CallOpts)
}

// GetReturnPercent is a free data retrieval call binding the contract method 0x8d49f7ec.
//
// Solidity: function getReturnPercent() view returns(uint256)
func (_Pool *PoolCaller) GetReturnPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getReturnPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReturnPercent is a free data retrieval call binding the contract method 0x8d49f7ec.
//
// Solidity: function getReturnPercent() view returns(uint256)
func (_Pool *PoolSession) GetReturnPercent() (*big.Int, error) {
	return _Pool.Contract.GetReturnPercent(&_Pool.CallOpts)
}

// GetReturnPercent is a free data retrieval call binding the contract method 0x8d49f7ec.
//
// Solidity: function getReturnPercent() view returns(uint256)
func (_Pool *PoolCallerSession) GetReturnPercent() (*big.Int, error) {
	return _Pool.Contract.GetReturnPercent(&_Pool.CallOpts)
}

// GetRewardAmount is a free data retrieval call binding the contract method 0x44a040f5.
//
// Solidity: function getRewardAmount(address _addr) view returns(uint256)
func (_Pool *PoolCaller) GetRewardAmount(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getRewardAmount", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardAmount is a free data retrieval call binding the contract method 0x44a040f5.
//
// Solidity: function getRewardAmount(address _addr) view returns(uint256)
func (_Pool *PoolSession) GetRewardAmount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetRewardAmount(&_Pool.CallOpts, _addr)
}

// GetRewardAmount is a free data retrieval call binding the contract method 0x44a040f5.
//
// Solidity: function getRewardAmount(address _addr) view returns(uint256)
func (_Pool *PoolCallerSession) GetRewardAmount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetRewardAmount(&_Pool.CallOpts, _addr)
}

// GetRewardLog is a free data retrieval call binding the contract method 0x68a2e8d3.
//
// Solidity: function getRewardLog(address nodeAddress) view returns((uint256,uint256,uint256)[])
func (_Pool *PoolCaller) GetRewardLog(opts *bind.CallOpts, nodeAddress common.Address) ([]MyStakeRewardInfo, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getRewardLog", nodeAddress)

	if err != nil {
		return *new([]MyStakeRewardInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]MyStakeRewardInfo)).(*[]MyStakeRewardInfo)

	return out0, err

}

// GetRewardLog is a free data retrieval call binding the contract method 0x68a2e8d3.
//
// Solidity: function getRewardLog(address nodeAddress) view returns((uint256,uint256,uint256)[])
func (_Pool *PoolSession) GetRewardLog(nodeAddress common.Address) ([]MyStakeRewardInfo, error) {
	return _Pool.Contract.GetRewardLog(&_Pool.CallOpts, nodeAddress)
}

// GetRewardLog is a free data retrieval call binding the contract method 0x68a2e8d3.
//
// Solidity: function getRewardLog(address nodeAddress) view returns((uint256,uint256,uint256)[])
func (_Pool *PoolCallerSession) GetRewardLog(nodeAddress common.Address) ([]MyStakeRewardInfo, error) {
	return _Pool.Contract.GetRewardLog(&_Pool.CallOpts, nodeAddress)
}

// GetRewardPercent is a free data retrieval call binding the contract method 0x0e8622d3.
//
// Solidity: function getRewardPercent() view returns(uint256)
func (_Pool *PoolCaller) GetRewardPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getRewardPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardPercent is a free data retrieval call binding the contract method 0x0e8622d3.
//
// Solidity: function getRewardPercent() view returns(uint256)
func (_Pool *PoolSession) GetRewardPercent() (*big.Int, error) {
	return _Pool.Contract.GetRewardPercent(&_Pool.CallOpts)
}

// GetRewardPercent is a free data retrieval call binding the contract method 0x0e8622d3.
//
// Solidity: function getRewardPercent() view returns(uint256)
func (_Pool *PoolCallerSession) GetRewardPercent() (*big.Int, error) {
	return _Pool.Contract.GetRewardPercent(&_Pool.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Pool *PoolCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Pool *PoolSession) GetTokenAddress() (common.Address, error) {
	return _Pool.Contract.GetTokenAddress(&_Pool.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Pool *PoolCallerSession) GetTokenAddress() (common.Address, error) {
	return _Pool.Contract.GetTokenAddress(&_Pool.CallOpts)
}

// GetUploadCount is a free data retrieval call binding the contract method 0xca27cf04.
//
// Solidity: function getUploadCount() view returns(uint256)
func (_Pool *PoolCaller) GetUploadCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getUploadCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUploadCount is a free data retrieval call binding the contract method 0xca27cf04.
//
// Solidity: function getUploadCount() view returns(uint256)
func (_Pool *PoolSession) GetUploadCount() (*big.Int, error) {
	return _Pool.Contract.GetUploadCount(&_Pool.CallOpts)
}

// GetUploadCount is a free data retrieval call binding the contract method 0xca27cf04.
//
// Solidity: function getUploadCount() view returns(uint256)
func (_Pool *PoolCallerSession) GetUploadCount() (*big.Int, error) {
	return _Pool.Contract.GetUploadCount(&_Pool.CallOpts)
}

// GetUploadTime is a free data retrieval call binding the contract method 0xa0a58af7.
//
// Solidity: function getUploadTime() view returns(uint256)
func (_Pool *PoolCaller) GetUploadTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getUploadTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUploadTime is a free data retrieval call binding the contract method 0xa0a58af7.
//
// Solidity: function getUploadTime() view returns(uint256)
func (_Pool *PoolSession) GetUploadTime() (*big.Int, error) {
	return _Pool.Contract.GetUploadTime(&_Pool.CallOpts)
}

// GetUploadTime is a free data retrieval call binding the contract method 0xa0a58af7.
//
// Solidity: function getUploadTime() view returns(uint256)
func (_Pool *PoolCallerSession) GetUploadTime() (*big.Int, error) {
	return _Pool.Contract.GetUploadTime(&_Pool.CallOpts)
}

// GetVerifyTimestamp is a free data retrieval call binding the contract method 0xc61513e7.
//
// Solidity: function getVerifyTimestamp(address verifyAddress) view returns(uint256)
func (_Pool *PoolCaller) GetVerifyTimestamp(opts *bind.CallOpts, verifyAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getVerifyTimestamp", verifyAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVerifyTimestamp is a free data retrieval call binding the contract method 0xc61513e7.
//
// Solidity: function getVerifyTimestamp(address verifyAddress) view returns(uint256)
func (_Pool *PoolSession) GetVerifyTimestamp(verifyAddress common.Address) (*big.Int, error) {
	return _Pool.Contract.GetVerifyTimestamp(&_Pool.CallOpts, verifyAddress)
}

// GetVerifyTimestamp is a free data retrieval call binding the contract method 0xc61513e7.
//
// Solidity: function getVerifyTimestamp(address verifyAddress) view returns(uint256)
func (_Pool *PoolCallerSession) GetVerifyTimestamp(verifyAddress common.Address) (*big.Int, error) {
	return _Pool.Contract.GetVerifyTimestamp(&_Pool.CallOpts, verifyAddress)
}

// IsExistEntry is a free data retrieval call binding the contract method 0x5e2f1a35.
//
// Solidity: function isExistEntry(address _addr) view returns(bool)
func (_Pool *PoolCaller) IsExistEntry(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "isExistEntry", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistEntry is a free data retrieval call binding the contract method 0x5e2f1a35.
//
// Solidity: function isExistEntry(address _addr) view returns(bool)
func (_Pool *PoolSession) IsExistEntry(_addr common.Address) (bool, error) {
	return _Pool.Contract.IsExistEntry(&_Pool.CallOpts, _addr)
}

// IsExistEntry is a free data retrieval call binding the contract method 0x5e2f1a35.
//
// Solidity: function isExistEntry(address _addr) view returns(bool)
func (_Pool *PoolCallerSession) IsExistEntry(_addr common.Address) (bool, error) {
	return _Pool.Contract.IsExistEntry(&_Pool.CallOpts, _addr)
}

// IsTokenSet is a free data retrieval call binding the contract method 0x9ee7138c.
//
// Solidity: function isTokenSet() view returns(bool)
func (_Pool *PoolCaller) IsTokenSet(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "isTokenSet")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenSet is a free data retrieval call binding the contract method 0x9ee7138c.
//
// Solidity: function isTokenSet() view returns(bool)
func (_Pool *PoolSession) IsTokenSet() (bool, error) {
	return _Pool.Contract.IsTokenSet(&_Pool.CallOpts)
}

// IsTokenSet is a free data retrieval call binding the contract method 0x9ee7138c.
//
// Solidity: function isTokenSet() view returns(bool)
func (_Pool *PoolCallerSession) IsTokenSet() (bool, error) {
	return _Pool.Contract.IsTokenSet(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolCallerSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolSession) Paused() (bool, error) {
	return _Pool.Contract.Paused(&_Pool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolCallerSession) Paused() (bool, error) {
	return _Pool.Contract.Paused(&_Pool.CallOpts)
}

// Random is a free data retrieval call binding the contract method 0x5ec01e4d.
//
// Solidity: function random() view returns(uint256)
func (_Pool *PoolCaller) Random(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "random")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Random is a free data retrieval call binding the contract method 0x5ec01e4d.
//
// Solidity: function random() view returns(uint256)
func (_Pool *PoolSession) Random() (*big.Int, error) {
	return _Pool.Contract.Random(&_Pool.CallOpts)
}

// Random is a free data retrieval call binding the contract method 0x5ec01e4d.
//
// Solidity: function random() view returns(uint256)
func (_Pool *PoolCallerSession) Random() (*big.Int, error) {
	return _Pool.Contract.Random(&_Pool.CallOpts)
}

// GetGas is a paid mutator transaction binding the contract method 0x66a83a4f.
//
// Solidity: function getGas() payable returns()
func (_Pool *PoolTransactor) GetGas(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "getGas")
}

// GetGas is a paid mutator transaction binding the contract method 0x66a83a4f.
//
// Solidity: function getGas() payable returns()
func (_Pool *PoolSession) GetGas() (*types.Transaction, error) {
	return _Pool.Contract.GetGas(&_Pool.TransactOpts)
}

// GetGas is a paid mutator transaction binding the contract method 0x66a83a4f.
//
// Solidity: function getGas() payable returns()
func (_Pool *PoolTransactorSession) GetGas() (*types.Transaction, error) {
	return _Pool.Contract.GetGas(&_Pool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pool.Contract.RenounceOwnership(&_Pool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pool.Contract.RenounceOwnership(&_Pool.TransactOpts)
}

// ReturnToken is a paid mutator transaction binding the contract method 0x6178efee.
//
// Solidity: function returnToken(uint256 _amount) payable returns()
func (_Pool *PoolTransactor) ReturnToken(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "returnToken", _amount)
}

// ReturnToken is a paid mutator transaction binding the contract method 0x6178efee.
//
// Solidity: function returnToken(uint256 _amount) payable returns()
func (_Pool *PoolSession) ReturnToken(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ReturnToken(&_Pool.TransactOpts, _amount)
}

// ReturnToken is a paid mutator transaction binding the contract method 0x6178efee.
//
// Solidity: function returnToken(uint256 _amount) payable returns()
func (_Pool *PoolTransactorSession) ReturnToken(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ReturnToken(&_Pool.TransactOpts, _amount)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_Pool *PoolTransactor) Reward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "reward")
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_Pool *PoolSession) Reward() (*types.Transaction, error) {
	return _Pool.Contract.Reward(&_Pool.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_Pool *PoolTransactorSession) Reward() (*types.Transaction, error) {
	return _Pool.Contract.Reward(&_Pool.TransactOpts)
}

// SetFactoryAddress is a paid mutator transaction binding the contract method 0x83c17c55.
//
// Solidity: function setFactoryAddress(address _factoryAddress) returns()
func (_Pool *PoolTransactor) SetFactoryAddress(opts *bind.TransactOpts, _factoryAddress common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setFactoryAddress", _factoryAddress)
}

// SetFactoryAddress is a paid mutator transaction binding the contract method 0x83c17c55.
//
// Solidity: function setFactoryAddress(address _factoryAddress) returns()
func (_Pool *PoolSession) SetFactoryAddress(_factoryAddress common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetFactoryAddress(&_Pool.TransactOpts, _factoryAddress)
}

// SetFactoryAddress is a paid mutator transaction binding the contract method 0x83c17c55.
//
// Solidity: function setFactoryAddress(address _factoryAddress) returns()
func (_Pool *PoolTransactorSession) SetFactoryAddress(_factoryAddress common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetFactoryAddress(&_Pool.TransactOpts, _factoryAddress)
}

// SetGasAmount is a paid mutator transaction binding the contract method 0x0c8f78fb.
//
// Solidity: function setGasAmount(uint256 ga) returns()
func (_Pool *PoolTransactor) SetGasAmount(opts *bind.TransactOpts, ga *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setGasAmount", ga)
}

// SetGasAmount is a paid mutator transaction binding the contract method 0x0c8f78fb.
//
// Solidity: function setGasAmount(uint256 ga) returns()
func (_Pool *PoolSession) SetGasAmount(ga *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetGasAmount(&_Pool.TransactOpts, ga)
}

// SetGasAmount is a paid mutator transaction binding the contract method 0x0c8f78fb.
//
// Solidity: function setGasAmount(uint256 ga) returns()
func (_Pool *PoolTransactorSession) SetGasAmount(ga *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetGasAmount(&_Pool.TransactOpts, ga)
}

// SetReturnPercent is a paid mutator transaction binding the contract method 0xa62c65ce.
//
// Solidity: function setReturnPercent(uint256 rt) returns()
func (_Pool *PoolTransactor) SetReturnPercent(opts *bind.TransactOpts, rt *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setReturnPercent", rt)
}

// SetReturnPercent is a paid mutator transaction binding the contract method 0xa62c65ce.
//
// Solidity: function setReturnPercent(uint256 rt) returns()
func (_Pool *PoolSession) SetReturnPercent(rt *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetReturnPercent(&_Pool.TransactOpts, rt)
}

// SetReturnPercent is a paid mutator transaction binding the contract method 0xa62c65ce.
//
// Solidity: function setReturnPercent(uint256 rt) returns()
func (_Pool *PoolTransactorSession) SetReturnPercent(rt *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetReturnPercent(&_Pool.TransactOpts, rt)
}

// SetRewardPercent is a paid mutator transaction binding the contract method 0xc1c5dd27.
//
// Solidity: function setRewardPercent(uint256 rp) returns()
func (_Pool *PoolTransactor) SetRewardPercent(opts *bind.TransactOpts, rp *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setRewardPercent", rp)
}

// SetRewardPercent is a paid mutator transaction binding the contract method 0xc1c5dd27.
//
// Solidity: function setRewardPercent(uint256 rp) returns()
func (_Pool *PoolSession) SetRewardPercent(rp *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetRewardPercent(&_Pool.TransactOpts, rp)
}

// SetRewardPercent is a paid mutator transaction binding the contract method 0xc1c5dd27.
//
// Solidity: function setRewardPercent(uint256 rp) returns()
func (_Pool *PoolTransactorSession) SetRewardPercent(rp *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetRewardPercent(&_Pool.TransactOpts, rp)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Pool *PoolTransactor) SetTokenAddress(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setTokenAddress", _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Pool *PoolSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetTokenAddress(&_Pool.TransactOpts, _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Pool *PoolTransactorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetTokenAddress(&_Pool.TransactOpts, _tokenAddress)
}

// SetUploadCount is a paid mutator transaction binding the contract method 0x27cf4264.
//
// Solidity: function setUploadCount(uint256 _uploadCount) returns()
func (_Pool *PoolTransactor) SetUploadCount(opts *bind.TransactOpts, _uploadCount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setUploadCount", _uploadCount)
}

// SetUploadCount is a paid mutator transaction binding the contract method 0x27cf4264.
//
// Solidity: function setUploadCount(uint256 _uploadCount) returns()
func (_Pool *PoolSession) SetUploadCount(_uploadCount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetUploadCount(&_Pool.TransactOpts, _uploadCount)
}

// SetUploadCount is a paid mutator transaction binding the contract method 0x27cf4264.
//
// Solidity: function setUploadCount(uint256 _uploadCount) returns()
func (_Pool *PoolTransactorSession) SetUploadCount(_uploadCount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetUploadCount(&_Pool.TransactOpts, _uploadCount)
}

// SetUploadTime is a paid mutator transaction binding the contract method 0x84e96291.
//
// Solidity: function setUploadTime(uint256 _uploadTime) returns()
func (_Pool *PoolTransactor) SetUploadTime(opts *bind.TransactOpts, _uploadTime *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setUploadTime", _uploadTime)
}

// SetUploadTime is a paid mutator transaction binding the contract method 0x84e96291.
//
// Solidity: function setUploadTime(uint256 _uploadTime) returns()
func (_Pool *PoolSession) SetUploadTime(_uploadTime *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetUploadTime(&_Pool.TransactOpts, _uploadTime)
}

// SetUploadTime is a paid mutator transaction binding the contract method 0x84e96291.
//
// Solidity: function setUploadTime(uint256 _uploadTime) returns()
func (_Pool *PoolTransactorSession) SetUploadTime(_uploadTime *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetUploadTime(&_Pool.TransactOpts, _uploadTime)
}

// StakeToken is a paid mutator transaction binding the contract method 0x9ee30600.
//
// Solidity: function stakeToken(uint256 _amount, address walletAddr) returns()
func (_Pool *PoolTransactor) StakeToken(opts *bind.TransactOpts, _amount *big.Int, walletAddr common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "stakeToken", _amount, walletAddr)
}

// StakeToken is a paid mutator transaction binding the contract method 0x9ee30600.
//
// Solidity: function stakeToken(uint256 _amount, address walletAddr) returns()
func (_Pool *PoolSession) StakeToken(_amount *big.Int, walletAddr common.Address) (*types.Transaction, error) {
	return _Pool.Contract.StakeToken(&_Pool.TransactOpts, _amount, walletAddr)
}

// StakeToken is a paid mutator transaction binding the contract method 0x9ee30600.
//
// Solidity: function stakeToken(uint256 _amount, address walletAddr) returns()
func (_Pool *PoolTransactorSession) StakeToken(_amount *big.Int, walletAddr common.Address) (*types.Transaction, error) {
	return _Pool.Contract.StakeToken(&_Pool.TransactOpts, _amount, walletAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.TransferOwnership(&_Pool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.TransferOwnership(&_Pool.TransactOpts, newOwner)
}

// Undeposit is a paid mutator transaction binding the contract method 0x60cac74d.
//
// Solidity: function undeposit() returns(bool)
func (_Pool *PoolTransactor) Undeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "undeposit")
}

// Undeposit is a paid mutator transaction binding the contract method 0x60cac74d.
//
// Solidity: function undeposit() returns(bool)
func (_Pool *PoolSession) Undeposit() (*types.Transaction, error) {
	return _Pool.Contract.Undeposit(&_Pool.TransactOpts)
}

// Undeposit is a paid mutator transaction binding the contract method 0x60cac74d.
//
// Solidity: function undeposit() returns(bool)
func (_Pool *PoolTransactorSession) Undeposit() (*types.Transaction, error) {
	return _Pool.Contract.Undeposit(&_Pool.TransactOpts)
}

// UploadAllInfo is a paid mutator transaction binding the contract method 0xfe6d50ca.
//
// Solidity: function uploadAllInfo(uint256 ts, string tsStr, bytes32 sign, address[] nodeAddress, uint256[][][] TokenIds, uint256[][][] TokenPrice) returns()
func (_Pool *PoolTransactor) UploadAllInfo(opts *bind.TransactOpts, ts *big.Int, tsStr string, sign [32]byte, nodeAddress []common.Address, TokenIds [][][]*big.Int, TokenPrice [][][]*big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "uploadAllInfo", ts, tsStr, sign, nodeAddress, TokenIds, TokenPrice)
}

// UploadAllInfo is a paid mutator transaction binding the contract method 0xfe6d50ca.
//
// Solidity: function uploadAllInfo(uint256 ts, string tsStr, bytes32 sign, address[] nodeAddress, uint256[][][] TokenIds, uint256[][][] TokenPrice) returns()
func (_Pool *PoolSession) UploadAllInfo(ts *big.Int, tsStr string, sign [32]byte, nodeAddress []common.Address, TokenIds [][][]*big.Int, TokenPrice [][][]*big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UploadAllInfo(&_Pool.TransactOpts, ts, tsStr, sign, nodeAddress, TokenIds, TokenPrice)
}

// UploadAllInfo is a paid mutator transaction binding the contract method 0xfe6d50ca.
//
// Solidity: function uploadAllInfo(uint256 ts, string tsStr, bytes32 sign, address[] nodeAddress, uint256[][][] TokenIds, uint256[][][] TokenPrice) returns()
func (_Pool *PoolTransactorSession) UploadAllInfo(ts *big.Int, tsStr string, sign [32]byte, nodeAddress []common.Address, TokenIds [][][]*big.Int, TokenPrice [][][]*big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UploadAllInfo(&_Pool.TransactOpts, ts, tsStr, sign, nodeAddress, TokenIds, TokenPrice)
}

// PoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pool contract.
type PoolOwnershipTransferredIterator struct {
	Event *PoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolOwnershipTransferred)
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
		it.Event = new(PoolOwnershipTransferred)
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
func (it *PoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolOwnershipTransferred represents a OwnershipTransferred event raised by the Pool contract.
type PoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pool *PoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoolOwnershipTransferredIterator{contract: _Pool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pool *PoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolOwnershipTransferred)
				if err := _Pool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pool *PoolFilterer) ParseOwnershipTransferred(log types.Log) (*PoolOwnershipTransferred, error) {
	event := new(PoolOwnershipTransferred)
	if err := _Pool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pool contract.
type PoolPausedIterator struct {
	Event *PoolPaused // Event containing the contract specifics and raw log

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
func (it *PoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolPaused)
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
		it.Event = new(PoolPaused)
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
func (it *PoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolPaused represents a Paused event raised by the Pool contract.
type PoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pool *PoolFilterer) FilterPaused(opts *bind.FilterOpts) (*PoolPausedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PoolPausedIterator{contract: _Pool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pool *PoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PoolPaused) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolPaused)
				if err := _Pool.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Pool *PoolFilterer) ParsePaused(log types.Log) (*PoolPaused, error) {
	event := new(PoolPaused)
	if err := _Pool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pool contract.
type PoolUnpausedIterator struct {
	Event *PoolUnpaused // Event containing the contract specifics and raw log

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
func (it *PoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUnpaused)
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
		it.Event = new(PoolUnpaused)
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
func (it *PoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUnpaused represents a Unpaused event raised by the Pool contract.
type PoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pool *PoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PoolUnpausedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PoolUnpausedIterator{contract: _Pool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pool *PoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUnpaused)
				if err := _Pool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Pool *PoolFilterer) ParseUnpaused(log types.Log) (*PoolUnpaused, error) {
	event := new(PoolUnpaused)
	if err := _Pool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
