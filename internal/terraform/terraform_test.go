package terraform_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/terraform"
	"testing"
)

func TestRunTerraformFmt(t *testing.T) {
	// test1 non formatted file
	filePath := "test_files/non_fmt.tf"

	expectedResult, expectedBool := "test_files/non_fmt.tf\n", false
	testResult, testBool := terraform.RunTerraformFmt(filePath)

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("RunTerraformFmt(non_fmt) mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(testBool, expectedBool); diff != "" {
		t.Errorf("RunTerraformFmt(non_fmt) mismatch (-want +got):\n%s", diff)
	}

	// test2 formatted file
	filePath = "test_files/fmt.tf"

	expectedResult, expectedBool = "", true
	testResult, testBool = terraform.RunTerraformFmt(filePath)

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("RunTerraformFmt(non_fmt) mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(testBool, expectedBool); diff != "" {
		t.Errorf("RunTerraformFmt(non_fmt) mismatch (-want +got):\n%s", diff)
	}
}
