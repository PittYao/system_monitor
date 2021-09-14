package util

import (
	"fmt"
	"math"
)

// Decimal 保留两位小数
func Decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}

// FormatByteSize 字节的单位转换 保留两位小数
func FormatByteSize(byteSize int64) (size string) {
	if byteSize < 1024 {
		//return strconv.FormatInt(byteSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(byteSize)/float64(1))
	} else if byteSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(byteSize)/float64(1024))
	} else if byteSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(byteSize)/float64(1024*1024))
	} else if byteSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(byteSize)/float64(1024*1024*1024))
	} else if byteSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(byteSize)/float64(1024*1024*1024*1024))
	} else { //if byteSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(byteSize)/float64(1024*1024*1024*1024*1024))
	}
}

func FormatByteSizeForByteAndKbAndMb(byteSize int64) (size []string) {
	byte := fmt.Sprintf("%.2fB", float64(byteSize)/float64(1))
	Kb := fmt.Sprintf("%.2fKB", float64(byteSize)/float64(1024))
	Mb := fmt.Sprintf("%.2fMB", float64(byteSize)/float64(1024*1024))
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
	return fmt.Sprintf("%.2fGB", float64(byteSize)/float64(1024*1024*1024))
}
