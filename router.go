package main

import (
	"dietpizza/portalfs/routes"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	routes.SetupStaticFilesGroup(app)
	routes.SetupHealthRoutes(app)
}
