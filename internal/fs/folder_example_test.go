package fs_test

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/fs"
)

func ExampleListTerraformFiles() {
	path := "../../testdata/ok_default_config"
	var fileList []string
	fileList = fs.ListTerraformFiles(path)
	fmt.Printf("%v", fileList)
}

func ExampleNewTerraformFolder() {
	path := "../../testdata/ok_default_config"
	folder := fs.NewTerraformFolder(path)

	fmt.Printf("%+v", folder)
}
