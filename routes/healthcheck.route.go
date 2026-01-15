package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"uptime": time.Since(startTime).String(),
	})
}

func HealthCheckGroup(router *gin.Engine) {
	router.GET("/health", HealthCheckHandler)
}
