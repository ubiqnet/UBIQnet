package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
	"ubiqnet/constants"
	"ubiqnet/store/localstore"

	nvidiasmi "github.com/c3sr/nvidia-smi"
	logging "github.com/ipfs/go-log"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

var deviceLog = logging.Logger("tools:deviceInfo")

type deviceHW struct {
	LogicCores    int
	TotalMemory   uint64
	FreeDisk      uint64
	DownloadSpeed float64
	UploadSpeed   float64
}
type deviceInfo struct {
	Status          int
	OS              string
	TotalMemory     uint64
	FreeMemory      uint64
	TotalDisk       uint64
	UsedDisk        uint64
	FreeDisk        uint64
	GpuInfo         []localstore.GpuInfo
	LogicCores      int
	NumOfCores      int
	ModelName       string
	CpuSpeed        float64
	HostName        string
	UpTime          uint64
	Platform        string
	HostId          string
	DownloadSpeed   float64
	UploadSpeed     float64
	Valid           bool
	ReadCount       uint64
	WriteCount      uint64
	Timestamp       int64
	MinimalRequire  bool
	Arch            string
	HW              int
	Netinfo         Netinfo
	containerUsages []ContainerUsage
	LowHW           deviceHW
	MediumHW        deviceHW
	HighHW          deviceHW
}

var (
	DeviceInfo deviceInfo

	LowHW = deviceHW{
		LogicCores:    2,
		TotalMemory:   uint64(4294967296),
		FreeDisk:      uint64(10737418240),
		DownloadSpeed: float64(50),
		UploadSpeed:   float64(5),
	}

	MediumHW = deviceHW{
		LogicCores:    2,
		TotalMemory:   uint64(4294967296),
		FreeDisk:      uint64(10737418240),
		DownloadSpeed: float64(50),
		UploadSpeed:   float64(5),
	}
	HighHW = deviceHW{
		LogicCores:    8,
		TotalMemory:   uint64(4294967296 * 4),
		FreeDisk:      uint64(10737418240 * 5),
		DownloadSpeed: float64(50 * 2),
		UploadSpeed:   float64(5 * 2),
	}
)

func DealWithErr(err error) {
	if err != nil {
		fmt.Println(err.Error())

	}
}

func SendRegion() {
	country := DeviceInfo.Netinfo.Country
	if country == "Unknown" {
		return
	}
	region := DeviceInfo.Netinfo.RegionName
	city := DeviceInfo.Netinfo.City
	ip := DeviceInfo.Netinfo.IP
	nodeAddress := localstore.Get(constants.ContractAddressKey)
	if nodeAddress == "" {
		return
	}
	client := http.Client{}
	data := "{\"country\":\"" + country + "\",\"region\":\"" + region + "\",\"city\":\"" + city + "\",\"ip\":\"" + ip + "\",\"nodeAddress\":\"" + nodeAddress + "\"}"
	_, err := client.Post("https://amino.world/ubiqnetServer", "application/json", strings.NewReader(data))
	if err != nil {
		println(err)
		return
	}
}

