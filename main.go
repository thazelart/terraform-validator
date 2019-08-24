package main

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/checks"
	"os"
)

var (
	version string
)

func init() {
	if version == "" {
		version = "dev"
	}
}

func main() {
	exitCode := 0
	defer func() {
		if exitCode == 0 {
			fmt.Println("INFO: terraform-validator ran successfully")
		}
		os.Exit(exitCode)
	}()

	rootDir := config.ParseArgs(version)

	ok := checks.MainChecks(config.DefaultTfvConfig(), rootDir)
	if !ok {
		exitCode = 1
	}
}
