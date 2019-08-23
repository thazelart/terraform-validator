package config_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/pkg/utils"
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
	// case1 test: using custom config from example
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
	expectedCustomResult.EnsureTerraformVersion = false
	expectedCustomResult.EnsureReadmeUpdated = false

	customConfigFile := fs.NewFile("testdata/case1/.terraform-validator.yaml")
	var testCustomResult config.TerraformConfig
	err := yaml.Unmarshal(customConfigFile.Content, &testCustomResult)
	utils.EnsureOrFatal(err)

	if diff := cmp.Diff(expectedCustomResult, testCustomResult); diff != "" {
		t.Errorf("TestUnmarshalYAML(custom) mismatch (-want +got):\n%s", diff)
	}

	// case2 test with the others possibility of custmization
	expectedCustomResult = config.DefaultTerraformConfig
	expectedCustomResult.EnsureProvidersVersion = false
	expectedCustomResult.BlockPatternName = "foo"

	customConfigFile = fs.NewFile("testdata/case2/.terraform-validator.yaml")
	var testCustomResult2 config.TerraformConfig
	err = yaml.Unmarshal(customConfigFile.Content, &testCustomResult2)
	utils.EnsureOrFatal(err)

	if diff := cmp.Diff(expectedCustomResult, testCustomResult2); diff != "" {
		t.Errorf("TestUnmarshalYAML(custom) mismatch (-want +got):\n%s", diff)
	}
}

func TestGetTerraformConfig(t *testing.T) {
	// case1 test case: with custom config
	WorkDir := "testdata/case1"
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
	expectedCustomResult.EnsureTerraformVersion = false
	expectedCustomResult.EnsureReadmeUpdated = false

	testCustomResult := config.GetTerraformConfig(WorkDir)

	if diff := cmp.Diff(expectedCustomResult, testCustomResult); diff != "" {
		t.Errorf("GetCustomConfig(custom) mismatch (-want +got):\n%s", diff)
	}

	// case2 test case: with custom config
	WorkDir = "testdata/case2"
	expectedCustomResult = config.DefaultTerraformConfig
	expectedCustomResult.EnsureProvidersVersion = false
	expectedCustomResult.BlockPatternName = "foo"

	testCustomResult = config.GetTerraformConfig(WorkDir)

	if diff := cmp.Diff(expectedCustomResult, testCustomResult); diff != "" {
		t.Errorf("GetCustomConfig(custom) mismatch (-want +got):\n%s", diff)
	}

	// case3 test case: no custom config
	WorkDir = "testdata/case3"
	expectedDefaultResult := config.DefaultTerraformConfig
	testDefaultResult := config.GetTerraformConfig(WorkDir)

	if diff := cmp.Diff(expectedDefaultResult, testDefaultResult); diff != "" {
		t.Errorf("GetCustomConfig(default) mismatch (-want +got):\n%s", diff)
	}
}

func TestGenerateGlobalConfig(t *testing.T) {
	// Case1: with custom config
	workDir := "../../testdata/ok_custom_config/"
	os.Args = []string{"terraform-validator", workDir}

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
	customConfig.EnsureTerraformVersion = false
	customConfig.EnsureReadmeUpdated = false
	expectedCustomConfig := config.GlobalConfig{WorkDir: workDir, TerraformConfig: customConfig}
	testCustomResult := config.GenerateGlobalConfig("dev")

	if diff := cmp.Diff(expectedCustomConfig, testCustomResult); diff != "" {
		t.Errorf("GetCustomConfig(custom) mismatch (-want +got):\n%s", diff)
	}

	// case3 test case: no custom config
	workDir = "testdata/case3"
	os.Args = []string{"terraform-validator", workDir}

	defaultConfig := config.DefaultTerraformConfig
	expectedDefaultConfig := config.GlobalConfig{WorkDir: workDir, TerraformConfig: defaultConfig}
	testDefaultResult := config.GenerateGlobalConfig("dev")

	if diff := cmp.Diff(expectedDefaultConfig, testDefaultResult); diff != "" {
		t.Errorf("GetCustomConfig(default) mismatch (-want +got):\n%s", diff)
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
	expectedResult = []string{}
	testResult, _ = testGC.GetAuthorizedBlocks("foo.tf")

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetAuthorizedBlocks(unknownFileNoDefault) mismatch (-want +got):\n%s", diff)
	}
}

func TestGetMandatoryFiles(t *testing.T) {
	var testGC config.GlobalConfig
	testGC.TerraformConfig = config.DefaultTerraformConfig

	// set default as mandatory, it must not be in expectedResult
	tmpDefault := testGC.TerraformConfig.Files["default"]
	tmpDefault.Mandatory = true
	testGC.TerraformConfig.Files["default"] = tmpDefault

	expectedResult := []string{"backend.tf", "main.tf", "outputs.tf", "providers.tf", "variables.tf"}

	testResult := testGC.GetMandatoryFiles()
	sort.Strings(testResult)

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetMandatoryFiles() mismatch (-want +got):\n%s", diff)
	}
}
