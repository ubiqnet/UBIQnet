package openapi

import (
	"encoding/json"
	"math/big"
	"strconv"
	"time"
	"ubiqnet/constants"
	"ubiqnet/service"
	"ubiqnet/store/localstore"

	"github.com/gin-gonic/gin"
)

type MintNtfRequest struct {
	NftType int    `json:"nftType"`
	Info    string `json:"info"`
}

type PutNftRequest struct {
	NftType  int    `json:"nftType"`
	TokenId  int    `json:"tokenId"`
	HardInfo string `json:"info"`
}
type NftItemData struct {
	NftTypeName string   `json:"nft"`
	NftType     int      `json:"nftType"`
	Info        string   `json:"configure"`
	Status      string   `json:"status"`
	TokenId     *big.Int `json:"tokenId"`
}

func MintNft(c *gin.Context) {

	var postData MintNtfRequest
	err := c.Bind(&postData)
	if err != nil {
		c.JSON(200, gin.H{
			"result": nil,
			"error":  err,
		})
		return
	}
	hash, err := service.Nft.MintNfts(postData.NftType, postData.Info)
	if err != nil {
		c.JSON(200, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"result": hash,
		"error":  nil,
	})

}

func PutNft(c *gin.Context) {
	var nftData PutNftRequest
	err := c.Bind(&nftData)
	if err != nil {
		c.JSON(200, gin.H{
			"result": nil,
			"error":  err,
		})
		return
	}
	var nftAddress string
	mintType := nftData.NftType
	if mintType == 1 {
		nftAddress = constants.NftCpuAddress

	} else if mintType == 2 {
		nftAddress = constants.NftMemAddress
	} else if mintType == 3 {
		nftAddress = constants.NftStorageAddress
	} else if mintType == 4 {
		nftAddress = constants.NftGpuAddress
	} else if mintType == 5 {
		nftAddress = constants.NftIpAddress
	} else if mintType == 6 {
		nftAddress = constants.NftBandWidth
	} else {

		c.JSON(200, gin.H{
			"result": nil,
			"error":  "nft type not exsit",
		})
		return

	}
	contractAddress := localstore.Get(constants.ContractAddressKey)
	t, err := rpcService.PutNft(contractAddress, nftAddress, nftData.TokenId)
	if err != nil {
		c.JSON(200, gin.H{
			"result": nil,
			"error":  err,
		})
		return
	}
	t_time := time.Now().UnixMilli()
	localstore.Put(constants.HASH_CHECK_KEY_PRIFIX+t.Hash().String()+"-"+"put", strconv.Itoa(mintType)+"-"+strconv.Itoa(nftData.TokenId))
	localstore.AddNftRecord(t_time, nftData.HardInfo, 2, t.Hash().String())
	c.JSON(200, gin.H{
		"result": t.Hash(),
		"error":  nil,
	})
}

func DeleteNft(c *gin.Context) {
	var nftData PutNftRequest
	err := c.Bind(&nftData)
	if err != nil {
		c.JSON(200, gin.H{
			"result": nil,
			"error":  err,
		})
		return
	}
	var nftAddress string
	mintType := nftData.NftType
	if mintType == 1 {
		nftAddress = constants.NftCpuAddress

	} else if mintType == 2 {
		nftAddress = constants.NftMemAddress
	} else if mintType == 3 {
		nftAddress = constants.NftStorageAddress
	} else if mintType == 4 {
		nftAddress = constants.NftGpuAddress
	} else if mintType == 5 {
		nftAddress = constants.NftIpAddress
	} else if mintType == 6 {
		nftAddress = constants.NftBandWidth
	} else {

		c.JSON(200, gin.H{
			"result": nil,
			"error":  "nft type not exsit",
		})
		return

	}

	contractAddress := localstore.Get(constants.ContractAddressKey)
	t, err := rpcService.DownNft(contractAddress, nftAddress, nftData.TokenId)
	if err != nil {
		c.JSON(200, gin.H{
			"result": nil,
			"error":  err,
		})
		return
	}
	t_time := time.Now().UnixMilli()
	localstore.Put(constants.HASH_CHECK_KEY_PRIFIX+t.Hash().String()+"-"+"del", strconv.Itoa(mintType)+"-"+strconv.Itoa(nftData.TokenId))
	localstore.AddNftRecord(t_time, nftData.HardInfo, 3, t.Hash().String())
	c.JSON(200, gin.H{
		"result": t.Hash(),
		"error":  nil,
	})
}

