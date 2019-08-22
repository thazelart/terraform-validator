package main

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/tfv"
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

	// Get the configuration
	globalConfig := config.GenerateGlobalConfig(version)

	ok := tfv.MainChecks(globalConfig)
	if !ok {
		exitCode = 1
	}
}
