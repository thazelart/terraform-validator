// Package hcl handle the parsing of your terraform file
package hcl

// GetBlockNamesByType return you a map[string][]string defining the blocktypes
// present in the given parsedfile and the blocknames for each blocktype
func (parsedfile ParsedFile) GetBlockNamesByType() map[string][]string {
	result := make(map[string][]string)
	for _, variable := range parsedfile.Blocks.Variables {
		result["variable"] = append(result["variable"], variable.Name)
	}
	for _, output := range parsedfile.Blocks.Outputs {
		result["output"] = append(result["output"], output.Name)
	}
	for _, module := range parsedfile.Blocks.Modules {
		result["module"] = append(result["module"], module.Name)
	}
	for _, provider := range parsedfile.Blocks.Providers {
		result["provider"] = append(result["provider"], provider.Name)
	}
	for _, resource := range parsedfile.Blocks.Resources {
		result["resource"] = append(result["resource"], resource.Name)
	}
	for _, data := range parsedfile.Blocks.Data {
		result["data"] = append(result["data"], data.Name)
	}
	for _, locals := range parsedfile.Blocks.Locals {
		for _, local := range locals {
			result["locals"] = append(result["locals"], local)
		}
	}
	if parsedfile.Blocks.Terraform != *new(Terraform) {
		result["terraform"] = []string{}
	}
	return result
}
