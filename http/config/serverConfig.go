package config

import "sync"

type ServerConfig struct {
	Name         string `json:"name"`
	IP           string `json:"ip"`
	Port         string `json:"port"`
	ReadTimeout  uint   `json:"readTimeout"`
	WriteTimeout uint   `json:writeTimeout`
}

type GlobalConfig struct {
	Server []ServerConfig `json:"servers"`
}

var config *GlobalConfig

var once sync.Once

func initConfig() {
	config = &GlobalConfig{
		Server: []ServerConfig{{IP: "0.0.0.0", Port: "9999", ReadTimeout: 10, WriteTimeout: 10}},
	}
}

func GetConfigInstance() *GlobalConfig {
	once.Do(initConfig)
	return config
}
