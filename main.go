package main

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/pkg/hcl"
	"os"
)

const (
	version     = "0.0.1"
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

func main() {
	exitCode := 0
	defer func() { os.Exit(exitCode) }()

	globalConfig := config.GenerateGlobalConfig(version)

	for _, file := range globalConfig.WorkDir.Content {
		var errors []error

		tfParsedContent := hcl.InitTerraformFileParsedContent(file)

		tfParsedContent.VerifyBlockNames(globalConfig, &errors)

		if len(errors) > 0 {
			exitCode = 1
			fmt.Printf("%s contains errors:\n", file.Path)
			for _, err := range errors {
				fmt.Println(err)
			}
		}

	}
}
