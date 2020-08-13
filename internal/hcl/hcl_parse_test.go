package hcl_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/internal/hcl"
	"testing"
)

func TestGetFolderParsedContents(t *testing.T) {
	testPath := "../checks/testdata/ok_default_config"

	expectedResult := []hcl.ParsedFile{
		{
			Name: "backend.tf",
			Blocks: hcl.TerraformBlocks{
				Terraform: hcl.Terraform{
					Version:           ">=0.12",
					Backend:           "",
					RequiredProviders: map[string]string{"aws": ">= 2.7.0", "newrelic": "~> 1.19"},
				},
			},
		},
		{
			Name:   "main.tf",
			Blocks: hcl.TerraformBlocks{},
		},
		{
			Name: "outputs.tf",
			Blocks: hcl.TerraformBlocks{
				Outputs: []hcl.Output{
					{
						Name:        "out1",
						Description: "",
					},
				},
			},
		},
		{
			Name: "providers.tf",
			Blocks: hcl.TerraformBlocks{
				Providers: []hcl.Provider{
					{
						Name:    "google",
						Version: "foo",
					},
					{
						Name:    "aws",
						Version: "foo2",
					},
				},
			},
		},
		{
			Name: "variables.tf",
			Blocks: hcl.TerraformBlocks{
				Variables: []hcl.Variable{
					{
						Name:        "var1",
						Description: "",
					},
				},
			},
		},
	}
	testResult := hcl.GetFolderParsedContents(testPath)

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetFolderParsedContents() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetParsedContent(t *testing.T) {
	// test1 with definition of all block types
	testFile := fs.File{Path: "/tmp/test.tf", Content: []byte(hcl.TestFileContent)}

	var expectedResult hcl.ParsedFile
	expectedResult.Name = "test.tf"
	expectedResult.Blocks = hcl.TestExpectedResult

	testResult := hcl.GetParsedContent(testFile)

	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("hclParse(all) mismatch (-want +got):\n%s", diff)
	}

	// test2 with only Variables
	TestFileContent := `
	variable "var_with_description" {
		description = "a var description"
		type = map(string)
	}`
	testFile = fs.File{Path: "/tmp/test.tf", Content: []byte(TestFileContent)}

	var expectedResult2 hcl.ParsedFile
	expectedResult2.Name = "test.tf"
	expectedResult2.Blocks.Variables = []hcl.Variable{
		{Name: "var_with_description", Description: "a var description"},
	}

	testResult2 := hcl.GetParsedContent(testFile)

	if diff := cmp.Diff(expectedResult2, testResult2); diff != "" {
		t.Errorf("hclParse(variable) mismatch (-want +got):\n%s", diff)
	}
}

func TestTerraformBlockIsEmpty(t *testing.T) {
	// test1 is Empty
	testTerraform := hcl.Terraform{
		Version:           "",
		Backend:           "",
		RequiredProviders: map[string]string{},
	}
	expectedResult := true

	testResult := hcl.TerraformBlockIsEmpty(testTerraform)

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("TerraformBlockIsEmpty() mismatch (-want +got):\n%s", diff)
	}

	// test2 only Backend
	testTerraform = hcl.Terraform{
		Version:           "",
		Backend:           "gcs",
		RequiredProviders: map[string]string{},
	}
	expectedResult = false

	testResult = hcl.TerraformBlockIsEmpty(testTerraform)

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("TerraformBlockIsEmpty(backend) mismatch (-want +got):\n%s", diff)
	}

	// test3 only version
	testTerraform = hcl.Terraform{
		Version:           "> 0.12",
		Backend:           "",
		RequiredProviders: map[string]string{},
	}
	expectedResult = false

	testResult = hcl.TerraformBlockIsEmpty(testTerraform)

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("TerraformBlockIsEmpty(version) mismatch (-want +got):\n%s", diff)
	}

	// test4 only requiredProvider
	testTerraform = hcl.Terraform{
		Version:           "",
		Backend:           "",
		RequiredProviders: map[string]string{"aws": ">= 2.7.0", "gcp": "", "newrelic": "~> 1.19"},
	}
	expectedResult = false

	testResult = hcl.TerraformBlockIsEmpty(testTerraform)

	if diff := cmp.Diff(testResult, expectedResult); diff != "" {
		t.Errorf("TerraformBlockIsEmpty(requiredVersion) mismatch (-want +got):\n%s", diff)
	}
}
