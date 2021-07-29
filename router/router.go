package router

import (
	"github.com/gin-gonic/gin"
	"sjw_system_monitor/api"
	"sjw_system_monitor/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Cors())

	router.GET("/cpu", api.CpuInfoHandler)
	router.GET("/memory", api.CpuMemoryHandler)
	router.GET("/bootTime", api.BootTimeHandler)
	router.GET("/netIo", api.NetIoHandler)

	return router
}
