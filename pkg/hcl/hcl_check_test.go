package hcl_test

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/pkg/hcl"
	"testing"
)

func TestVerifyBlockNames(t *testing.T) {
	var testResult []error
	var globalConfig config.GlobalConfig

	tfParsedContent := hcl.TerraformFileParsedContent{
		"terraform": []string{"required_version"},
		"variable":  []string{"one_map", "a_list"},
		"output":    []string{"my_ip", "my_name"},
		"resource":  []string{"bucket_1", "bucket42"},
		"locals":    []string{"package_name", "creator"},
		"data":      []string{"centos7"},
		"module":    []string{"module_instance_name"},
		"provider":  []string{"google", "github"},
	}

	// First case, no error
	var expectedResult []error
	globalConfig.TerraformConfig.BlockPatternName = "^[a-z0-9_]*$"

	tfParsedContent.VerifyBlockNames(globalConfig, &testResult)
	if diff := cmp.Diff(len(expectedResult), len(testResult)); diff != "" {
		t.Errorf("VerifyBlockNames mismatch (-want +got):\n%s", diff)
	}

	// Second case, no numbers => 3 errors
	globalConfig.TerraformConfig.BlockPatternName = "^[a-z_]*$"
	expectedResult = []error{
		errors.New("  resource block \"bucket_1\" does not match \"^[a-z_]*$\""),
		errors.New("  resource block \"bucket42\" does not match \"^[a-z_]*$\""),
		errors.New("  data block \"centos7\" does not match \"^[a-z_]*$\""),
	}

	tfParsedContent.VerifyBlockNames(globalConfig, &testResult)
	if diff := cmp.Diff(len(expectedResult), len(testResult)); diff != "" {
		t.Errorf("VerifyBlockNames mismatch (-want +got):\n%s", diff)
	}
}

func TestVerifyBlocksInFiles(t *testing.T) {
	var testResult []error
	var testFile fs.File
	var globalConfig config.GlobalConfig
	globalConfig.TerraformConfig = config.DefaultTerraformConfig

	tfParsedContent := hcl.TerraformFileParsedContent{
		"variable": []string{"one_map", "a_list"},
	}

	// First with known filename
	var expectedResult []error
	testFile.Path = "/path/variables.tf"
	tfParsedContent.VerifyBlocksInFiles(globalConfig, testFile, &testResult)

	if diff := cmp.Diff(len(expectedResult), len(testResult)); diff != "" {
		t.Errorf("VerifyBlocksInFiles(variables.tf) mismatch (-want +got):\n%s", diff)
	}

	// test2 with unknown filename
	expectedResult = append(expectedResult, fmt.Errorf("  variables blocks are not authorized"))
	testFile.Path = "/path/main.tf"
	tfParsedContent.VerifyBlocksInFiles(globalConfig, testFile, &testResult)

	if diff := cmp.Diff(len(expectedResult), len(testResult)); diff != "" {
		t.Errorf("VerifyBlocksInFiles(main.tf) mismatch (-want +got):\n%s", diff)
	}

	// test3 with unknown filename and no default
	delete(globalConfig.TerraformConfig.Files, "default")
	expectedResult = append(expectedResult,
		fmt.Errorf("  cannot check authorized blocks, their is no file configuration for foo.tf nor default"))
	testFile.Path = "/path/foo.tf"
	tfParsedContent.VerifyBlocksInFiles(globalConfig, testFile, &testResult)

	if diff := cmp.Diff(len(expectedResult), len(testResult)); diff != "" {
		t.Errorf("VerifyBlocksInFiles(foo.tf) mismatch (-want +got):\n%s", diff)
	}
}
