package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	HealthCheckGroup(router)
	SetupStaticFilesGroup(router)
	SetupListingGroup(router)
}
