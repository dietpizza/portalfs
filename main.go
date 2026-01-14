package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
)

type AppConfig struct {
	App      AppSettings `mapstructure:"app"`
	Projects []Volume    `mapstructure:"volumes"`
}

type AppSettings struct {
	Port int `mapstructure:"port"`
}

type Volume struct {
	Name string `mapstructure:"name"`
	Path string `mapstructure:"path"`
}

func SetupAppConfig() (AppConfig, error) {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config AppConfig
	// "unmarshal" / parse yaml code
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to parse config file")
		return AppConfig{}, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return config, nil
}

func main() {
	// Load config yaml
	config, err := SetupAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	server := fiber.New()
	SetupRoutes(server)

	log.Fatal(server.Listen(fmt.Sprintf(":%d", config.App.Port)))
}
