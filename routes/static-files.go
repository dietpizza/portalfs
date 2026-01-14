package routes

import (
	"dietpizza/portalfs/config"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func GetVolumeName(path string) string {
	segments := strings.Split(path, "/")

	return segments[len(segments)-1]
}

func SetupStaticFilesGroup(app *fiber.App) {
	static_routes := app.Group("/static")

	volumes := config.GetConfig().Volumes

	for _, v := range volumes {

		if DoesFolderExist(v) {
			volume_name := GetVolumeName(v)

			static_routes.Get("/"+volume_name+"*", static.New("", static.Config{
				FS:     os.DirFS(v),
				Browse: true,
			}))
		}
	}
}

func DoesFolderExist(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}

	if file.IsDir() {
		return true
	}

	return false
}
