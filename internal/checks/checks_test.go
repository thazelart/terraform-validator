package checks_test

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/kami-zh/go-capturer"
	"github.com/thazelart/terraform-validator/internal/checks"
	"github.com/thazelart/terraform-validator/internal/hcl"
	"sort"
	"testing"
)

var parsedFileOk = hcl.ParsedFile{
	Name: "main.tf",
	Blocks: hcl.TerraformBlocks{
		Variables: []hcl.Variable{
			{Name: "var_with_description", Description: "a var description"},
			{Name: "var_2", Description: "2"},
		},
		Outputs: []hcl.Output{
			{Name: "out_with_description", Description: "a output description"},
			{Name: "out_2", Description: "2"},
		},
		Resources: []hcl.Resource{
			{Name: "a_resource", Type: "google_sql_database"},
		},
		Locals: []hcl.Locals{
			{"a_local", "another_local"},
			{"third_local"},
		},
		Data: []hcl.Data{
			{Name: "a_data", Type: "consul_key_prefix"},
		},
		Providers: []hcl.Provider{
			{Name: "google", Version: "=1.28.0"},
		},
		Terraform: hcl.Terraform{Version: "> 0.12.0", Backend: "gcs"},
		Modules: []hcl.Module{
			{Name: "consul", Version: "0.0.5"},
			{Name: "network", Version: "1.2.3"},
		},
	},
}

var parsedFileKo = hcl.ParsedFile{
	Name: "main.tf",
	Blocks: hcl.TerraformBlocks{
		Variables: []hcl.Variable{
			{Name: "var_with_description", Description: "a var description"},
			{Name: "var_2", Description: ""},
		},
		Outputs: []hcl.Output{
			{Name: "out_with_description", Description: "a output description"},
			{Name: "out_2", Description: ""},
		},
		Resources: []hcl.Resource{
			{Name: "a_resource", Type: "google_sql_database"},
		},
		Locals: []hcl.Locals{
			{"a_local", "another_local"},
			{"third_local"},
		},
		Data: []hcl.Data{
			{Name: "a_data", Type: "consul_key_prefix"},
			{Name: "a-data", Type: "consul_key_prefix"},
		},
		Providers: []hcl.Provider{
			{Name: "google", Version: "=1.28.0"},
		},
		Terraform: hcl.Terraform{Version: "> 0.12.0", Backend: "gcs"},
		Modules: []hcl.Module{
			{Name: "consul", Version: "0.0.5"},
			{Name: "network", Version: "1.2.3"},
		},
	},
}

func TestVerifyVariablesOutputsDescritions(t *testing.T) {
	parsedFileTest := parsedFileOk

	// without error
	var expectedResult []error
	testResult := checks.VerifyVariablesOutputsDescritions(parsedFileTest, true, true)

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("VerifyVariablesOutputsDescritions(ok) mismatch (-want +got):\n%s", diff)
	}

	// test2 with errors
	parsedFileTest = parsedFileKo
	expectedResult = append(expectedResult, fmt.Errorf("out_2 (output)"))
	expectedResult = append(expectedResult, fmt.Errorf("var_2 (variable)"))

	testResult = checks.VerifyVariablesOutputsDescritions(parsedFileTest, true, true)

	var stringExpectedResult, stringTestResult []string
	for i := range expectedResult {
		stringExpectedResult = append(stringExpectedResult, expectedResult[i].Error())
		stringTestResult = append(stringTestResult, testResult[i].Error())
	}
	sort.Strings(stringExpectedResult)
	sort.Strings(stringTestResult)
	if diff := cmp.Diff(stringExpectedResult, stringTestResult); diff != "" {
		t.Errorf("VerifyVariablesOutputsDescritions(ko) mismatch (-want +got):\n%s", diff)
	}
}

func TestVerifyFile(t *testing.T) {
	// test1 well formed ParsedFile
	parsedFileTest := parsedFileOk
	authorizedBocks := []string{"module", "locals", "provider", "data",
		"variable", "resource", "output", "terraform"}
	expectedOut := ""
	testOut := capturer.CaptureStdout(func() {
		checks.VerifyFile(parsedFileTest, "^[a-z0-9_]*$", authorizedBocks, true, true)
	})

	if diff := cmp.Diff(expectedOut, testOut); diff != "" {
		t.Errorf("VerifyFile(ok) mismatch (-want +got):\n%s", diff)
	}

	// test2 with unwanted blocks and misnamed bloc and no descrition
	parsedFileTest = parsedFileKo
	authorizedBocks = []string{"module", "locals", "provider", "data",
		"variable", "output", "terraform"}

	expectedOut = `ERROR: main.tf misformed:
  Unmatching "^[a-z0-9_]*$" pattern blockname(s):
    - a-data (data)
  Unauthorized block(s):
    - resource
  Undescribed variables(s) and/or output(s):
    - var_2 (variable)
    - out_2 (output)

`

	testOut = capturer.CaptureStdout(func() {
		checks.VerifyFile(parsedFileTest, "^[a-z0-9_]*$", authorizedBocks, true, true)
	})

	if diff := cmp.Diff(expectedOut, testOut); diff != "" {
		t.Errorf("VerifyFile(ko) mismatch (-want +got):\n%s", diff)
	}
}

