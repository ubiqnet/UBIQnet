// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"os"
	"strconv"
	"time"
	"ubiqnet/abi"
	"ubiqnet/constants"
	"ubiqnet/store/localstore"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	logging "github.com/ipfs/go-log"
)

var rpcLog = logging.Logger("service:rpc")

type RpcService struct {
}
type NodeData struct {
	DataDir       string
	RpcAddress    string
	PrivateKey    *ecdsa.PrivateKey
	PublicKey     []byte
	NetWorkId     int
	WalletAddress string
	DebugApi      string
	MinerStatus   bool
	TotalDiskFree uint64
	CpuCore       int
	DepositGas    uint64
	HarvestGas    int64
	Hw            int
	NodeType      int
	VerifyPleage  int
	NormalPleage  int
	StablePleage  int
	PeerId        string
}

var (
	Node   NodeData
	client *ethclient.Client
)

func (service *RpcService) ConnectRPC() *ethclient.Client {
	if client == nil {
		logging.SetLogLevelRegex("service:rpc", "info")
		rpcClient, err := rpc.Dial(Node.RpcAddress)
		if err != nil {
			rpcLog.Error(err.Error())
			os.Exit(1)
		}
		conn := ethclient.NewClient(rpcClient)
		client = conn
	}

	return client

}

type WalletData struct {
	Address string `json:"address"`
	Crypto  string `json:"crypto"`
	Salt    string `json:"salt"`
	Version string `json:"version"`
}
type IncomeData struct {
	DepositAmount   *big.Int
	AllRewardAmount *big.Int
	AllReturnAmount *big.Int
	RewardAmount    *big.Int
}

func (s *RpcService) Deploy() (string, error) {

	token, err := abi.NewMinerFactory(common.HexToAddress(constants.FactoryAddress), client)

	if err != nil {
		return "", err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	if err != nil {
		return "", err
	}

	price = price.Mul(price, big.NewInt(3))
	opts.GasLimit = Node.DepositGas
	opts.GasPrice = price
	if err != nil {
		return "", err
	}
	t, err := token.DeployMiner(opts, common.HexToAddress(Node.WalletAddress))
	if err != nil {
		return "", err
	}
	hash := t.Hash()
	return hash.String(), nil

}

func (s *RpcService) BalanceOfEth() (*big.Int, error) {
	client = s.ConnectRPC()
	balance, err := client.BalanceAt(context.TODO(), common.HexToAddress(Node.WalletAddress), nil)
	if err != nil {
		rpcLog.Error(err)
		return nil, err
	}
	return balance, nil

}

func (s *RpcService) BalanceOfUbq() *big.Int {
	client = s.ConnectRPC()
	token, err := abi.NewUbq(common.HexToAddress(constants.UbqAddress), client)
	if err != nil {
		return nil
	}
	amount, err := token.BalanceOf(nil, common.HexToAddress(Node.WalletAddress))
	if err != nil {
		return big.NewInt(0)
	}
	return amount
}

func (s *RpcService) BalanceOfMiner(contractAddress string) (*big.Int, error) {
	client = s.ConnectRPC()
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := bind.CallOpts{From: common.HexToAddress(Node.WalletAddress), Context: ctx}
	amount, err := token.GetBalance(&opts)
	if err != nil {
		return nil, err
	}

	return amount, nil
}

func (s *RpcService) HinDecimal() uint8 {
	client = s.ConnectRPC()
	token, err := abi.NewUbq(common.HexToAddress(constants.UbqAddress), client)
	if err != nil {
		return 0
	}
	amount, err := token.Decimals(nil)
	if err != nil {
		return 0
	}

	return amount
}

func (s *RpcService) checkTransaction(hash string) (bool, error) {
	client = s.ConnectRPC()
	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return false, err
	}
	if receipt.Status == uint64(1) {
		return true, nil
	}
	if receipt.Status == uint64(0) {
		return false, nil
	}
	return false, errors.New("no res!")
}

