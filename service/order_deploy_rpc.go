package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/url"
	"strconv"
	"time"
	"ubiqnet/abi"
	"ubiqnet/constants"
	"ubiqnet/store/localstore"
)

type NodeSuiteInfo struct {
	TemplateAddress      string  `json:"templateAddress"`
	OrderId              int     `json:"orderId"`
	IpTokenIds           []int64 `json:"ipTokenIds"`
	BandwidthTokenIds    []int64 `json:"bandwidthTokenIds"`
	GpuTokenIds          []int64 `json:"gpuTokenIds"`
	StorageTokenIds      []int64 `json:"storageTokenIds"`
	CpuTokenIds          []int64 `json:"cpuTokenIds"`
	MemTokenIds          []int64 `json:"memTokenIds"`
	NodeAddress          string  `json:"nodeAddress"`
	SuitUtcTimestampNano int64   `json:"SuitUtcTimestampNano"`
	Index                int     `json:"index"`
	ExpireTimestamp      int64   `json:"expireTimestamp"`
}

func (info NodeSuiteInfo) ToString() string {
	res, err := json.Marshal(info)
	if err != nil {
		return err.Error()
	}

	return string(res)
}

func ToNodeSuitInfo(res string) NodeSuiteInfo {
	var p2 NodeSuiteInfo
	err := json.Unmarshal([]byte(res), &p2)
	if err != nil {
		return NodeSuiteInfo{}
	}
	return p2
}
func (info ValidatorSuitInfo) ToString() string {
	res, err := json.Marshal(info)
	if err != nil {
		return err.Error()
	}

	return string(res)
}

func ToValidatorSuitInfo(res string) ValidatorSuitInfo {
	var p2 ValidatorSuitInfo
	err := json.Unmarshal([]byte(res), &p2)
	if err != nil {
		return ValidatorSuitInfo{}
	}
	return p2
}

type ValidatorSuitInfo struct {
	AppAddress            string          `json:"appAddress"`
	AppOwnerAddress       string          `json:"appOwnerAddress"`
	OrderId               int             `json:"orderId"`
	FinishedTimestampNano int64           `json:"finishedTimestampNano"`
	NodeSuitInfos         []NodeSuiteInfo `json:"nodeSuitInfos"`
	ValidateAddress       string          `json:"validateAddress"`
	OrderExpireTime       int64           `json:"orderExpireTime"`
}

func (info ImageInfo) ToString() string {
	res, err := json.Marshal(info)
	if err != nil {
		return err.Error()
	}

	return string(res)
}

func ToImageInfo(res string) ImageInfo {
	var p2 ImageInfo
	err := json.Unmarshal([]byte(res), &p2)
	if err != nil {
		return ImageInfo{}
	}
	return p2
}

type ImageInfo struct {
	OrderId         int               `json:"orderId"`
	IpTokenCount    int               `json:"ipTokenCount"`
	BandwidthCount  int               `json:"bandwidthCount"`
	GpuCount        int               `json:"gpuCount"`
	StorageCount    int               `json:"storageCount"`
	CpuCount        int               `json:"cpuCount"`
	MemCount        int               `json:"memCount"`
	CreateTimeStamp int64             `json:"createTimeStamp"`
	ExpireTimeStamp int64             `json:"expireTimeStamp"`
	TemplateAddress string            `json:"templateAddress"`
	OwnerAddress    string            `json:"ownerAddress"`
	ImageName       string            `json:"imageName"`
	Cmd             string            `json:"cmd"`
	Env             []string          `json:"env"`
	Location        string            `json:"location"`
	PortMap         map[string]string `json:"portMap"`
	NodeAddress     string            `json:"nodeAddress"`
	IpTokenIds      []int64           `json:"ipTokenIds"`
	BandwidthIds    []int64           `json:"bandwidthIds"`
	GpuIds          []int64           `json:"gpuIds"`
	StorageIds      []int64           `json:"storageIds"`
	CpuIds          []int64           `json:"cpuIds"`
	MemIds          []int64           `json:"memIds"`
	Index           int               `json:"index"`
}

