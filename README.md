# terraform-validator [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#validation)

[![GoDoc](https://godoc.org/github.com/thazelart/terraform-validator?status.svg)](https://godoc.org/github.com/thazelart/terraform-validator) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)          
[![Build Status](https://travis-ci.com/thazelart/terraform-validator.svg?branch=master)](https://travis-ci.com/thazelart/terraform-validator) [![CodeCov](https://codecov.io/gh/thazelart/terraform-validator/branch/master/graph/badge.svg)](https://codecov.io/gh/thazelart/terraform-validator) [![Go Report Card](https://goreportcard.com/badge/github.com/thazelart/terraform-validator)](https://goreportcard.com/report/github.com/thazelart/terraform-validator)      
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/thazelart/terraform-validator.svg)](https://hub.docker.com/r/thazelart/terraform-validator) [![Docker Pulls](https://img.shields.io/docker/pulls/thazelart/terraform-validator)](https://hub.docker.com/r/thazelart/terraform-validator)                 

Terraform is a Go library that help you ensure that a terraform folder answer to your norms and conventions rules. This can be really useful in several cases :
* You're a team that want to have a clean and maintainable code
* You're a lonely developer that develop a lot of modules and you want to have a certain consistency between them               

**Features:**         
 * [x] make sure that the block names match a certain pattern
 * [x] s'assurer que le code est bien dispatché. Pour ça vous pouvez décider quel type de bloc peut contenir chaque fichier (for example output blocks must be in `outputs.tf`).
 * [x] ensure that mandatory `.tf` files are present
 * [x] ensure that the terraform version has been defined
 * [x] ensure that the providers' version has been defined
 * [x] make sure that the variables and/or outputs blocks have the description argument filled in.
 * [x] layered terraform folders (test recursively)

**Next features:**                    
 * [ ] possibility of prohibiting certain types of resources
 * [ ] ensure Readme was updated (using [terraform-docs](https://github.com/segmentio/terraform-docs) as a first step)
 * [ ] ensure `terraform fmt` is ok

:warning: **Terraform 0.12+ is supported only by the versions 2.0.0 and higher**.

## Install

Prerequisite: install [Go 1.11+](https://golang.org/).

### Get the last version from releases
You can [download from here](https://github.com/thazelart/terraform-validator/releases) the binary. move it into a directory in your `$PATH` to use it. For example:
```
mv ~/download/terraform-validator /usr/local/bin
```

### Install from code:
To add terraform-validator, clone this repository and then get :
then you can build it:
```bash
go build
```
move it into a directory in your `$PATH` to use it. For example:
```
mv terraform-validator /usr/local/bin
```

## Getting Started

Show help information:

``` bash
terraform-validator --help
```
Show terraform-validator version:

``` bash
terraform-validator --version
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
