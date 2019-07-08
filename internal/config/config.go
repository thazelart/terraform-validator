package config

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/internal/utils"
	"gopkg.in/yaml.v3"
	"path"
	"strconv"
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
		"providers.tf": FileConfig{
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
	BlockPatternName:       "^[a-z0-9_]*$",
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
	EnsureTerraformVersion bool
	EnsureProvidersVersion bool
	EnsureReadmeUpdated    bool
	BlockPatternName       string
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

// UnmarshalYAML is a custom yaml unmarshaller for TerraformConfig
func (c *TerraformConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var customO struct {
		Files                  map[string]FileConfig `yaml:"files"`
		EnsureTerraformVersion string                `yaml:"ensure_terraform_version"`
		EnsureProvidersVersion string                `yaml:"ensure_providers_version"`
		EnsureReadmeUpdated    string                `yaml:"ensure_readme_updated"`
		BlockPatternName       string                `yaml:"block_pattern_name"`
	}
	if err := unmarshal(&customO); err != nil {
		return err
	}

	if len(customO.Files) != 0 {
		c.Files = customO.Files
	} else {
		c.Files = DefaultTerraformConfig.Files
	}

	if customO.EnsureTerraformVersion != "" {
		typedEnsureTerraformVersion, err := strconv.ParseBool(customO.EnsureTerraformVersion)
		utils.EnsureOrFatal(err)
		c.EnsureTerraformVersion = typedEnsureTerraformVersion
	} else {
		c.EnsureTerraformVersion = DefaultTerraformConfig.EnsureTerraformVersion
	}

	if customO.EnsureProvidersVersion != "" {
		typedEnsureProvidersVersion, err := strconv.ParseBool(customO.EnsureProvidersVersion)
		utils.EnsureOrFatal(err)
		c.EnsureProvidersVersion = typedEnsureProvidersVersion
	} else {
		c.EnsureProvidersVersion = DefaultTerraformConfig.EnsureProvidersVersion
	}

	if customO.EnsureReadmeUpdated != "" {
		typedEnsureReadmeUpdated, err := strconv.ParseBool(customO.EnsureReadmeUpdated)
		utils.EnsureOrFatal(err)
		c.EnsureReadmeUpdated = typedEnsureReadmeUpdated
	} else {
		c.EnsureReadmeUpdated = DefaultTerraformConfig.EnsureReadmeUpdated
	}

	if customO.BlockPatternName != "" {
		c.BlockPatternName = customO.BlockPatternName
	} else {
		c.BlockPatternName = DefaultTerraformConfig.BlockPatternName
	}

	return nil
}

// GetTerraformConfig get the terraform-validator config. If terraform-validator.yaml
// exists it merge the default and the custom config
func GetTerraformConfig(workDir fs.Folder) TerraformConfig {
	customConfigFile := path.Join(workDir.Path, ".terraform-validator.yaml")

	if !utils.FileExists(customConfigFile) {
		fmt.Println("INFO: using default configuration")
		return DefaultTerraformConfig
	}
	fmt.Println("INFO: using custom configuration")
	tempFile := fs.NewFile(customConfigFile)

	var customConfig TerraformConfig
	err := yaml.Unmarshal(tempFile.Content, &customConfig)
	utils.EnsureOrFatal(err)

	return customConfig
}

// GenerateGlobalConfig generates the terraform-validator global config.
// It takes the WorkDir needed informations and the TerraformConfig (default or custom)
func GenerateGlobalConfig(version string) GlobalConfig {
	// get folder information
	workDir := ParseArgs(version)
	workFolder := fs.NewTerraformFolder(workDir)

	// get config
	conf := GetTerraformConfig(workFolder)

	_, ok := conf.Files["default"]
	utils.OkOrFatal(ok, "FATAL Config.Files must contains at leat \"default\" !")

	return GlobalConfig{WorkDir: workFolder, TerraformConfig: conf}
}

// GetAuthorizedBlocks gets you the authorized blocks for the given filename.
// If the filename is not configure it gets you the dfault configuration.
// If their is no default either, return you an error.
func (globalConfig GlobalConfig) GetAuthorizedBlocks(filename string) ([]string, error) {
	_, ok := globalConfig.TerraformConfig.Files[filename]
	if ok {
		return globalConfig.TerraformConfig.Files[filename].AuthorizedBlocks, nil
	} else {
		_, ok := globalConfig.TerraformConfig.Files["default"]
		if ok {
			return globalConfig.TerraformConfig.Files["default"].AuthorizedBlocks, nil
		} else {
			return nil, fmt.Errorf("  cannot check authorized blocks, their is no file configuration for %s nor default", filename)
		}
	}
}

// GetMandatoryFiles get the mandatory file list from the globalConfig
func (globalConfig GlobalConfig) GetMandatoryFiles() []string {
	var mandatoryFiles []string

	for filename, fileInfos := range globalConfig.TerraformConfig.Files {
		if filename == "default" {
			continue
		}
		if fileInfos.Mandatory {
			mandatoryFiles = append(mandatoryFiles, filename)
		}
	}

	return mandatoryFiles
}

// GetFileNameList get the list of filename present in the working directory
func (globalConfig GlobalConfig) GetFileNameList() []string {
	var filesList []string

	for _, file := range globalConfig.WorkDir.Content {
		filesList = append(filesList, file.GetFilename())
	}

	return filesList
}
