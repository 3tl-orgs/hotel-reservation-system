package core

import (
	"path/filepath"
	"runtime"
	"strings"
)

func ToFileURL(path string) string {
	if runtime.GOOS == "windows" {
		// Clean and use forward slashes
		path = strings.ReplaceAll(filepath.ToSlash(path), " ", "\\ ")
		// Remove drive letter and ensure triple slash
		if strings.HasPrefix(path, "/") {
			return "file://" + path
		}
		// Handle C:/... → /C:/...
		drive := strings.ToUpper(string(path[0]))
		rest := path[2:]
		return "file:///" + drive + ":" + rest
	}

	// Unix-like: just ensure forward slashes
	return "file://" + filepath.ToSlash(path)
}
