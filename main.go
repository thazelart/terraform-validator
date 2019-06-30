package main

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/utils"
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
		var blockNamesErrors []error
		var blocksInFilesErrors []error

		tfParsedContent := hcl.InitTerraformFileParsedContent(file)

		tfParsedContent.VerifyBlockNames(globalConfig, &blockNamesErrors)
		tfParsedContent.VerifyBlocksInFiles(globalConfig, file, &blocksInFilesErrors)

		if len(blockNamesErrors) > 0 || len(blocksInFilesErrors) > 0 {
			exitCode = 1
			fmt.Printf("\nERROR: %s misformed:\n", file.Path)
			if len(blockNamesErrors) > 0 {
				fmt.Printf("  Unmatching \"%s\" pattern block names:\n",
					globalConfig.TerraformConfig.BlockPatternName)
				for _, err := range blockNamesErrors {
					fmt.Printf("    - %s\n", err.Error())
				}
			}
			if len(blocksInFilesErrors) > 0 {
				fmt.Println("  Unauthorized blocks:")
				for _, err := range blocksInFilesErrors {
					fmt.Printf("    - %s\n", err.Error())
				}
			}
		}
	}

	// Check mandatory files
	TfFileList := globalConfig.GetFileNameList()
	var mandatoryErrors []error
	for _, mandatoryFile := range globalConfig.GetMandatoryFiles() {
		mandatoryFilePresent := utils.Contains(TfFileList, mandatoryFile)
		if !mandatoryFilePresent {
			mandatoryErrors = append(mandatoryErrors,
				fmt.Errorf("%s", mandatoryFile))
		}
	}
	if len(mandatoryErrors) > 0 {
		exitCode = 1
		fmt.Println("\nERROR: Missing mandatory file(s):")
		for _, err := range mandatoryErrors {
			fmt.Printf("  - %s\n", err.Error())
		}
	}
}
