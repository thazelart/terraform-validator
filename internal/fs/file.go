package fs

import (
	"bytes"
	"github.com/thazelart/terraform-validator/pkg/utils"
	"io/ioutil"
	"path/filepath"
)

// File is a simple structure to permit fs function overriding in others
// terraform-validator subpackages. It contains the file Path and it's content
// in []byte.
type File struct {
	Path    string
	Content []byte
}

// NewFile create a new File from the given path
func NewFile(path string) File {
	content, err := ioutil.ReadFile(path)
	utils.EnsureOrFatal(err)

	return File{Path: path, Content: content}
}

// GetFilename return you the filename instead of it's full path
func (file File) GetFilename() string {
	return filepath.Base(file.Path)
}

// FileEqual ensures that the two files have the same content
func (file File) FileEqual(file2 File) bool {
	return bytes.Equal(file.Content, file2.Content)
}
