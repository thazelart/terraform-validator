package hcl

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/internal/utils"
	"regexp"
)

// VerifyBlockNames ensure that all the terraform blocks are well named
func (terraformFileParsedContent TerraformFileParsedContent) VerifyBlockNames(
	config config.GlobalConfig, errs *[]error) {
	blockPatternName := config.TerraformConfig.BlockPatternName

	for blockType, subBlocks := range terraformFileParsedContent {
		if utils.Contains([]string{"provider", "terraform"}, blockType) {
			continue
		}
		for _, blockName := range subBlocks {
			matched, _ := regexp.MatchString(blockPatternName, blockName)
			if !matched {
				*errs = append(*errs, fmt.Errorf("%s (%s)", blockName, blockType))
			}
		}
	}
}

// VerifyBlocksInFiles verify that all the blocks in you file are allowed
func (terraformFileParsedContent TerraformFileParsedContent) VerifyBlocksInFiles(
	config config.GlobalConfig, file fs.File, errs *[]error) {
	filename := file.GetFilename()

	authorizedBlocks, err := config.GetAuthorizedBlocks(filename)
	if err != nil {
		*errs = append(*errs, err)
		return
	}

	for blockType, _ := range terraformFileParsedContent {
		authorized := utils.Contains(authorizedBlocks, blockType)
		if !authorized {
			*errs = append(*errs, fmt.Errorf("%s", blockType))
		}
	}
}

// ContainsTerraformVersion return true if the terraform version was set
func (terraformFileParsedContent TerraformFileParsedContent) ContainsTerraformVersion() bool {
	terraformBlock, terraformBlockPresent := terraformFileParsedContent["terraform"]
	if !terraformBlockPresent {
		return false
	}
	return utils.Contains(terraformBlock, "required_version")
}

// ContainsProvidersVersion return nil if all providers versions are set else a slice of error
func (terraformFileParsedContent TerraformFileParsedContent) ContainsProvidersVersion(file fs.File, errs *[]error) {
	_, providerBlockPresent := terraformFileParsedContent["provider"]
	if !providerBlockPresent {
		return
	}
	providerContent := GetProviderConfiguration(file)
	for provider, configurations := range providerContent {
		if !utils.Contains(configurations, "version") {
			*errs = append(*errs, fmt.Errorf("%s", provider))
		}
	}
}
