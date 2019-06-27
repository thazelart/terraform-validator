package hcl_test

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/pkg/hcl"
)

func ExampleParseContent() {
	testFile := fs.File{Path: "/tmp/main.tf", Content: []byte(fileContent)}

	hclContent := hcl.ParseContent(testFile)

	fmt.Printf("Your parsed HCL content is:\n%+v\n", hclContent)
}

func ExampleGetTerraformBlockTypes() {
	testFile := fs.File{Path: "/tmp/main.tf", Content: []byte(fileContent)}

	hclContent := hcl.ParseContent(testFile)

	blockTypes := hcl.GetTerraformBlockTypes(hclContent)

	fmt.Printf("The blocks types contained in that file are:\n%+v\n", blockTypes)
}

func ExampleGetBlockNamesByType() {
	testFile := fs.File{Path: "/tmp/main.tf", Content: []byte(fileContent)}
	hclContent := hcl.ParseContent(testFile)

	for _, blockT := range hcl.TerraformBlockTypes {
		blockNames := hcl.GetBlockNamesByType(hclContent, blockT)
		fmt.Printf("Block type %s contains those elements : %v", blockT, blockNames)
	}
}
