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
