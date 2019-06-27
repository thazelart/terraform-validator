package utils_test

import (
	"errors"
	"github.com/thazelart/terraform-validator/internal/utils"
	"testing"
)

func TestEnsureOrFatal(t *testing.T) {
	// After this test, replace the original fatal function
	origLogFatal := utils.LogFatal
	defer func() { utils.LogFatal = origLogFatal }()

	errs := []string{}
	utils.LogFatal = func(args ...interface{}) {
		errs = append(errs, "failed")
	}

	utils.EnsureOrFatal(errors.New("failed"))

	if len(errs) != 1 {
		t.Errorf("Got: %d error, wanted: 1", len(errs))
	}
}

func TestOkOrFatal(t *testing.T) {
	// After this test, replace the original fatal function
	origLogFatal := utils.LogFatal
	defer func() { utils.LogFatal = origLogFatal }()

	errs := []string{}
	utils.LogFatal = func(args ...interface{}) {
		errs = append(errs, "failed")
	}

	utils.OkOrFatal(false, "failed")
	utils.OkOrFatal(true, "failed")

	if len(errs) != 1 {
		t.Errorf("OkOrFatal() Got: %d error, wanted: 1", len(errs))
	}
}

func TestFileExists(t *testing.T) {

	result1 := utils.FileExists("utils_test.go")
	result2 := utils.FileExists("file_that_does_not_exist")

	// ensure result1 is true (equal)
	if !result1 {
		t.Errorf("Got: %v, wanted: true", result1)
	}
	// ensure result2 is false (not equal)
	if result2 {
		t.Errorf("Got: %v, wanted: false", result2)
	}
}
