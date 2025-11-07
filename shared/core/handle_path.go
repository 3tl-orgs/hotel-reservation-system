package core

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func ToFileURL(path string) string {
	path = filepath.ToSlash(path)
	if runtime.GOOS == "windows" {
		// Windows: D:/... → file://D:/...
		if strings.HasPrefix(path, "/") {
			path = strings.TrimPrefix(path, "/")
		}
		return fmt.Sprintf("file://%s", path)
	}
	// Unix-like: /home/... → file:///home/...
	return fmt.Sprintf("file:///%s", path)
}
