package config

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/pkg/utils"
	"gopkg.in/yaml.v3"
	"path"
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

// FileConfig is the configuration for a .tf file
// AuthorizedBlocks is the list of authorized blocks in that file (for example
// "variables", "output"...).
// Mandatory is a boolean that define if the file is mandatory or not.
type FileConfig struct {
	AuthorizedBlocks []string `yaml:"authorized_blocks"`
	Mandatory        bool     `yaml:"mandatory"`
}

// FolderConfigClass is a type that define a class of config for a folder.
type FolderConfigClass struct {
	Files                  map[string]FileConfig `yaml:"files"`
	EnsureTerraformVersion bool                  `yaml:"ensure_terraform_version"`
	EnsureProvidersVersion bool                  `yaml:"ensure_providers_version"`
	BlockPatternName       string                `yaml:"block_pattern_name"`
}

// TfvConfig is the full configuration of terraform validator
// CurrentFolderClass is the current folder applied class
// Classes is a map of FolderConfigClass
type TfvConfig struct {
	CurrentFolderClass string
	Classes            map[string]FolderConfigClass
}

// DefaultFolderConfigClass return you the default FolderConfigClass
func DefaultFolderConfigClass() FolderConfigClass {
	return FolderConfigClass{
		Files: map[string]FileConfig{
			"main.tf": {
				Mandatory:        true,
				AuthorizedBlocks: nil,
			},
			"variables.tf": {
				Mandatory:        true,
				AuthorizedBlocks: []string{"variable"},
			},
			"outputs.tf": {
				Mandatory:        true,
				AuthorizedBlocks: []string{"output"},
			},
			"providers.tf": {
				Mandatory:        true,
				AuthorizedBlocks: []string{"provider"},
			},
			"backend.tf": {
				Mandatory:        true,
				AuthorizedBlocks: []string{"terraform"},
			},
			"default": {
				Mandatory:        false,
				AuthorizedBlocks: []string{"resource", "module", "data", "locals"},
			},
		},
		EnsureTerraformVersion: false,
		EnsureProvidersVersion: false,
		BlockPatternName:       "^[a-z0-9_]*$",
	}
}

// DefaultTfvConfig returns you the default TfvConfig
func DefaultTfvConfig() TfvConfig {
	return TfvConfig{
		CurrentFolderClass: "default",
		Classes: map[string]FolderConfigClass{
			"default": DefaultFolderConfigClass(),
		},
	}
}

// ParseArgs get the path given as os argument
func ParseArgs(version string) string {
	args, err := docopt.ParseArgs(usage, nil, version)
	utils.EnsureOrFatal(err)

	return args["<path>"].(string)
}

// UnmarshalYAML is a custom yaml unmarshaller for TerraformConfig
func (c *TfvConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var customO struct {
		CurrentFolderClass string                       `yaml:"current_folder_class"`
		Classes            map[string]FolderConfigClass `yaml:"classes"`
	}
	err := unmarshal(&customO)
	utils.EnsureOrFatal(err)

	if customO.CurrentFolderClass != "" {
		c.CurrentFolderClass = customO.CurrentFolderClass
	} else {
		c.CurrentFolderClass = "default"
	}

	if c.Classes == nil {
		c.Classes = make(map[string]FolderConfigClass)
		c.Classes["default"] = DefaultFolderConfigClass()
	}
	for key, class := range customO.Classes {
		if len(class.Files) == 0 {
			class.Files = DefaultFolderConfigClass().Files
		}

		if class.BlockPatternName == "" {
			class.BlockPatternName = DefaultFolderConfigClass().BlockPatternName
		}

		c.Classes[key] = class
	}

	return nil
}

// GetTerraformConfig get the terraform-validator config. If .terraform-validator.yaml
// exists it merge the default and the custom config
func (c TfvConfig) GetTerraformConfig(workDir string) TfvConfig {
	customConfigFile := path.Join(workDir, ".terraform-validator.yaml")

	if utils.FileExists(customConfigFile) {
		tempFile := fs.NewFile(customConfigFile)
		err := yaml.Unmarshal(tempFile.Content, &c)
		utils.EnsureOrFatal(err)
	}

	fmt.Printf("INFO: using %s configuration in %s\n", c.CurrentFolderClass, workDir)
	return c
}

// GetFolderConfigClass get the applied FolderConfigClass
func (c TfvConfig) GetFolderConfigClass() FolderConfigClass {
	folderConfigClass, ok := c.Classes[c.CurrentFolderClass]

	utils.OkOrFatal(ok,
		fmt.Sprintf("FATAL: terraform-validation configuration does not contain %s class",
			c.CurrentFolderClass,
		),
	)

	return folderConfigClass
}

// GetAuthorizedBlocks gets you the authorized blocks for the given filename.
// If the filename is not configure it gets you the dfault configuration.
// If their is no default either, return you an error.
func (folderConfigClass FolderConfigClass) GetAuthorizedBlocks(filename string) ([]string, error) {
	file, ok := folderConfigClass.Files[filename]
	if ok {
		return file.AuthorizedBlocks, nil
	}

	file, ok = folderConfigClass.Files["default"]
	if ok {
		return file.AuthorizedBlocks, nil
	}

	return []string{}, fmt.Errorf("  cannot check authorized blocks, their is no file configuration for %s nor default", filename)
}

// GetMandatoryFiles get the mandatory file list from the globalConfig
func (folderConfigClass FolderConfigClass) GetMandatoryFiles() []string {
	var mandatoryFiles []string

	for filename, fileInfos := range folderConfigClass.Files {
		if filename == "default" {
			continue
		}
		if fileInfos.Mandatory {
			mandatoryFiles = append(mandatoryFiles, filename)
		}
	}

	return mandatoryFiles
}
