package config

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/internal/utils"
	"gopkg.in/yaml.v3"
	"path"
	"strings"
)

// usage is the terraform-validator help
const usage = `
  Usage:
    terraform-validator <path>
    terraform-validator -h | --help
  Examples:
    # Check configuration in /tmp/my-module
    $ terraform-validator /tmp/my-module

    # Check configuration here
    $ terraform-validator .

  Options:
    -h, --help        show help information
    -v, --version     show version and exit
`

// DefaultTerraformConfig is the default TerraformConfig configuration
var DefaultTerraformConfig = TerraformConfig{
	Files: map[string]FileConfig{
		"main.tf": FileConfig{
			Mandatory:        true,
			AuthorizedBlocks: nil,
		},
		"variables.tf": FileConfig{
			Mandatory:        true,
			AuthorizedBlocks: []string{"variable"},
		},
		"outputs.tf": FileConfig{
			Mandatory:        true,
			AuthorizedBlocks: []string{"output"},
		},
		"provider.tf": FileConfig{
			Mandatory:        true,
			AuthorizedBlocks: []string{"provider"},
		},
		"backend.tf": FileConfig{
			Mandatory:        true,
			AuthorizedBlocks: []string{"terraform"},
		},
		"default": FileConfig{
			Mandatory:        false,
			AuthorizedBlocks: []string{"resource", "module", "data", "locals"},
		},
	},
	EnsureTerraformVersion: true,
	EnsureProvidersVersion: true,
	EnsureReadmeUpdated:    true,
	BlockPatternName:       "^[a-z_]*$",
}

// FileConfig is the configuration for a .tf file
// AuthorizedBlocks is the list of authorized blocks in that file (for example
// "variables", "output"...).
// Mandatory is a boolean that define if the file is mandatory or not.
type FileConfig struct {
	AuthorizedBlocks []string `yaml:"authorized_blocks"`
	Mandatory        bool     `yaml:"mandatory"`
}

// TerraformConfig is the full configuration of terraform validator
// Files is the map of .tf files defines with the FileConfig type.
// EnsureTerraformVersion define if the terraform version has to be set or not.
// EnsureProvidersVersion define if the providers versions has to be set or not.
// EnsureReadmeUpdated define if we care or not if the documentation has been updated.
// BlockPatternName is the pattern that must match all the terraform blocks name.
type TerraformConfig struct {
	Files                  map[string]FileConfig
	EnsureTerraformVersion bool   `yaml:"ensure_terraform_version"`
	EnsureProvidersVersion bool   `yaml:"ensure_providers_version"`
	EnsureReadmeUpdated    bool   `yaml:"ensure_readme_updated"`
	BlockPatternName       string `yaml:"block_pattern_name"`
}

// GlobalConfig is the global terraform validator config
type GlobalConfig struct {
	WorkDir         fs.Folder
	TerraformConfig TerraformConfig
}

// ParseArgs get the path given as os argument
func ParseArgs(version string) string {
	args, err := docopt.ParseArgs(usage, nil, version)
	utils.EnsureOrFatal(err)

	return args["<path>"].(string)
}

// ReadYaml take a TerraformConfig and a path to the yaml file and return the
// fulfilled TerraformConfig.
// If the given TerraformConfig is empty then is take the full yaml from the
// file in parameter. otherwise it merge them.
func (terraformConfig TerraformConfig) ReadYaml(pathFile string) TerraformConfig {
	tempFile := fs.NewFile(pathFile)
	// TODO: à renommer mergeCustomAndDefault
	// pour si tempContent terraformConfig.Files = content.Files
	if strings.Contains(string(tempFile.Content), "files:") {
		terraformConfig.Files = nil
	}

	err := yaml.Unmarshal(tempFile.Content, &terraformConfig)
	utils.EnsureOrFatal(err)

	return terraformConfig
}

// NewTerraformConfig return TerraformConfig with the default value (DefaultTerraformConfig)
func NewTerraformConfig() TerraformConfig {
	return DefaultTerraformConfig
}

// GetCustomConfig take a TerraformConfig (generally the default one) and get the custom
// config if set. If the custom config is set it merge the default and the custom configurations.
func (terraformConfig TerraformConfig) GetCustomConfig(workDir fs.Folder) TerraformConfig {
	customConfigFile := path.Join(workDir.Path, "terraform-validator.yaml")
	if !utils.FileExists(customConfigFile) {
		fmt.Printf("Working on %s with default configuration\n", workDir.Path)
		return terraformConfig
	}
	fmt.Printf("Working on %s with custom configuration\n", workDir.Path)
	terraformConfig = terraformConfig.ReadYaml(customConfigFile)
	return terraformConfig
}

// GenerateConfig generates the terraform-validator global config.
// It takes the WorkDir needed informations and the TerraformConfig (default or custom)
func GenerateGlobalConfig(version string) GlobalConfig {
	// get folder information
	workDir := ParseArgs(version)
	workFolder := fs.NewTerraformFolder(workDir)

	// get config
	conf := NewTerraformConfig()
	conf = conf.GetCustomConfig(workFolder)

	_, ok := conf.Files["default"]
	utils.OkOrFatal(ok, "FATAL Config.Files must contains at leat \"default\" !")

	return GlobalConfig{WorkDir: workFolder, TerraformConfig: conf}
}
