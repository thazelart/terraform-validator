package hcl

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/config"
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
				*errs = append(*errs,
					fmt.Errorf("  %s block \"%s\" does not match \"%s\"",
						blockType, blockName, blockPatternName))
			}
		}
	}
}
