package config

import (
	"errors"
	"sync"

	"github.com/BurntSushi/toml"
)

const (
	defaultServerPort              = 80
	defaultShutdownTimeoutInSecond = 60
)

type config struct {
	Server ServerConfig `toml:"server"`
}

type ServerConfig struct {
	Port                    int `toml:"port"`
	ShutdownTimeoutInSecond int `toml:"shutdown_timeout_in_second"`
}

var conf *config
var once sync.Once

func newDefaultConfig() *config {
	return &config{
		Server: ServerConfig{
			Port:                    defaultServerPort,
			ShutdownTimeoutInSecond: defaultShutdownTimeoutInSecond,
		},
	}
}

func InitializeConfig(path string) {
	once.Do(func() {
		conf = newDefaultConfig()
		if path != "" {
			if _, err := toml.DecodeFile(path, &conf); err != nil {
				panic(err)
			}
		}
	})
}

func GetConfig() *config {
	if conf == nil {
		panic(errors.New("config is nil"))
	}
	return conf
}
