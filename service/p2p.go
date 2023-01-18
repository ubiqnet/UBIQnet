package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"time"
	"ubiqnet/constants"
	"ubiqnet/store/localstore"

	logging "github.com/ipfs/go-log"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
)

type P2pService struct {
}
type P2pData struct {
	Sender    string `json:"sender"`
	Event     string `json:"event"`
	Extra     string `json:"extra"`
	TimeStamp int64  `json:"timestamp"`
}

type ChildNode struct {
	ContractAddress string `json:"contractAddress"`
	PeerId          string `peerId`
}

type UploadData struct {
	ContractAddress string `json:"contractAddress"`
	TimeStamp       int64  `json:"timestamp"`

	CpuIds        []*big.Int `json:"cpuIds"`
	IpIds         []*big.Int `json:"ipIds"`
	BrandWidthIds []*big.Int `json:"brandWidthIds"`
	GpuIds        []*big.Int `json:"gpuIds"`
	MemIds        []*big.Int `json:"memIds"`
	StorageIds    []*big.Int `json:"storageIds"`

	CpuPrices        []*big.Int `json:"cpuPrices"`
	IpPrices         []*big.Int `json:"ipPrices"`
	BrandWidthPrices []*big.Int `json:"brandWidthPrices"`
	GpuPrices        []*big.Int `json:"gpuPrices"`
	MemPrices        []*big.Int `json:"memPrices"`
	StoragePrices    []*big.Int `json:"storagePrices"`
	Sign             [32]byte   `json:"sign"`
}

var P2p P2pService
var API coreiface.CoreAPI

var p2pLog = logging.Logger("service:p2p")
var peerId string

func (p2p *P2pService) New(id string, api coreiface.CoreAPI) {
	logging.SetLogLevelRegex("service:p2p", "info")
	API = api
	peerId = id
}

func (p2p *P2pService) SubscribeBroardCastData() {

	sub, err := API.PubSub().Subscribe(context.Background(), "broardcast")

	if err != nil {
		p2pLog.Error(err)
		return
	}
	defer sub.Close()
	for {

		msg, err := sub.Next(context.Background())
		if err == io.EOF || err == context.Canceled {
			p2pLog.Error(err)
		}
		p2p.handlerP2pMsg(msg.Data())
	}

}

func (p2p *P2pService) SubscribeP2pData() {

	p2pSub, err := API.PubSub().Subscribe(context.Background(), peerId)
	if err != nil {
		p2pLog.Error(err)
		return
	}

	defer p2pSub.Close()
	for {

		p2pMsg, err := p2pSub.Next(context.Background())
		if err == io.EOF || err == context.Canceled {
			p2pLog.Error(err)
		}
		p2p.handlerP2pMsg(p2pMsg.Data())

	}
}

func (p2p *P2pService) handlerP2pMsg(p2pMsg []byte) {
	var p2pData P2pData
	err := json.Unmarshal(p2pMsg, &p2pData)
	if err == nil {

		if (p2pData.Sender == peerId) && (p2pData.Event == "register" || p2pData.Event == "registerResult") {
			return
		}
		p2pLog.Info("receive p2p data: ", p2pData)
		switch p2pData.Event {
		case "register":
			if Node.NodeType == constants.NodeTypeVerify {
				p2p.registerReqEvent(p2pData.Sender)
			}

		case "registerResult":
			p2p.registerResEvent(p2pData.Sender, p2pData.Extra)

		case "uploadInfo":
			p2p.uploadInfoReqEvent(p2pData.Sender, p2pData.Extra)

		case "uploadInfoResult":
			p2p.uploadInfoResEvent(p2pData.Sender, p2pData.Extra)

		case NewOrderTemplateTopic:
			p2pLog.Info(NewOrderTemplateTopic)
			if p2pData.Extra != "" {
				TemplateInfoQueue.Enqueue(ToImageInfo(p2pData.Extra))
			}

		case TemplateSuited:
			p2pLog.Info(TemplateSuited)
			if p2pData.Extra != "" {
				nodeSuiteInfo := ToNodeSuitInfo(p2pData.Extra)
				OtherTemplateQueue.Enqueue(nodeSuiteInfo)
				NodeSuitedQueue.Enqueue(nodeSuiteInfo)
			}
			break
		case OrderSuitComplete:
			p2pLog.Info(OrderSuitComplete)
			if p2pData.Extra != "" {
				OrderSuiteComplete.Enqueue(ToValidatorSuitInfo(p2pData.Extra))
			}
			break
		case StartContainer:
			p2pLog.Info(StartContainer)
			if p2pData.Extra != "" {
				StartImageQueue.Enqueue(ToImageInfo(p2pData.Extra))
			}
			break
		case StopContainer:
			p2pLog.Info(StopContainer)
			if p2pData.Extra != "" {
				StopImageQueue.Enqueue(ToImageInfo(p2pData.Extra))
			}
			break
		}

	}
}

