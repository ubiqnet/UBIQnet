package service

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
	"ubiqnet/constants"
	"ubiqnet/store/localstore"
	"ubiqnet/tools"
)

var OrderSuitComplete = "ORDER_SUIT_COMPLETE"
var NodeSuitedQueue = tools.NewQueue()
var OrderSuiteComplete = tools.NewQueue()

func ReceiveNodeTemplateSuitQueueConsume() {
	NodeSuitedQueue.Run(func(item tools.QueueItem) {
		nodeSuiteInfo, ok := (item).(NodeSuiteInfo)
		if ok {
			receiveNormalNodeSuit(nodeSuiteInfo)
		}
	})
}

func ReceiveOrderSuitedCompleteQueueConsume() {
	OrderSuiteComplete.Run(func(item tools.QueueItem) {
		validatorSuitInfo, ok := (item).(ValidatorSuitInfo)
		if ok {
			receiveValidatorOrderComplete(validatorSuitInfo)
		}
	})
}

func receiveNormalNodeSuit(info NodeSuiteInfo) {
	lastUpdateTimestamp := localstore.Get(constants.LastUploadTimeKey)
	if lastUpdateTimestamp != "" {
		timestamp, err := strconv.ParseInt(lastUpdateTimestamp, 10, 64)
		if err != nil {
			return
		}
		if (timestamp + 2*3600) < time.Now().UTC().Unix() {
			return
		}
	}
	validatorSuitInfoStr := localstore.MemGet(constants.MEM_ORDER_CHECK_PREFIX + strconv.Itoa(info.OrderId))

	if validatorSuitInfoStr != "" {

		validatorSuitInfo := ToValidatorSuitInfo(validatorSuitInfoStr)
		if validatorSuitInfo.FinishedTimestampNano == int64(0) {
			finished := true
			var nodesSuiteInfo = validatorSuitInfo.NodeSuitInfos[info.Index]
			var index = info.Index
			if nodesSuiteInfo.NodeAddress == "" {

				validatorSuitInfo.NodeSuitInfos[index] = info
				print(validatorSuitInfo.ToString())
			} else {
				if nodesSuiteInfo.SuitUtcTimestampNano == info.SuitUtcTimestampNano {
					if nodesSuiteInfo.NodeAddress > info.NodeAddress {

						validatorSuitInfo.NodeSuitInfos[index] = info
					}
				} else if nodesSuiteInfo.SuitUtcTimestampNano > info.SuitUtcTimestampNano {

					validatorSuitInfo.NodeSuitInfos[index] = info
				}
			}
			for _, i := range validatorSuitInfo.NodeSuitInfos {
				if i.NodeAddress == "" {

					finished = false
				}
			}
			if finished {
				validatorSuitInfo.FinishedTimestampNano = time.Now().UTC().UnixNano()
				p2pData := P2pData{Sender: peerId, Event: OrderSuitComplete, Extra: validatorSuitInfo.ToString(), TimeStamp: time.Now().UnixMilli()}
				p2pBytesData, _ := json.Marshal(&p2pData)
				err := P2p.PublishBroadCastData(p2pBytesData)
				if err != nil {
					taskLog.Error(err)
				}
				timer := time.NewTimer(1 * time.Minute)
				select {
				case <-timer.C:
					if localstore.MemContains(constants.MEM_ORDER_CHECK_PREFIX + strconv.Itoa(info.OrderId)) {

						imageInfos, err := rpcService.GetAppInfo(info.OrderId, common.HexToAddress(validatorSuitInfo.AppAddress))
						if err != nil {
							taskLog.Error("fail to start")
						}
						var allTokenLen []int
						for _, suiteInfo := range validatorSuitInfo.NodeSuitInfos {
							allTokenLen = append(allTokenLen, len(suiteInfo.IpTokenIds))
							allTokenLen = append(allTokenLen, len(suiteInfo.CpuTokenIds))
							allTokenLen = append(allTokenLen, len(suiteInfo.MemTokenIds))
							allTokenLen = append(allTokenLen, len(suiteInfo.StorageTokenIds))
							allTokenLen = append(allTokenLen, len(suiteInfo.GpuTokenIds))
							allTokenLen = append(allTokenLen, len(suiteInfo.BandwidthTokenIds))
						}
						//sort.Ints(allTokenLen)
						//max := allTokenLen[len(allTokenLen)-1]
						a := len(imageInfos)
						nodeAddress := make([]common.Address, a)
						tokenIds := make([][][]*big.Int, a)
						for i := 0; i < a; i++ {
							tokenIds[i] = make([][]*big.Int, 6)
							//for j := 0; j < 6; j++ {
							//	tokenIds[i][j] = make([]*big.Int, max)
							//
							//}
						}

						for i := 0; i < len(imageInfos); i++ {

							suiteInfo := validatorSuitInfo.NodeSuitInfos[i]
							nodeAddress[i] = common.HexToAddress(suiteInfo.NodeAddress)
							setTokenIds(suiteInfo.IpTokenIds, tokenIds, i, 0)
							setTokenIds(suiteInfo.BandwidthTokenIds, tokenIds, i, 1)
							setTokenIds(suiteInfo.GpuTokenIds, tokenIds, i, 2)
							setTokenIds(suiteInfo.StorageTokenIds, tokenIds, i, 3)
							setTokenIds(suiteInfo.CpuTokenIds, tokenIds, i, 4)
							setTokenIds(suiteInfo.MemTokenIds, tokenIds, i, 5)
						}

						trade, err := rpcService.SetTokenTrade(common.HexToAddress(validatorSuitInfo.AppAddress), nodeAddress, tokenIds, int64(validatorSuitInfo.OrderId))
						if err != nil {
							taskLog.Error(err)
							return
						}
						go func() {
							client := http.Client{}
							data := "{\"orderId\":\"" + strconv.Itoa(info.OrderId) + "\",\"hash\":\"" + trade + "\"}"
							err := tools.Retry(10, 10*time.Second, func() error {
								_, err := client.Post("https://amino.world/ubiqnetServer/orderHash", "application/json", strings.NewReader(data))
								if err != nil {
									return err
								}
								return nil
							})
							if err != nil {
								return
							}

						}()
						go checkOrder(validatorSuitInfo.OrderId, validatorSuitInfo.AppOwnerAddress)
						taskLog.Info("rent NFT result:" + trade)
						timer.Stop()
					}

				}
			}
			if localstore.MemContains(constants.MEM_ORDER_CHECK_PREFIX + strconv.Itoa(info.OrderId)) {
				localstore.MemPut(constants.MEM_ORDER_CHECK_PREFIX+strconv.Itoa(info.OrderId), validatorSuitInfo.ToString())
			}
		}
	}
}

