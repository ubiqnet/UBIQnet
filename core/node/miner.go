package node

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"os"
	"path"
	"strconv"
	"ubiqnet/config"
	"ubiqnet/constants"
	"ubiqnet/service"
	"ubiqnet/store/localstore"
	"ubiqnet/tools"

	"ubiqnet/task"

	version "ubiqnet"
	"ubiqnet/openapi"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	logging "github.com/ipfs/go-log"
	"github.com/robfig/cron/v3"
	"golang.org/x/crypto/sha3"
)

var (
	c *cron.Cron

	rpcService service.RpcService
)

type MinerNode struct {
	RpcAddress string

	NetWorkId     int
	WalletAddress string
	MinerStatus   bool
	TotalDiskFree uint64
	CpuCore       int
	Hw            int
	NodeType      int
}
type RpcService struct {
}

type MinerWalletData struct {
	Address   string `json:"address"`
	PublicKey string `json:"publicKey"`
	Cipher    string `json:"cipher"`
}

type BroadCastData struct {
	SenderNodeType int `json:"senderNodeType"`
	Msg            string
	TimeStamp      int64
}

var minerLog = logging.Logger("core:miner")
var client *ethclient.Client

// var rechargeAmount big.Int

func (miner *MinerNode) NewMiner(nodeType int) {
}
func (miner *MinerNode) IsVerifyNode() bool {

	return miner.NodeType == constants.NodeTypeVerify

}

func (miner *MinerNode) IsNormalNode() bool {

	return miner.NodeType == constants.NodeTypeNomal
}

func (miner *MinerNode) IsStableNode() bool {

	return miner.NodeType == constants.NodeTypeStable
}

