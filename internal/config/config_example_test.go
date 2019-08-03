package config_test

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/config"
)

func ExampleParseArgs() {
	workDir := config.ParseArgs("version")
	fmt.Printf("The path given in argument is: %s", workDir)
	// workDir is the path given by the os.Args
}

func ExampleGetTerraformConfig() {
	// get folder information
	workDir := config.ParseArgs("dev")

	// get config
	terraformConf := config.GetTerraformConfig(workDir)

	fmt.Printf("The terraform config': %+v", terraformConf)
}

func ExampleGenerateGlobalConfig() {
	version := "dev"
	globalConfig := config.GenerateGlobalConfig(version)

	fmt.Printf("The global config': %+v", globalConfig)
}
