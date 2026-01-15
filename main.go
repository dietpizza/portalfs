package main

import (
	"dietpizza/portalfs/config"
	"dietpizza/portalfs/middleware"
	"dietpizza/portalfs/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.GetConfig()
	server := ConfigureServer()

	err := server.Run(fmt.Sprintf(":%d", config.App.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func ConfigureServer() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(middleware.CorsAllowAll())

	routes.SetupRoutes(router)

	return router
}
