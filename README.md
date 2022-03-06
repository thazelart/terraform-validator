:warning: **This repository is not maintained anymore. Sorry for the inconvenience**.


[![Terraform-Validator](docs/source/_static/terraform-validator.svg)](https://thazelart.github.io/terraform-validator/)

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#validation) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)    
[![Documentation Status](https://readthedocs.org/projects/terraform-validator/badge/?version=latest)](https://terraform-validator.readthedocs.io/en/latest/?badge=latest) [![GoDoc](https://godoc.org/github.com/thazelart/terraform-validator?status.svg)](https://godoc.org/github.com/thazelart/terraform-validator)      
[![Build Status](https://travis-ci.com/thazelart/terraform-validator.svg?branch=main)](https://travis-ci.com/thazelart/terraform-validator) [![CodeCov](https://codecov.io/gh/thazelart/terraform-validator/branch/main/graph/badge.svg)](https://codecov.io/gh/thazelart/terraform-validator) [![Go Report Card](https://goreportcard.com/badge/github.com/thazelart/terraform-validator)](https://goreportcard.com/report/github.com/thazelart/terraform-validator)      
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/thazelart/terraform-validator.svg)](https://hub.docker.com/r/thazelart/terraform-validator) [![Docker Pulls](https://img.shields.io/docker/pulls/thazelart/terraform-validator)](https://hub.docker.com/r/thazelart/terraform-validator)                 

This tool will help you ensure that a terraform folder answer to your norms and conventions rules. This can be really useful in several cases :
* You're a team that want to have a clean and maintainable code.
* You're a lonely developer that develop a lot of modules and you want to have a certain consistency between them.               

**Features:**         
 * [x] make sure that the block names match a certain pattern.
 * [x] make sure that the code is properly dispatched. To do this you can decide what type of block can contain each file (for example output blocks must be in `outputs.tf`).
 * [x] ensure that mandatory `.tf` files are present.
 * [x] ensure that the terraform version has been defined.
 * [x] ensure that the providers' version has been defined.
 * [x] make sure that the variables and/or outputs blocks have the description argument filled in.
 * [x] layered terraform folders (test recursively).

:warning: **Terraform 0.12+ is supported only by the versions 2.0.0 and higher**.

## Documentation
Please find the full documentation [here (ReadTheDocs)](https://terraform-validator.readthedocs.io).

## Authors
[Thibault Hazelart](https://github.com/thazelart)                   
Logo by [Alexis Normand](https://github.com/alexis-n)

## License
[Apache 2.0](/LICENSE)
