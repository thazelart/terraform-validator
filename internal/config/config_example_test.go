package config_test

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/fs"
)

func ExampleParseArgs() {
	workDir := config.ParseArgs("version")
	fmt.Printf("The path given in argument is: %s", workDir)
	// workDir is the path given by the os.Args
}

func ExampleGetTerraformConfig() {
	// get folder information
	workDir := config.ParseArgs("dev")
	workFolder := fs.NewTerraformFolder(workDir)

	// get config
	terraformConf := config.GetTerraformConfig(workFolder)

	fmt.Printf("The terraform config': %+v", terraformConf)
}

func ExampleGenerateGlobalConfig() {
	version := "dev"
	globalConfig := config.GenerateGlobalConfig(version)

	fmt.Printf("The global config': %+v", globalConfig)
}
