package tfv

import (
	"github.com/thazelart/terraform-validator/internal/checks"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/hcl"
)

// MainChecks in the main function that will check an entire folder
func MainChecks(globalConfig config.GlobalConfig) bool {
	isSuccess := true

	// Get the terraform files informations
	folderParsedContent := hcl.GetFolderParsedContents(globalConfig.WorkDir)

	// Verify files normes and conventions
	for _, fileParsedContent := range folderParsedContent {
		authorizedBlocks, _ := globalConfig.GetAuthorizedBlocks(fileParsedContent.Name)

		ok := checks.VerifyFile(fileParsedContent,
			globalConfig.TerraformConfig.BlockPatternName,
			authorizedBlocks)

		if !ok {
			isSuccess = false
		}
	}

	// Ensure mandatory files are present
	mandatoryFiles := globalConfig.GetMandatoryFiles()
	ok := checks.VerifyMandatoryFilesPresent(folderParsedContent, mandatoryFiles)
	if !ok {
		isSuccess = false
	}

	// Ensure Providers version is set
	if globalConfig.TerraformConfig.EnsureProvidersVersion {
		ok := checks.VerifyProvidersVersion(folderParsedContent)
		if !ok {
			isSuccess = false
		}
	}

	// Ensure Terraform version is set
	if globalConfig.TerraformConfig.EnsureTerraformVersion {
		ok := checks.VerifyTerraformVersion(folderParsedContent)
		if !ok {
			isSuccess = false
		}
	}

	return isSuccess
}
