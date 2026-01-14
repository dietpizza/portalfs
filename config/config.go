package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App     AppSettings `mapstructure:"app"`
	Volumes []string    `mapstructure:"volumes"`
}

type AppSettings struct {
	Port int `mapstructure:"port"`
}

var (
	config *AppConfig
	once   sync.Once
)

func GetConfig() *AppConfig {
	once.Do(func() {
		viper.SetConfigName("default")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		config = &AppConfig{}
		// "unmarshal" / parse yaml code
		err = viper.Unmarshal(config)
		if err != nil {
			log.Fatal("Unable to parse config file")
		}
	})

	return config
}