func (s *RpcService) GetRewardAmount(address string) *big.Int {
	client = s.ConnectRPC()
	token, err := abi.NewPool(common.HexToAddress(constants.PoolAddress), client)
	if err != nil {
		rpcLog.Error(err.Error())
		return nil
	}
	amount, err := token.GetRewardAmount(nil, common.HexToAddress(address))
	if err != nil {
		rpcLog.Error(err.Error())
		return nil
	}
	return amount

}

func (s *RpcService) GetGasAmount() (*big.Int, error) {
	client := s.ConnectRPC()
	token, err := abi.NewPool(common.HexToAddress(constants.PoolAddress), client)
	if err != nil {
		return nil, err
	}
	amount, err := token.GetGasAmount(nil)
	if err != nil {

		return nil, err
	}
	return amount, nil

}

func (s *RpcService) GetReward(amount *big.Int, contractAddress string, gas *big.Int) (*types.Transaction, error) {
	client = s.ConnectRPC()
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {

		return nil, err
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}
	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{From: common.HexToAddress(Node.WalletAddress), GasPrice: price.Mul(price, big.NewInt(3))}
	limit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, err
	}
	opts.From = common.HexToAddress(Node.WalletAddress)

	opts.Value = gas
	opts.GasPrice = price
	opts.GasLimit = limit * 100

	t, err := token.GetReward(opts, amount)
	if err != nil {
		return nil, err

	}

	return t, nil

}

func (s *RpcService) WithDraw(contractAddress string, amount *big.Int) (*types.Transaction, error) {
	client = s.ConnectRPC()
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {

		return nil, err
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {

		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	if err != nil {

		return nil, err
	}
	t, err := token.Withdraw(opts, amount)
	if err != nil {

		return nil, err
	}
	return t, nil

}

func (s *RpcService) GetAllInfo(address string) (*IncomeData, error) {
	client = s.ConnectRPC()
	token, err := abi.NewPool(common.HexToAddress(constants.PoolAddress), client)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := bind.CallOpts{From: common.HexToAddress(Node.WalletAddress), Context: ctx}
	res, err := token.GetAllInfo(&opts, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	var data IncomeData
	data.DepositAmount = res.DepositAmount
	data.AllRewardAmount = res.AllRewardAmount
	data.AllReturnAmount = res.AllReturnAmount
	data.RewardAmount = res.RewardAmount
	return &data, nil
}

func (s *RpcService) UnDeposit(contractAddress string) (*types.Transaction, error) {
	client = s.ConnectRPC()
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}

	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{From: common.HexToAddress(Node.WalletAddress), GasPrice: price}
	limit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, err
	}
	opts.GasLimit = limit * 100

	t, err := token.Undeposit(opts)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *RpcService) Recharge(amount *big.Int, contractAddress string) (*types.Transaction, error) {
	client = s.ConnectRPC()
	ubq, err := abi.NewUbq(common.HexToAddress(constants.UbqAddress), client)
	if err != nil {
		return nil, err
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}
	t, err := ubq.Approve(opts, common.HexToAddress(contractAddress), amount)
	if err != nil {
		return nil, err
	}
	bind.WaitMined(context.Background(), client, t)
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}

	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{From: common.HexToAddress(Node.WalletAddress), GasPrice: price.Mul(price, big.NewInt(3))}
	limit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, err
	}
	opts.GasLimit = limit * 100

	t, err = token.Stake(opts, amount)
	if err != nil {
		return nil, err
	}
	return t, nil

}

func (s *RpcService) MintNft(contactAddress string, nftContactAddress string, hardInfo string) (*types.Transaction, error) {

	client = s.ConnectRPC()
	token, err := abi.NewMinerFactory(common.HexToAddress(constants.FactoryAddress), client)
	if err != nil {
		return nil, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}

	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{From: common.HexToAddress(Node.WalletAddress), GasPrice: price.Mul(price, big.NewInt(int64(3)))}
	limit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, err
	}
	opts.GasLimit = limit * 100

	t, err := token.MintNfts(opts, common.HexToAddress(contactAddress), common.HexToAddress(nftContactAddress), hardInfo)
	if err != nil {
		return nil, err
	}
	return t, nil

}

