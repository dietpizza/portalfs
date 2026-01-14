package main

import (
	"dietpizza/portalfs/config"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Load config yaml
	config := config.GetConfig()

	server := fiber.New()
	SetupRoutes(server)

	log.Fatal(server.Listen(fmt.Sprintf(":%d", config.App.Port)))
}
