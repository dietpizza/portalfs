package util

import (
	"os"
	"strings"
)

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
