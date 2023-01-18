package service

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
	"ubiqnet/constants"
	"ubiqnet/store/localstore"
	"ubiqnet/tools"

	logging "github.com/ipfs/go-log"
)

var NewOrderTemplateTopic string = "newOrderTemplate"

var taskLog = logging.Logger("service:validator-newOrder")
var rpcService RpcService

type ImageConfig struct {
	ImageName string            `json:"imageName"`
	Cmd       string            `json:"cmd"`
	Location  string            `json:"location"`
	Env       []string          `json:"env"`
	Portmap   map[string]string `json:"portmap"`
}

func ToImageConfig(str string) ImageConfig {
	var p2 ImageConfig
	err := json.Unmarshal([]byte(str), &p2)
	if err != nil {
		taskLog.Error(err)
		return ImageConfig{}
	}
	return p2
}

func CheckOrder() {
	idstr := localstore.MemGet(constants.MEM_CurrentOrderIdKey)
	if idstr == "" {
		idstr = strconv.Itoa(rpcService.GetLatestOrderId())
		localstore.MemPut(constants.MEM_CurrentOrderIdKey, idstr)
	}
	id, _ := strconv.Atoi(idstr)
	var orderInfo, err = rpcService.GetOrderById(int64(id))
	if err != nil || orderInfo == nil {
		if err != nil {
			taskLog.Error(err)
		}
		return
	}

	if orderInfo.CreateTime.Int64() == int64(0) {

		return
	}
	orderExpireTime := orderInfo.CreateTime.Int64() + int64(5*60)
	if orderExpireTime < time.Now().UTC().Unix() || orderInfo.CreateTime.Int64() == int64(0) {

		id = id + 1
		localstore.MemPut(constants.MEM_CurrentOrderIdKey, strconv.Itoa(id))
		return
	}
	forCheckId := id

	go func() {
		imageInfos, err := rpcService.GetAppInfo(forCheckId, orderInfo.AppAddress)
		if err != nil {
			taskLog.Error(err)
			return
		}
		ownerAddress, err := rpcService.GetAppOwner(orderInfo.AppAddress)
		if err != nil {
			taskLog.Error(err)
			return
		}
		var validateSuitInfo = new(ValidatorSuitInfo)
		validateSuitInfo.OrderId = forCheckId
		validateSuitInfo.AppOwnerAddress = ownerAddress
		validateSuitInfo.AppAddress = orderInfo.AppAddress.String()
		validateSuitInfo.OrderExpireTime = orderExpireTime
		validateSuitInfo.NodeSuitInfos = make([]NodeSuiteInfo, len(imageInfos))
		for i := 0; i < len(imageInfos); i++ {
			imageInfo := imageInfos[i]
			var nodeSuitInfo = new(NodeSuiteInfo)
			nodeSuitInfo.OrderId = imageInfo.OrderId
			nodeSuitInfo.TemplateAddress = imageInfo.TemplateAddress
			imageInfo.CreateTimeStamp = orderInfo.CreateTime.Int64()
			nodeSuitInfo.Index = i
			validateSuitInfo.NodeSuitInfos[i] = *nodeSuitInfo
			validateSuitInfo.FinishedTimestampNano = int64(0)
			validateSuitInfo.ValidateAddress = localstore.Get(constants.ContractAddressKey)
			p2pData := P2pData{Sender: peerId, Event: NewOrderTemplateTopic, Extra: imageInfo.ToString(), TimeStamp: time.Now().UnixMilli()}
			p2pBytesData, _ := json.Marshal(&p2pData)
			err := P2p.PublishBroadCastData(p2pBytesData)
			if err != nil {
				taskLog.Error(err)
				return
			}

		}
		localstore.MemPut(constants.MEM_ORDER_CHECK_PREFIX+strconv.Itoa(forCheckId), validateSuitInfo.ToString())
	}()
	id = id + 1
	localstore.MemPut(constants.MEM_CurrentOrderIdKey, strconv.Itoa(id))
}

func checkOrder(orderId int, ownerAddress string) {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		var orderInfo, err = rpcService.GetOrderById(int64(orderId))
		if err != nil || orderInfo == nil {
			taskLog.Error(err)
			continue
		}
		orderExpireTime := orderInfo.CreateTime.Int64() + int64(5*60)

		if orderExpireTime < time.Now().UTC().Unix() {

			ticker.Stop()
			return
		}
		if orderInfo.StartTime.Int64() == int64(0) {

			continue
		}
		res, err := rpcService.GetAppInfoDetail(orderId, *orderInfo, ownerAddress)
		if err == nil {
			for _, i := range res {
				p2pData := P2pData{Sender: peerId, Event: StartContainer, Extra: i.ToString(), TimeStamp: time.Now().UnixMilli()}
				p2pBytesData, _ := json.Marshal(&p2pData)
				err := P2p.PublishBroadCastData(p2pBytesData)
				if err != nil {
					taskLog.Error(err)
				}
			}
		} else {
			taskLog.Error(err)
		}
		ticker.Stop()
		return
	}
}

