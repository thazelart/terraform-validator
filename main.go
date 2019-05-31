package main

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/fs"
)

func main() {
	workDir := fs.File{Path: "/tmp"}
	fmt.Println("Work in progress")

	fileList := workDir.ListTerraformFiles()
	fmt.Printf("We will work on that list of files: %v\n", fileList)
}
