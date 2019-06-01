package fs_test

import (
	"github.com/thazelart/terraform-validator/internal/fs"
	"reflect"
	"testing"
)

func TestListTerraformFiles(t *testing.T) {
	testPath := "../../examples/default_config"

	expectedFileFlist := []string{"../../examples/default_config/main.tf", "../../examples/default_config/provider.tf"}
	resultFileList := fs.ListTerraformFiles(testPath)
	if !reflect.DeepEqual(expectedFileFlist, resultFileList) {
		t.Errorf("Got: %v, wanted: %v", resultFileList, expectedFileFlist)
	}
}

func TestNewTerraformFolder(t *testing.T) {
	testPath := "../../examples/default_config"
	// Create the expected resultFileList
	mainFile := fs.NewFile("../../examples/default_config/main.tf")
	providerFile := fs.NewFile("../../examples/default_config/provider.tf")
	expectedResult := fs.Folder{Path: testPath, Content: []fs.File{mainFile, providerFile}}

	testResult := fs.NewTerraformFolder(testPath)

	if !reflect.DeepEqual(testResult, expectedResult) {
		t.Errorf("Got: %v, wanted: %v", testResult, expectedResult)
	}
}
