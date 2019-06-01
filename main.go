package main

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/fs"
)

func main() {
	workDir := fs.NewTerraformFolder("examples/default_config")
	fmt.Println("Work in progress")

	fmt.Printf("We will work on that folder: %s\n", workDir.Path)
}
