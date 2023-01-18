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

// MinerMetaData contains all meta data concerning the Miner contract.
var MinerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_createTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bandwidth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gpu\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_storage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cpu\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mem\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"approveNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIssuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendGas\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"appAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_nodeAddress\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][][]\",\"name\":\"_tokenIds\",\"type\":\"uint256[][][]\"}],\"name\":\"setTokenIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tsStr\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"sign\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"nodeAddress\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][][]\",\"name\":\"TokenIds\",\"type\":\"uint256[][][]\"},{\"internalType\":\"uint256[][][]\",\"name\":\"TokenPrice\",\"type\":\"uint256[][][]\"}],\"name\":\"uploadInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MinerABI is the input ABI used to generate the binding from.
// Deprecated: Use MinerMetaData.ABI instead.
var MinerABI = MinerMetaData.ABI

// Miner is an auto generated Go binding around an Ethereum contract.
type Miner struct {
	MinerCaller     // Read-only binding to the contract
	MinerTransactor // Write-only binding to the contract
	MinerFilterer   // Log filterer for contract events
}

// MinerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinerSession struct {
	Contract     *Miner            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinerCallerSession struct {
	Contract *MinerCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MinerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinerTransactorSession struct {
	Contract     *MinerTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinerRaw struct {
	Contract *Miner // Generic contract binding to access the raw methods on
}

// MinerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinerCallerRaw struct {
	Contract *MinerCaller // Generic read-only contract binding to access the raw methods on
}

// MinerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinerTransactorRaw struct {
	Contract *MinerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMiner creates a new instance of Miner, bound to a specific deployed contract.
func NewMiner(address common.Address, backend bind.ContractBackend) (*Miner, error) {
	contract, err := bindMiner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Miner{MinerCaller: MinerCaller{contract: contract}, MinerTransactor: MinerTransactor{contract: contract}, MinerFilterer: MinerFilterer{contract: contract}}, nil
}

// NewMinerCaller creates a new read-only instance of Miner, bound to a specific deployed contract.
func NewMinerCaller(address common.Address, caller bind.ContractCaller) (*MinerCaller, error) {
	contract, err := bindMiner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinerCaller{contract: contract}, nil
}

// NewMinerTransactor creates a new write-only instance of Miner, bound to a specific deployed contract.
func NewMinerTransactor(address common.Address, transactor bind.ContractTransactor) (*MinerTransactor, error) {
	contract, err := bindMiner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinerTransactor{contract: contract}, nil
}

// NewMinerFilterer creates a new log filterer instance of Miner, bound to a specific deployed contract.
func NewMinerFilterer(address common.Address, filterer bind.ContractFilterer) (*MinerFilterer, error) {
	contract, err := bindMiner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinerFilterer{contract: contract}, nil
}

// bindMiner binds a generic wrapper to an already deployed contract.
func bindMiner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Miner *MinerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Miner.Contract.MinerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Miner *MinerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Miner.Contract.MinerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Miner *MinerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Miner.Contract.MinerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Miner *MinerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Miner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Miner *MinerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Miner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Miner *MinerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Miner.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Miner *MinerCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Miner.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Miner *MinerSession) GetBalance() (*big.Int, error) {
	return _Miner.Contract.GetBalance(&_Miner.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Miner *MinerCallerSession) GetBalance() (*big.Int, error) {
	return _Miner.Contract.GetBalance(&_Miner.CallOpts)
}

// GetCreateTime is a free data retrieval call binding the contract method 0x7f845655.
//
// Solidity: function getCreateTime() view returns(uint256)
func (_Miner *MinerCaller) GetCreateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Miner.contract.Call(opts, &out, "getCreateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCreateTime is a free data retrieval call binding the contract method 0x7f845655.
//
// Solidity: function getCreateTime() view returns(uint256)
func (_Miner *MinerSession) GetCreateTime() (*big.Int, error) {
	return _Miner.Contract.GetCreateTime(&_Miner.CallOpts)
}

// GetCreateTime is a free data retrieval call binding the contract method 0x7f845655.
//
// Solidity: function getCreateTime() view returns(uint256)
func (_Miner *MinerCallerSession) GetCreateTime() (*big.Int, error) {
	return _Miner.Contract.GetCreateTime(&_Miner.CallOpts)
}

// GetIssuer is a free data retrieval call binding the contract method 0x52556421.
//
// Solidity: function getIssuer() view returns(address)
func (_Miner *MinerCaller) GetIssuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Miner.contract.Call(opts, &out, "getIssuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIssuer is a free data retrieval call binding the contract method 0x52556421.
//
// Solidity: function getIssuer() view returns(address)
func (_Miner *MinerSession) GetIssuer() (common.Address, error) {
	return _Miner.Contract.GetIssuer(&_Miner.CallOpts)
}

// GetIssuer is a free data retrieval call binding the contract method 0x52556421.
//
// Solidity: function getIssuer() view returns(address)
func (_Miner *MinerCallerSession) GetIssuer() (common.Address, error) {
	return _Miner.Contract.GetIssuer(&_Miner.CallOpts)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0xda1140bc.
//
// Solidity: function approveNFT(address trade, address nftAddress, uint256[] tokenIds) returns()
func (_Miner *MinerTransactor) ApproveNFT(opts *bind.TransactOpts, trade common.Address, nftAddress common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Miner.contract.Transact(opts, "approveNFT", trade, nftAddress, tokenIds)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0xda1140bc.
//
// Solidity: function approveNFT(address trade, address nftAddress, uint256[] tokenIds) returns()
func (_Miner *MinerSession) ApproveNFT(trade common.Address, nftAddress common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Miner.Contract.ApproveNFT(&_Miner.TransactOpts, trade, nftAddress, tokenIds)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0xda1140bc.
//
// Solidity: function approveNFT(address trade, address nftAddress, uint256[] tokenIds) returns()
func (_Miner *MinerTransactorSession) ApproveNFT(trade common.Address, nftAddress common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Miner.Contract.ApproveNFT(&_Miner.TransactOpts, trade, nftAddress, tokenIds)
}

// GetReward is a paid mutator transaction binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 amount) payable returns()
func (_Miner *MinerTransactor) GetReward(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Miner.contract.Transact(opts, "getReward", amount)
}

// GetReward is a paid mutator transaction binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 amount) payable returns()
func (_Miner *MinerSession) GetReward(amount *big.Int) (*types.Transaction, error) {
	return _Miner.Contract.GetReward(&_Miner.TransactOpts, amount)
}

// GetReward is a paid mutator transaction binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 amount) payable returns()
func (_Miner *MinerTransactorSession) GetReward(amount *big.Int) (*types.Transaction, error) {
	return _Miner.Contract.GetReward(&_Miner.TransactOpts, amount)
}

// SendGas is a paid mutator transaction binding the contract method 0xed38f6b0.
//
// Solidity: function sendGas() payable returns()
func (_Miner *MinerTransactor) SendGas(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Miner.contract.Transact(opts, "sendGas")
}

// SendGas is a paid mutator transaction binding the contract method 0xed38f6b0.
//
// Solidity: function sendGas() payable returns()
func (_Miner *MinerSession) SendGas() (*types.Transaction, error) {
	return _Miner.Contract.SendGas(&_Miner.TransactOpts)
}

// SendGas is a paid mutator transaction binding the contract method 0xed38f6b0.
//
// Solidity: function sendGas() payable returns()
func (_Miner *MinerTransactorSession) SendGas() (*types.Transaction, error) {
	return _Miner.Contract.SendGas(&_Miner.TransactOpts)
}

// SetTokenIds is a paid mutator transaction binding the contract method 0x2884a39d.
//
// Solidity: function setTokenIds(address appAddress, address[] _nodeAddress, uint256[][][] _tokenIds) returns()
func (_Miner *MinerTransactor) SetTokenIds(opts *bind.TransactOpts, appAddress common.Address, _nodeAddress []common.Address, _tokenIds [][][]*big.Int) (*types.Transaction, error) {
	return _Miner.contract.Transact(opts, "setTokenIds", appAddress, _nodeAddress, _tokenIds)
}

// SetTokenIds is a paid mutator transaction binding the contract method 0x2884a39d.
//
// Solidity: function setTokenIds(address appAddress, address[] _nodeAddress, uint256[][][] _tokenIds) returns()
func (_Miner *MinerSession) SetTokenIds(appAddress common.Address, _nodeAddress []common.Address, _tokenIds [][][]*big.Int) (*types.Transaction, error) {
	return _Miner.Contract.SetTokenIds(&_Miner.TransactOpts, appAddress, _nodeAddress, _tokenIds)
}

// SetTokenIds is a paid mutator transaction binding the contract method 0x2884a39d.
//
// Solidity: function setTokenIds(address appAddress, address[] _nodeAddress, uint256[][][] _tokenIds) returns()
func (_Miner *MinerTransactorSession) SetTokenIds(appAddress common.Address, _nodeAddress []common.Address, _tokenIds [][][]*big.Int) (*types.Transaction, error) {
	return _Miner.Contract.SetTokenIds(&_Miner.TransactOpts, appAddress, _nodeAddress, _tokenIds)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Miner *MinerTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Miner.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Miner *MinerSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Miner.Contract.Stake(&_Miner.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Miner *MinerTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Miner.Contract.Stake(&_Miner.TransactOpts, amount)
}

// Undeposit is a paid mutator transaction binding the contract method 0x60cac74d.
//
// Solidity: function undeposit() returns()
func (_Miner *MinerTransactor) Undeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Miner.contract.Transact(opts, "undeposit")
}

// Undeposit is a paid mutator transaction binding the contract method 0x60cac74d.
//
// Solidity: function undeposit() returns()
func (_Miner *MinerSession) Undeposit() (*types.Transaction, error) {
	return _Miner.Contract.Undeposit(&_Miner.TransactOpts)
}

// Undeposit is a paid mutator transaction binding the contract method 0x60cac74d.
//
// Solidity: function undeposit() returns()
func (_Miner *MinerTransactorSession) Undeposit() (*types.Transaction, error) {
	return _Miner.Contract.Undeposit(&_Miner.TransactOpts)
}

// UploadInfo is a paid mutator transaction binding the contract method 0x652aa583.
//
// Solidity: function uploadInfo(uint256 ts, string tsStr, bytes32 sign, address[] nodeAddress, uint256[][][] TokenIds, uint256[][][] TokenPrice) returns()
func (_Miner *MinerTransactor) UploadInfo(opts *bind.TransactOpts, ts *big.Int, tsStr string, sign [32]byte, nodeAddress []common.Address, TokenIds [][][]*big.Int, TokenPrice [][][]*big.Int) (*types.Transaction, error) {
	return _Miner.contract.Transact(opts, "uploadInfo", ts, tsStr, sign, nodeAddress, TokenIds, TokenPrice)
}

// UploadInfo is a paid mutator transaction binding the contract method 0x652aa583.
//
// Solidity: function uploadInfo(uint256 ts, string tsStr, bytes32 sign, address[] nodeAddress, uint256[][][] TokenIds, uint256[][][] TokenPrice) returns()
func (_Miner *MinerSession) UploadInfo(ts *big.Int, tsStr string, sign [32]byte, nodeAddress []common.Address, TokenIds [][][]*big.Int, TokenPrice [][][]*big.Int) (*types.Transaction, error) {
	return _Miner.Contract.UploadInfo(&_Miner.TransactOpts, ts, tsStr, sign, nodeAddress, TokenIds, TokenPrice)
}

// UploadInfo is a paid mutator transaction binding the contract method 0x652aa583.
//
// Solidity: function uploadInfo(uint256 ts, string tsStr, bytes32 sign, address[] nodeAddress, uint256[][][] TokenIds, uint256[][][] TokenPrice) returns()
func (_Miner *MinerTransactorSession) UploadInfo(ts *big.Int, tsStr string, sign [32]byte, nodeAddress []common.Address, TokenIds [][][]*big.Int, TokenPrice [][][]*big.Int) (*types.Transaction, error) {
	return _Miner.Contract.UploadInfo(&_Miner.TransactOpts, ts, tsStr, sign, nodeAddress, TokenIds, TokenPrice)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Miner *MinerTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Miner.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Miner *MinerSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Miner.Contract.Withdraw(&_Miner.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Miner *MinerTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Miner.Contract.Withdraw(&_Miner.TransactOpts, amount)
}
