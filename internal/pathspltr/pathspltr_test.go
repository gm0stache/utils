package pathspltr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilePathSplitting(t *testing.T) {
	cases := map[string]struct {
		givenPath        string
		expectedFilename string
		expectedDirPath  string
	}{
		"filename with one file extension": {
			givenPath:        "/example/test.txt",
			expectedFilename: "test.txt",
			expectedDirPath:  "/example",
		},
		"filename with double file extension": {
			givenPath:        "/example/test.dev.txt",
			expectedFilename: "test.dev.txt",
			expectedDirPath:  "/example",
		},
		"filename without extension": {
			givenPath:        "/example/Makefile",
			expectedFilename: "Makefile",
			expectedDirPath:  "/example",
		},
		"filename with leading dot": {
			givenPath:        "~/.vimrc",
			expectedFilename: ".vimrc",
			expectedDirPath:  "~",
		},
	}

	for testname, testData := range cases {
		t.Run(testname, func(t *testing.T) {
			// act
			pathComponents := GetFilePathParts(testData.givenPath)

			// assert
			assert.EqualValues(t, pathComponents.DirPath, testData.expectedDirPath)
			assert.EqualValues(t, pathComponents.Filename, testData.expectedFilename)
		})
	}
}