func GetNftList(c *gin.Context) {
	result := localstore.GetNftRecordList()
	c.JSON(200, gin.H{
		"result": result,
		"error":  nil,
	})

}
func GetAllNftList(c *gin.Context) {
	tokenMap, err := rpcService.GetAllTokenIds()
	if err != nil {
		c.JSON(200, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}
	var resultList = []NftItemData{}

	cpuIds := tokenMap["cpu"]
	memIds := tokenMap["mem"]
	storageIds := tokenMap["storage"]
	ipIds := tokenMap["ip"]
	brandWidthIds := tokenMap["brandWidth"]
	gpuIds := tokenMap["gpu"]
	if len(cpuIds) > 0 {
		for i := 0; i < len(cpuIds); i++ {

			id := cpuIds[i]
			info, _ := rpcService.GetHardInfo(constants.NftCpuAddress, id)
			status, _ := rpcService.GetNftStatus(constants.NftCpuAddress, id)

			item := NftItemData{NftTypeName: "CPU", Info: info, Status: status, TokenId: id, NftType: 1}
			resultList = append(resultList, item)

		}

	}
	if len(memIds) > 0 {
		for i := 0; i < len(memIds); i++ {

			id := memIds[i]
			info, _ := rpcService.GetHardInfo(constants.NftMemAddress, id)
			status, _ := rpcService.GetNftStatus(constants.NftMemAddress, id)

			item := NftItemData{NftTypeName: "RAM", Info: info, Status: status, TokenId: id, NftType: 2}
			resultList = append(resultList, item)

		}

	}
	if len(storageIds) > 0 {
		for i := 0; i < len(storageIds); i++ {

			id := storageIds[i]
			info, _ := rpcService.GetHardInfo(constants.NftStorageAddress, id)
			status, _ := rpcService.GetNftStatus(constants.NftStorageAddress, id)

			item := NftItemData{NftTypeName: "STORAGE", Info: info, Status: status, TokenId: id, NftType: 3}
			resultList = append(resultList, item)

		}

	}
	if len(ipIds) > 0 {
		for i := 0; i < len(ipIds); i++ {

			id := ipIds[i]
			info, _ := rpcService.GetHardInfo(constants.NftIpAddress, id)
			status, _ := rpcService.GetNftStatus(constants.NftIpAddress, id)

			item := NftItemData{NftTypeName: "IP", Info: info, Status: status, TokenId: id, NftType: 5}
			resultList = append(resultList, item)

		}

	}
	if len(brandWidthIds) > 0 {
		for i := 0; i < len(brandWidthIds); i++ {

			id := brandWidthIds[i]
			info, _ := rpcService.GetHardInfo(constants.NftBandWidth, id)
			status, _ := rpcService.GetNftStatus(constants.NftBandWidth, id)

			item := NftItemData{NftTypeName: "BANDWIDTH", Info: info, Status: status, TokenId: id, NftType: 6}
			resultList = append(resultList, item)

		}

	}
	if len(gpuIds) > 0 {
		for i := 0; i < len(gpuIds); i++ {

			id := gpuIds[i]
			info, _ := rpcService.GetHardInfo(constants.NftGpuAddress, id)
			status, _ := rpcService.GetNftStatus(constants.NftGpuAddress, id)

			item := NftItemData{NftTypeName: "GPU", Info: info, Status: status, TokenId: id, NftType: 4}
			resultList = append(resultList, item)

		}

	}

	c.JSON(200, gin.H{
		"result": resultList,
		"error":  nil,
	})
}

type Info struct {
	Info    localstore.HardwareInfo `json:"info"`
	OrderId string                  `json:"orderId"`
}

func (info Info) ToString() string {
	res, err := json.Marshal(info)
	if err != nil {
		return err.Error()
	}

	return string(res)
}

func GetNftDetail(c *gin.Context) {
	resultStr := localstore.Get(constants.DEVICE_INFO_KEY)
	hardwareInfo := localstore.ToInfo(resultStr)
	info := Info{Info: hardwareInfo, OrderId: localstore.MemGet(constants.MEM_CurrentOrderIdKey)}
	c.JSON(200, gin.H{
		"result": info.ToString(),
		"error":  nil,
	})
}

func DelOrder(c *gin.Context) {
	ls := localstore.FindListRange(constants.ORDER_TEMPLATE_ID_RECEIVED)
	for k, _ := range ls {
		localstore.Remove(k)
	}
	c.JSON(200, gin.H{
		"result": "1",
		"error":  nil,
	})
}

func ReleaseAll(c *gin.Context) {
	service.ReleaseAll()
	c.JSON(200, gin.H{
		"result": "1",
		"error":  nil,
	})
}
