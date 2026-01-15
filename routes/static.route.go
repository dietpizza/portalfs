package routes

import (
	"dietpizza/portalfs/config"
	"dietpizza/portalfs/util"

	"github.com/gin-gonic/gin"
)

func CreateStaticRoute(router *gin.RouterGroup, path string, osPath string) {
	relPath := "/" + path

	router.StaticFS(relPath, gin.Dir(osPath, false))
}

func SetupStaticFilesGroup(router *gin.Engine) {
	volumes := config.GetConfig().Volumes
	staticGroup := router.Group("/static")

	for _, volume := range volumes {
		folderExists := util.DoesDirExist(volume.Path)

		if folderExists {
			CreateStaticRoute(staticGroup, volume.Name, volume.Path)
		}
	}
}
