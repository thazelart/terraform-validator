package checks

import (
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/hcl"
)

// MainChecks in the main function that will check an entire folder
func MainChecks(conf config.TfvConfig, workDir string) bool {
	isSuccess := true

	// Get current folder configuration
	currentFolderConfig := conf.GetTerraformConfig(workDir)

	// Get current folder configuration class
	classConfig := currentFolderConfig.GetFolderConfigClass()

	// Get the terraform files informations
	folderParsedContent := hcl.GetFolderParsedContents(workDir)

	// Verify files normes and conventions
	for _, fileParsedContent := range folderParsedContent {
		authorizedBlocks, _ := classConfig.GetAuthorizedBlocks(fileParsedContent.Name)

		ok := VerifyFile(fileParsedContent,
			classConfig.BlockPatternName,
			authorizedBlocks)

		if !ok {
			isSuccess = false
		}
	}

	// Ensure mandatory files are present
	mandatoryFiles := classConfig.GetMandatoryFiles()
	ok := VerifyMandatoryFilesPresent(folderParsedContent, mandatoryFiles)
	if !ok {
		isSuccess = false
	}

	// Ensure Providers version is set
	if classConfig.EnsureProvidersVersion {
		ok := VerifyProvidersVersion(folderParsedContent)
		if !ok {
			isSuccess = false
		}
	}

	// Ensure Terraform version is set
	if classConfig.EnsureTerraformVersion {
		ok := VerifyTerraformVersion(folderParsedContent)
		if !ok {
			isSuccess = false
		}
	}

	return isSuccess
}