func (s *RpcService) PutNft(contractAddress string, nftContractAddress string, tokenId int) (*types.Transaction, error) {

	client = s.ConnectRPC()
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}

	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{From: common.HexToAddress(Node.WalletAddress), GasPrice: price.Mul(price, big.NewInt(int64(2)))}
	limit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, err
	}
	opts.GasLimit = limit * 100
	tokenIds := []*big.Int{big.NewInt(int64(tokenId))}

	t, err := token.ApproveNFT(opts, common.HexToAddress(constants.TradeAddress), common.HexToAddress(nftContractAddress), tokenIds)
	if err != nil {
		return nil, err
	}

	return t, nil

}

func (s *RpcService) DownNft(contractAddress string, nftAddress string, tokenId int) (*types.Transaction, error) {

	client = s.ConnectRPC()
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}

	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{From: common.HexToAddress(Node.WalletAddress), GasPrice: price.Mul(price, big.NewInt(int64(2)))}
	limit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, err
	}
	opts.GasLimit = limit * 100
	tokenIds := []*big.Int{big.NewInt(int64(tokenId))}

	t, err := token.ApproveNFT(opts, common.HexToAddress("0x0000000000000000000000000000000000000000"), common.HexToAddress(nftAddress), tokenIds)
	if err != nil {
		return nil, err
	}
	return t, nil

}

func (s *RpcService) GetAllTokenIds() (map[string][]*big.Int, error) {

	client = s.ConnectRPC()
	contractAddress := localstore.Get(constants.ContractAddressKey)
	token, err := abi.NewMinerFactory(common.HexToAddress(constants.FactoryAddress), client)
	tokenMap := make(map[string][]*big.Int)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := bind.CallOpts{From: common.HexToAddress(Node.WalletAddress), Context: ctx}
	res, err := token.GetIdsByMiner(&opts, common.HexToAddress(contractAddress))
	if err != nil {
		return nil, err
	}

	tokenMap["cpu"] = res.CpuTokenIds
	tokenMap["mem"] = res.MemTokenIds
	tokenMap["storage"] = res.StorageTokenIds
	tokenMap["ip"] = res.IpTokenIds
	tokenMap["brandWidth"] = res.BandwidthTokenIds
	tokenMap["gpu"] = res.GpuTokenIds

	return tokenMap, nil

}

func (s *RpcService) UploadInfo(nodeAddressList []common.Address, tokenIds [][][]*big.Int, tokenPrices [][][]*big.Int, sign [32]byte, timestamp int64) (*types.Transaction, error) {

	contractAddress := localstore.Get(constants.ContractAddressKey)
	client = s.ConnectRPC()
	token, err := abi.NewMiner(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {

		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}
	opts.From = common.HexToAddress(Node.WalletAddress)
	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		rpcLog.Error(err.Error())
		return nil, err
	}
	msg := ethereum.CallMsg{From: common.HexToAddress(Node.WalletAddress), GasPrice: price.Mul(price, big.NewInt(3))}
	limit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		rpcLog.Error(err.Error())
		return nil, err
	}
	opts.GasLimit = limit * 100
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts.Context = ctx
	t, err := token.UploadInfo(opts, big.NewInt(timestamp), strconv.FormatInt(timestamp, 10), sign, nodeAddressList, tokenIds, tokenPrices)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *RpcService) GetHardInfo(contractAddress string, tokenId *big.Int) (string, error) {
	client = s.ConnectRPC()
	token, err := abi.NewNft(common.HexToAddress(contractAddress), client)
	if err != nil {
		return "", err
	}
	res, err := token.GetHardwareInfo(nil, tokenId)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *RpcService) GetNftStatus(contractAddress string, tokenId *big.Int) (string, error) {
	client = s.ConnectRPC()
	token, err := abi.NewNft(common.HexToAddress(contractAddress), client)
	if err != nil {
		return "", err
	}
	res, err := token.GetApproved(nil, tokenId)
	if err != nil {
		return "", err
	}

	if res.String()==""{
		return "error",nil
	}
	if res.String()=="0x0000000000000000000000000000000000000000"{
		return "local",nil
	}


	return "online", nil
}
