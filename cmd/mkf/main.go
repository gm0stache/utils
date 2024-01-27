package main

import (
	"fmt"
	"os"

	"github.com/gm0stache/utils/internal/pathspltr"
)

const (
	default_dir_permissons_code = 0700
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("invalid invocation. Usage: mkf file_path")
		os.Exit(-1)
	}
	filePath := args[1]
	pathComponents := pathspltr.GetFilePathParts(filePath)
	err := os.MkdirAll(pathComponents.DirPath, default_dir_permissons_code)
	if err != nil {
		fmt.Printf("Error:\n%s", err.Error())
		os.Exit(-1)
	}
	_, err = os.Create(filePath)
	if err != nil {
		fmt.Printf("Error:\n%s", err.Error())
		os.Exit(-1)
	}
}
