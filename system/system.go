package system

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
	"sjw_system_monitor/util"
	"time"
)

var (
	TimeFormatter = "2006-01-02 15:04:05"
)

// GetCpuPercent cpu使用lv
func GetCpuPercent() CpuInfo {

	//cpu 核数
	counts, _ := cpu.Counts(true)
	log.Println("cpu核数：", counts)

	percent, _ := cpu.Percent(time.Second, true)
	log.Println("cpu各个核使用率：", percent)

	total := 0.0
	for _, value := range percent {
		total += value
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
		UsePercent:   percent,
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

// GetIO 网络下载速度
func GetIO() NetIoInfo {
	counters, _ := net.IOCounters(false)
	//log.Println(counters[0])
	log.Println("发送数据大小：", util.FormatByteSize(int64(counters[0].BytesSent)))
	log.Println("接收数据大小：", util.FormatByteSize(int64(counters[0].BytesRecv)))

	time.Sleep(time.Second * 1)

	newCounters, _ := net.IOCounters(false)
	//log.Println(newCounters[0])
	log.Println("发送数据大小：", util.FormatByteSize(int64(newCounters[0].BytesSent)))
	log.Println("接收数据大小：", util.FormatByteSize(int64(newCounters[0].BytesRecv)))

	spcSent := util.FormatByteSize(int64(newCounters[0].BytesSent - counters[0].BytesSent))
	spcRecv := util.FormatByteSize(int64(newCounters[0].BytesRecv - counters[0].BytesRecv))

	log.Println("1秒内上传的差值：", spcSent, "/S")
	log.Println("1秒内下载的差值：", spcRecv, "/S")

	spcSentStr := fmt.Sprintf("1秒内发送的差值:%s/S", spcSent)
	spcRecvStr := fmt.Sprintf("1秒内接收的差值:%s/S", spcRecv)

	netIoInfo := NetIoInfo{
		SentSpc: spcSentStr,
		RecvSpc: spcRecvStr,
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
