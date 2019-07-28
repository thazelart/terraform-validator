# terraform-validator [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#validation)

[![GoDoc](https://godoc.org/github.com/thazelart/terraform-validator?status.svg)](https://godoc.org/github.com/thazelart/terraform-validator) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)          
[![Build Status](https://travis-ci.com/thazelart/terraform-validator.svg?branch=master)](https://travis-ci.com/thazelart/terraform-validator) [![CodeCov](https://codecov.io/gh/thazelart/terraform-validator/branch/master/graph/badge.svg)](https://codecov.io/gh/thazelart/terraform-validator) [![Go Report Card](https://goreportcard.com/badge/github.com/thazelart/terraform-validator)](https://goreportcard.com/report/github.com/thazelart/terraform-validator)      
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/thazelart/terraform-validator.svg)](https://hub.docker.com/r/thazelart/terraform-validator) [![Docker Pulls](https://img.shields.io/docker/pulls/thazelart/terraform-validator)](https://hub.docker.com/r/thazelart/terraform-validator)                 

Terraform is a Go library that help you ensure that a terraform folder answer to your norms and conventions rules.
**Features:**         
 * [x] ensure that the terraform blocknames follow the given pattern             
 * [x] ensure that blocks are in a wanted files (for example output blocks must be in `outputs.tf`)               
 * [x] ensure that mandatory `.tf` files are present               
 * [x] ensure that a terraform version has been set               
 * [x] ensure that the providers version has been set           

**Next features:**                    
 * [ ] ensure Readme was updated (if you are using [terraform-docs](https://github.com/segmentio/terraform-docs))               
 * [ ] ensure `terraform fmt` is ok               

## Install

Prerequisite: install [Go 1.11+](https://golang.org/).

### Install from code:
To add terraform-validator, I recommend using a Go dependency manager such as
[dep](https://github.com/golang/dep):

```bash
go mod init github.com/thazelart/terraform-validator
```

then you can install it :

```bash
go mod download
go install
```

### Get the last version from releases
You can [download from here](https://github.com/thazelart/terraform-validator/releases) the binary. move it into a directory in your `$PATH` to use it:
```
mv ~/download/terraform-validator /usr/local/bin
```

## Getting Started

Show help information:

``` bash
terraform-validator --help
```

Validate a terraform folder located in `./examples`:

```bash
terraform-validator ./examples
```

## Configuration
The [default](/internal/config/default_config.yaml) configuration is:
```yaml
---
files:
  main.tf:
    mandatory: true
    authorized_blocks:
  variables.tf:
    mandatory: true
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
ensure_terraform_version: true
ensure_providers_version: true
ensure_readme_updated: true
block_pattern_name: "^[a-z0-9_]*$"
```

You can set you own configuration by adding a `.terraform-validator.yaml` file at the root of your terraform folder.
For example :
```yaml
---
files:
  # you don't have any rule about files, every file can have every block type
  default:
    authorized_blocks:
      - variable
      - output
      - provider
      - terraform
      - resource
      - module
      - data
      - locals
# you don't want to check the terraform version
ensure_terraform_version: false
# you prefere kebab case for terraform blocknames
block_pattern_name: "^[a-z0-9-]*$"
```
The non present parameters in your `.terraform-validator.yaml` will take the default value.

## CI/CD integration
You can run directly terraform-validator inside you build pipeline thanks to the [terraform-validator docker image](https://hub.docker.com/r/thazelart/terraform-validator) !
This docker file run automatically terraform in the root directory. If you want to run it in your build pipeline, you can configure if very easily.                  

For example, in cloud build :
```yaml
steps:
- name: 'thazelart/terraform-validator'
```

## Authors
[Thibault Hazelart](https://github.com/thazelart)

## License
[Apache 2.0](/LICENSE)