func CheckHash() {
	hashList := localstore.FindListRange(constants.HASH_CHECK_KEY_PRIFIX)
	if len(hashList) == 0 {
		return
	}
	for key, value := range hashList {
		var hashActionstr = strings.Replace(key, constants.HASH_CHECK_KEY_PRIFIX, "", 1)
		hashAction := strings.Split(hashActionstr, "-")
		typeTokenid := strings.Split(value, "-")
		var nftType, err = strconv.Atoi(typeTokenid[0])
		var hash = hashAction[0]
		var action = hashAction[1]
		if err != nil {
			continue
		}
		tokenId, err := strconv.Atoi(typeTokenid[1])
		if err != nil {
			continue
		}

		res, err := rpcService.checkTransaction(hash)
		if err != nil {
			continue
		}
		if res {
			onOffNFT(action, nftType, tokenId)
		}
		localstore.Remove(key)
	}
}
func onOffNFT(action string, nftType int, tokenId int) {
	Lock.Lock()
	defer Lock.Unlock()
	var hardwareInfo = new(localstore.HardwareInfo)
	deviceInfoStr := localstore.Get(constants.DEVICE_INFO_KEY)
	if deviceInfoStr != "" {
		*hardwareInfo = localstore.ToInfo(deviceInfoStr)
	} else {
		hardwareInfo.GpuInfo = tools.DeviceInfo.GpuInfo
		hardwareInfo.City = tools.DeviceInfo.Netinfo.City
		hardwareInfo.CountryCode = tools.DeviceInfo.Netinfo.CountryCode
		hardwareInfo.IP = tools.DeviceInfo.Netinfo.IP
		hardwareInfo.OS = tools.DeviceInfo.OS
		hardwareInfo.Region = tools.DeviceInfo.Netinfo.RegionName
		hardwareInfo.CpuSpeed = tools.DeviceInfo.CpuSpeed
		hardwareInfo.IsPublic = tools.DeviceInfo.Netinfo.IsPublic
		hardwareInfo.TotalCpuCore = tools.DeviceInfo.NumOfCores
		hardwareInfo.TotalDisk = tools.DeviceInfo.FreeDisk
		hardwareInfo.TotalMemory = tools.DeviceInfo.TotalMemory
	}
	switch nftType {
	case 1:
		if action == "put" {
			hardwareInfo.CpuTokens = append(hardwareInfo.CpuTokens, int64(tokenId))
			hardwareInfo.CpuTokensFree = append(hardwareInfo.CpuTokensFree, int64(tokenId))
		} else {
			hardwareInfo.CpuTokens = tools.DeleteSlice(hardwareInfo.CpuTokens, int64(tokenId))
			hardwareInfo.CpuTokensFree = tools.DeleteSlice(hardwareInfo.CpuTokensFree, int64(tokenId))
		}
		break
	case 2:
		if action == "put" {
			hardwareInfo.MemTokens = append(hardwareInfo.MemTokens, int64(tokenId))
			hardwareInfo.MemTokensFree = append(hardwareInfo.MemTokensFree, int64(tokenId))
		} else {
			hardwareInfo.MemTokens = tools.DeleteSlice(hardwareInfo.MemTokens, int64(tokenId))
			hardwareInfo.MemTokensFree = tools.DeleteSlice(hardwareInfo.MemTokensFree, int64(tokenId))
		}
		break
	case 3:
		if action == "put" {
			hardwareInfo.DiskTokens = append(hardwareInfo.DiskTokens, int64(tokenId))
			hardwareInfo.DiskTokensFree = append(hardwareInfo.DiskTokensFree, int64(tokenId))
		} else {
			hardwareInfo.DiskTokens = tools.DeleteSlice(hardwareInfo.DiskTokens, int64(tokenId))
			hardwareInfo.DiskTokensFree = tools.DeleteSlice(hardwareInfo.DiskTokensFree, int64(tokenId))
		}
		break
	case 4:
		if action == "put" {
			hardwareInfo.GpuTokens = append(hardwareInfo.GpuTokens, int64(tokenId))
			hardwareInfo.GpuTokensFree = append(hardwareInfo.GpuTokensFree, int64(tokenId))
		} else {
			hardwareInfo.GpuTokens = tools.DeleteSlice(hardwareInfo.GpuTokens, int64(tokenId))
			hardwareInfo.GpuTokensFree = tools.DeleteSlice(hardwareInfo.GpuTokensFree, int64(tokenId))
		}
		break
	case 5:
		if action == "put" {
			hardwareInfo.IpTokens = append(hardwareInfo.IpTokens, int64(tokenId))
			hardwareInfo.IpTokensFree = append(hardwareInfo.IpTokensFree, int64(tokenId))
		} else {
			hardwareInfo.IpTokens = tools.DeleteSlice(hardwareInfo.IpTokens, int64(tokenId))
			hardwareInfo.IpTokensFree = tools.DeleteSlice(hardwareInfo.IpTokensFree, int64(tokenId))
		}
		break
	case 6:
		if action == "put" {
			hardwareInfo.BandwidthTokens = append(hardwareInfo.BandwidthTokens, int64(tokenId))
			hardwareInfo.BandwidthTokensFree = append(hardwareInfo.BandwidthTokensFree, int64(tokenId))
		} else {
			hardwareInfo.BandwidthTokens = tools.DeleteSlice(hardwareInfo.BandwidthTokens, int64(tokenId))
			hardwareInfo.BandwidthTokensFree = tools.DeleteSlice(hardwareInfo.BandwidthTokensFree, int64(tokenId))
		}
		break

	}
	localstore.Put(constants.DEVICE_INFO_KEY, hardwareInfo.ToString())

}
