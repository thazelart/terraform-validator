package hcl

// TerraformBlockTypes = []string{
// 	"terraform",
// 	"module",
// }

type nameDescription struct {
	Name        string
	Description string
}

// Variable get all needed information of variable blocks for terraform-validator
type Variable nameDescription

// Output get all needed information of output blocks for terraform-validator
type Output nameDescription

type nameType struct {
	Name string
	Type string
}

// Resource get all needed information of resource blocks for terraform-validator
type Resource nameType

// Data get all needed information of data blocks for terraform-validator
type Data nameType

type nameVersion struct {
	Name    string
	Version string
}

// Provider get all needed information of provider blocks for terraform-validator
type Provider nameVersion

// Module get all needed information of module blocks for terraform-validator
type Module nameVersion

// Terraform get all needed information of terraform config blocks for terraform-validator
type Terraform struct {
	Version string
	Backend string
}

// Locals get all needed information of Locals config blocks for terraform-validator
type Locals []string

// TerraformBlocks is the structure that define a whole terraform file and
// contains all the blocks informations needed by terraform-validator
type TerraformBlocks struct {
	Variables []Variable
	Terraform Terraform
	Modules   []Module
	Providers []Provider
	Data      []Data
	Resources []Resource
	Locals    []Locals
	Outputs   []Output
}

// ParsedFile get all needed information of all blocks for terraform-validator
type ParsedFile struct {
	Name   string
	Blocks TerraformBlocks
}
