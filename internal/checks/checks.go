package checks

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/hcl"
	"github.com/thazelart/terraform-validator/pkg/utils"
	"regexp"
	"sort"
)

// VerifyBlockNames ensure that all the terraform blocks are well named
func verifyBlockNames(blocks map[string][]string, pattern string) (errs []error) {
	for blockType, blockNames := range blocks {
		if blockType == "provider" {
			// provider names are not chosen by the user
			continue
		}
		for _, blockName := range blockNames {
			matched, _ := regexp.MatchString(pattern, blockName)
			if !matched {
				errs = append(errs, fmt.Errorf("%s (%s)", blockName, blockType))
			}
		}
	}
	return errs
}

// VerifyBlockNames ensure that all the terraform blocks are well named
func verifyAuthorizedBlocktypes(blocks map[string][]string, authorizedBlocks []string) (errs []error) {
	for blockType := range blocks {
		if utils.Contains(authorizedBlocks, blockType) {
			continue
		}
		errs = append(errs, fmt.Errorf("%s", blockType))
	}
	return errs
}

// VerifyVariablesOutputsDescritions ensure that all the variables and
// outputs blocks have a descrition
func VerifyVariablesOutputsDescritions(parsedFile hcl.ParsedFile,
	verifyVariables bool, verifyOutputs bool) (errs []error) {
	if verifyVariables {
		for _, variable := range parsedFile.Blocks.Variables {
			if variable.Description == "" {
				errs = append(errs, fmt.Errorf("%s (variable)", variable.Name))
			}
		}
	}
	if verifyOutputs {
		for _, output := range parsedFile.Blocks.Outputs {
			if output.Description == "" {
				errs = append(errs, fmt.Errorf("%s (output)", output.Name))
			}
		}
	}
	return errs
}

// VerifyFile launch every check that are file dependant (block names and
// authorized blocks)
func VerifyFile(parsedFile hcl.ParsedFile, pattern string,
	authorizedBlocks []string, verifyVariables bool, verifyOutputs bool) bool {

	blocks := parsedFile.GetBlockNamesByType()

	bnErrs := verifyBlockNames(blocks, pattern)
	btErrs := verifyAuthorizedBlocktypes(blocks, authorizedBlocks)
	ioErrs := VerifyVariablesOutputsDescritions(parsedFile, verifyVariables, verifyOutputs)

	hasBnErrs := len(bnErrs) > 0
	hasBtErrs := len(btErrs) > 0
	hasIoErrs := len(ioErrs) > 0

	if hasBnErrs || hasBtErrs || hasIoErrs {
		fmt.Printf("ERROR: %s misformed:\n", parsedFile.Name)
		if hasBnErrs {
			fmt.Printf("  Unmatching \"%s\" pattern blockname(s):\n", pattern)
			for _, err := range bnErrs {
				fmt.Printf("    - %s\n", err.Error())
			}
		}
		if hasBtErrs {
			fmt.Println("  Unauthorized block(s):")
			for _, err := range btErrs {
				fmt.Printf("    - %s\n", err.Error())
			}
		}
		if hasIoErrs {
			fmt.Println("  Undescribed variables(s) and/or output(s):")
			for _, err := range ioErrs {
				fmt.Printf("    - %s\n", err.Error())
			}
		}
		fmt.Println()
		return false
	}
	return true
}

// VerifyProvidersVersion ensure that all providers have a version
func VerifyProvidersVersion(parsedFolder []hcl.ParsedFile) bool {
	var errs []error
	for _, parsedFile := range parsedFolder {
		for _, provider := range parsedFile.Blocks.Providers {
			if provider.Version == "" {
				errs = append(errs, fmt.Errorf("%s", provider.Name))
			}
		}
	}
	if len(errs) > 0 {
		fmt.Println("ERROR: Provider's version not set:")
		for _, err := range errs {
			fmt.Printf("  - %s\n", err.Error())
		}
		fmt.Println()
		return false
	}
	return true
}

// VerifyTerraformVersion ensure that the terraform version is set
func VerifyTerraformVersion(parsedFolder []hcl.ParsedFile) bool {
	isTerraformVersionSet := false
	for _, parsedFile := range parsedFolder {
		if parsedFile.Blocks.Terraform.Version != "" {
			isTerraformVersionSet = true
		}
	}

	if !isTerraformVersionSet {
		fmt.Println("ERROR: Terraform's version not set")
		fmt.Println()
		return false
	}
	return true
}

// VerifyMandatoryFilesPresent ensure that the mandatory files are present
func VerifyMandatoryFilesPresent(parsedFolder []hcl.ParsedFile,
	mandatoryFiles []string) bool {

	var files, missingFiles []string

	// Get terraform file list
	for _, parsedFile := range parsedFolder {
		files = append(files, parsedFile.Name)
	}

	// Ensure mandatory ones are present
	for _, mandatoryFile := range mandatoryFiles {
		if !utils.Contains(files, mandatoryFile) {
			missingFiles = append(missingFiles, mandatoryFile)
		}
	}

	sort.Strings(missingFiles)

	// Print errors
	if len(missingFiles) > 0 {
		fmt.Println("ERROR: missing mandatory file(s):")
		for _, missingFile := range missingFiles {
			fmt.Printf("  - %s\n", missingFile)
		}
		fmt.Println()
		return false
	}

	return true
}
