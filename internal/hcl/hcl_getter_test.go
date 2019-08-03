package hcl_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/hcl"
	"testing"
)

func TestGetBlockNamesByType(t *testing.T) {
	var testGame hcl.ParsedFile
	testGame.Name = "test.tf"
	testGame.Blocks = hcl.TestExpectedResult

	expectedResult := map[string][]string{
		"data":      {"a_data"},
		"locals":    {"a_local", "another_local", "third_local"},
		"module":    {"consul", "network"},
		"output":    {"out_with_description", "out_without_description"},
		"provider":  {"google"},
		"resource":  {"a_resource"},
		"terraform": {},
		"variable":  {"var_with_description", "var_without_description"},
	}

	testResult := testGame.GetBlockNamesByType()
	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetVariableBlocks() mismatch (-want +got):\n%s", diff)
	}

	// test2 with no resource
	expectedResult = map[string][]string{}
	testGame.Blocks = *new(hcl.TerraformBlocks)

	testResult = testGame.GetBlockNamesByType()
	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetVariableBlocks() mismatch (-want +got):\n%s", diff)
	}
}
