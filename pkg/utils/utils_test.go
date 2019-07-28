package utils_test

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/kami-zh/go-capturer"
	"github.com/thazelart/terraform-validator/pkg/utils"
	"testing"
	"fmt"
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

func TestContains(t *testing.T) {
	list := []string{"foo", "bar"}
	// test1 true
	expectedResult := true
	testResult := utils.Contains(list, "foo")

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("Contains(foo) mismatch (-want +got):\n%s", diff)
	}
	// test2 false
	expectedResult = false
	testResult = utils.Contains(list, "foobar")

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("Contains(foo) mismatch (-want +got):\n%s", diff)
	}
}

func TestEnsureProgramInstalled(t *testing.T) {
	// After this test, replace the original fatal function
	origLogFatalf := utils.LogFatalf
	defer func() { utils.LogFatalf = origLogFatalf }()

	testResult := []string{}
	utils.LogFatalf = func(format string, a ...interface{}) {
		testResult = append(testResult, fmt.Sprintf(format, a[0]))
	}

	// test1 no error, ls program exists
	expectedResult := []string{}
	utils.EnsureProgramInstalled("ls")

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("EnsureProgramInstalled(ls) mismatch (-want +got):\n%s", diff)
	}

	// test2 error, dontExist program does not exist
	expectedResult = append(expectedResult, "FATAL: dontExist is not installed\n")
	utils.EnsureProgramInstalled("dontExist")

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("EnsureProgramInstalled(ls) mismatch (-want +got):\n%s", diff)
	}
}

func TestRunSystemCommand(t *testing.T) {
	// After this test, replace the original fatal function
	origLogFatalf := utils.LogFatalf
	defer func() { utils.LogFatalf = origLogFatalf }()

	errResult := []string{}
	utils.LogFatalf = func(format string, a ...interface{}) {
		errResult = append(errResult, fmt.Sprintf(format, a[0]))
	}

	// test1 ok: ls -1 utils.go
	expectedResult := "out: utils.go\n\nerr: "
	testResult := capturer.CaptureStdout(func() {
		utils.RunSystemCommand("ls", "-1", "utils.go")
	})

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("RunSystemCommand(ls -1 utils.go) mismatch (-want +got):\n%s", diff)
	}

	// test2 not ok: ls -1 utils.foo
	expectedResult = "out: \nerr: ls: cannot access 'utils.foo': No such file or directory\n"
	testResult = capturer.CaptureStdout(func() {
		utils.RunSystemCommand("ls", "-1", "utils.foo")
	})

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("RunSystemCommand(ls -1 utils.go) mismatch (-want +got):\n%s", diff)
	}
}
