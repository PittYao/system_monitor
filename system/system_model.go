package system

import "fmt"

type CpuInfo struct {
	CpuCount     int      `json:"cpuCount"`
	UsePercent   []string `json:"usePercent"`
	TotalPercent string   `json:"totalPercent"`
	Mhz          string   `json:"mhz"`
}

type MemoryInfo struct {
	FreeMemory  string `json:"freeMemory"`
	TotalMemory string `json:"totalMemory"`
	UsedPercent string `json:"usedPercent"`
}

type NetIoInfo struct {
	SentSpc           string            `json:"sentSpc"`
	RecvSpc           string            `json:"recvSpc"`
	NetIoInfoSentInfo NetIoInfoSentInfo `json:"netIoInfoSentInfo"`
	NetIoInfoRecvInfo NetIoInfoRecvInfo `json:"netIoInfoRecvInfo"`
}

type NetIoInfoSentInfo struct {
	B  string `json:"B"`
	KB string `json:"KB"`
	MB string `json:"MB"`
}

type DiskInfo struct {
	Device      string `json:"device"`
	UsedPercent string `json:"usedPercent"`
	Total       string `json:"total"`
	Used        string `json:"used"`
	Free        string `json:"free"`
}

func (sent *NetIoInfoSentInfo) Formatter(size []string) {
	if len(size) >= 3 {
		sent.B = fmt.Sprintf("%s/S", size[0])
		sent.KB = fmt.Sprintf("%s/S", size[1])
		sent.MB = fmt.Sprintf("%s/S", size[2])
	}
}

type NetIoInfoRecvInfo struct {
	B  string `json:"B"`
	KB string `json:"KB"`
	MB string `json:"MB"`
}

func (recv *NetIoInfoRecvInfo) Formatter(size []string) {
	if len(size) >= 3 {
		recv.B = fmt.Sprintf("%s/S", size[0])
		recv.KB = fmt.Sprintf("%s/S", size[1])
		recv.MB = fmt.Sprintf("%s/S", size[2])
	}
}

type WsModel struct {
	BootTime   string     `json:"bootTime"`
	CpuInfo    CpuInfo    `json:"cpuInfo"`
	MemoryInfo MemoryInfo `json:"memoryInfo"`
	Ip         string     `json:"ip"`
	NetIoInfo  NetIoInfo  `json:"netIoInfo"`
}
