package checks

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/internal/hcl"
)

// MainChecks in the main function that will check an entire folder
func MainChecks(conf config.TfvConfig, workDir string) bool {
	isSuccess := true

	// Get current folder configuration
	globalConfig := conf.GetTerraformConfig(workDir)

	// Get current folder configuration class
	currentConfig := globalConfig.GetFolderConfigClass()

	// Get the terraform files informations
	folderParsedContent := hcl.GetFolderParsedContents(workDir)

	// Verify files normes and conventions
	if len(folderParsedContent) > 0 {
		fmt.Printf("INFO: running on %s with %s configuration\n",
			workDir, globalConfig.CurrentFolderClass)
		ok := FolderChecks(folderParsedContent, currentConfig)
		if !ok {
			isSuccess = false
		}
	}

	// Run inside sub-directories
	subfolders := fs.GetSubFolderList(workDir)
	for _, subfolder := range subfolders {
		ok := MainChecks(globalConfig, subfolder)
		if !ok {
			isSuccess = false
		}
	}

	return isSuccess
}

// FolderChecks run the check inside the given folder
func FolderChecks(folder []hcl.ParsedFile, config config.FolderConfigClass) bool {
	isSuccess := true

	for _, fileParsedContent := range folder {
		authorizedBlocks, _ := config.GetAuthorizedBlocks(fileParsedContent.Name)

		ok := VerifyFile(fileParsedContent,
			config.BlockPatternName,
			authorizedBlocks)

		if !ok {
			isSuccess = false
		}
	}

	// Ensure mandatory files are present
	mandatoryFiles := config.GetMandatoryFiles()
	ok := VerifyMandatoryFilesPresent(folder, mandatoryFiles)
	if !ok {
		isSuccess = false
	}

	// Ensure Providers version is set
	if config.EnsureProvidersVersion {
		ok := VerifyProvidersVersion(folder)
		if !ok {
			isSuccess = false
		}
	}

	// Ensure Terraform version is set
	if config.EnsureTerraformVersion {
		ok := VerifyTerraformVersion(folder)
		if !ok {
			isSuccess = false
		}
	}

	return isSuccess
}
