package test

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"sjw_system_monitor/util"
	"testing"
)

func TestGetDisk(t *testing.T) {
	parts, err := disk.Partitions(true)
	if err != nil {
		panic(err)
	}
	for _, part := range parts {
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf(
			"disk %s usedPercent:%v total:%v used:%v free:%v\n",
			part.Device,
			util.Float642StringWith2Point(diskInfo.UsedPercent),
			util.FormatByteSizeForGb(diskInfo.Total),
			util.FormatByteSizeForGb(diskInfo.Used),
			util.FormatByteSizeForGb(diskInfo.Free))
	}
}
