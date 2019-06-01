package fs_test

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/fs"
)

func ExampleFileEqual() {
	file1 := fs.NewFile("my_file.txt")
	file2 := fs.NewFile("another_one.txt")

	result := file1.FileEqual(file2)

	fmt.Printf("It is %t to say that my files are equals !", result)
}
