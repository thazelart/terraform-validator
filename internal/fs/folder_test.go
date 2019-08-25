package fs_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/fs"
	"testing"
)

func TestListTerraformFiles(t *testing.T) {
	testPath := "testdata"

	expectedFileFlist := []string{
		"testdata/main.tf",
		"testdata/outputs.tf",
	}
	resultFileList := fs.ListTerraformFiles(testPath)
	if diff := cmp.Diff(expectedFileFlist, resultFileList); diff != "" {
		t.Errorf("ListTerraformFiles() mismatch (-want +got):\n%s", diff)
	}
}

func TestNewTerraformFolder(t *testing.T) {
	testPath := "testdata"
	// Create the expected resultFileList
	mainFile := fs.NewFile("testdata/main.tf")
	outputsFile := fs.NewFile("testdata/outputs.tf")
	expectedResult := fs.Folder{Path: testPath, Content: []fs.File{mainFile, outputsFile}}

	testResult := fs.NewTerraformFolder(testPath)

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("NewTerraformFolder() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetSubFolderList(t *testing.T) {
	expectedResult := []string{"testdata/modules"}

	testResult := fs.GetSubFolderList("testdata")

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetSubFolderList() mismatch (-want +got):\n%s", diff)
	}
}
