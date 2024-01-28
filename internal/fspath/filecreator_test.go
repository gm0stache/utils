package fspath

import (
	"io/fs"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

type fakeFsCreator struct {
	fs *afero.MemMapFs
}

func (rc *fakeFsCreator) MkdirAll(path string, perm fs.FileMode) error {
	return rc.fs.MkdirAll(path, perm)
}

func (rc *fakeFsCreator) Create(name string) error {
	_, err := rc.fs.Create(name)
	return err
}

func TestFsCreator(t *testing.T) {
	cases := map[string]string{
		"simple file":                   "./exampleDir/example.txt",
		"dotfile, relative to user dir": "~/.config/.vimrc",
		"multi-dot-file, deeply nested": "./test/test/test/testing.dev.env",
	}

	// arrange
	fakeFS := fakeFsCreator{
		fs: new(afero.MemMapFs),
	}

	for name, path := range cases {
		t.Run(name, func(t *testing.T) {
			// act
			err := CreateFile(&fakeFS, path)

			// assert
			assert.NoError(t, err)
			// fileExists, _ := fakeFS.fs.Exists(path)
			// assert.True(t, fileExists, "file should be created")
		})
	}
}
