package routes

import (
	"dietpizza/portalfs/config"
	"dietpizza/portalfs/util"
	"net/http"
	"path"
	"slices"

	"github.com/gin-gonic/gin"
)

func findVolume(name string) (config.Volume, bool) {
	volumes := config.GetConfig().Volumes
	for _, v := range volumes {
		if v.Name == name {
			return v, true
		}
	}

	return config.Volume{}, false
}

func ListingHandler(c *gin.Context) {
	appConfig := config.GetConfig()

	webPath := util.NormalizePath(c.Query("path"))
	volumeName := c.Param("volume")

	index := slices.IndexFunc(appConfig.Volumes, func(e config.Volume) bool {
		return e.Name == volumeName
	})

	if index == -1 {
		util.SendError(c, http.StatusNotFound, "volume not found")
		return
	}

	volume, ok := findVolume(volumeName)
	if !ok {
		util.SendError(c, http.StatusNotFound, "volume not found")
	}

	finalOsPath := path.Join(volume.Path, webPath)

	if !util.DoesDirExist(finalOsPath) {
		util.SendError(c, http.StatusNotFound, "folder not found")
		return
	}

	dirContents, err := util.GetDirectoryListing(finalOsPath, webPath)
	if err != nil {
		util.SendError(c, http.StatusInternalServerError, "error reading directory")
		return
	}

	c.JSON(http.StatusOK, dirContents)
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
