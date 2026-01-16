package routes

import (
	"dietpizza/portalfs/config"
	"dietpizza/portalfs/util"
	"net/http"
	"path"
	"slices"

	"github.com/gin-gonic/gin"
)

func ListingHandler(c *gin.Context) {
	appConfig := config.GetConfig()

	webPath := util.NormalizePath(c.Query("path"))
	volumeName := c.Param("volume")

	status, error := http.StatusInternalServerError, "unknown error"

	index := slices.IndexFunc(appConfig.Volumes, func(e config.Volume) bool {
		return e.Name == volumeName
	})

	if index == -1 {
		status, error = http.StatusNotFound, "volume not found"
	} else {
		volumeInfo := appConfig.Volumes[index]
		finalOsPath := path.Join(volumeInfo.Path, webPath)

		if util.DoesDirExist(finalOsPath) {
			dirContents, err := util.GetDirectoryListing(finalOsPath, webPath)

			if err != nil {
				status, error = http.StatusInternalServerError, "error reading directory"
			} else {
				c.JSON(http.StatusOK, dirContents)
				return
			}
		} else {
			status, error = http.StatusNotFound, "folder not found"
		}
	}

	c.JSON(status, error)
}

func VolumeListHandler(c *gin.Context) {
	appConfig := config.GetConfig()
	c.JSON(http.StatusOK, appConfig.Volumes)
}

func SetupListingGroup(router *gin.Engine) {
	group := router.Group("/listing")

	group.GET("/:volume", ListingHandler)
	group.GET("/", VolumeListHandler)
}
