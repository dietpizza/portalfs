package main

import (
	"dietpizza/portalfs/routes"

	"github.com/gofiber/fiber/v3"
)

func SetupHealthGroup(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", routes.HealthHandler)
}

func SetupRoutes(app *fiber.App) {
	SetupHealthGroup(app)

	app.Get("/", routes.HealthHandler)
}
