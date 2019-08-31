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

// ConfigurationLayer is a type that define a layer of config for a folder.
type ConfigurationLayer struct {
	Files                     map[string]FileConfig `yaml:"files"`
	EnsureTerraformVersion    bool                  `yaml:"ensure_terraform_version"`
	EnsureProvidersVersion    bool                  `yaml:"ensure_providers_version"`
	EnsureVariablesDescrition bool                  `yaml:"ensure_variables_description"`
	EnsureOutputsDescrition   bool                  `yaml:"ensure_outputs_description"`
	BlockPatternName          string                `yaml:"block_pattern_name"`
}

// TfvConfig is the full configuration of terraform validator
// CurrentLayer is the current folder applied layer
// Layers is a map of ConfigurationLayer
type TfvConfig struct {
	CurrentLayer string
	Layers       map[string]ConfigurationLayer
}

// DefaultConfigurationLayer return you the default ConfigurationLayer
func DefaultConfigurationLayer() ConfigurationLayer {
	return ConfigurationLayer{
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
		CurrentLayer: "default",
		Layers: map[string]ConfigurationLayer{
			"default": DefaultConfigurationLayer(),
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
		CurrentLayer string                        `yaml:"current_layer"`
		Layers       map[string]ConfigurationLayer `yaml:"layers"`
	}
	err := unmarshal(&customO)
	utils.EnsureOrFatal(err)

	if customO.CurrentLayer != "" {
		c.CurrentLayer = customO.CurrentLayer
	} else {
		c.CurrentLayer = "default"
	}

	if c.Layers == nil {
		c.Layers = make(map[string]ConfigurationLayer)
		c.Layers["default"] = DefaultConfigurationLayer()
	}
	for key, layer := range customO.Layers {
		if len(layer.Files) == 0 {
			layer.Files = DefaultConfigurationLayer().Files
		}

		if layer.BlockPatternName == "" {
			layer.BlockPatternName = DefaultConfigurationLayer().BlockPatternName
		}

		c.Layers[key] = layer
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

	return c
}

// GetConfigurationLayer get the applied ConfigurationLayer
func (c TfvConfig) GetConfigurationLayer() ConfigurationLayer {
	configLayer, ok := c.Layers[c.CurrentLayer]

	utils.OkOrFatal(ok,
		fmt.Sprintf("FATAL: terraform-validation configuration does not contain %s layer",
			c.CurrentLayer,
		),
	)

	return configLayer
}

// GetAuthorizedBlocks gets you the authorized blocks for the given filename.
// If the filename is not configure it gets you the dfault configuration.
// If their is no default either, return you an error.
func (configLayer ConfigurationLayer) GetAuthorizedBlocks(filename string) ([]string, error) {
	file, ok := configLayer.Files[filename]
	if ok {
		return file.AuthorizedBlocks, nil
	}

	file, ok = configLayer.Files["default"]
	if ok {
		return file.AuthorizedBlocks, nil
	}

	return []string{}, fmt.Errorf("  cannot check authorized blocks, their is no file configuration for %s nor default", filename)
}

// GetMandatoryFiles get the mandatory file list from the globalConfig
func (configLayer ConfigurationLayer) GetMandatoryFiles() []string {
	var mandatoryFiles []string

	for filename, fileInfos := range configLayer.Files {
		if filename == "default" {
			continue
		}
		if fileInfos.Mandatory {
			mandatoryFiles = append(mandatoryFiles, filename)
		}
	}

	return mandatoryFiles
}
