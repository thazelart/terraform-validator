---
id: layers
title: Layers
sidebar_label: Layers
---

Layers is a map of layer. By default, layers contains one layer named `default` :
```yaml
layers:
  default:
    files:
      main.tf:
        mandatory: true
        authorized_blocks:
      variables.tf:
        mandatory: true,
        authorized_blocks:
          - variable
      outputs.tf:
        mandatory: true
        authorized_blocks:
          - output
      providers.tf:
        mandatory: true
        authorized_blocks:
          - provider
      backend.tf:
        mandatory: true
        authorized_blocks:
          - terraform
      default:
        mandatory: false
        authorized_blocks:
          - resource
          - module
          - data
          - locals
    ensure_terraform_version: false
    ensure_providers_version: false
    ensure_variables_description: false
    ensure_outputs_description: false
    block_pattern_name: "^[a-z0-9_]*$"
```
## What is a layer ?
A layer is a set of parameter that define the complete configuration of terraform-validator.

## files
files is a map of filename that contains what rules define each files.
* `mandatory`: set if the file is mandatory.                            
  *Type: `boolean`*                          
  *Default: `false`*
* `authorized_blocks`: is a list of authorized terraform block types.                          
  *Type: `List`*                    
  *Default: `<empty>`*                    
  available terraform block types:
    * variable
    * output
    * provider
    * terraform
    * resource
    * module
    * data
    * locals

If a match does not match exactly one of the files, the configuration will be taken from `default`.

```yaml
# .terraform-validator.yaml
...
      # main.tf is mandatory with no block inside
      main.tf:
        mandatory: true
        authorized_blocks:
      # variables.tf is not mandatory with only variable blocks inside
      variables.tf:
        mandatory: true,
        authorized_blocks:
          - variable
      # other files will match default config
      default:
        mandatory: false
        authorized_blocks:
          - resource
          - module
          - data
          - locals
...
```

## ensure_terraform_version
*Type: `boolean`*                 
*Default: `false`*                  
Configure terraform-validator in order to ensure (or not) if the terraform version has been set.

## ensure_providers_version
*Type: `boolean`*                 
*Default: `false`*          
Configure terraform-validator in order to ensure (or not) if the providers versions has been set.


## ensure_variables_description
*Type: `boolean`*                 
*Default: `false`*          
configures terraform-validator to check whether or not the variable blocks are described.

## ensure_outputs_description
*Type: `boolean`*                 
*Default: `false`*           
configures terraform-validator to check whether or not the output blocks are described.

## block_pattern_name
*Type: `string`*                 
*Default: `^[a-z0-9_]+$` *    
Configure the pattern that should match each terraform resources.
