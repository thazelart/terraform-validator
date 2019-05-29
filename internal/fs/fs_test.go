package fs_test

import (
	"github.com/thazelart/terraform-validator/internal/fs"
	"os"
	"path"
	"reflect"
	"testing"
	"errors"
)

func TestEnsureOrFatal(t *testing.T) {
	// After this test, replace the original fatal function
	origLogFatal := fs.LogFatal
	defer func() { fs.LogFatal = origLogFatal }()

	errs := []string{}
	fs.LogFatal = func(args ...interface{}) {
		errs = append(errs, "failed")
	}

	fs.EnsureOrFatal(errors.New("lol"))

	if len(errs) != 1 {
		t.Errorf("Got: %d error, wanted: 1", len(errs))
	}
}

func TestListTerraformFiles(t *testing.T) {
	testPath := ".."
	testFileName := "sample.tf"
	testFilePath := path.Join(testPath, testFileName)

	// Prepare the test by adding a .tf file
	var testFile, _ = os.Create(testFilePath)
	testFile.Close()
	defer os.Remove(testFilePath)

	expectedFileFlist := []string{testFilePath}
	resultFileList := fs.File{testPath}.ListTerraformFiles()
	if !reflect.DeepEqual(expectedFileFlist, resultFileList) {
		t.Errorf("Got: %v, wanted: %v", expectedFileFlist, resultFileList)
	}
}

func TestFileEqual(t *testing.T) {
	file1 := fs.File{"fs_test.go"}
	file2 := fs.File{"fs_example_test.go"}

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
