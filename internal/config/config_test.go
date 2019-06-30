package config_test

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/internal/utils"
	"gopkg.in/yaml.v3"
	"os"
	"sort"
	"testing"
)

func TestParseArgs(t *testing.T) {
	expectedResult := "/tmp"
	os.Args = []string{"terraform-validator", expectedResult}

	testResult := config.ParseArgs("version")

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("ParseArgs() mismatch (-want +got):\n%s", diff)
	}
}

func TestUnmarshalYAML(t *testing.T) {
	// first test: using custom config from example
	expectedCustomResult := config.DefaultTerraformConfig
	expectedCustomResult.Files = map[string]config.FileConfig{
		"default": {
			Mandatory: false,
			AuthorizedBlocks: []string{
				"variable",
				"output",
				"provider",
				"terraform",
				"resource",
				"module",
				"data",
				"locals",
			},
		},
	}
	expectedCustomResult.EnsureProvidersVersion = false
	expectedCustomResult.EnsureReadmeUpdated = false

	customConfigFile := fs.NewFile("../../examples/custom_config/terraform-validator.yaml")
	var testCustomResult config.TerraformConfig
	err := yaml.Unmarshal(customConfigFile.Content, &testCustomResult)
	utils.EnsureOrFatal(err)

	if diff := cmp.Diff(expectedCustomResult, testCustomResult); diff != "" {
		t.Errorf("GetCustomConfig(custom) mismatch (-want +got):\n%s", diff)
	}

	// second test with the others possibility of custmization
	expectedCustomResult2 := config.DefaultTerraformConfig
	expectedCustomResult2.EnsureTerraformVersion = false
	expectedCustomResult2.BlockPatternName = "foo"

	customConfigContent := []byte("ensure_terraform_version: false\nblock_pattern_name: 'foo'")
	var testCustomResult2 config.TerraformConfig
	err = yaml.Unmarshal(customConfigContent, &testCustomResult2)
	utils.EnsureOrFatal(err)

	if diff := cmp.Diff(expectedCustomResult2, testCustomResult2); diff != "" {
		t.Errorf("GetCustomConfig(custom) mismatch (-want +got):\n%s", diff)
	}
}

func TestGetTerraformConfig(t *testing.T) {
	// First test case: no custom config
	defaultFolder := fs.NewTerraformFolder("../../examples/default_config/")
	expectedDefaultResult := config.DefaultTerraformConfig
	testDefaultResult := config.GetTerraformConfig(defaultFolder)

	if diff := cmp.Diff(expectedDefaultResult, testDefaultResult); diff != "" {
		t.Errorf("GetCustomConfig(default) mismatch (-want +got):\n%s", diff)
	}

	// Second test case: with custom config
	expectedCustomResult := config.DefaultTerraformConfig
	expectedCustomResult.Files = map[string]config.FileConfig{
		"default": {
			Mandatory: false,
			AuthorizedBlocks: []string{
				"variable",
				"output",
				"provider",
				"terraform",
				"resource",
				"module",
				"data",
				"locals",
			},
		},
	}
	expectedCustomResult.EnsureProvidersVersion = false
	expectedCustomResult.EnsureReadmeUpdated = false

	customFolder := fs.NewTerraformFolder("../../examples/custom_config/")
	testCustomResult := config.GetTerraformConfig(customFolder)

	if diff := cmp.Diff(expectedCustomResult, testCustomResult); diff != "" {
		t.Errorf("GetCustomConfig(custom) mismatch (-want +got):\n%s", diff)
	}
}

func TestGenerateGlobalConfig(t *testing.T) {
	// First test case: no custom config
	os.Args = []string{"terraform-validator", "../../examples/default_config/"}

	defaultConfig := config.DefaultTerraformConfig
	defaultFolder := fs.NewTerraformFolder("../../examples/default_config/")
	expectedDefaultConfig := config.GlobalConfig{WorkDir: defaultFolder, TerraformConfig: defaultConfig}
	testDefaultResult := config.GenerateGlobalConfig("dev")

	if diff := cmp.Diff(expectedDefaultConfig, testDefaultResult); diff != "" {
		t.Errorf("GetCustomConfig(default) mismatch (-want +got):\n%s", diff)
	}

	// Second test case: with custom config
	os.Args = []string{"terraform-validator", "../../examples/custom_config/"}

	customConfig := config.DefaultTerraformConfig
	customConfig.Files = map[string]config.FileConfig{
		"default": {
			AuthorizedBlocks: []string{
				"variable",
				"output",
				"provider",
				"terraform",
				"resource",
				"module",
				"data",
				"locals",
			},
		},
	}
	customConfig.EnsureProvidersVersion = false
	customConfig.EnsureReadmeUpdated = false
	customFolder := fs.NewTerraformFolder("../../examples/custom_config/")
	expectedCustomConfig := config.GlobalConfig{WorkDir: customFolder, TerraformConfig: customConfig}
	testCustomResult := config.GenerateGlobalConfig("dev")

	if diff := cmp.Diff(expectedCustomConfig, testCustomResult); diff != "" {
		t.Errorf("GetCustomConfig(custom) mismatch (-want +got):\n%s", diff)
	}
}

func TestGetAuthorizedBlocks(t *testing.T) {
	var testGC config.GlobalConfig
	testGC.TerraformConfig = config.DefaultTerraformConfig

	// test1 with known filename
	expectedResult := []string{"variable"}
	testResult, _ := testGC.GetAuthorizedBlocks("variables.tf")

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetAuthorizedBlocks(knownFile) mismatch (-want +got):\n%s", diff)
	}

	// test2 with unknown filename
	expectedResult = []string{"resource", "module", "data", "locals"}
	testResult, _ = testGC.GetAuthorizedBlocks("foo.tf")

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetAuthorizedBlocks(unknownFile) mismatch (-want +got):\n%s", diff)
	}

	// test3 with unknown filename and no default
	delete(testGC.TerraformConfig.Files, "default")
	expectedErrorResult := errors.New("  cannot check authorized blocks, their is no file configuration for foo.tf nor default")
	_, testErrorResult := testGC.GetAuthorizedBlocks("foo.tf")

	if expectedErrorResult == nil {
		t.Errorf("GetAuthorizedBlocks(unknownFileNoDefault) mismatch (-want +got):\n- %#v\n+ %+v",
			expectedErrorResult, testErrorResult)
	}
}

func TestGetMandatoryFiles(t *testing.T) {
	var testGC config.GlobalConfig
	testGC.TerraformConfig = config.DefaultTerraformConfig

	// set default as mandatory, it must not be in expectedResult
	tmpDefault := testGC.TerraformConfig.Files["default"]
	tmpDefault.Mandatory = true
	testGC.TerraformConfig.Files["default"] = tmpDefault

	expectedResult := []string{"backend.tf", "main.tf", "outputs.tf", "provider.tf", "variables.tf"}

	testResult := testGC.GetMandatoryFiles()
	sort.Strings(testResult)

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetMandatoryFiles() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetFileNameList(t *testing.T) {
	var testGC config.GlobalConfig

	foo := fs.File{Path: "/path/to/foo.tf", Content: []byte{}}
	bar := fs.File{Path: "/path/to/bar.tf", Content: []byte{}}

	var fileList []fs.File
	fileList = append(fileList, foo)
	fileList = append(fileList, bar)

	testGC.WorkDir = fs.Folder{Path: "/path/to", Content: fileList}

	expectedResult := []string{"bar.tf", "foo.tf"}

	testResult := testGC.GetFileNameList()
	sort.Strings(testResult)

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetMandatoryFiles() mismatch (-want +got):\n%s", diff)
	}
}
