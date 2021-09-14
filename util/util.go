package util

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"math"
	"strings"
)

const (
	B = 1

	KB = 1024 * B

	MB = 1024 * KB

	GB = 1024 * MB

	TB = 1024 * GB

	EB = 1024 * TB
)

// Decimal 保留两位小数
func Decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}

// FormatByteSize 字节的单位转换 保留两位小数
func FormatByteSize(byteSize int64) (size string) {
	if byteSize < B {
		//return strconv.FormatInt(byteSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(byteSize)/float64(B))
	} else if byteSize < (KB) {
		return fmt.Sprintf("%.2fKB", float64(byteSize)/float64(KB))
	} else if byteSize < (MB) {
		return fmt.Sprintf("%.2fMB", float64(byteSize)/float64(MB))
	} else if byteSize < (GB) {
		return fmt.Sprintf("%.2fGB", float64(byteSize)/float64(GB))
	} else if byteSize < (TB) {
		return fmt.Sprintf("%.2fTB", float64(byteSize)/float64(TB))
	} else { //if byteSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(byteSize)/float64(EB))
	}
}

func FormatByteSizeForByteAndKbAndMb(byteSize int64) (size []string) {
	byte := fmt.Sprintf("%.2fB", float64(byteSize)/float64(B))
	Kb := fmt.Sprintf("%.2fKB", float64(byteSize)/float64(KB))
	Mb := fmt.Sprintf("%.2fMB", float64(byteSize)/float64(MB))
	return append(size, byte, Kb, Mb)
}

// ResolveTime 秒转换为日时分秒和补零操作
func ResolveTime(seconds int64) (day, hour, minute, second int64) {
	day = seconds / (24 * 3600)
	hour = (seconds - day*3600*24) / 3600
	minute = (seconds - day*24*3600 - hour*3600) / 60
	second = seconds - day*24*3600 - hour*3600 - minute*60
	return
}

func Float642String(data float64) string {
	return fmt.Sprintf("%g", data)
}

// Float642StringWith2Point float64转string保留两位小数
func Float642StringWith2Point(data float64) string {
	return fmt.Sprintf("%.2f%%", data)
}

// FormatByteSizeForGb byte字节转为Gb
func FormatByteSizeForGb(byteSize uint64) (size string) {
	return fmt.Sprintf("%.2fGB", float64(byteSize)/float64(GB))
}

// JudgePlatformIsWindows 判断系统平台是否是Windows
func JudgePlatformIsWindows() bool {
	information, _, _, _ := host.PlatformInformation()
	return strings.Contains(information, "Windows")
}