func (s *RpcService) GetLatestOrderId() int {
	client = s.ConnectRPC()
	token, err := abi.NewTrade(common.HexToAddress(constants.TradeAddress), client)
	if err != nil {
		return 0
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := bind.CallOpts{From: common.HexToAddress(Node.WalletAddress), Context: ctx}
	res, err := token.GetLatestOrderId(&opts)
	if err != nil {
		return 0
	}
	orderId, err := strconv.Atoi(res.String())
	if err != nil {
		return 0
	}
	return orderId
}

func timeOutContex() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("timeOut")
		fmt.Println(ctx.Err()) // prints "context deadline exceeded" }
	}

}

func (s *RpcService) GetOrderById(orderId int64) (*abi.TradeNodeInfo, error) {
	client = s.ConnectRPC()
	token, err := abi.NewTrade(common.HexToAddress(constants.TradeAddress), client)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	opts := bind.CallOpts{From: common.HexToAddress(Node.WalletAddress), Context: ctx}
	res, err := token.GetOrderInfo(&opts, big.NewInt(orderId))
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *RpcService) GetAppInfo(orderId int, appAddress common.Address) ([]ImageInfo, error) {
	client = s.ConnectRPC()
	token, err := abi.NewApp(appAddress, client)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := bind.CallOpts{From: common.HexToAddress(Node.WalletAddress), Context: ctx}
	tokenNeeded, err := token.GetTokenNeed(&opts)
	if err != nil {
		return nil, err
	}
	configs, err := token.GetImage(&opts)
	if err != nil {
		return nil, err
	}

	var imageInfos []ImageInfo
	for i := 0; i < len(tokenNeeded); i++ {
		var imageInfo = new(ImageInfo)
		configStr := configs[i]
		urlSafeConfig, _ := url.QueryUnescape(configStr)
		imageConfig := ToImageConfig(urlSafeConfig)
		imageInfo.Cmd = imageConfig.Cmd
		imageInfo.Index = i
		imageInfo.ImageName = imageConfig.ImageName
		imageInfo.Env = imageConfig.Env
		imageInfo.Location = imageConfig.Location
		imageInfo.PortMap = imageConfig.Portmap
		tokenCount := tokenNeeded[i]
		imageInfo.OrderId = orderId
		imageInfo.IpTokenCount, _ = strconv.Atoi(tokenCount[0].String())
		imageInfo.BandwidthCount, _ = strconv.Atoi(tokenCount[1].String())
		imageInfo.GpuCount, _ = strconv.Atoi(tokenCount[2].String())
		imageInfo.StorageCount, _ = strconv.Atoi(tokenCount[3].String())
		imageInfo.CpuCount, _ = strconv.Atoi(tokenCount[4].String())
		imageInfo.MemCount, _ = strconv.Atoi(tokenCount[5].String())
		imageInfos = append(imageInfos, *imageInfo)
	}
	return imageInfos, nil
}

