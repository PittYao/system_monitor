package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//Config global
var Config = loadConfig()

//ConfigST struct
type ConfigST struct {
	Server ServerST `json:"server"`
	Cron   CronST   `json:"cron"`
}

//ServerST struct
type ServerST struct {
	HTTPPort   string `json:"http_port"`
	HTTPSPort  string `json:"https_port"`
	Ssl        bool   `json:"ssl"`
	SslCrtPath string `json:"ssl_crt_path"`
	SslKeyPath string `json:"ssl_key_path"`
}

//CronST struct
type CronST struct {
	TimeInterval string `json:"time_interval"`
}

func loadConfig() *ConfigST {
	var tmp ConfigST
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(data, &tmp)
	return &tmp
}
