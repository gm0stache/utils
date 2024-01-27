package pathspltr

import (
	"path/filepath"
)

type PathComponent struct {
	Filename string
	DirPath  string
}

func GetFilePathParts(path string) PathComponent {
	return PathComponent{
		Filename: filepath.Base(path),
		DirPath:  filepath.Dir(path),
	}
}
