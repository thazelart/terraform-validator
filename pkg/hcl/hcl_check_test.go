package hcl_test

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/config"
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
