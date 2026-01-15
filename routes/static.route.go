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

	for _, osPath := range volumes {
		folderExists := util.DoesDirExist(osPath)

		if folderExists {
			urlPath := util.GetDirName(osPath)
			CreateStaticRoute(staticGroup, urlPath, osPath)
		}
	}
}
