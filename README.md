# terraform-validator
[![GoDoc](https://godoc.org/github.com/thazelart/terraform-validator?status.svg)](https://godoc.org/github.com/thazelart/terraform-validator) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)          
[![Build Status](https://travis-ci.com/thazelart/terraform-validator.svg?branch=master)](https://travis-ci.com/thazelart/terraform-validator) [![CodeCov](https://codecov.io/gh/thazelart/terraform-validator/branch/master/graph/badge.svg)](https://codecov.io/gh/thazelart/terraform-validator) [![Go Report Card](https://goreportcard.com/badge/github.com/thazelart/terraform-validator)](https://goreportcard.com/report/github.com/thazelart/terraform-validator)

Terraform is a Go library that help you ensure that a terraform folder answer to your normes and conventions rules. The features that are developped (:heavy_check_mark:), in progress (:construction_worker:) or ready to dev (:x:) are :             
:heavy_check_mark: ensure that the terraform block name follow the given pattern             
:heavy_check_mark: ensure that blocks are in a wanted files (for example output blocks must be in `outputs.tf`)               
:heavy_check_mark: ensure that mandatory `.tf` files are present               
:x: ensure that a terraform version has been set               
:x: ensure that the providers version has been set               
:x: ensure Readme was updated (if you are using [terraform-docs](https://github.com/segmentio/terraform-docs))               
:x: ensure `terraform fmt` is ok               

## Install

Prerequisite: install [Go](https://golang.org/).

To add terraform-validator, we recommend using a Go dependency manager such as
[dep](https://github.com/golang/dep):

```bash
dep ensure -add github.com/thazelart/terraform-validator
```

Alternatively, you can use `go get`:

```bash
go get github.com/thazelart/terraform-validator
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
[The default configuration is: ](/internal/config/default_config.yaml)
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
  provider.tf:
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

You can set you own configuration by adding a `terraform-validator.yaml` file at the root of your terraform folder.
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
# you prefere snake case for terraform block names
block_pattern_name: "^[a-z0-9_]*$"
```
The non present parameters in your `terraform-validator.yaml` will take the default value.

## Authors
[Thibault Hazelart](https://github.com/thazelart)

## License
[Apache 2.0](/LICENSE)
