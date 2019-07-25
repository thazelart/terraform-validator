package main

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/utils"
	"github.com/thazelart/terraform-validator/pkg/hcl"
	"os"
)

const (
	version = "1.3.4"
)

var (
	isTerraformVersionSet = false
)

func main() {
	exitCode := 0
	defer func() {
		if exitCode == 0 {
			fmt.Println("INFO: terraform-validator ran successfully")
		}
		os.Exit(exitCode)
	}()

	globalConfig := config.GenerateGlobalConfig(version)

	for _, file := range globalConfig.WorkDir.Content {
		var blockNamesErrors []error
		var blocksInFilesErrors []error
		var providersVersionErrors []error

		tfParsedContent := hcl.InitTerraformFileParsedContent(file)

		tfParsedContent.VerifyBlockNames(globalConfig, &blockNamesErrors)
		tfParsedContent.VerifyBlocksInFiles(globalConfig, file, &blocksInFilesErrors)

		// if terraform version not yet set verify if that file contain it
		if !isTerraformVersionSet && globalConfig.TerraformConfig.EnsureTerraformVersion {
			isTerraformVersionSet = tfParsedContent.ContainsTerraformVersion()
		}

		if globalConfig.TerraformConfig.EnsureProvidersVersion {
			tfParsedContent.ContainsProvidersVersion(file, &providersVersionErrors)
		}

		if len(blockNamesErrors) > 0 || len(blocksInFilesErrors) > 0 || len(providersVersionErrors) > 0 {
			exitCode = 1
			fmt.Printf("\nERROR: %s misformed:\n", file.Path)
			if len(providersVersionErrors) > 0 {
				fmt.Printf("  Unversioned provider(s):\n")
				for _, err := range providersVersionErrors {
					fmt.Printf("    - %s\n", err.Error())
				}
			}
			if len(blockNamesErrors) > 0 {
				fmt.Printf("  Unmatching \"%s\" pattern blockname(s):\n",
					globalConfig.TerraformConfig.BlockPatternName)
				for _, err := range blockNamesErrors {
					fmt.Printf("    - %s\n", err.Error())
				}
			}
			if len(blocksInFilesErrors) > 0 {
				fmt.Println("  Unauthorized block(s):")
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

	if !isTerraformVersionSet && globalConfig.TerraformConfig.EnsureTerraformVersion {
		fmt.Println("\nERROR: Terraform version was not set")
	}
}
