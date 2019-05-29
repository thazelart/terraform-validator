package fs

import (
	"bytes"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

var LogFatal = log.Fatal

type File struct {
	Path string
}

// Ensure the error in nil or uses log.Fatal
func EnsureOrFatal(err error) {
	if err != nil {
		LogFatal(err)
	}
}

// Read the given file
func (file File) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(file.Path)
}

// Get the file list of the given path folder
func (folder File) ListTerraformFiles() []string {
	var files []string

	filesInfo, err := ioutil.ReadDir(folder.Path)
	EnsureOrFatal(err)

	for _, f := range filesInfo {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".tf") {
			files = append(files, path.Join(folder.Path, f.Name()))
		}
	}

	return files
}

// Ensure that the two files are equals
func (file1 File) FileEqual(file2 File) bool {
	content1, err1 := file1.ReadFile()
	EnsureOrFatal(err1)

	content2, err2 := file2.ReadFile()
	EnsureOrFatal(err2)

	return bytes.Equal(content1, content2)
}
