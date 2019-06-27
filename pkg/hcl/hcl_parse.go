package hcl

import (
	hashicorp_hcl "github.com/hashicorp/hcl"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/internal/utils"
)

// TerraformFileParsedContent is a simple type created to contain all block
// names by blockTypes
type TerraformFileParsedContent map[string][]string

var (
	// TerraformBlockTypes contains all the blockTypes names
	TerraformBlockTypes = []string{
		"provider",
		"terraform",
		"variable",
		"output",
		"resource",
		"locals",
		"data",
		"module",
	}
)

// ParseContent parse a file and resturn the HCL parsed content
func ParseContent(file fs.File) map[string]interface{} {
	var output map[string]interface{}
	err := hashicorp_hcl.Unmarshal(file.Content, &output)
	utils.EnsureOrFatal(err)

	return output
}

// GetTerraformBlockTypes get all the terraform blockTypes contained in the
// given hcl content
func GetTerraformBlockTypes(hclContent map[string]interface{}) []string {
	var keys []string

	for k := range hclContent {
		if len(k) > 0 {
			keys = append(keys, k)
		}
	}

	return keys
}

// getSubBlock permit to find the sublocks in hclContent
func getSubBlock(hclContent interface{}) ([]string, interface{}) {
	typedHclContent, ok := hclContent.([]map[string]interface{})
	utils.OkOrFatal(ok, "FATAL Unable to decode hclContent into []map[string]interface{}")

	var keys []string
	//blocks :=  make(map[string]interface{})
	var blocks []map[string]interface{}

	// For each superBlock inside blockType
	for _, superBlock := range typedHclContent {
		// for each block inside the superBlock, get the block name
		for key, block := range superBlock {
			if block != nil {
				keys = append(keys, key)
				typedBlock, ok := block.([]map[string]interface{})
				if ok {
					blocks = append(blocks, typedBlock...)
				} else {
					typedBlock, ok := block.(string)
					utils.OkOrFatal(ok, "FATAL Unable to decode block into []map[string]interface{} neither string")
					blocks = append(blocks, map[string]interface{}{key: typedBlock})
				}
			}
		}
	}

	return keys, blocks
}

// GetBlockNamesByType return you the sublocks names inside the given blockType
// for example: gets you all the resources blocks names inside block blockType
func GetBlockNamesByType(hclContent map[string]interface{}, blockType string) []string {
	var keys []string

	if blockType == "resource" || blockType == "data" {
		_, blocks := getSubBlock(hclContent[blockType])
		keys, _ = getSubBlock(blocks)
	} else {
		keys, _ = getSubBlock(hclContent[blockType])
	}

	return keys
}

// InitTerraformFileParsedContent init you the TerraformFileParsedContent. This
// means that it gets you all the subblock names inside each blockTypes
func InitTerraformFileParsedContent(file fs.File) TerraformFileParsedContent {
	terraformFileParsedContent := make(map[string][]string)

	hclContent := ParseContent(file)

	blockTypes := GetTerraformBlockTypes(hclContent)

	for _, key := range blockTypes {
		terraformFileParsedContent[key] = GetBlockNamesByType(hclContent, key)
	}

	return terraformFileParsedContent
}
