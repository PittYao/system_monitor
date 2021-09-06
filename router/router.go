package router

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"log"
	"net/http"
	"sjw_system_monitor/api"
	. "sjw_system_monitor/config"
	"sjw_system_monitor/middleware"
	"sjw_system_monitor/ws"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Cors())

	router.GET("/ws", ws.WsHandler)
	router.GET("/cpu", api.CpuInfoHandler)
	router.GET("/memory", api.MemoryHandler)
	router.GET("/bootTime", api.BootTimeHandler)
	router.GET("/netIo", api.NetIoHandler)
	router.GET("/ip", api.IPHandler)

	// 静态文件代理
	router.StaticFS("/static", http.Dir("ws"))

	if Config.Server.Ssl == true {
		// 启动https
		router.Use(TlsHandler())
		err := router.RunTLS(Config.Server.HTTPSPort, Config.Server.SslCrtPath, Config.Server.SslKeyPath)
		if err != nil {
			log.Fatalln("启动https失败 ", err)
		}
	} else {
		// 启动http
		err := router.Run(Config.Server.HTTPPort)
		if err != nil {
			log.Fatalln("启动http失败 ", err)
		}
	}

	return router
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     Config.Server.HTTPSPort,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
