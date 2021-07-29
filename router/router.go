package router

import (
	"github.com/gin-gonic/gin"
	"sjw_system_monitor/api"
	"sjw_system_monitor/middleware"
	"sjw_system_monitor/ws"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Cors())

	router.GET("/ws", ws.WsHandler)
	router.GET("/cpu", api.CpuInfoHandler)
	router.GET("/memory", api.CpuMemoryHandler)
	router.GET("/bootTime", api.BootTimeHandler)
	router.GET("/netIo", api.NetIoHandler)
	router.GET("/ip", api.IPHandler)

	return router
}
