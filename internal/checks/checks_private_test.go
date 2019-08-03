package checks

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"sort"
	"testing"
)

func TestVerifyBlockNames(t *testing.T) {
	var testMap = map[string][]string{
		"data":      {"a_data"},
		"locals":    {"a_local", "another_local", "third_local"},
		"module":    {"consul", "network"},
		"output":    {"out_with_description", "out_without_description"},
		"provider":  {"google"},
		"resource":  {"a_resource"},
		"terraform": {},
		"variable":  {"var_with_description", "var_without_description"},
	}

	var expectedResult []error
	testResult := verifyBlockNames(testMap, "^[a-z0-9_]*$")
	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("verifyBlockNames() mismatch (-want +got):\n%s", diff)
	}

	// test2 with errors
	testMap["output"] = []string{"out_with_description", "out-description"}
	testMap["resource"] = []string{"a-resource"}

	expectedResult = append(expectedResult, fmt.Errorf("out-description (output)"))
	expectedResult = append(expectedResult, fmt.Errorf("a-resource (resource)"))
	testResult = verifyBlockNames(testMap, "^[a-z0-9_]*$")

	var stringExpectedResult, stringTestResult []string
	for i := range expectedResult {
		stringExpectedResult = append(stringExpectedResult, expectedResult[i].Error())
		stringTestResult = append(stringTestResult, testResult[i].Error())
	}
	sort.Strings(stringExpectedResult)
	sort.Strings(stringTestResult)
	if diff := cmp.Diff(stringExpectedResult, stringTestResult); diff != "" {
		t.Errorf("verifyAuthorizedBlocktypes(error) mismatch (-want +got):\n%s", diff)
	}
}

func TestVerifyAuthorizedBlocktypes(t *testing.T) {
	var testMap = map[string][]string{
		"data":   {"a_data"},
		"locals": {"a_local", "another_local", "third_local"},
		"module": {"consul", "network"},
	}
	authorizedBlocks := []string{"data", "locals", "module"}

	var expectedResult []error
	testResult := verifyAuthorizedBlocktypes(testMap, authorizedBlocks)
	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("verifyBlockNames() mismatch (-want +got):\n%s", diff)
	}

	// test2 with errors
	authorizedBlocks = []string{"module", "locals"}

	expectedResult = append(expectedResult, fmt.Errorf("data"))
	testResult = verifyAuthorizedBlocktypes(testMap, authorizedBlocks)

	var stringExpectedResult, stringTestResult []string
	for i := range expectedResult {
		stringExpectedResult = append(stringExpectedResult, expectedResult[i].Error())
		stringTestResult = append(stringTestResult, testResult[i].Error())
	}
	sort.Strings(stringExpectedResult)
	sort.Strings(stringTestResult)
	if diff := cmp.Diff(stringExpectedResult, stringTestResult); diff != "" {
		t.Errorf("verifyAuthorizedBlocktypes(error) mismatch (-want +got):\n%s", diff)
	}

	// test3 with errors because no blocks authorized
	authorizedBlocks = []string{}

	expectedResult = append(expectedResult, fmt.Errorf("locals"))
	expectedResult = append(expectedResult, fmt.Errorf("module"))
	testResult = verifyAuthorizedBlocktypes(testMap, authorizedBlocks)

	for i := range expectedResult {
		stringExpectedResult = append(stringExpectedResult, expectedResult[i].Error())
		stringTestResult = append(stringTestResult, testResult[i].Error())
	}
	sort.Strings(stringExpectedResult)
	sort.Strings(stringTestResult)
	if diff := cmp.Diff(stringExpectedResult, stringTestResult); diff != "" {
		t.Errorf("verifyAuthorizedBlocktypes(error) mismatch (-want +got):\n%s", diff)
	}
}
