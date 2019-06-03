package fs_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/fs"
	"testing"
)

func TestListTerraformFiles(t *testing.T) {
	testPath := "../../examples/default_config"

	expectedFileFlist := []string{"../../examples/default_config/main.tf", "../../examples/default_config/provider.tf"}
	resultFileList := fs.ListTerraformFiles(testPath)
	if diff := cmp.Diff(expectedFileFlist, resultFileList); diff != "" {
		t.Errorf("ListTerraformFiles() mismatch (-want +got):\n%s", diff)
	}
}

func TestNewTerraformFolder(t *testing.T) {
	testPath := "../../examples/default_config"
	// Create the expected resultFileList
	mainFile := fs.NewFile("../../examples/default_config/main.tf")
	providerFile := fs.NewFile("../../examples/default_config/provider.tf")
	expectedResult := fs.Folder{Path: testPath, Content: []fs.File{mainFile, providerFile}}

	testResult := fs.NewTerraformFolder(testPath)

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("NewTerraformFolder() mismatch (-want +got):\n%s", diff)
	}
}
