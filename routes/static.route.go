package routes

import (
	"dietpizza/portalfs/config"
	"dietpizza/portalfs/util"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func CreateStaticRoute(group fiber.Router, path string, osPath string) {
	group.Get("/"+path+"*", static.New("", static.Config{
		FS:        os.DirFS(osPath),
		Browse:    false,
		ByteRange: true,
		Download:  true,
	}))
}

func SetupStaticFilesGroup(app *fiber.App) {
	static_routes := app.Group("/static")

	volumes := config.GetConfig().Volumes

	for _, osPath := range volumes {
		folderExists := util.DoesDirExist(osPath)

		if folderExists {
			urlPath := util.GetDirName(osPath)
			CreateStaticRoute(static_routes, urlPath, osPath)
		}
	}
}
