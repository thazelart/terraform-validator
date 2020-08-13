package hcl

import (
	hcl "github.com/hashicorp/hcl/v2"
)

type hclNameDescription struct {
	Name    string         `hcl:"name,label"`
	Default hcl.Attributes `hcl:"default,remain"`
}
type hclVariable hclNameDescription
type hclOutput hclNameDescription
type hclModule hclNameDescription

type hclProvider struct {
	Name    string   `hcl:"name,label"`
	Alias   *string  `hcl:"alias,attr"`
	Version *string  `hcl:"version,attr"`
	Config  hcl.Body `hcl:",remain"`
}

type hclNameType struct {
	Type   string   `hcl:"type,label"`
	Name   string   `hcl:"name,label"`
	Config hcl.Body `hcl:",remain"`
}
type hclResource hclNameType
type hclData hclNameType

type configOnly struct {
	Config hcl.Attributes `hcl:",remain"`
}
type hclLocals configOnly

type backend struct {
	Type   string   `hcl:"type,label"`
	Config hcl.Body `hcl:",remain"`
}

type requiredProviders struct {
	Config hcl.Attributes `hcl:",remain"`
}

type hclTerraform struct {
	RequiredVersion   *string            `hcl:"required_version,attr"`
	Backend           *backend           `hcl:"backend,block"`
	RequiredProviders *requiredProviders `hcl:"required_providers,block"`
}

type hclRoot struct {
	Variables []*hclVariable `hcl:"variable,block"`
	Outputs   []*hclOutput   `hcl:"output,block"`
	Resources []*hclResource `hcl:"resource,block"`
	Locals    []*hclLocals   `hcl:"locals,block"`
	Data      []*hclData     `hcl:"data,block"`
	Providers []*hclProvider `hcl:"provider,block"`
	Terraform *hclTerraform  `hcl:"terraform,block"`
	Modules   []*hclModule   `hcl:"module,block"`
}
