package hcl

import (
	"github.com/google/go-cmp/cmp"
	"github.com/thazelart/terraform-validator/internal/fs"
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

func TestGetSubBlock(t *testing.T) {
	testFile := fs.File{Path: "/tmp/main.tf", Content: []byte(fileContent)}
	hclContent := ParseContent(testFile)

	expectedKeys := TerraformFileParsedContent{
		"terraform": []string{"required_version"},
		"variable":  []string{"one_map", "a_list"},
		"output":    []string{"my_ip", "my_name"},
		"resource":  []string{"google_storage_bucket", "google_storage_bucket"},
		"locals":    []string{"package_name", "creator"},
		"data":      []string{"google_compute_image"},
		"module":    []string{"module_instance_name"},
		"provider":  []string{"google", "github"},
	}

	expectedBlocks := map[string]interface{}{
		"terraform": []map[string]interface{}{
			{"required_version": string("~> 1.0")},
		},
		"variable": []map[string]interface{}{
			{"type": string("map")},
			{"type": string("list")},
		},
		"output": []map[string]interface{}{
			{"value": string("127.0.0.1")},
			{"value": string("thazelart")},
		},
		"resource": []map[string]interface{}{
			{"bucket_1": []map[string]interface{}{{"name": string("bucket1")}}},
			{"bucket42": []map[string]interface{}{{"name": string("bucket1")}}},
		},

		"locals": []map[string]interface{}{
			{"package_name": string("terraform-validator")},
			{"creator": string("thazelart")},
		},
		"data": []map[string]interface{}{
			{"centos7": []map[string]interface{}{{"family": string("centos-7")}}},
		},
		"module": []map[string]interface{}{
			{"source": string("git@github.com:thazelart/tf_mod_test")},
		},
		"provider": []map[string]interface{}{
			{"project": string("my-rpoject-xdfg"), "version": string("~> 1.0")},
			{"organization": string("thazelart")},
		},
	}

	// ensure result1 is true (equal)
	for _, blockT := range TerraformBlockTypes {
		testkeys, testBlocks := getSubBlock(hclContent[blockT])
		if diff := cmp.Diff(expectedKeys[blockT], testkeys); diff != "" {
			t.Errorf("getSubBlock(%s-key) mismatch (-want +got):\n%s", blockT, diff)
		}
		if diff := cmp.Diff(expectedBlocks[blockT], testBlocks); diff != "" {
			t.Errorf("getSubBlock(%s-block) mismatch (-want +got):\n%s", blockT, diff)
		}
	}
}
