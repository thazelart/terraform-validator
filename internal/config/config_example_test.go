package config_test

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/fs"
	"testing"
)

func ExampleParseArgs(t *testing.T) {
	workDir := config.ParseArgs("version")
	fmt.Printf("The path given in argument is: %s", workDir)
	// workDir is the path given by the os.Args
}

func ExampleReadYaml(t *testing.T) {
	var tfConfig config.TerraformConfig
	tfConfig = tfConfig.ReadYaml("default_config.yaml")

	fmt.Printf("Terraform config from yaml file: %+v", tfConfig)
}

func ExampleNewTerraformConfig(t *testing.T) {
	var tfConfig config.TerraformConfig
	tfConfig = config.NewTerraformConfig()

	fmt.Printf("Default terraform config: %+v", tfConfig.Files["main.tf"])
}

func ExampleGetCustomConfig(t *testing.T) {
	customFolder := fs.NewTerraformFolder("../../examples/custom_config/")
	tfConfig := config.DefaultTerraformConfig.GetCustomConfig(customFolder)

	fmt.Printf("Custom terraform config: %+v", tfConfig)
	// return the defaultTerraformCOnfig modified by the custom config
}

func ExampleGenerateGlobalConfig(t *testing.T) {
	version := "dev"
	globalConfig := config.GenerateGlobalConfig(version)

	fmt.Printf("The global config': %+v", globalConfig)
}
