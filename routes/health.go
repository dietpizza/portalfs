package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func HealthHandler(c fiber.Ctx) {
	c.SendStatus(http.StatusOK)
}
