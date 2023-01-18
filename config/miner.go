package config

type MinerConfig struct {
	SwapEndPoint string `json:"swap_end_point"`
	NetWorkId    int    `json:"network_id"`
	DebugApi     string `json:"debug_api"`
	DepositGas   uint64 `json:"swap_initial_deposit"`
	HarvestGas   int64 `json:"harvest_gas"`
}
