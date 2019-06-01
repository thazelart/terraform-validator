package fs

import (
	"github.com/thazelart/terraform-validator/internal/utils"
	"io/ioutil"
	"path"
	"strings"
)

// File is a simple structure to permit fs function overriding in others
// terraform-validator subpackages
type Folder struct {
	Path    string
	Content []File
}

// listTerraformFiles get the terraform file list in the given pathF
func ListTerraformFiles(pathF string) []string {
	var files []string

	filesInfo, err := ioutil.ReadDir(pathF)
	utils.EnsureOrFatal(err)

	for _, f := range filesInfo {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".tf") {
			files = append(files, path.Join(pathF, f.Name()))
		}
	}

	return files
}

// NewTerraformFolder return you a Folder var that contains all the Terraform files in the given pathF
func NewTerraformFolder(pathF string) Folder {
	files := ListTerraformFiles(pathF)
	var fileList []File

	for _, f := range files {
		fileList = append(fileList, NewFile(f))
	}

	return Folder{Path: pathF, Content: fileList}
}