func (miner *MinerNode) CreateMinerWallet(rootDir string) {
	err := logging.SetLogLevelRegex("core:miner", "info")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		minerLog.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	// privateKeyStr := hexutil.Encode(privateKeyBytes)[2:]
	privateKeyStr := hexutil.Encode(privateKeyBytes)
	minerLog.Info("privateKey: ", privateKeyStr[2:])
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		minerLog.Debug("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyStr := hexutil.Encode(publicKeyBytes)[4:]
	minerLog.Info("The node wallet was successfully generated, please record the information")
	minerLog.Info("publicKey: ", publicKeyStr)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	address := hexutil.Encode(hash.Sum(nil)[12:])
	minerLog.Info("wallet address: ", address)
	result := tools.PswEncrypt(privateKeyStr)
	p := path.Join(rootDir, "key.json")
	filePtr, err := os.Create(p)
	if err != nil {
		minerLog.Error(err)
		return
	}
	defer filePtr.Close()
	encoder := json.NewEncoder(filePtr)
	walletData := MinerWalletData{Address: address, PublicKey: publicKeyStr, Cipher: string(result)}
	err = encoder.Encode(walletData)
	if err != nil {
		minerLog.Error(err)
		return
	}

}

func (miner *MinerNode) GetECDSAPrivateKey(rootDir string) (*ecdsa.PrivateKey, error) {
	minerData := miner.GetMinerWallet(rootDir)
	privateStr := miner.GetMinerWalletPrivateKey(*minerData)
	privateBytes, err := hexutil.Decode(privateStr)
	if err != nil {
		minerLog.Error(err)
		return nil, err
	}
	r, err := crypto.ToECDSA(privateBytes)
	if err != nil {
		return nil, err
	}

	return r, nil

}

func (miner *MinerNode) GetWalletAddressFromPrivateKey(privateKey ecdsa.PrivateKey) string {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		minerLog.Error("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return ""
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	address := hexutil.Encode(hash.Sum(nil)[12:])

	return address
}

func (miner *MinerNode) GetMinerWallet(rootDir string) *MinerWalletData {

	p := path.Join(rootDir, "key.json")
	filePtr, err := os.Open(p)
	if err != nil {
		minerLog.Error(err)
		return nil
	}
	var minerWalletData MinerWalletData
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&minerWalletData)
	if err != nil {
		minerLog.Error(err)
		return nil
	}

	return &minerWalletData

}

func (miner *MinerNode) GetMinerWalletPrivateKey(data MinerWalletData) string {
	privateKey := tools.PswDecrypt(data.Cipher)

	return privateKey
}

func (miner *MinerNode) GetMinerWalletPublicKey(privateKey *ecdsa.PrivateKey) []byte {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		minerLog.Info("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	return publicKeyBytes
}

func (miner *MinerNode) Daemon(cfgRoot string, cfg config.MinerConfig) {

	err := logging.SetLogLevelRegex("core:miner", "info")
	if err != nil {
		panic(err)
	}

	localstore.InitDb(cfgRoot)
	localstore.Put(constants.NODE_DATA_DIR_KEY, cfgRoot)
	service.Node.DataDir = cfgRoot
	service.Node.RpcAddress = cfg.SwapEndPoint
	service.Node.DepositGas = cfg.DepositGas
	privateKey, err := miner.GetECDSAPrivateKey(cfgRoot)

	if err != nil {
		panic(err)
	}
	service.Node.PrivateKey = privateKey
	service.Node.PublicKey = miner.GetMinerWalletPublicKey(privateKey)
	service.Node.WalletAddress = miner.GetWalletAddressFromPrivateKey(*privateKey)
	service.Node.DebugApi = cfg.DebugApi
	deviceinfo := tools.GetHardwareData()

	if deviceinfo.MinimalRequire {
		minerLog.Error(constants.SYSTEM_NOT_SUITABLE)
		// os.Exit(0)
	}
	tools.DeviceInfo.HW = 1

	miner.deploy()
	s := openapi.HttpServer{}
	go s.InitHttpServer()
	go openapi.StartWebsocket()

}

func (miner *MinerNode) deploy() {
	minerLog.Info("The wallet address for generating chequebook: ", service.Node.WalletAddress)
	minerLog.Info("ubiqnet version: ", version.CurrentVersionNumber)
	minerLog.Info("starting with ubiqnet")
	client = rpcService.ConnectRPC()
	// defer client.Close()
	contractAddress := localstore.Get(constants.ContractAddressKey)
	contractHash := localstore.Get(constants.ContractHashKey)
	if len(contractAddress) > 0 {

		minerLog.Info("The chequebook address is: " + contractAddress)
		service.Node.NodeType, _ = strconv.Atoi(localstore.Get(constants.NodeTypeKey))
		ubqAmount := rpcService.BalanceOfUbq()
		if ubqAmount.Uint64() == 0 && localstore.Get(constants.PledgeStatusKey) == "" {
			miner.checkBlockTask()
			return
		}
		miner.nftCheck()

		return
	}

	if len(contractHash) > 0 {
		service.Node.NodeType, _ = strconv.Atoi(localstore.Get(constants.NodeTypeKey))
		minerLog.Info("The chequebook has been generated with a hash of: " + contractHash)
		miner.checkBlockTask()
		return
	}
	blockNum, _ := client.BlockNumber(context.Background())
	localstore.Put(constants.DeployBlockNumKey, strconv.FormatUint(blockNum, 10))
	hash, err := rpcService.Deploy()
	if err != nil {
		minerLog.Error("Deployment error: ", err)
		os.Exit(0)
	}

	minerLog.Info("successfully deposited to chequebook")
	localstore.Put(constants.ContractHashKey, hash)
	minerLog.Info("The chequebook is in generation with a hash of: " + hash)
	localstore.Put(constants.NodeTypeKey, fmt.Sprint(constants.NodeTypeNomal))
	service.Node.NodeType = constants.NodeTypeNomal
	miner.checkBlockTask()
}

func (miner *MinerNode) checkBlockTask() {

	c = cron.New(cron.WithSeconds())
	c.AddFunc("*/20 * * * * ?", miner.checkBlockStatus)
	c.Start()

}

func (miner *MinerNode) checkBlockStatus() {
	txHash := localstore.Get(constants.ContractHashKey)
	startBlock := new(big.Int)
	if localstore.Get(constants.DeployBlockNumKey) != "" {

		startBlock.SetString(localstore.Get(constants.DeployBlockNumKey), 10)

	}
	endBlock := big.NewInt(10000)
	if startBlock.Uint64() > 0 {
		endBlock = endBlock.Add(startBlock, endBlock)
	} else {
		endBlock = nil
	}

	query := ethereum.FilterQuery{FromBlock: startBlock, ToBlock: endBlock, Addresses: []common.Address{common.HexToAddress(constants.FactoryAddress)}}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		minerLog.Error(err.Error())

		return
	}
	flag := false
	// address := ""
	for _, log := range logs {
		if txHash == log.TxHash.String() {

			addr := common.Bytes2Hex(log.Data)
			if addr == "" {

				localstore.Remove(constants.ContractAddressKey)
				minerLog.Info("Transaction failed")
				os.Exit(0)
				return
			}
			addr = "0x" + string(addr[len(addr)-40:])

			localstore.Put(constants.ContractAddressKey, addr)
			go tools.SendRegion()
			minerLog.Info("The chequebook address is: " + addr)
			flag = true
			// address = addr
			c.Stop()
			break
		}

	}

	if flag {
		// status := true
		// var amountStr string
		// for status {
		// 	fmt.Println("Whether to pledge, pledge UBQ to get more income (enter 0 to the next step, enter the number to pledge)")
		// 	fmt.Scan(&amountStr)
		// 	amount := new(big.Int)
		// 	amount.SetString(amountStr, 10)
		// 	IsDigit := false
		// 	for _, r := range amountStr {
		// 		IsDigit = unicode.IsDigit(r)
		// 		if !IsDigit {
		// 			break
		// 		}

		// 	}
		// amount = amount.Mul(big.NewInt(constants.PledgeSymbol), amount)
		// if amount.Uint64() > 0 {
		// 	miner.recharge(amount, address)
		// 	localstore.Put(constants.PledgeStatusKey, "1")
		// 	status = false
		// 	break
		// } else if IsDigit && amount.Uint64() == 0 {
		// 	localstore.Put(constants.PledgeStatusKey, "1")
		// 	status = false
		// 	miner.nftCheck()
		// 	break
		// } else {
		// 	localstore.Put(constants.PledgeStatusKey, "")
		// 	fmt.Println("input error")
		// }
		localstore.Put(constants.PledgeStatusKey, "1")
		miner.nftCheck()

		// }

	}

}

func (miner *MinerNode) mintNfts() {
	gpuTotal := len(tools.DeviceInfo.GpuInfo)
	var gpuJoin bool
	var gpuJoinStr string
	flag := true
	if gpuTotal > 0 {
		for flag {
			fmt.Println("Your GPU meets the rental standards, please input 'y' to join the rental list, while 'n' to cancel.")
			fmt.Scan(&gpuJoinStr)
			if gpuJoinStr == "y" || gpuJoinStr == "Y" {
				gpuJoin = true
				break
			} else if gpuJoinStr == "n" || gpuJoinStr == "N" {
				gpuJoin = false
				break
			}

		}

	}

	var cpu int
	for flag {
		fmt.Printf("The cores of your CPU is %d, please input the amount to rent out (e.g.'4', minimum '2')\n", tools.DeviceInfo.NumOfCores)

		fmt.Scan(&cpu)
		if cpu < 2 || cpu > tools.DeviceInfo.NumOfCores {
			fmt.Println("Invalid input ")

		} else {
			break
		}
	}

	memTotal := int(math.Ceil(float64(tools.DeviceInfo.FreeMemory/1024/1024) / 1024))
	var mem int
	for flag {

		fmt.Printf("Your free memory available is %d G, please input the amount to rent out(e.g.'8', minimum '4')\n", memTotal)
		fmt.Scan(&mem)
		if mem < 2 || mem > memTotal {
			fmt.Println("Invalid input ")
		} else {
			break
		}
	}
	cpu = cpu / 2
	mem = mem / 4
	if cpu > mem {
		cpu = mem
	} else {
		mem = cpu
	}

	storageTotal := int(math.Ceil(float64(tools.DeviceInfo.FreeDisk/1024/1024) / 1024))
	var storage int
	for flag {
		fmt.Printf("Your free storage available is %d G,  please input the amount to rent out (e.g.'40', minimum '10')\n", storageTotal)

		fmt.Scan(&storage)
		if storage < 10 || storage > storageTotal {
			fmt.Println("Invalid input ")
		} else {
			break
		}
	}
	var gpu int

	if gpuTotal > 0 && gpuJoin {

		for flag {
			fmt.Printf("The cores of your GPU is %d, please input the amount to rent out (e.g.'2', minimum '1')\n", gpuTotal)

			fmt.Scan(&gpu)
			if gpu < 0 || gpu > gpuTotal {
				fmt.Println("Invalid input ")

			} else {
				break
			}
		}
	}
	for i := 0; i < cpu; i++ {
		hash, err := service.Nft.MintNfts(1, fmt.Sprintf("CPU %d cores", 2))
		if err != nil {
			minerLog.Error(err.Error())

		}
		minerLog.Info("Cpu: ", hash)
	}

	localstore.Put(constants.NftMintKey, "1")

	for i := 0; i < mem; i++ {
		hash, err := service.Nft.MintNfts(2, fmt.Sprintf("MEM %d G", 4))
		if err != nil {
			minerLog.Error(err.Error())

		}
		minerLog.Info("Memory: ", hash)
	}

	for i := 0; i < storage/10; i++ {
		hash, err := service.Nft.MintNfts(3, fmt.Sprintf("Storage %d G", 10))
		if err != nil {
			minerLog.Error(err.Error())

		}
		minerLog.Info("Storage: ", hash)
	}

	hash, err := service.Nft.MintNfts(5, fmt.Sprintf("Ip %s 1", tools.DeviceInfo.Netinfo.IP))
	if err != nil {
		minerLog.Error(err.Error())

	}
	minerLog.Info("Ip: ", hash)

	hash, err = service.Nft.MintNfts(6, fmt.Sprintf("BrandWidth %f 1", tools.DeviceInfo.DownloadSpeed))
	if err != nil {
		minerLog.Error(err.Error())

	}
	minerLog.Info("BrandWidth: ", hash)

	if gpuJoin && gpu > 0 {
		for i := 0; i < gpu; i++ {
			hash, err = service.Nft.MintNfts(4, fmt.Sprintf("Gpu %s 1 ", tools.DeviceInfo.GpuInfo[i].TotalMemory))
			if err != nil {
				minerLog.Error(err.Error())

			}
			minerLog.Info("Gpu: ", hash)
		}
	}

}

func (miner *MinerNode) recharge(amount *big.Int, addr string) {

	walletAmount := rpcService.BalanceOfUbq()

	if walletAmount.Cmp(amount) < 0 {
		minerLog.Error("Insufficient UBQ staking, and minimum of ", amount.Div(amount, big.NewInt(constants.PledgeSymbol)), "UBQ is required for the normal operation of the node.")

		os.Exit(0)
	}

	t, err := rpcService.Recharge(amount, addr)
	if err != nil {
		minerLog.Error(err.Error())
		return
	}
	minerLog.Info("sent deposit transaction " + t.Hash().String())
	bind.WaitMined(context.Background(), client, t)
	minerLog.Info("successfully deposited to chequebook")
	miner.nftCheck()
}
func (miner *MinerNode) nftCheck() {
	nftMintStr := localstore.Get(constants.NftMintKey)
	if nftMintStr == "" {
		miner.mintNfts()
	}

	task.InitTask()
}
