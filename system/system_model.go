package system

type CpuInfo struct {
	CpuCount     int       `json:"cpuCount"`
	UsePercent   []float64 `json:"usePercent"`
	TotalPercent string    `json:"totalPercent"`
	Mhz          string    `json:"mhz"`
}

type MemoryInfo struct {
	FreeMemory  string `json:"freeMemory"`
	TotalMemory string `json:"totalMemory"`
	UsedPercent string `json:"usedPercent"`
}

type NetIoInfo struct {
	SentSpc string `json:"sentSpc"`
	RecvSpc string `json:"recvSpc"`
}