func (s *RpcService) GetAppInfoDetail(orderId int, info abi.TradeNodeInfo, ownerAddress string) ([]ImageInfo, error) {
	client = s.ConnectRPC()
	token, err := abi.NewApp(info.AppAddress, client)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := bind.CallOpts{From: common.HexToAddress(Node.WalletAddress), Context: ctx}
	res, err := token.GetTokenIds(&opts)
	templateAddress, err := token.GetTemplates(&opts)
	nodes, err := token.GetNodeAddress(&opts)
	if err != nil {
		return nil, err
	}
	imagesLen := len(nodes)
	var imageInfos = make([]ImageInfo, imagesLen)
	configs, err := token.GetImage(&opts)
	if err != nil {
		return nil, err
	}

	for i := 0; i < imagesLen; i++ {
		infos := res[i]
		var imageInfo = new(ImageInfo)
		imageInfo.NodeAddress = nodes[i].String()
		imageInfo.TemplateAddress = templateAddress[i].String()
		imageInfo.OrderId = orderId
		configStr := configs[i]
		urlSafeConfig, _ := url.QueryUnescape(configStr)
		imageConfig := ToImageConfig(urlSafeConfig)
		imageInfo.Cmd = imageConfig.Cmd
		imageInfo.Location = imageConfig.Location
		imageInfo.OwnerAddress = ownerAddress
		imageInfo.ImageName = imageConfig.ImageName
		imageInfo.Env = imageConfig.Env
		imageInfo.Index = i
		if len(imageInfo.IpTokenIds) > 0 {
			imageInfo.PortMap = imageConfig.Portmap
		}
		imageInfo.ExpireTimeStamp = info.ExpiresTime.Int64()
		if len(infos[0]) > 0 {
			for _, b := range infos[0] {
				imageInfo.IpTokenIds = append(imageInfo.IpTokenIds, b.Int64())
			}

		}
		if len(infos[1]) > 0 {
			for _, b := range infos[1] {
				imageInfo.BandwidthIds = append(imageInfo.BandwidthIds, b.Int64())
			}
		}
		if len(infos[2]) > 0 {
			for _, b := range infos[2] {
				imageInfo.GpuIds = append(imageInfo.GpuIds, b.Int64())
			}
		}
		if len(infos[3]) > 0 {
			for _, b := range infos[3] {
				imageInfo.StorageIds = append(imageInfo.StorageIds, b.Int64())
			}
		}
		if len(infos[4]) > 0 {
			for _, b := range infos[4] {
				imageInfo.CpuIds = append(imageInfo.CpuIds, b.Int64())
			}
		}
		if len(infos[5]) > 0 {
			for _, b := range infos[5] {
				imageInfo.MemIds = append(imageInfo.MemIds, b.Int64())
			}
		}
		if len(imageInfo.IpTokenIds) > 0 {
			imageInfo.PortMap = imageConfig.Portmap
		}
		imageInfos[i] = *imageInfo
	}
	return imageInfos, nil
}

func (s *RpcService) GetAppOwner(appAddress common.Address) (string, error) {
	client = s.ConnectRPC()
	token, err := abi.NewApp(appAddress, client)
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := bind.CallOpts{From: common.HexToAddress(Node.WalletAddress), Context: ctx}
	address, err := token.GetMyOwner(&opts)
	if err != nil {
		return "", err
	}
	return address.String(), nil
}

func (s *RpcService) SetTokenTrade(appAddress common.Address, nodeAddress []common.Address, tokenIds [][][]*big.Int, orderId int64) (string, error) {
	client = s.ConnectRPC()
	token, err := abi.NewTrade(common.HexToAddress(constants.TradeAddress), client)
	if err != nil {
		return "", err
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {

		return "", err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(Node.PrivateKey, chainId)
	if err != nil {
		return "", err
	}
	opts.From = common.HexToAddress(Node.WalletAddress)
	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		rpcLog.Error(err.Error())
		return "", err
	}

	msg := ethereum.CallMsg{From: common.HexToAddress(Node.WalletAddress), GasPrice: price.Mul(price, big.NewInt(int64(3)))}
	limit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		rpcLog.Error(err.Error())
		return "", err
	}
	opts.GasLimit = limit * 100
	transaction, err := token.SetTokenIds(opts, appAddress, nodeAddress, tokenIds, big.NewInt(orderId), common.HexToAddress(localstore.Get(constants.ContractAddressKey)))
	if err != nil {
		return "", err
	}
	return transaction.Hash().String(), nil
}

func (s *RpcService) CheckTokenIsRent(tokenId int64, tokenContract string) (bool, error) {
	client = s.ConnectRPC()
	token, err := abi.NewNft(common.HexToAddress(tokenContract), client)
	if err != nil {
		return false, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := bind.CallOpts{From: common.HexToAddress(Node.WalletAddress), Context: ctx}
	exp, err := token.UserExpires(&opts, big.NewInt(tokenId))
	if err != nil {
		println(err)
		return false, err
	}
	times := time.Now().UTC().Unix()
	exps := exp.Int64()
	if times >= exps {
		return false, nil
	} else {
		return true, nil
	}

}
