# terraform-validator [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#validation)

[![GoDoc](https://godoc.org/github.com/thazelart/terraform-validator?status.svg)](https://godoc.org/github.com/thazelart/terraform-validator) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)          
[![Build Status](https://travis-ci.com/thazelart/terraform-validator.svg?branch=master)](https://travis-ci.com/thazelart/terraform-validator) [![CodeCov](https://codecov.io/gh/thazelart/terraform-validator/branch/master/graph/badge.svg)](https://codecov.io/gh/thazelart/terraform-validator) [![Go Report Card](https://goreportcard.com/badge/github.com/thazelart/terraform-validator)](https://goreportcard.com/report/github.com/thazelart/terraform-validator)      
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/thazelart/terraform-validator.svg)](https://hub.docker.com/r/thazelart/terraform-validator) [![Docker Pulls](https://img.shields.io/docker/pulls/thazelart/terraform-validator)](https://hub.docker.com/r/thazelart/terraform-validator)                 

Terraform is a Go library that help you ensure that a terraform folder answer to your norms and conventions rules. This can be really useful in several cases :
* You're a team that want to have a clean and maintainable code
* You're a lonely developer that develop a lot of modules and you want to have a certain consistency between them               

**Features:**         
 * [x] ensure that the terraform blocknames follow the given pattern
 * [x] ensure that blocks are in a wanted files (for example output blocks must be in `outputs.tf`)
 * [x] ensure that mandatory `.tf` files are present
 * [x] ensure that a terraform version has been set
 * [x] ensure that the providers version has been set

**Next features:**                    
 * [ ] layered terraform folders (test recursively)
 * [ ] ensure Readme was updated (if you are using [terraform-docs](https://github.com/segmentio/terraform-docs))
 * [ ] ensure `terraform fmt` is ok

:warning: **Terraform 0.12+ is supported only by the versions 2.0.0 and higher**.

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
Even if a default configuration is provided by terraform validator, you can customize it in a `.terraform-validator.yaml` file.

[Full configuration documentation here](docs/Configuration.md)

## CI/CD integration
You can run directly terraform-validator inside you build pipeline thanks to the [terraform-validator docker image](https://hub.docker.com/r/thazelart/terraform-validator) !

[Full CI/CD documentation here](docs/CICD.md)

## Authors
[Thibault Hazelart](https://github.com/thazelart)

## License
[Apache 2.0](/LICENSE)
