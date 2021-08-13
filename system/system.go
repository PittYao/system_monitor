package system

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
	localNet "net"
	"sjw_system_monitor/util"
	"time"
)

var (
	TimeFormatter = "2006-01-02 15:04:05"
)

// GetCpuPercent cpu使用率
func GetCpuPercent() CpuInfo {

	//cpu 核数
	counts, _ := cpu.Counts(true)
	log.Println("cpu核数：", counts)

	percent, _ := cpu.Percent(time.Second, true)
	log.Println("cpu各个核使用率：", percent)

	var usePercent []string

	total := 0.0
	for _, value := range percent {
		total += value
		usePercentItem := util.Float642String(value)
		usePercent = append(usePercent, usePercentItem)
	}
	// 总核数
	totalPercent := len(percent) * 100
	f := (total / float64(totalPercent)) * 100
	e := util.Decimal(f)
	strTotalPercent := util.Float642String(e) + "%"
	log.Println("cpu总使用率：", strTotalPercent)

	infoStats, _ := cpu.Info()
	decimal := infoStats[0].Mhz / 1000
	strMhz := util.Float642String(decimal) + "GHz"
	log.Println("cpu赫兹： ", strMhz)

	cpuInfo := CpuInfo{
		CpuCount:     counts,
		UsePercent:   usePercent,
		TotalPercent: strTotalPercent,
		Mhz:          strMhz,
	}
	return cpuInfo

}

// GetMemPercent 内存使用率
func GetMemPercent() MemoryInfo {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	log.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	log.Println(v)

	memoryInfo := MemoryInfo{
		FreeMemory:  util.FormatByteSize(int64(v.Free)),
		TotalMemory: util.FormatByteSize(int64(v.Total)),
		UsedPercent: util.Float642String(v.UsedPercent) + "%",
	}
	return memoryInfo
}

// GetDiskPercent 磁盘占用率
func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[1].Mountpoint)
	return diskInfo.UsedPercent
}

// GetNetIO 网络下载速度
func GetNetIO() NetIoInfo {
	counters, _ := net.IOCounters(false)
	//log.Println(counters[0])
	log.Println("发送数据大小：", util.FormatByteSize(int64(counters[0].BytesSent)))
	log.Println("接收数据大小：", util.FormatByteSize(int64(counters[0].BytesRecv)))

	time.Sleep(time.Second * 1)

	newCounters, _ := net.IOCounters(false)
	//log.Println(newCounters[0])
	log.Println("发送数据大小：", util.FormatByteSize(int64(newCounters[0].BytesSent)))
	log.Println("接收数据大小：", util.FormatByteSize(int64(newCounters[0].BytesRecv)))

	sendSize := int64(newCounters[0].BytesSent - counters[0].BytesSent)
	spcSent := util.FormatByteSize(sendSize)

	recvSize := int64(newCounters[0].BytesRecv - counters[0].BytesRecv)
	spcRecv := util.FormatByteSize(recvSize)

	log.Println("1秒内上传的差值：", spcSent, "/S")
	log.Println("1秒内下载的差值：", spcRecv, "/S")

	// 格式化为 Kb
	spcSentStr := fmt.Sprintf("%s/S", spcSent)
	spcRecvStr := fmt.Sprintf("%s/S", spcRecv)

	// 格式化为 b kb mb 三种格式
	sendByteAndKbAndMb := util.FormatByteSizeForByteAndKbAndMb(sendSize)
	recvByteAndKbAndMb := util.FormatByteSizeForByteAndKbAndMb(recvSize)

	netIoInfoSentInfo := NetIoInfoSentInfo{}
	netIoInfoSentInfo.Formatter(sendByteAndKbAndMb)

	netIoInfoRecvInfo := NetIoInfoRecvInfo{}
	netIoInfoRecvInfo.Formatter(recvByteAndKbAndMb)

	netIoInfo := NetIoInfo{
		SentSpc:           spcSentStr,
		RecvSpc:           spcRecvStr,
		NetIoInfoSentInfo: netIoInfoSentInfo,
		NetIoInfoRecvInfo: netIoInfoRecvInfo,
	}

	return netIoInfo
}

// GetBootTime 开机时长
func GetBootTime() string {
	bootTime, _ := host.BootTime()
	//log.Println(bootTime)
	tm := time.Unix(int64(bootTime), 0)
	log.Println("系统启动时间:", tm.Format(TimeFormatter))

	now := time.Now().Unix()
	//log.Println(now)
	nowTime := time.Unix(now, 0)
	log.Println("现在时间:", nowTime.Format(TimeFormatter))

	spc := now - int64(bootTime)
	day, hour, minute, second := util.ResolveTime(spc)
	bootTimeStr := fmt.Sprintf("%d天%d时%d分%d秒", day, hour, minute, second)
	return bootTimeStr
}

// GetOutboundIP  获取ip地址
func GetOutboundIP() string {
	conn, err := localNet.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*localNet.UDPAddr)
	fmt.Println(localAddr.String())
	return localAddr.IP.String()
}

// WsGetSystemInfo ws获取所有资源信息
func WsGetSystemInfo() WsModel {
	wsModel := WsModel{
		BootTime:   GetBootTime(),
		CpuInfo:    GetCpuPercent(),
		MemoryInfo: GetMemPercent(),
		Ip:         GetOutboundIP(),
		NetIoInfo:  GetNetIO(),
	}

	return wsModel
}
