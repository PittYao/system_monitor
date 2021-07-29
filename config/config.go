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
}

//ServerST struct
type ServerST struct {
	HTTPPort string `json:"http_port"`
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
