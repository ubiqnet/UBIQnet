package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strconv"
	"strings"
	"sync"
	"time"
	"ubiqnet/constants"
	"ubiqnet/store/localstore"
	"ubiqnet/tools"
)

var TemplateSuited = "TEMPLATE_SUITED"
var StartContainer = "START_CONTAINER"
var StopContainer = "STOP_CONTAINER"
var TemplateInfoQueue = tools.NewQueue()
var OtherTemplateQueue = tools.NewQueue()
var StartImageQueue = tools.NewQueue()
var StopImageQueue = tools.NewQueue()

var Lock sync.Mutex

func ReceiveStartDockerQueueConsumer() {
	StartImageQueue.Run(func(item tools.QueueItem) {
		imageInfo, ok := (item).(ImageInfo)
		if ok {
			err := StartImage(imageInfo)
			if err != nil {
				taskLog.Info("start container failed", err)
				return
			}
		}
	})
}

func ReceiveStopDockerQueueConsumer() {
	StopImageQueue.Run(func(item tools.QueueItem) {
		imageInfo, ok := (item).(ImageInfo)
		if ok {
			CloseImage(imageInfo)
		}
	})
}

func ReceiveQueueConsumer() {
	TemplateInfoQueue.Run(func(item tools.QueueItem) {
		imageInfo, ok := (item).(ImageInfo)
		if ok {
			receiveOrder(imageInfo)
		}

	})
}
func ReceiveOtherTemplateQueueConsumer() {
	OtherTemplateQueue.Run(func(item tools.QueueItem) {
		nodeSuitInfo, ok := (item).(NodeSuiteInfo)
		if ok {
			receiveOtherSuite(nodeSuitInfo)
		}

	})
}

func CheckContainerNeedStop() {
	currentTimestamp := time.Now().UTC().Unix()
	infoS := localstore.FindListRange(constants.STOP_CONTAONER_KEY_PRIFIX)
	if len(infoS) != 0 {
		for _, infoStr := range infoS {
			info := ToImageInfo(infoStr)
			if info.ExpireTimeStamp <= currentTimestamp {
				//var name = strings.Replace(key, constants.STOP_CONTAONER_KEY_PRIFIX, "", 1)
				CloseImage(info)
			}
		}
	}
}

func receiveOtherSuite(info NodeSuiteInfo) {

	if strings.EqualFold(info.NodeAddress, localstore.Get(constants.ContractAddressKey)) {
		return
	}
	timestampNanoStr := localstore.MemGet(strconv.Itoa(info.OrderId) + info.TemplateAddress)
	if timestampNanoStr != "" {
		currentTimestampNano, err := strconv.ParseInt(timestampNanoStr, 10, 64)
		if err == nil {
			if currentTimestampNano == info.SuitUtcTimestampNano {
				if localstore.Get(constants.ContractAddressKey) > info.NodeAddress {
					releaseLocked(info.OrderId, info.Index)
				}
			} else if currentTimestampNano > info.SuitUtcTimestampNano {
				releaseLocked(info.OrderId, info.Index)
			}
		}
	}

}

func arrayTake(freeArray []int64, takeCount *int) ([]int64, []int64) {
	var takeArray = make([]int64, *takeCount)
	if *takeCount == 0 {
		return freeArray, takeArray
	}
	takeArray = freeArray[:*takeCount]
	newFreeArray := freeArray[*takeCount:]
	return newFreeArray, takeArray

}