func setTokenIds(oneTokenIds []int64, tokenIds [][][]*big.Int, i int, j int) {
	if len(oneTokenIds) == 0 {
		tokenIds[i][j] = make([]*big.Int, 0)
	} else {
		tokenIds[i][j] = make([]*big.Int, len(oneTokenIds))
		for i2, id := range oneTokenIds {
			if id != 0 {
				if big.NewInt(id) != nil {
					tokenIds[i][j][i2] = big.NewInt(id)
				}
			}

		}
	}

}

func receiveValidatorOrderComplete(info ValidatorSuitInfo) {
	if strings.EqualFold(info.ValidateAddress, localstore.Get(constants.ContractAddressKey)) {
		return
	}
	validatorSuitInfoStr := localstore.MemGet(constants.MEM_ORDER_CHECK_PREFIX + strconv.Itoa(info.OrderId))
	if validatorSuitInfoStr != "" {
		validatorSuitInfo := ToValidatorSuitInfo(validatorSuitInfoStr)
		currentTimestampNano := validatorSuitInfo.FinishedTimestampNano
		if currentTimestampNano == info.FinishedTimestampNano {
			if localstore.Get(constants.ContractAddressKey) > info.ValidateAddress {

				localstore.MemRemove(constants.MEM_ORDER_CHECK_PREFIX + strconv.Itoa(info.OrderId))
			}
		} else if currentTimestampNano > info.FinishedTimestampNano {

			localstore.MemRemove(constants.MEM_ORDER_CHECK_PREFIX + strconv.Itoa(info.OrderId))
		}
	}
}
