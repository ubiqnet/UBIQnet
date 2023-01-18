package service

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
	"strings"
	"time"
	"ubiqnet/constants"
	"ubiqnet/store/localstore"
	"ubiqnet/tools"
)

type NftService struct {
}

var Nft NftService

type NftCount struct {
	Cpu        int `json:"cpu"`
	Mem        int `json:"mem"`
	Storage    int `json:"storage"`
	Gpu        int `json:"gpu"`
	Ip         int `json:"ip"`
	BrandWidth int `json:"brandWidth"`
}

func (service NftService) MintNfts(mintType int, info string) (string, error) {

	contractAddress := localstore.Get(constants.ContractAddressKey)
	var nftCount NftCount
	nftCountStr := localstore.Get(constants.NftCountKey)
	json.Unmarshal([]byte(nftCountStr), &nftCount)

	cpu := 0
	mem := 0
	gpu := 0
	storage := 0
	ip := 0
	brandWidth := 0

	infos := strings.Split(info, " ")
	cpuTotal := tools.DeviceInfo.NumOfCores
	memTotal := int(math.Ceil(float64(tools.DeviceInfo.FreeMemory/1024/1024) / 1024))
	storageTotal := int(math.Ceil(float64(tools.DeviceInfo.FreeDisk/1024/1024) / 1024))

	gpuTotal := len(tools.DeviceInfo.GpuInfo)

	ipTotal := 1
	brandWidthTotal := math.Ceil(tools.DeviceInfo.DownloadSpeed)

	nftAddress := ""
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
	}
	if nftAddress == "" {
		return "", errors.New("nft type not exsit")
	}
	if mintType == 1 {
		cpu, _ = strconv.Atoi(infos[1])
		if nftCount.Cpu+cpu > cpuTotal {
			return "", errors.New("not enough cpu resources")
		}
	} else if mintType == 2 {
		mem, _ = strconv.Atoi(infos[1])
		if nftCount.Mem+mem > memTotal {
			return "", errors.New("not enough Memory resources")
		}
	} else if mintType == 3 {
		storage, _ = strconv.Atoi(infos[1])
		if nftCount.Storage+storage > storageTotal {

			return "", errors.New("not enough Storage resources")
		}
	} else if mintType == 4 {
		// gpu, _ = strconv.Atoi(infos[1])
		gpu = 1
		if nftCount.Gpu+gpu > gpuTotal {
			return "", errors.New("not enough Gpu resources")
		}
	} else if mintType == 5 {
		ip = 1
		info = tools.DeviceInfo.Netinfo.IP
		if nftCount.Ip+ip > ipTotal {
			return "", errors.New("not enough ip resources")
		}
	} else if mintType == 6 {
		brandWidth, _ = strconv.Atoi(infos[1])
		if nftCount.BrandWidth+brandWidth > int(brandWidthTotal) {

			return "", errors.New("not enough bandwidth resources")
		}
	}

	t, err := rpcService.MintNft(contractAddress, nftAddress, info)
	if err != nil {

		return "", err
	}

	if mintType == 1 {
		nftCount.Cpu += cpu
	} else if mintType == 2 {
		nftCount.Mem += mem
	} else if mintType == 3 {
		nftCount.Storage += storage
	} else if mintType == 4 {
		nftCount.Gpu += gpu
	} else if mintType == 5 {
		nftCount.Ip += ip
	} else if mintType == 6 {
		nftCount.BrandWidth += brandWidth
	}
	data, _ := json.Marshal(&nftCount)
	localstore.Put(constants.NftCountKey, string(data))

	t_time := time.Now().UnixMilli()
	localstore.AddNftRecord(t_time, info, 1, t.Hash().String())
	return t.Hash().String(), nil
}
