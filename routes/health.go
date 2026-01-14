package routes

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
)

var startTime = time.Now()

func HealthHandler(c fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "healthy",
		"uptime": time.Since(startTime).String(),
	})
}

func SetupHealthRoutes(app *fiber.App) {
	app.Get("/health", HealthHandler)
}
