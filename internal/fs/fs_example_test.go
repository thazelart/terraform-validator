package fs_test

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/fs"
)

func ExampleListTerraformFiles() {
	path := "/tmp"
	var fileList []string
	fileList = fs.File{Path: path}.ListTerraformFiles()
	fmt.Printf("%v", fileList)
}

func ExampleFileEqual() {
	file1 := fs.File{Path: "my_file.txt"}
	file2 := fs.File{Path: "another_one.txt"}

	result := file1.FileEqual(file2)

	fmt.Printf("It is %t to say that my files are equals !", result)
}
