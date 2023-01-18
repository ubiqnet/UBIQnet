package localstore

import (
	"encoding/json"
)

type HardwareInfo struct {
	Hash                string    `json:"Hash"`
	WalletAddress       string    `json:"WalletAddress"`
	IpTokens            []int64   `json:"IpTokens"`
	CpuTokens           []int64   `json:"CpuTokens"`
	GpuTokens           []int64   `json:"GpuTokens"`
	BandwidthTokens     []int64   `json:"bandwidthTokens"`
	MemTokens           []int64   `json:"MemTokens"`
	DiskTokens          []int64   `json:"DiskTokens"`
	IpTokensFree        []int64   `json:"IpTokensFree"`
	CpuTokensFree       []int64   `json:"CpuTokensFree"`
	GpuTokensFree       []int64   `json:"GpuTokensFree"`
	MemTokensFree       []int64   `json:"MemTokensFree"`
	DiskTokensFree      []int64   `json:"DiskTokensFree"`
	BandwidthTokensFree []int64   `json:"bandwidthTokensFree"`
	IP                  string    `json:"IP"`
	CountryCode         string    `json:"CountryCode"`
	Region              string    `json:"Region"`
	City                string    `json:"City"`
	IsPublic            bool      `json:"IsPublic"`
	OS                  string    `json:"OS"`
	TotalMemory         uint64    `json:"TotalMemory"`
	TotalDisk           uint64    `json:"TotalDisk"`
	GpuInfo             []GpuInfo `json:"GpuInfo"`
	CpuSpeed            float64   `json:"CpuSpeed"`
	TotalCpuCore        int       `json:"TotalCpuCore"`
}
type GpuInfo struct {
	ProductName string `json:"productName"`
	TotalMemory string `json:"totalMemory"`
}

func (info HardwareInfo) ToString() string {
	res, err := json.Marshal(info)
	if err != nil {
		return err.Error()
	}

	return string(res)
}

func ToInfo(res string) HardwareInfo {
	var p2 HardwareInfo
	err := json.Unmarshal([]byte(res), &p2)
	if err != nil {
		return HardwareInfo{}
	}
	return p2
}
