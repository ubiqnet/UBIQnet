package task

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"
	"ubiqnet/constants"
	"ubiqnet/service"
	"ubiqnet/store/localstore"
	"ubiqnet/tools"

	"github.com/ethereum/go-ethereum/common"
	logging "github.com/ipfs/go-log"
	"github.com/robfig/cron/v3"
)

type HeartData struct {
	Status    int
	Timestamp int64
	Info      string
}

type UploadData struct {
	ContractAddress string     `json:"contractAddress"`
	TimeStamp       int64      `json:"timestamp"`
	CpuIds          []*big.Int `json:"cpuIds"`
	IpIds           []*big.Int `json:"ipIds"`
	BrandWidthIds   []*big.Int `json:"brandWidthIds"`
	GpuIds          []*big.Int `json:"gpuIds"`
	MemIds          []*big.Int `json:"memIds"`
	StorageIds      []*big.Int `json:"storageIds"`

	CpuPrices        []*big.Int `json:"cpuPrices"`
	IpPrices         []*big.Int `json:"ipPrices"`
	BrandWidthPrices []*big.Int `json:"brandWidthPrices"`
	GpuPrices        []*big.Int `json:"gpuPrices"`
	MemPrices        []*big.Int `json:"memPrices"`
	StoragePrices    []*big.Int `json:"storagePrices"`
	Sign             [32]byte   `json:"sign"`
}

var (
	blockNum uint64 = 0
	IdKey           = "id_key"
	MaxId           = 100000
)
var taskLog = logging.Logger("task:task")

var rpcService service.RpcService

func InitTask() {
	logging.SetLogLevelRegex("task:task", "info")

	// logging.SetAllLoggers(log.LevelInfo)
	logging.SetLogLevelRegex("service:validator-newOrder", "debug")

	go service.P2p.SubscribeBroardCastData()
	go service.P2p.SubscribeP2pData()

	c := cron.New(cron.WithSeconds())
	c2 := cron.New(cron.WithSeconds())
	c.AddFunc("0 */5 * * * ?", heart)
	c2.AddFunc("*/10 * * * * ?", service.CheckContainerNeedStop)
	c2.AddFunc("0 */1 * * * ?", service.CheckHash)
	go checkBlock()
	service.ReceiveStartDockerQueueConsumer()
	service.ReceiveStopDockerQueueConsumer()
	service.ReceiveQueueConsumer()
	service.ReceiveOtherTemplateQueueConsumer()
	go service.AutoRelease()

	if service.Node.NodeType == constants.NodeTypeVerify {
		service.ReceiveNodeTemplateSuitQueueConsume()
		service.ReceiveOrderSuitedCompleteQueueConsume()
		c.AddFunc("0 */20 * * * ?", uploadInfo)
		c2.AddFunc("*/10 * * * * ?", service.CheckOrder)
	}

	go c.Start()
	go c2.Start()

}

func checkBlock() {
	client := rpcService.ConnectRPC()
	defer client.Close()
	for {

		num, _ := client.BlockNumber(context.Background())
		if num > blockNum {
			for {
				blockNum += 1
				if blockNum == num {
					break
				}
			}

			taskLog.Info("harvest service: updated block height to " + strconv.FormatUint(num, 10))

		}
		v := tools.GetRandomInt(10, 30)
		time.Sleep(time.Duration(v) * time.Second)

	}

}

