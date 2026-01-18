package util

import (
	"errors"
	"mime"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type FileMeta struct {
	Filename     string    `json:"filename"`
	RelativePath string    `json:"relative_path"`
	Mtime        time.Time `json:"mtime"`
	Ctime        time.Time `json:"ctime"`
	MimeType     string    `json:"mimetype"`
	Size         int64     `json:"size"`
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
			RelativePath: path.Join(webPath, info.Name()),
			MimeType:     GetMimeType(info),
			Mtime:        info.ModTime(),
			Ctime:        info.ModTime(),
			Size:         info.Size(),
		}

		listing = append(listing, meta)

		sort.Slice(listing, func(i, j int) bool {
			return listing[i].Mtime.After(listing[j].Mtime)
		})
	}

	return listing, nil
}
