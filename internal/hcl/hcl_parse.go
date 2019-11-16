// Package hcl handle the parsing of your terraform file
package hcl

import (
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/thazelart/terraform-validator/internal/fs"
	"github.com/thazelart/terraform-validator/pkg/utils"
	"sort"
	"strings"
)

// GetFolderParsedContents returns you a []ParsedFile that define every file
// inside the given path
func GetFolderParsedContents(path string) []ParsedFile {
	var result []ParsedFile

	folder := fs.NewTerraformFolder(path)

	for _, file := range folder.Content {
		result = append(result, GetParsedContent(file))
	}

	return result
}

// GetParsedContent will parse a file for you and return the associated ParsedFile
func GetParsedContent(f fs.File) ParsedFile {
	var result ParsedFile
	root := hclParse(f)

	result.Name = f.GetFilename()

	if variables := root.getVariablesInfomation(); variables != nil {
		result.Blocks.Variables = variables
	}
	if outputs := root.getOutputsInfomation(); outputs != nil {
		result.Blocks.Outputs = outputs
	}
	if resources := root.getResourcesInfomation(); resources != nil {
		result.Blocks.Resources = resources
	}
	if providers := root.getProvidersInfomation(); providers != nil {
		result.Blocks.Providers = providers
	}
	if data := root.getDataInfomation(); data != nil {
		result.Blocks.Data = data
	}
	if locals := root.getLocalsInfomation(); locals != nil {
		result.Blocks.Locals = locals
	}
	if modules := root.getModulesInfomation(); modules != nil {
		result.Blocks.Modules = modules
	}
	if terraform := root.getTerraformInfomation(); terraform != *new(Terraform) {
		result.Blocks.Terraform = terraform
	}

	return result
}

func hclParse(f fs.File) hclRoot {
	var result hclRoot

	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL(f.Content, f.Path)
	utils.NoDiagsOrFatal(diags)

	diags = gohcl.DecodeBody(file.Body, nil, &result)
	utils.NoDiagsOrFatal(diags)

	return result
}

func (hclroot hclRoot) getVariablesInfomation() []Variable {
	var result []Variable
	for _, variable := range hclroot.Variables {
		var description string
		for _, arg := range variable.Default {
			if arg.Name == "description" {
				gohcl.DecodeExpression(arg.Expr, nil, &description)
			}
		}
		result = append(result, Variable{Name: variable.Name, Description: description})
	}

	return result
}

func (hclroot hclRoot) getOutputsInfomation() []Output {
	var result []Output

	for _, output := range hclroot.Outputs {
		var description string
		for _, arg := range output.Default {
			if arg.Name == "description" {
				gohcl.DecodeExpression(arg.Expr, nil, &description)
			}
		}
		result = append(result, Output{Name: output.Name, Description: description})
	}

	return result
}

func (hclroot hclRoot) getResourcesInfomation() []Resource {
	var result []Resource

	for _, resource := range hclroot.Resources {
		result = append(result, Resource{Name: resource.Name, Type: resource.Type})
	}

	return result
}

func (hclroot hclRoot) getDataInfomation() []Data {
	var result []Data

	for _, data := range hclroot.Data {
		result = append(result, Data{Name: data.Name, Type: data.Type})
	}

	return result
}

func (hclroot hclRoot) getLocalsInfomation() []Locals {
	var result []Locals

	for _, locals := range hclroot.Locals {
		var localsInBlock Locals
		for local := range locals.Config {
			localsInBlock = append(localsInBlock, local)
		}
		sort.Strings(localsInBlock)
		result = append(result, localsInBlock)
	}

	return result
}

func (hclroot hclRoot) getProvidersInfomation() []Provider {
	var result []Provider

	for _, provider := range hclroot.Providers {
		var version string
		for _, arg := range provider.Default {
			if arg.Name == "version" {
				gohcl.DecodeExpression(arg.Expr, nil, &version)
			}
		}
		result = append(result, Provider{Name: provider.Name, Version: version})
	}

	return result
}

func (hclroot hclRoot) getModulesInfomation() []Module {
	var result []Module

	for _, module := range hclroot.Modules {
		var version string
		for _, arg := range module.Default {
			if arg.Name == "version" {
				gohcl.DecodeExpression(arg.Expr, nil, &version)
				break
			}
			if arg.Name == "source" {
				var source string
				gohcl.DecodeExpression(arg.Expr, nil, &source)
				if strings.Contains(source, "?href=") {
					version = strings.Split(source, "=")[1]
					break
				}
			}
		}
		result = append(result, Module{Name: module.Name, Version: version})
	}

	return result
}

func (hclroot hclRoot) getTerraformInfomation() Terraform {
	terraformConfig := hclroot.Terraform
	var version, backend string

	if terraformConfig == nil {
		return *new(Terraform)
	}

	if terraformConfig.RequiredVersion != nil {
		version = *terraformConfig.RequiredVersion
	}
	if terraformConfig.Backend != nil {
		backend = terraformConfig.Backend.Type
	}
	return Terraform{
		Version: version,
		Backend: backend,
	}
}
