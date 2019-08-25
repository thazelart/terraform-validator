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

func TestDefaultTfvConfig(t *testing.T) {
	expectedResult := config.TfvConfig{
		CurrentFolderClass: "default",
		Classes: map[string]config.FolderConfigClass{
			"default": config.DefaultFolderConfigClass(),
		},
	}

	testResult := config.DefaultTfvConfig()

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("() mismatch (-want +got):\n%s", diff)
	}
}

func TestDefaultFolderConfigClass(t *testing.T) {
	expectedResult := config.FolderConfigClass{
		Files: map[string]config.FileConfig{
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

	testResult := config.DefaultFolderConfigClass()

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("DefaultFolderConfigClass() mismatch (-want +got):\n%s", diff)
	}
}

func TestUnmarshalYAML(t *testing.T) {
	// case1 test: using custom config from example
	expectedCustomResult := config.DefaultTfvConfig()
	expectedCustomResult.CurrentFolderClass = "cust1"
	expectedCustomResult.Classes["cust1"] = config.FolderConfigClass{
		Files: map[string]config.FileConfig{
			"backend.tf":   {AuthorizedBlocks: []string{"terraform"}, Mandatory: true},
			"default":      {AuthorizedBlocks: []string{"resource", "module", "data", "locals"}},
			"main.tf":      {Mandatory: true},
			"outputs.tf":   {AuthorizedBlocks: []string{"output"}, Mandatory: true},
			"providers.tf": {AuthorizedBlocks: []string{"provider"}, Mandatory: true},
			"variables.tf": {AuthorizedBlocks: []string{"variable"}, Mandatory: true},
		},
		BlockPatternName: "^[a-z0-9_]*$",
	}
	expectedCustomResult.Classes["cust2"] = config.FolderConfigClass{
		Files: map[string]config.FileConfig{
			"backend.tf":   {AuthorizedBlocks: []string{"terraform"}, Mandatory: true},
			"default":      {AuthorizedBlocks: []string{"resource", "module", "data", "locals"}},
			"main.tf":      {Mandatory: true},
			"outputs.tf":   {AuthorizedBlocks: []string{"output"}, Mandatory: true},
			"providers.tf": {AuthorizedBlocks: []string{"provider"}, Mandatory: true},
			"variables.tf": {AuthorizedBlocks: []string{"variable"}, Mandatory: true},
		},
		BlockPatternName: `^[a-z0-9]*$"`,
	}

	customConfigFile := fs.NewFile("testdata/case1/.terraform-validator.yaml")
	testCustomResult := config.DefaultTfvConfig()
	err := yaml.Unmarshal(customConfigFile.Content, &testCustomResult)
	utils.EnsureOrFatal(err)

	if diff := cmp.Diff(testCustomResult, expectedCustomResult); diff != "" {
		t.Errorf("TestUnmarshalYAML(custom) mismatch (-want +got):\n%s", diff)
	}

	// case2 same test with empty TfvConfig
	customConfigFile = fs.NewFile("testdata/case2/.terraform-validator.yaml")
	expectedCustomResult.CurrentFolderClass = "default"
	var testCustomResult2 config.TfvConfig
	err = yaml.Unmarshal(customConfigFile.Content, &testCustomResult2)
	utils.EnsureOrFatal(err)

	if diff := cmp.Diff(testCustomResult2, expectedCustomResult); diff != "" {
		t.Errorf("TestUnmarshalYAML(custom) mismatch (-want +got):\n%s", diff)
	}
}

func TestGetTerraformConfig(t *testing.T) {
	// case1 test case: with custom config
	WorkDir := "testdata/case1"
	expectedResult := config.DefaultTfvConfig()
	expectedResult.CurrentFolderClass = "cust1"
	expectedResult.Classes["cust1"] = config.FolderConfigClass{
		Files: map[string]config.FileConfig{
			"backend.tf":   {AuthorizedBlocks: []string{"terraform"}, Mandatory: true},
			"default":      {AuthorizedBlocks: []string{"resource", "module", "data", "locals"}},
			"main.tf":      {Mandatory: true},
			"outputs.tf":   {AuthorizedBlocks: []string{"output"}, Mandatory: true},
			"providers.tf": {AuthorizedBlocks: []string{"provider"}, Mandatory: true},
			"variables.tf": {AuthorizedBlocks: []string{"variable"}, Mandatory: true},
		},
		BlockPatternName: "^[a-z0-9_]*$",
	}
	expectedResult.Classes["cust2"] = config.FolderConfigClass{
		Files: map[string]config.FileConfig{
			"backend.tf":   {AuthorizedBlocks: []string{"terraform"}, Mandatory: true},
			"default":      {AuthorizedBlocks: []string{"resource", "module", "data", "locals"}},
			"main.tf":      {Mandatory: true},
			"outputs.tf":   {AuthorizedBlocks: []string{"output"}, Mandatory: true},
			"providers.tf": {AuthorizedBlocks: []string{"provider"}, Mandatory: true},
			"variables.tf": {AuthorizedBlocks: []string{"variable"}, Mandatory: true},
		},
		BlockPatternName: `^[a-z0-9]*$"`,
	}

	testResult := config.DefaultTfvConfig().GetTerraformConfig(WorkDir)

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetCustomConfig(custom) mismatch (-want +got):\n%s", diff)
	}

	// case2 test case: no custom config
	WorkDir = "testdata/case3"
	expectedResult = config.DefaultTfvConfig()
	testResult = config.DefaultTfvConfig().GetTerraformConfig(WorkDir)

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetCustomConfig(default) mismatch (-want +got):\n%s", diff)
	}
}

func TestGetFolderConfigClass(t *testing.T) {
	testData := config.DefaultTfvConfig()
	expectedResult := config.DefaultFolderConfigClass()

	testResult := testData.GetFolderConfigClass()

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetFolderConfigClass() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetAuthorizedBlocks(t *testing.T) {
	testGC := config.DefaultTfvConfig()
	config := testGC.Classes[testGC.CurrentFolderClass]

	// test1 with known filename
	expectedResult := []string{"variable"}
	testResult, _ := config.GetAuthorizedBlocks("variables.tf")

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetAuthorizedBlocks(knownFile) mismatch (-want +got):\n%s", diff)
	}

	// test2 with unknown filename
	expectedResult = []string{"resource", "module", "data", "locals"}
	testResult, _ = config.GetAuthorizedBlocks("foo.tf")

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetAuthorizedBlocks(unknownFile) mismatch (-want +got):\n%s", diff)
	}

	// test3 with unknown filename and no default
	delete(config.Files, "default")
	expectedResult = []string{}
	testResult, _ = config.GetAuthorizedBlocks("foo.tf")

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetAuthorizedBlocks(unknownFileNoDefault) mismatch (-want +got):\n%s", diff)
	}
}

func TestGetMandatoryFiles(t *testing.T) {
	testGC := config.DefaultTfvConfig()
	conf := testGC.Classes[testGC.CurrentFolderClass]

	// set default as mandatory, it must not be in expectedResult
	tmpDefault := conf.Files["default"]
	tmpDefault.Mandatory = true
	conf.Files["default"] = tmpDefault

	expectedResult := []string{"backend.tf", "main.tf", "outputs.tf", "providers.tf", "variables.tf"}

	testResult := conf.GetMandatoryFiles()
	sort.Strings(testResult)

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetMandatoryFiles() mismatch (-want +got):\n%s", diff)
	}
}
