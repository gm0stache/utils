package fspath

import (
	"path/filepath"
)

type pathComponents struct {
	Filename string
	DirPath  string
}

func getPathComponents(path string) pathComponents {
	return pathComponents{
		Filename: filepath.Base(path),
		DirPath:  filepath.Dir(path),
	}
}
