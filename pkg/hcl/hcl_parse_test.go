package hcl_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/pkg/hcl"
	"sort"
	"testing"
)

const (
	fileContent = `variable "one_map" {
	  type        = "map"
	}
	variable "a_list" {
	  type        = "list"
	}

	output "my_ip" {
	  value       = "127.0.0.1"
	}
	output "my_name" {
	  value       = "thazelart"
	}


	resource "google_storage_bucket" "bucket_1" {
	  name     = "bucket1"
	}
	resource "google_storage_bucket" "bucket42" {
	  name     = "bucket1"
	}

	locals {
	  package_name = "terraform-validator"
	}
	locals {
	  creator = "thazelart"
	}

	data "google_compute_image" "centos7" {
	  family  = "centos-7"
	}

	module "module_instance_name" {
	  source               = "git@github.com:thazelart/tf_mod_test"
	}

	provider "google" {
	  project = "my-rpoject-xdfg"
	  version = "~> 1.0"
	}

	provider "github" {
	  organization = "thazelart"
	}

	terraform {
	  required_version = "~> 1.0"
	}`
)

func TestParseContent(t *testing.T) {
	testFile := fs.File{Path: "/tmp/main.tf", Content: []byte(fileContent)}

	hclContent := hcl.ParseContent(testFile)

	for _, blockT := range hcl.TerraformBlockTypes {
		if _, ok := hclContent[blockT]; !ok {
			t.Errorf("ParseContent() %s block type is not defined", blockT)
		}
	}
}

func TestGetTerraformBlockTypes(t *testing.T) {
	testFile := fs.File{Path: "/tmp/main.tf", Content: []byte(fileContent)}
	expectedResult := hcl.TerraformBlockTypes
	sort.Strings(expectedResult)

	hclContent := hcl.ParseContent(testFile)
	testResult := hcl.GetTerraformBlockTypes(hclContent)
	sort.Strings(testResult)
	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetTerraformBlockTypes() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetBlockNamesByType(t *testing.T) {
	testFile := fs.File{Path: "/tmp/main.tf", Content: []byte(fileContent)}
	hclContent := hcl.ParseContent(testFile)

	expectedResult := hcl.TerraformFileParsedContent{
		"terraform": []string{"required_version"},
		"variable":  []string{"one_map", "a_list"},
		"output":    []string{"my_ip", "my_name"},
		"resource":  []string{"bucket_1", "bucket42"},
		"locals":    []string{"package_name", "creator"},
		"data":      []string{"centos7"},
		"module":    []string{"module_instance_name"},
		"provider":  []string{"google", "github"},
	}

	// ensure result1 is true (equal)
	for _, blockT := range hcl.TerraformBlockTypes {
		testResult := hcl.GetBlockNamesByType(hclContent, blockT)
		if diff := cmp.Diff(expectedResult[blockT], testResult); diff != "" {
			t.Errorf("GetBlockNamesByType(%s) mismatch (-want +got):\n%s", blockT, diff)
		}
	}
}

func TestInitTerraformFileParsedContent(t *testing.T) {
	testFile := fs.File{Path: "/tmp/main.tf", Content: []byte(fileContent)}

	expectedResult := hcl.TerraformFileParsedContent{
		"terraform": []string{"required_version"},
		"variable":  []string{"one_map", "a_list"},
		"output":    []string{"my_ip", "my_name"},
		"resource":  []string{"bucket_1", "bucket42"},
		"locals":    []string{"package_name", "creator"},
		"data":      []string{"centos7"},
		"module":    []string{"module_instance_name"},
		"provider":  []string{"google", "github"},
	}

	testResult := hcl.InitTerraformFileParsedContent(testFile)
	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("InitTerraformFileParsedContent mismatch (-want +got):\n%s", diff)
	}
}

func TestGetProviderConfiguration(t *testing.T) {
	testFile := fs.File{Path: "/tmp/main.tf", Content: []byte(fileContent)}
	expectedResult := map[string][]string{
		"google": {"project", "version"},
		"github": {"organization"},
	}

	testResult := hcl.GetProviderConfiguration(testFile)
	sort.Strings(testResult["google"])
	if diff := cmp.Diff(expectedResult, testResult); diff != "" {
		t.Errorf("GetProviderConfiguration mismatch (-want +got):\n%s", diff)
	}
}
