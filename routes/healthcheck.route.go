package routes

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
)

var startTime = time.Now()

func HealthCheckHandler(c fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "healthy",
		"uptime": time.Since(startTime).String(),
	})
}

func SetupHealthCheckRoutes(app *fiber.App) {
	app.Get("/health", HealthCheckHandler)
}
