package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sjw_system_monitor/system"
)

func CpuInfoHandler(c *gin.Context) {
	cpuInfo := system.GetCpuPercent()

	var responseDTO ResponseDTO

	c.JSON(http.StatusOK, responseDTO.SuccessWithData("获取cpu信息", cpuInfo))
}

func MemoryHandler(c *gin.Context) {
	memoryInfo := system.GetMemPercent()

	var responseDTO ResponseDTO

	c.JSON(http.StatusOK, responseDTO.SuccessWithData("获取内存信息", memoryInfo))
}

func DiskHandler(c *gin.Context) {
	diskInfos := system.GetDiskPercent()

	var responseDTO ResponseDTO

	c.JSON(http.StatusOK, responseDTO.SuccessWithData("获取磁盘信息", diskInfos))
}

func BootTimeHandler(c *gin.Context) {
	bootTimeInfo := system.GetBootTime()

	var responseDTO ResponseDTO

	c.JSON(http.StatusOK, responseDTO.SuccessWithData("获取系统启动时长信息", bootTimeInfo))
}

func NetIoHandler(c *gin.Context) {
	ioInfo := system.GetNetIO()

	var responseDTO ResponseDTO

	c.JSON(http.StatusOK, responseDTO.SuccessWithData("获取上传下载信息", ioInfo))
}

func IPHandler(c *gin.Context) {
	ipInfo := system.GetOutboundIP()

	var responseDTO ResponseDTO

	c.JSON(http.StatusOK, responseDTO.SuccessWithData("获取ip信息", ipInfo))
}