func heart() {
	if service.Node.NodeType == constants.NodeTypeNomal || service.Node.NodeType == constants.NodeTypeStable {
		verifyNode := localstore.Get(constants.VerifyPeerIdKey)
		if verifyNode == "" {
			peerList := service.P2p.GetPeers()
			for k := range peerList {

				verifyNode = localstore.Get(constants.VerifyPeerIdKey)
				if verifyNode == "" {

					service.P2p.BindVerifyNode(k)
					t := time.NewTimer(10 * time.Second)
					<-t.C
				} else {
					break
				}

			}
		}

	} else {
		localstore.Put(constants.VerifyPeerIdKey, service.Node.PeerId)
	}

	if localstore.Get(constants.VerifyPeerIdKey) == "" {
		return
	}

	amount, err := rpcService.BalanceOfEth()
	if err != nil {
		taskLog.Error(err)
		return
	}
	if amount.Cmp(big.NewInt(0)) == 0 {
		taskLog.Error("The operation failed due to the insufficient ETH in your wallet: ", service.Node.WalletAddress)
		return
	}

	contractAddress := localstore.Get(constants.ContractAddressKey)

	amount, err = rpcService.BalanceOfMiner(contractAddress)
	if err != nil {
		taskLog.Error(err)
		return
	}
	if service.Node.NodeType != constants.NodeTypeNomal && amount.Cmp(big.NewInt(0)) == 0 {
		taskLog.Error("Insufficient UBQ staking, please stake UBQ to the chequebook first.")
		return
	}

	taskLog.Info("harvest service: harvest packing data ...")

	timestamp := time.Now().UnixMilli()

	sign := getSign(timestamp)
	peerId := localstore.Get(constants.VerifyPeerIdKey)
	tokenIds, err := rpcService.GetAllTokenIds()
	if err != nil {
		taskLog.Error(err)
		return
	}
	cpuPriceList := []*big.Int{}
	gpuPriceList := []*big.Int{}
	memPriceList := []*big.Int{}
	storagePriceList := []*big.Int{}
	ipPriceList := []*big.Int{}
	brandWidthPriceList := []*big.Int{}

	cpuPrice := big.NewInt(constants.CpuPrice)
	gpuPrice := big.NewInt(constants.GpuPrice)
	memPrice := big.NewInt(constants.MemPrice)
	storagePrice := big.NewInt(constants.StoragePrice)
	ipPrice := big.NewInt(constants.IpPrice)
	brandWidthPrice := big.NewInt(constants.BrandWidthPrice)
	for i := 0; i < len(tokenIds["cpu"]); i++ {
		cpuPriceList = append(cpuPriceList, cpuPrice)
	}
	for i := 0; i < len(tokenIds["gpu"]); i++ {
		gpuPriceList = append(gpuPriceList, gpuPrice)
	}
	for i := 0; i < len(tokenIds["mem"]); i++ {
		memPriceList = append(memPriceList, memPrice)
	}
	for i := 0; i < len(tokenIds["storage"]); i++ {

		storagePriceList = append(storagePriceList, storagePrice)
	}

	for i := 0; i < len(tokenIds["ip"]); i++ {

		ipPriceList = append(ipPriceList, ipPrice)
	}

	for i := 0; i < len(tokenIds["brandWidth"]); i++ {
		brandWidthPriceList = append(brandWidthPriceList, brandWidthPrice)
	}

	uploadData := UploadData{
		ContractAddress:  contractAddress,
		TimeStamp:        timestamp,
		CpuIds:           tokenIds["cpu"],
		MemIds:           tokenIds["mem"],
		GpuIds:           tokenIds["gpu"],
		BrandWidthIds:    tokenIds["brandWidth"],
		IpIds:            tokenIds["ip"],
		StorageIds:       tokenIds["storage"],
		CpuPrices:        cpuPriceList,
		MemPrices:        memPriceList,
		StoragePrices:    storagePriceList,
		GpuPrices:        gpuPriceList,
		IpPrices:         ipPriceList,
		BrandWidthPrices: brandWidthPriceList,

		Sign: sign}
	uploadBytes, _ := json.Marshal(&uploadData)
	p2pData := service.P2pData{Sender: service.Node.PeerId, Event: "uploadInfo", Extra: string(uploadBytes)}
	sendBytes, _ := json.Marshal(&p2pData)
	service.P2p.PublishP2pData(peerId, sendBytes)

	//end

}

func uploadInfo() {

	uploadMap := localstore.FindAllHeart()
	fmt.Println(uploadMap)

	nodeAddressList := []common.Address{}
	tokenIdsList := [][][]*big.Int{}
	tokenPricesList := [][][]*big.Int{}
	if len(uploadMap) == 0 {
		return
	}
	var tokenIdMap = make(map[string][][]*big.Int)
	var tokenPriceMap = make(map[string][][]*big.Int)
	for contractAddress, infos := range uploadMap {
		nodeAddressList = append(nodeAddressList, common.HexToAddress(contractAddress))
		if len(infos) >= constants.UploadNum {

			for k, v := range infos {
				var uploadInfo UploadData
				json.Unmarshal([]byte(v), &uploadInfo)
				tokenIds := [][]*big.Int{uploadInfo.IpIds,
					uploadInfo.BrandWidthIds,
					uploadInfo.GpuIds,
					uploadInfo.StorageIds,
					uploadInfo.CpuIds,
					uploadInfo.MemIds}

				tokenPrices := [][]*big.Int{uploadInfo.IpPrices,
					uploadInfo.BrandWidthPrices,
					uploadInfo.GpuPrices,
					uploadInfo.StoragePrices,
					uploadInfo.CpuPrices,
					uploadInfo.MemPrices}

				tokenIdMap[contractAddress] = tokenIds
				tokenPriceMap[contractAddress] = tokenPrices
				localstore.Remove(k)

			}

		}

		tokenIdsList = append(tokenIdsList, tokenIdMap[contractAddress])
		if len(tokenIdsList) > 0 {
			tokenPricesList = append(tokenPricesList, tokenPriceMap[contractAddress])
		}

	}

	timestamp := time.Now().UnixMilli() / 1000
	sign := getSign(timestamp)
	taskLog.Info("uploadInfo", nodeAddressList, tokenIdsList, tokenPricesList, fmt.Sprintf("%x", sign), timestamp)
	t, err := rpcService.UploadInfo(nodeAddressList, tokenIdsList, tokenPricesList, sign, timestamp)
	if err != nil {
		taskLog.Error(err.Error())
		return
	}
	taskLog.Info("nodeAddress: ", nodeAddressList, "tokenIds: ", tokenIdsList, "tokenPrice: ", tokenPricesList, "timestamp: ", timestamp)

	// localstore.Put(constants.HeartHashKeyPrefix, t.Hash().String())
	localstore.Put(constants.LastUploadTimeKey, strconv.FormatInt(timestamp, 10))
	taskLog.Info("harvest service: harvest  height: " + strconv.FormatInt(int64(blockNum), 10) + ",tx " + t.Hash().String())
}

func getSign(timestamp int64) [32]byte {
	return sha256.Sum256([]byte(constants.EncryptKey + strconv.FormatInt(timestamp, 10)))
}
