package http

import (
	"log"
	"sjw_system_monitor/config"
	"sjw_system_monitor/router"
)

type RtspUrlDTO struct {
	RtspUrl string `json:"rtspUrl" binding:"required"`
}

func ServeHTTP() {
	router := router.InitRouter()

	err := router.Run(config.Config.Server.HTTPPort)
	if err != nil {
		log.Fatalln("Start HTTP Server error", err)
	}
}
