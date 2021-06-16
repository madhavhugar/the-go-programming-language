package __basic_data_types

import (
	"strings"
)

func basename1(path string) string {
	var lastBackSlashIndex int
	for i, c := range path {
		if c == '/' {
			lastBackSlashIndex = i
		}
	}

	dotIndex := len(path)
	for i := dotIndex-1; i >= 0; i-- {
		if path[i] == '.' {
			dotIndex = i
			break
		}
	}

	return path[lastBackSlashIndex+1:dotIndex]
}

func basename2(path string) string {
	lastBackSlashIndex := strings.LastIndex(path, "/")

	lastDotIndex := strings.LastIndex(path, ".")
	if lastDotIndex == -1 {
		return path[lastBackSlashIndex+1:]
	}

	return path[lastBackSlashIndex+1: lastDotIndex]
}