func StartImage(info ImageInfo) error {
	if strings.EqualFold(info.NodeAddress, localstore.Get(constants.ContractAddressKey)) {
		localstore.Remove(constants.MEM_DEVICE_LOCKED_TOKEN_PREFIX + strconv.Itoa(info.OrderId) + "-" + strconv.Itoa(info.Index))
		localstore.MemRemove(constants.MEM_DEVICE_LOCKED_PREFIX + strconv.Itoa(info.OrderId) + "-" + strconv.Itoa(info.Index))
		imageName := strconv.Itoa(info.OrderId) + info.TemplateAddress + strconv.Itoa(info.Index)
		defer localstore.Put(constants.STOP_CONTAONER_KEY_PRIFIX+imageName, info.ToString())
		err := tools.Retry(10, 1*time.Second, func() error {
			return tools.PullImage(info.ImageName)

		})
		if err != nil {
			return err
		}
		containerPS := tools.GetContainers()
		for _, i := range containerPS {
			if i.Names[0] == imageName {
				if strings.EqualFold(i.Status, "exited") {
					err := tools.RemoveContainer(i.Id)
					if err != nil {
						taskLog.Error(err)
						return err
					}
				} else {
					return errors.New("already has same container")
				}

			}
		}
		err = tools.Retry(3, 2*time.Second, func() error {
			taskLog.Info("now start container ", info.ToString())
			_, err = tools.CreateContainer(info.ImageName, imageName, info.Cmd, info.Env, info.MemCount*2, info.CpuCount*2, info.PortMap, info.StorageCount*10, info.GpuCount)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}

	}
	return nil
}

func CloseImage(info ImageInfo) {

	imageName := strconv.Itoa(info.OrderId) + info.TemplateAddress + strconv.Itoa(info.Index)
	tools.RemoveContainerByName(imageName)
	var hardwareInfo = new(localstore.HardwareInfo)
	Lock.Lock()
	defer Lock.Unlock()
	deviceInfoStr := localstore.Get(constants.DEVICE_INFO_KEY)
	*hardwareInfo = localstore.ToInfo(deviceInfoStr)
	taskLog.Info("close image:", info.ToString())
	if len(info.CpuIds) > 0 {
		for _, id := range info.CpuIds {
			if hardwareInfo.CpuTokensFree == nil {
				hardwareInfo.CpuTokensFree = make([]int64, 0)
			}
			if tools.InSlice(hardwareInfo.CpuTokens, id) {
				hardwareInfo.CpuTokensFree = tools.RemoveRep(append(hardwareInfo.CpuTokensFree, id))
			}
		}
	}
	if len(info.IpTokenIds) > 0 {
		for _, id := range info.IpTokenIds {
			if hardwareInfo.IpTokensFree == nil {
				hardwareInfo.IpTokensFree = make([]int64, 0)
			}
			if tools.InSlice(hardwareInfo.IpTokens, id) {
				hardwareInfo.IpTokensFree = tools.RemoveRep(append(hardwareInfo.IpTokensFree, id))
			}
		}
	}
	if len(info.BandwidthIds) > 0 {
		if hardwareInfo.BandwidthTokensFree == nil {
			hardwareInfo.BandwidthTokensFree = make([]int64, 0)
		}
		for _, id := range info.BandwidthIds {
			if tools.InSlice(hardwareInfo.BandwidthTokens, id) {
				hardwareInfo.BandwidthTokensFree = tools.RemoveRep(append(hardwareInfo.BandwidthTokensFree, id))
			}
		}
	}
	if len(info.StorageIds) > 0 {
		if hardwareInfo.DiskTokensFree == nil {
			hardwareInfo.DiskTokensFree = make([]int64, 0)
		}
		for _, id := range info.StorageIds {
			if tools.InSlice(hardwareInfo.DiskTokens, id) {
				hardwareInfo.DiskTokensFree = tools.RemoveRep(append(hardwareInfo.DiskTokensFree, id))
			}
		}
	}
	if len(info.GpuIds) > 0 {
		if hardwareInfo.GpuTokensFree == nil {
			hardwareInfo.GpuTokensFree = make([]int64, 0)
		}
		for _, id := range info.GpuIds {
			if tools.InSlice(hardwareInfo.GpuTokens, id) {
				hardwareInfo.GpuTokensFree = tools.RemoveRep(append(hardwareInfo.GpuTokensFree, id))
			}
		}
	}
	if len(info.MemIds) > 0 {
		if hardwareInfo.MemTokensFree == nil {
			hardwareInfo.MemTokensFree = make([]int64, 0)
		}
		for _, id := range info.MemIds {
			if tools.InSlice(hardwareInfo.MemTokens, id) {
				hardwareInfo.MemTokensFree = tools.RemoveRep(append(hardwareInfo.MemTokensFree, id))
			}
		}
	}
	localstore.Put(constants.DEVICE_INFO_KEY, hardwareInfo.ToString())
	localstore.Remove(constants.STOP_CONTAONER_KEY_PRIFIX + imageName)
}

func receiveOrder(imageInfo ImageInfo) {
	taskLog.Info("receive Order", imageInfo.ToString())
	b, err := localstore.Contains(constants.ORDER_TEMPLATE_ID_RECEIVED + strconv.Itoa(imageInfo.OrderId) + "-" + strconv.Itoa(imageInfo.Index))
	if err != nil {
		taskLog.Error(err)
	}
	if b {

		return
	}
	localstore.Put(constants.ORDER_TEMPLATE_ID_RECEIVED+strconv.Itoa(imageInfo.OrderId)+"-"+strconv.Itoa(imageInfo.Index), "1")
	localtimestamp := time.Now().UTC().Unix()
	if (imageInfo.CreateTimeStamp + int64(5*60)) < localtimestamp {
		taskLog.Info("ReceiveOrder timeout")
		return
	}
	Lock.Lock()
	defer Lock.Unlock()
	deviceInfoStr := localstore.Get(constants.DEVICE_INFO_KEY)
	var hardwareInfo = localstore.ToInfo(deviceInfoStr)
	var ipSuit = len(hardwareInfo.IpTokensFree) >= imageInfo.IpTokenCount
	if !ipSuit {
		return
	}
	var bandSuit = len(hardwareInfo.BandwidthTokensFree) >= imageInfo.BandwidthCount
	if !bandSuit {
		return
	}
	if imageInfo.Location != "" && imageInfo.Location != "Unknown" {
		if imageInfo.Location != tools.DeviceInfo.Netinfo.RegionName {
			return
		}
	}
	var gpuSuit = len(hardwareInfo.GpuTokensFree) >= imageInfo.GpuCount
	if !gpuSuit {
		return
	}
	var storageSuit = len(hardwareInfo.DiskTokensFree) >= imageInfo.StorageCount
	if !storageSuit {
		return
	}
	var cpuSuit = len(hardwareInfo.CpuTokensFree) >= imageInfo.CpuCount
	if !cpuSuit {
		return
	}
	var memSuit = len(hardwareInfo.MemTokensFree) >= imageInfo.MemCount
	if !memSuit {
		return
	}

	ipTokenLock := make([]int64, imageInfo.IpTokenCount)
	hardwareInfo.IpTokensFree, ipTokenLock = arrayTake(hardwareInfo.IpTokensFree, &imageInfo.IpTokenCount)

	bandTokenLock := make([]int64, imageInfo.BandwidthCount)
	hardwareInfo.BandwidthTokensFree, bandTokenLock = arrayTake(hardwareInfo.BandwidthTokensFree, &imageInfo.BandwidthCount)

	gpuTokenLock := make([]int64, imageInfo.GpuCount)
	hardwareInfo.GpuTokensFree, gpuTokenLock = arrayTake(hardwareInfo.GpuTokensFree, &imageInfo.GpuCount)

	storageTokenLock := make([]int64, imageInfo.StorageCount)
	hardwareInfo.DiskTokensFree, storageTokenLock = arrayTake(hardwareInfo.DiskTokensFree, &imageInfo.StorageCount)

	cpuTokenLock := make([]int64, imageInfo.CpuCount)
	hardwareInfo.CpuTokensFree, cpuTokenLock = arrayTake(hardwareInfo.CpuTokensFree, &imageInfo.CpuCount)

	memTokenLock := make([]int64, imageInfo.MemCount)
	hardwareInfo.MemTokensFree, memTokenLock = arrayTake(hardwareInfo.MemTokensFree, &imageInfo.MemCount)

	var nodeSuitInfo = new(NodeSuiteInfo)
	nodeSuitInfo.OrderId = imageInfo.OrderId
	nodeSuitInfo.NodeAddress = localstore.Get(constants.ContractAddressKey)
	nodeSuitInfo.IpTokenIds = ipTokenLock
	nodeSuitInfo.BandwidthTokenIds = bandTokenLock
	nodeSuitInfo.GpuTokenIds = gpuTokenLock
	nodeSuitInfo.StorageTokenIds = storageTokenLock
	nodeSuitInfo.CpuTokenIds = cpuTokenLock
	nodeSuitInfo.MemTokenIds = memTokenLock
	nodeSuitInfo.SuitUtcTimestampNano = time.Now().UTC().UnixNano()
	nodeSuitInfo.ExpireTimestamp = time.Now().UTC().Unix() + int64(5*60)
	localstore.Put(constants.DEVICE_INFO_KEY, hardwareInfo.ToString())
	localstore.Put(constants.MEM_DEVICE_LOCKED_TOKEN_PREFIX+strconv.Itoa(imageInfo.OrderId)+"-"+strconv.Itoa(imageInfo.Index), nodeSuitInfo.ToString())
	localstore.MemPut(constants.MEM_DEVICE_LOCKED_PREFIX+strconv.Itoa(imageInfo.OrderId)+imageInfo.TemplateAddress, strconv.FormatInt(nodeSuitInfo.SuitUtcTimestampNano, 10))

	p2pData := P2pData{Sender: peerId, Event: TemplateSuited, Extra: nodeSuitInfo.ToString(), TimeStamp: time.Now().UnixMilli()}
	p2pBytesData, _ := json.Marshal(&p2pData)
	err = P2p.PublishBroadCastData(p2pBytesData)

	if err != nil {
		taskLog.Error(err)
	}
	go checkRelease(imageInfo)

}

func checkRelease(imageInfo ImageInfo) {
	timer := time.NewTimer(5 * time.Minute)
	select {
	case <-timer.C:
		con := localstore.MemContains(constants.MEM_DEVICE_LOCKED_PREFIX + strconv.Itoa(imageInfo.OrderId) + imageInfo.TemplateAddress)
		if con {
			releaseLocked(imageInfo.OrderId, imageInfo.Index)
		}

	}
}

func ReleaseAll() {
	deviceInfoStr := localstore.Get(constants.DEVICE_INFO_KEY)
	var hardwareInfo = localstore.ToInfo(deviceInfoStr)
	hardwareInfo.IpTokensFree = tools.RemoveRep(hardwareInfo.IpTokens)
	hardwareInfo.IpTokens = hardwareInfo.IpTokensFree

	hardwareInfo.BandwidthTokensFree = tools.RemoveRep(hardwareInfo.BandwidthTokens)
	hardwareInfo.BandwidthTokens = hardwareInfo.BandwidthTokensFree

	hardwareInfo.GpuTokensFree = tools.RemoveRep(hardwareInfo.GpuTokens)
	hardwareInfo.GpuTokens = hardwareInfo.GpuTokensFree

	hardwareInfo.DiskTokensFree = tools.RemoveRep(hardwareInfo.DiskTokens)
	hardwareInfo.DiskTokens = hardwareInfo.DiskTokensFree

	hardwareInfo.CpuTokensFree = tools.RemoveRep(hardwareInfo.CpuTokens)
	hardwareInfo.CpuTokens = hardwareInfo.CpuTokensFree

	hardwareInfo.MemTokensFree = tools.RemoveRep(hardwareInfo.MemTokens)
	hardwareInfo.MemTokens = hardwareInfo.MemTokensFree

	localstore.Put(constants.DEVICE_INFO_KEY, hardwareInfo.ToString())
}

func AutoRelease() {
	maps := localstore.FindListRange(constants.MEM_DEVICE_LOCKED_TOKEN_PREFIX)
	if len(maps) != 0 {
		for k, v := range maps {
			nodeSuiteInfo := ToNodeSuitInfo(v)
			if nodeSuiteInfo.ExpireTimestamp != int64(0) && nodeSuiteInfo.ExpireTimestamp < time.Now().UTC().Unix() {
				k = strings.Replace(k, constants.MEM_DEVICE_LOCKED_TOKEN_PREFIX, "", 1)
				ids := strings.Split(k, "-")
				orderId, _ := strconv.Atoi(ids[0])
				index, _ := strconv.Atoi(ids[1])
				releaseLocked(orderId, index)
			}
		}
	}
	deviceInfoStr := localstore.Get(constants.DEVICE_INFO_KEY)
	Lock.Lock()
	defer Lock.Unlock()
	var hardwareInfo = localstore.ToInfo(deviceInfoStr)
	var rpcService RpcService
	hardwareInfo.IpTokensFree = checkTokenRent(rpcService, hardwareInfo.IpTokensFree, constants.NftIpAddress)
	hardwareInfo.BandwidthTokensFree = checkTokenRent(rpcService, hardwareInfo.BandwidthTokensFree, constants.NftBandWidth)
	hardwareInfo.CpuTokensFree = checkTokenRent(rpcService, hardwareInfo.CpuTokensFree, constants.NftCpuAddress)
	hardwareInfo.GpuTokensFree = checkTokenRent(rpcService, hardwareInfo.GpuTokensFree, constants.NftGpuAddress)
	hardwareInfo.DiskTokensFree = checkTokenRent(rpcService, hardwareInfo.DiskTokensFree, constants.NftStorageAddress)
	hardwareInfo.MemTokensFree = checkTokenRent(rpcService, hardwareInfo.MemTokensFree, constants.NftMemAddress)
	localstore.Put(constants.DEVICE_INFO_KEY, hardwareInfo.ToString())
}

func checkTokenRent(rpcService RpcService, tokenFree []int64, contractAddress string) []int64 {
	var freeTokens = tokenFree
	for _, i2 := range tokenFree {
		isRent, err := rpcService.CheckTokenIsRent(i2, contractAddress)
		if err != nil {
			taskLog.Error(err)
			return freeTokens
		}
		if isRent {
			freeTokens = tools.DeleteSlice(freeTokens, i2)
		}
	}
	return freeTokens
}

func releaseLocked(orderId int, index int) {

	Lock.Lock()
	defer Lock.Unlock()
	deviceInfoStr := localstore.Get(constants.DEVICE_INFO_KEY)
	nodeSuiteInfoStr := localstore.Get(constants.MEM_DEVICE_LOCKED_TOKEN_PREFIX + strconv.Itoa(orderId) + "-" + strconv.Itoa(index))
	if nodeSuiteInfoStr != "" {

		nodeSuiteInfo := ToNodeSuitInfo(nodeSuiteInfoStr)
		var hardwareInfo = localstore.ToInfo(deviceInfoStr)
		hardwareInfo.IpTokensFree = tools.RemoveRep(append(hardwareInfo.IpTokensFree, nodeSuiteInfo.IpTokenIds...))
		hardwareInfo.BandwidthTokensFree = tools.RemoveRep(append(hardwareInfo.BandwidthTokensFree, nodeSuiteInfo.BandwidthTokenIds...))
		hardwareInfo.GpuTokensFree = tools.RemoveRep(append(hardwareInfo.GpuTokensFree, nodeSuiteInfo.GpuTokenIds...))
		hardwareInfo.DiskTokensFree = tools.RemoveRep(append(hardwareInfo.DiskTokensFree, nodeSuiteInfo.StorageTokenIds...))
		hardwareInfo.CpuTokensFree = tools.RemoveRep(append(hardwareInfo.CpuTokensFree, nodeSuiteInfo.CpuTokenIds...))
		hardwareInfo.MemTokensFree = tools.RemoveRep(append(hardwareInfo.MemTokensFree, nodeSuiteInfo.MemTokenIds...))
		localstore.Remove(constants.MEM_DEVICE_LOCKED_TOKEN_PREFIX + strconv.Itoa(orderId) + "-" + strconv.Itoa(index))
		localstore.Put(constants.DEVICE_INFO_KEY, hardwareInfo.ToString())
		localstore.MemRemove(constants.MEM_DEVICE_LOCKED_PREFIX + strconv.Itoa(orderId) + "-" + strconv.Itoa(index))
	}
}

func VerifyMessage(message string, signedMessage string) (string, error) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("[ERROR]: handle error: %s\n", x)
		}
	}()
	// Hash the unsigned message using EIP-191
	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
	hash := crypto.Keccak256Hash(hashedMessage)

	// Get the bytes of the signed message
	decodedMessage := hexutil.MustDecode(signedMessage)

	// Handles cases where EIP-115 is not implemented (most wallets don't implement it)
	if decodedMessage[64] == 27 || decodedMessage[64] == 28 {
		decodedMessage[64] -= 27
	}

	// Recover a public key from the signed message
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), decodedMessage)
	if sigPublicKeyECDSA == nil {
		err = errors.New("Could not get a public address from the message signature ")
	}
	if err != nil {
		return "", err
	}

	return crypto.PubkeyToAddress(*sigPublicKeyECDSA).String(), nil
}
