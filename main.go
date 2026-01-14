package main

import (
	"dietpizza/portalfs/config"
	"dietpizza/portalfs/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config := config.GetConfig()
	host_port := fmt.Sprintf("0.0.0.0:%d", config.App.Port)

	server := fiber.New()
	routes.SetupRoutes(server)

	log.Fatal(server.Listen(host_port))
}
