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
	Port int `mapstructure:"port"`
}

type Volume struct {
	Name string `mapstructure:"name"`
	Path string `mapstructure:"path"`
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
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		config = &DefaultAppConfig
		err = viper.Unmarshal(config)
		if err != nil {
			log.Fatal("Unable to parse config file")
		}
	})

	return config
}