func (p2p *P2pService) registerReqEvent(senderId string) {
	p2pData := P2pData{Event: "registerResult", Sender: peerId, TimeStamp: time.Now().UnixMilli(), Extra: peerId}

	data, _ := json.Marshal(&p2pData)
	p2p.PublishP2pData(senderId, data)
}

func (p2p *P2pService) registerResEvent(senderId string, verifyPeerId string) {
	// jsonStr := localstore.Get(constants.ChildNodeKey)
	// var childNodes []ChildNode
	// err := json.Unmarshal([]byte(jsonStr), &childNodes)
	// if err == nil {
	// 	child := ChildNode{PeerId: senderId, ContractAddress: contractAddress}
	// 	childNodes = append(childNodes, child)
	// }

	// r, _ := json.Marshal(&childNodes)

	// localstore.Put(constants.ChildNodeKey, string(r))
	localstore.Put(constants.VerifyPeerIdKey, verifyPeerId)

}

func (p2p *P2pService) uploadInfoReqEvent(senderId string, info string) {
	idStr := localstore.Get(constants.UploadIdKey)
	if idStr == "" {
		idStr = "0"
	}
	id, _ := strconv.Atoi(idStr)
	id += 1
	localstore.Put(constants.UploadIdKey, fmt.Sprint(id))

	localstore.Put(fmt.Sprintf("%s%d", constants.ChildNodeKey, id), info)

	var uploadData UploadData
	json.Unmarshal([]byte(info), &uploadData)
	uploadId := fmt.Sprintf("%s-%s-%d", constants.HeartKeyPrefix, uploadData.ContractAddress, uploadData.TimeStamp)
	localstore.Put(uploadId, info)

	// p2pData := P2pData{Event: "uploadInfoResult", Sender: peerId, TimeStamp: time.Now().UnixMilli(), Extra: uploadId}

	// data, _ := json.Marshal(&p2pData)

	// p2p.PublishP2pData(senderId, data)

}

func (p2p *P2pService) uploadInfoResEvent(senderId string, id string) {

	localstore.Remove(id)

}

func (p2p *P2pService) GetPeers() map[string]interface{} {
	var peersMap = make(map[string]interface{})
	peers, err := API.Swarm().Peers(context.Background())
	if err != nil {
		p2pLog.Error(err)
		return peersMap
	}

	for i := 0; i < len(peers); i++ {
		peer := peers[i].ID().String()
		if _, ok := peersMap[peer]; !ok {
			if peer != peerId {
				peersMap[peer] = peers[i]
			}

		}

	}
	return peersMap

}

func (p2p *P2pService) PublishP2pData(peerId string, data []byte) error {

	err := API.PubSub().Publish(context.Background(), peerId, data)

	return err

}

func (p2p *P2pService) PublishBroadCastData(data []byte) error {
	err := API.PubSub().Publish(context.Background(), "broardcast", data)

	return err
}

func (p2p *P2pService) BindVerifyNode(verifyNodeId string) {
	p2pData := P2pData{Sender: peerId, Event: "register", Extra: "", TimeStamp: time.Now().UnixMilli()}
	data, err := json.Marshal(&p2pData)
	if err != nil {

		return
	}
	p2p.PublishP2pData(verifyNodeId, data)
}
