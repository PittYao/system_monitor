package ws

import (
	"encoding/json"
	"fmt"
	rcron "github.com/robfig/cron/v3"
	"log"
	"sjw_system_monitor/config"
	"sjw_system_monitor/system"
)

// SystemMonitorCron 定时获取信息资源
func SystemMonitorCron() {
	// 有ws的任一连接启动定时获取任务，ws中0个连接则停止定时任务
	c := rcron.New()

	// 定时任务执行间隔时间
	timeInterval := config.Config.Cron.TimeInterval
	time := "@every " + timeInterval + "s"
	log.Println("定时任务间隔是:", timeInterval, "s")

	// 执行任务
	c.AddFunc(time, func() {
		fmt.Println("tick every " + timeInterval + " second")
		// todo 获取资源信息
		systemInfo := system.WsGetSystemInfo()
		jsonData, _ := json.Marshal(systemInfo)
		// 发送到所有ws连接
		Manager.Broadcast <- jsonData

		if len(Manager.Clients) == 0 {
			// 停止定时任务
			log.Println("ws没有任何连接,停止定时任务")
			c.Stop()
		}
	})

	c.Start()
}
