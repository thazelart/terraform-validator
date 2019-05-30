package fs_test

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/fs"
	"testing"
)

func ExampleListTerraformFiles() {
	path := "/tmp"
	var fileList []string
	fileList = fs.File{path}.ListTerraformFiles()
	fmt.Printf("%v", fileList)
}

func ExampleFileEqual(t *testing.T) {
	file1 := fs.File{"my_file.txt"}
	file2 := fs.File{"another_one.txt"}

	result := file1.FileEqual(file2)

	fmt.Printf("It is %t to say that my files are equals !", result)
}
