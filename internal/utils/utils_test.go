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
