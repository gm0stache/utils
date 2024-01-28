package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/gm0stache/utils/internal/fspath"
)

const (
	default_dir_permissons_code = 0700
)

type RealFsCreator struct {
}

func (rc *RealFsCreator) MkdirAll(path string, perm fs.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (rc *RealFsCreator) Create(name string) error {
	_, err := os.Create(name)
	return err
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("invalid invocation.\nusage: mkf file_path")
		os.Exit(-1)
	}
	creator := RealFsCreator{}
	filePath := args[1]
	err := fspath.CreateFile(&creator, filePath)
	if err != nil {
		fmt.Printf("Error:\n%s", err.Error())
		os.Exit(-1)
	}
}
