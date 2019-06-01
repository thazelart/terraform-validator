package fs

import (
	"bytes"
	"github.com/thazelart/terraform-validator/internal/utils"
	"io/ioutil"
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

// FileEqual ensures that the two files have the same content
func (file File) FileEqual(file2 File) bool {
	return bytes.Equal(file.Content, file2.Content)
}
