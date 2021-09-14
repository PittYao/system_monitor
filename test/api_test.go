package test

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/v3/disk"
	"sjw_system_monitor/util"
	"strings"
	"testing"
)

func TestGetWindowsDisk(t *testing.T) {
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

func TestPlatform(t *testing.T) {
	information, s, s2, _ := host.PlatformInformation()

	contains := strings.Contains(information, "Windows")

	fmt.Println(contains)

	fmt.Println(information)
	fmt.Println(s)
	fmt.Println(s2)
}

func TestGetLinuxDisk(t *testing.T) {
	diskInfo, _ := disk.Usage("/")
	fmt.Printf(
		"disk %s usedPercent:%v total:%v used:%v free:%v\n",
		"/",
		util.Float642StringWith2Point(diskInfo.UsedPercent),
		util.FormatByteSizeForGb(diskInfo.Total),
		util.FormatByteSizeForGb(diskInfo.Used),
		util.FormatByteSizeForGb(diskInfo.Free))

}
