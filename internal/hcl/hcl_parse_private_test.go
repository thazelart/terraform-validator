package hcl

import (
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/fs"
	"testing"
)

const (
	TestFileContent = `
variable "var_with_description" {
	description = "a var description"
	type = map(string)
}

variable "var_without_description" {
	type        = string
}

output "out_with_description" {
	description = "a output description"
	type = map(string)
}

output "out_without_description" {
	type        = string
}

resource "google_sql_database" "a_resource" {
	name      = "a_resource_not_name"
	instance  = a_resource_instance
	charset   = "UTF8"
	collation = "en_US.UTF8"
}

locals {
  a_local = "foo"
	another_local = "bar"
}

locals {
  third_local = "foobar"
}

data "consul_key_prefix" "a_data" {
  path = "apps/example/env"
}

provider "google" {
  version = "=1.28.0"
}

terraform {
	required_version = "> 0.12.0"
	backend "gcs" {}
}

module "consul" {
  source  = "hashicorp/consul/aws"
  version = "0.0.5"

  servers = 3
}

module "network" {
  source  = "github.com/thazelart/a_module?href=1.2.3"
}
`
)

var TestExpectedResult = TerraformBlocks{
	Variables: []Variable{
		{Name: "var_with_description", Description: "a var description"},
		{Name: "var_without_description", Description: ""},
	},
	Outputs: []Output{
		{Name: "out_with_description", Description: "a output description"},
		{Name: "out_without_description", Description: ""},
	},
	Resources: []Resource{
		{Name: "a_resource", Type: "google_sql_database"},
	},
	Locals: []Locals{
		{"a_local", "another_local"},
		{"third_local"},
	},
	Data: []Data{
		{Name: "a_data", Type: "consul_key_prefix"},
	},
	Providers: []Provider{
		{Name: "google", Version: "=1.28.0"},
	},
	Terraform: Terraform{Version: "> 0.12.0", Backend: "gcs"},
	Modules: []Module{
		{Name: "consul", Version: "0.0.5"},
		{Name: "network", Version: "1.2.3"},
	},
}

func TestHclParse(t *testing.T) {
	testFile := fs.File{Path: "test.tf", Content: []byte(TestFileContent)}
	parsedContent := hclParse(testFile)

	expectedVariablesLen := 2
	expectedOutputsLen := 2
	expectedResourcesLen := 1

	resultVariablesLen := len(parsedContent.Variables)
	resultOutputsLen := len(parsedContent.Outputs)
	resultResourcesLen := len(parsedContent.Resources)

	if diff := cmp.Diff(resultVariablesLen, expectedVariablesLen); diff != "" {
		t.Errorf("hclParse(variable) mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(resultOutputsLen, expectedOutputsLen); diff != "" {
		t.Errorf("hclParse(Outputs) mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(resultResourcesLen, expectedResourcesLen); diff != "" {
		t.Errorf("hclParse(Resources) mismatch (-want +got):\n%s", diff)
	}
}

func TestGetVariablesInfomation(t *testing.T) {
	testFile := fs.File{Path: "test.tf", Content: []byte(TestFileContent)}
	parsedContent := hclParse(testFile)

	testResult := parsedContent.getVariablesInfomation()

	if diff := cmp.Diff(testResult, TestExpectedResult.Variables); diff != "" {
		t.Errorf("getVariablesInfomation() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetOutputsInfomation(t *testing.T) {
	testFile := fs.File{Path: "test.tf", Content: []byte(TestFileContent)}
	parsedContent := hclParse(testFile)

	testResult := parsedContent.getOutputsInfomation()

	if diff := cmp.Diff(testResult, TestExpectedResult.Outputs); diff != "" {
		t.Errorf("getOutputsInfomation() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetResourcesInfomation(t *testing.T) {
	testFile := fs.File{Path: "test.tf", Content: []byte(TestFileContent)}
	parsedContent := hclParse(testFile)

	testResult := parsedContent.getResourcesInfomation()

	if diff := cmp.Diff(testResult, TestExpectedResult.Resources); diff != "" {
		t.Errorf("getResourcesInfomation() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetDataInfomation(t *testing.T) {
	testFile := fs.File{Path: "test.tf", Content: []byte(TestFileContent)}
	parsedContent := hclParse(testFile)

	testResult := parsedContent.getDataInfomation()

	if diff := cmp.Diff(testResult, TestExpectedResult.Data); diff != "" {
		t.Errorf("getDataInfomation() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetLocalsInfomation(t *testing.T) {
	testFile := fs.File{Path: "test.tf", Content: []byte(TestFileContent)}
	parsedContent := hclParse(testFile)

	testResult := parsedContent.getLocalsInfomation()

	if diff := cmp.Diff(testResult, TestExpectedResult.Locals); diff != "" {
		t.Errorf("TestGetLocalsInfomation() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetProvidersInfomation(t *testing.T) {
	testFile := fs.File{Path: "test.tf", Content: []byte(TestFileContent)}
	parsedContent := hclParse(testFile)
	testResult := parsedContent.getProvidersInfomation()

	if diff := cmp.Diff(testResult, TestExpectedResult.Providers); diff != "" {
		t.Errorf("getProvidersInfomation() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetModulesInfomation(t *testing.T) {
	testFile := fs.File{Path: "test.tf", Content: []byte(TestFileContent)}
	parsedContent := hclParse(testFile)

	testResult := parsedContent.getModulesInfomation()

	if diff := cmp.Diff(testResult, TestExpectedResult.Modules); diff != "" {
		t.Errorf("getModulesInfomation() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetTerraformInfomation(t *testing.T) {
	testFile := fs.File{Path: "test.tf", Content: []byte(TestFileContent)}
	parsedContent := hclParse(testFile)

	testResult := parsedContent.getTerraformInfomation()

	if diff := cmp.Diff(testResult, TestExpectedResult.Terraform); diff != "" {
		t.Errorf("getTerraformInfomation() mismatch (-want +got):\n%s", diff)
	}
}
