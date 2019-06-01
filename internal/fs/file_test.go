package fs_test

import (
	"github.com/thazelart/terraform-validator/internal/fs"
	"reflect"
	"testing"
)

func TestNewFile(t *testing.T) {
	filePath := "../../examples/default_config/main.tf"
	fileContent := []byte("/* here would be a part of your terraform code */\n")

	expectedResult := fs.File{Path: filePath, Content: fileContent}
	testResult := fs.NewFile(filePath)

	if !reflect.DeepEqual(testResult, expectedResult) {
		t.Errorf("Got: %v, wanted: %v", testResult, expectedResult)
	}
}

func TestFileEqual(t *testing.T) {
	file1 := fs.NewFile("file_test.go")
	file2 := fs.NewFile("file_example_test.go")

	result1 := file1.FileEqual(file1)
	result2 := file1.FileEqual(file2)

	// ensure result1 is true (equal)
	if !result1 {
		t.Errorf("Got: %v, wanted: true", result1)
	}
	// ensure result2 is false (not equal)
	if result2 {
		t.Errorf("Got: %v, wanted: false", result2)
	}
}
