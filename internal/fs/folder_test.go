package fs_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/fs"
	"testing"
)

func TestListTerraformFiles(t *testing.T) {
	testPath := "../../testdata/ok_default_config"

	expectedFileFlist := []string{
		"../../testdata/ok_default_config/backend.tf",
		"../../testdata/ok_default_config/main.tf",
		"../../testdata/ok_default_config/outputs.tf",
		"../../testdata/ok_default_config/providers.tf",
		"../../testdata/ok_default_config/variables.tf",
	}
	resultFileList := fs.ListTerraformFiles(testPath)
	if diff := cmp.Diff(expectedFileFlist, resultFileList); diff != "" {
		t.Errorf("ListTerraformFiles() mismatch (-want +got):\n%s", diff)
	}
}

func TestNewTerraformFolder(t *testing.T) {
	testPath := "../../testdata/ok_custom_config"
	// Create the expected resultFileList
	mainFile := fs.NewFile("../../testdata/ok_custom_config/main.tf")
	providerFile := fs.NewFile("../../testdata/ok_custom_config/provider.tf")
	expectedResult := fs.Folder{Path: testPath, Content: []fs.File{mainFile, providerFile}}

	testResult := fs.NewTerraformFolder(testPath)

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("NewTerraformFolder() mismatch (-want +got):\n%s", diff)
	}
}