func GetHardwareData() deviceInfo {
	runtimeOS := runtime.GOOS
	DeviceInfo.LowHW = LowHW
	DeviceInfo.MediumHW = MediumHW
	DeviceInfo.HighHW = HighHW
	vmStat, err := mem.VirtualMemory()
	DealWithErr(err)
	b, err2 := nvidiasmi.New()

	if err2 != nil {
		fmt.Printf("Error getting GPU info: %v", err)
	} else {
		if b.HasGPU() {
			DeviceInfo.GpuInfo = make([]localstore.GpuInfo, 0)
			for _, gpu := range b.GPUS {
				var gpuInfo = localstore.GpuInfo{}
				gpuInfo.ProductName = gpu.ProductName
				gpuInfo.TotalMemory = gpu.TotalFbMemoryUsageGpu
				DeviceInfo.GpuInfo = append(DeviceInfo.GpuInfo, gpuInfo)
			}
		} else {
			println("no GPU")
		}
	}

	var totalDisk uint64 = 0
	var totalDiskFree uint64 = 0
	var totalDiskUsed uint64 = 0
	path := localstore.Get(constants.NODE_DATA_DIR_KEY)
	diskStat, err := disk.Usage(path)
	if diskStat != nil {
		totalDiskUsed = totalDiskUsed + diskStat.Used
		totalDisk = totalDisk + diskStat.Total
		totalDiskFree = totalDiskFree + diskStat.Free
		ioStat, _ := disk.IOCounters(path)
		ioStatCount, ok := ioStat[path]
		if ok {
			DeviceInfo.ReadCount = ioStatCount.ReadBytes * uint64(1000) / uint64(1024) / uint64(1024) / ioStatCount.ReadTime
			DeviceInfo.WriteCount = ioStatCount.WriteBytes * uint64(1000) / uint64(1024) / uint64(1024) / ioStatCount.WriteTime
		}

	}
	DealWithErr(err)
	cpuStat, err := cpu.Info()
	logicCore, err := cpu.Counts(true)
	DealWithErr(err)
	physicsCore, err := cpu.Counts(false)
	DealWithErr(err)
	hostStat, err := host.Info()
	DealWithErr(err)
	DeviceInfo.OS = runtimeOS
	DeviceInfo.Timestamp = time.Now().UTC().Unix()
	DeviceInfo.TotalMemory = vmStat.Total
	DeviceInfo.FreeMemory = vmStat.Free
	DeviceInfo.TotalDisk = totalDisk
	DeviceInfo.UsedDisk = totalDiskUsed
	DeviceInfo.FreeDisk = totalDiskFree
	DeviceInfo.LogicCores = logicCore
	DeviceInfo.NumOfCores = physicsCore
	DeviceInfo.ModelName = cpuStat[0].ModelName
	DeviceInfo.CpuSpeed = cpuStat[0].Mhz
	DeviceInfo.HostName = hostStat.Hostname
	DeviceInfo.HostId = hostStat.HostID
	DeviceInfo.UpTime = hostStat.Uptime
	DeviceInfo.Platform = hostStat.Platform
	DeviceInfo.Arch = hostStat.KernelArch
	var wg sync.WaitGroup

	if DeviceInfo.DownloadSpeed <= 0 || time.Now().Minute() == 0 {
		DeviceInfo.containerUsages, err = statsContainer()
		wg.Add(1)
		go func() {
			DeviceInfo.Netinfo = checkAccess()
			SendRegion()
			wg.Done()
		}()
		DealWithErr(err)
		startTestSpeed()

	}
	hw := getHW(DeviceInfo.LogicCores, DeviceInfo.TotalMemory, DeviceInfo.FreeDisk, DeviceInfo.DownloadSpeed, DeviceInfo.UploadSpeed)
	DeviceInfo.HW = hw
	if hw > 0 {
		DeviceInfo.Valid = true
		DeviceInfo.MinimalRequire = true
	} else {
		DeviceInfo.Valid = false
		DeviceInfo.MinimalRequire = false
		deviceLog.Error(constants.SYSTEM_NOT_SUITABLE)
		deviceLog.Info("Minimal requirements are:")
		lowS, err := json.MarshalIndent(LowHW, "", "\t")
		DealWithErr(err)
		deviceLog.Info(string(lowS))
	}
	dataString, err := json.MarshalIndent(DeviceInfo, "", "\t")
	DealWithErr(err)
	wg.Wait()
	deviceLog.Info("finished check and check data:" + string(dataString))
	DeviceInfo.Status = 1
	return DeviceInfo
}

func startTestSpeed() {
	err := retry(4, 1*time.Second, func(i int) error {
		return TestSpeed(i)
	})
	DealWithErr(err)
}

func TestSpeed(i int) error {
	defer func() {
		err := recover()
		if err != nil {
			print(err)
			fmt.Println("request timeout! ")
		}
	}()
	down, up := float64(0), float64(0)
	workDone := make(chan struct{}, 1)
	go func() {

		timeout := make(chan bool)
		go func() {
			time.Sleep(60 * time.Second)
			timeout <- true
		}()
		down, up = STestSpeed(i)
		workDone <- struct{}{}
	}()
	select {
	case <-workDone:
		if down <= 0 || up <= 0 {
			return errors.New("timeout speed 0! ")
		}
		DeviceInfo.DownloadSpeed = down
		DeviceInfo.UploadSpeed = up
		return nil
	case <-time.After(120 * time.Second):
		return fmt.Errorf("request timeout! ")
	}

}
func retry(attempts int, sleep time.Duration, f func(index int) error) (err error) {
	for i := 0; i < attempts; i++ {
		if i > 0 {
			deviceLog.Info("retrying after error:")
			DealWithErr(err)
			time.Sleep(sleep)
		}
		err = f(i)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func getHW(cpu int, memory uint64, disk uint64, downloadSpeed float64, uploadSpeed float64) int {
	a := suitLevel(cpu, memory, disk, downloadSpeed, uploadSpeed, HighHW)
	if a {
		return 3
	}
	b := suitLevel(cpu, memory, disk, downloadSpeed, uploadSpeed, MediumHW)
	if b {
		return 2
	}
	c := suitLevel(cpu, memory, disk, downloadSpeed, uploadSpeed, LowHW)
	if c {
		return 1
	} else {
		return 0
	}
}

func suitLevel(cpu int, memory uint64, disk uint64, downloadSpeed float64, uploadSpeed float64, hw deviceHW) bool {
	return cpu >= hw.LogicCores && memory >= hw.TotalMemory && (disk+DirSize(localstore.Get(constants.NODE_DATA_DIR_KEY))) >= hw.FreeDisk && downloadSpeed >= hw.DownloadSpeed && uploadSpeed >= hw.UploadSpeed
}
func GetHwDiskFree(hw int) uint64 {
	if hw == 2 {
		return MediumHW.FreeDisk
	} else {
		//	return HighHW.FreeDisk
		//} else {
		return LowHW.FreeDisk
	}

}
