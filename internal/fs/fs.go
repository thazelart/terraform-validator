// fs package handle the fileSystem part of terraform-validator (Files and Directories)
package fs

import (
	"bytes"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

// LogFatal is the log.Fatal go built-in function by default. Permit to change
// the behaviour of that variable in order to test the EnsureOrFatal function
var LogFatal = log.Fatal

// File is a simple structure to permit fs function overriding in others
// terraform-validator subpackages
type File struct {
	Path string
}

// EnsureOrFatal ensures the error in nil or uses LogFatal
func EnsureOrFatal(err error) {
	if err != nil {
		LogFatal(err)
	}
}

// ReadFile reads the file named by filename and returns the contents.
// A successful call returns err == nil, not err == EOF. Because ReadFile
// reads the whole file, it does not treat an EOF from Read as an error
// to be reported.
func (file File) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(file.Path)
}

// ListTerraformFiles get the terraform file list in the given path folder
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

// FileEqual ensures that the two files are equals
func (file1 File) FileEqual(file2 File) bool {
	content1, err1 := file1.ReadFile()
	EnsureOrFatal(err1)

	content2, err2 := file2.ReadFile()
	EnsureOrFatal(err2)

	return bytes.Equal(content1, content2)
}
