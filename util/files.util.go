package util

import (
	"errors"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

type FileMeta struct {
	Filename     string `json:"filename"`
	RelativePath string `json:"relative_path"`
	Mtime        int64  `json:"mtime"`
	Ctime        int64  `json:"ctime"`
	MimeType     string `json:"mimetype"`
	Size         int64  `json:"size"`
}

func GetDirName(path string) string {
	segments := strings.Split(path, "/")

	return segments[len(segments)-1]
}

func DoesDirExist(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}

	if file.IsDir() {
		return true
	}

	return false
}

func NormalizePath(path string) string {
	if path == "." || path == "" {
		return "/"
	}

	if path[0] != '/' {
		return "/" + path
	}

	return path
}

func GetMimeType(info os.FileInfo) string {
	mimeType := mime.TypeByExtension(filepath.Ext(info.Name()))
	if mimeType == "" {
		mimeType = "application/octet-stream" // default for unknown types
	}
	return mimeType
}

func GetDirectoryListing(osPath string, webPath string) ([]FileMeta, error) {
	listing := []FileMeta{}

	content, err := os.ReadDir(osPath)
	if err != nil {
		return listing, errors.New("error reading directory")
	}

	for i := range content {
		info, err := content[i].Info()
		if err != nil {
			continue
		}

		meta := FileMeta{
			Filename:     info.Name(),
			RelativePath: fmt.Sprintf("%s/%s", webPath, info.Name()),
			MimeType:     GetMimeType(info),
			Mtime:        info.ModTime().Unix(),
			Ctime:        2000,
			Size:         info.Size(),
		}

		listing = append(listing, meta)
	}

	return listing, nil
}