func TestVerifyProvidersVersion(t *testing.T) {
	// test1: ok
	var parsedFolder = []hcl.ParsedFile{
		{
			Name: "one.tf",
			Blocks: hcl.TerraformBlocks{
				Providers: []hcl.Provider{
					{Name: "google", Version: "=1.28.0"},
				},
			},
		},
		{
			Name: "other.tf",
			Blocks: hcl.TerraformBlocks{
				Providers: []hcl.Provider{
					{Name: "aws", Version: "=1.2.0"},
				},
			},
		},
	}

	expectedOut := ""
	testOut := capturer.CaptureStdout(func() {
		checks.VerifyProvidersVersion(parsedFolder)
	})
	if diff := cmp.Diff(expectedOut, testOut); diff != "" {
		t.Errorf("VerifyProvidersVersion(ok) mismatch (-want +got):\n%s", diff)
	}

	// test2: aws version not set
	parsedFolder[1].Blocks.Providers[0].Version = ""

	expectedOut = "ERROR: Provider's version not set:\n  - aws\n\n"
	testOut = capturer.CaptureStdout(func() {
		checks.VerifyProvidersVersion(parsedFolder)
	})
	if diff := cmp.Diff(expectedOut, testOut); diff != "" {
		t.Errorf("VerifyProvidersVersion(ko) mismatch (-want +got):\n%s", diff)
	}
}

func TestVerifyTerraformVersion(t *testing.T) {
	// test1: ok
	var parsedFolder = []hcl.ParsedFile{
		{
			Name: "one.tf",
			Blocks: hcl.TerraformBlocks{
				Providers: []hcl.Provider{
					{Name: "google", Version: "=1.28.0"},
				},
			},
		},
		{
			Name: "other.tf",
			Blocks: hcl.TerraformBlocks{
				Terraform: hcl.Terraform{Version: "> 0.12.0", Backend: "gcs"},
			},
		},
	}

	expectedOut := ""
	testOut := capturer.CaptureStdout(func() {
		checks.VerifyTerraformVersion(parsedFolder)
	})
	if diff := cmp.Diff(expectedOut, testOut); diff != "" {
		t.Errorf("VerifyTerraformVersion(ok) mismatch (-want +got):\n%s", diff)
	}

	// test2: aws version not set
	parsedFolder[1].Blocks.Terraform.Version = ""

	expectedOut = "ERROR: Terraform's version not set\n\n"
	testOut = capturer.CaptureStdout(func() {
		checks.VerifyTerraformVersion(parsedFolder)
	})
	if diff := cmp.Diff(expectedOut, testOut); diff != "" {
		t.Errorf("VerifyTerraformVersion(ko) mismatch (-want +got):\n%s", diff)
	}
}

func TestVerifyMandatoryFilesPresent(t *testing.T) {
	var parsedFolder = []hcl.ParsedFile{
		{
			Name: "one.tf",
			Blocks: hcl.TerraformBlocks{
				Providers: []hcl.Provider{
					{Name: "google", Version: "=1.28.0"},
				},
			},
		},
	}

	// test1 all mandatory files are present
	mandatoryFiles := []string{"one.tf"}
	expectedOut := ""

	testOut := capturer.CaptureStdout(func() {
		checks.VerifyMandatoryFilesPresent(parsedFolder, mandatoryFiles)
	})
	if diff := cmp.Diff(expectedOut, testOut); diff != "" {
		t.Errorf("VerifyMandatoryFilesPresent(ok) mismatch (-want +got):\n%s", diff)
	}

	// test2 one missing file
	mandatoryFiles = []string{"one.tf", "two.tf"}
	expectedOut = `ERROR: missing mandatory file(s):
  - two.tf

`

	testOut = capturer.CaptureStdout(func() {
		checks.VerifyMandatoryFilesPresent(parsedFolder, mandatoryFiles)
	})
	if diff := cmp.Diff(expectedOut, testOut); diff != "" {
		t.Errorf("VerifyMandatoryFilesPresent(ko) mismatch (-want +got):\n%s", diff)
	}
}
