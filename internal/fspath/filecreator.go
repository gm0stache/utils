package fspath

import (
	"io/fs"
)

const (
	default_dir_permissons_code = 0700
)

type FsCreator interface {
	MkdirAll(path string, perm fs.FileMode) error
	Create(name string) error
}

// CreateFile creates all missing parent dir's in the file and the file itself.
func CreateFile(creator FsCreator, path string) error {
	pathComponents := getPathComponents(path)
	err := creator.MkdirAll(pathComponents.DirPath, default_dir_permissons_code)
	if err != nil {
		return err
	}
	err = creator.Create(path)
	return err
}
