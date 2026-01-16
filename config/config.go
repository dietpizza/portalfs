package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App     AppSettings `mapstructure:"app"`
	Volumes []Volume    `mapstructure:"volumes"`
}
type AppSettings struct {
	Port int `json:"port"`
}
type Volume struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

var DefaultAppConfig = AppConfig{
	App: AppSettings{
		Port: 3000,
	},
	Volumes: []Volume{},
}
var (
	config *AppConfig
	once   sync.Once
)

func GetConfig() *AppConfig {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		viper.AddConfigPath(".")                      // dev config
		viper.AddConfigPath("$HOME/.config/portalfs") // user-specific
		viper.AddConfigPath("/etc/portalfs")          // system-wide

		err := viper.ReadInConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Println("config file not found -> using defaults ")
			} else {
				log.Fatalf("failed to load config: %v", err)
			}
		}

		config = &DefaultAppConfig

		err = viper.Unmarshal(config)
		if err != nil {
			log.Fatalf("failed to parse config %v", err)
		}
	})

	return config
}
