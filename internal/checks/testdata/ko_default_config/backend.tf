terraform {
  required_version = ">=0.12"

  required_providers {
    aws = {
      version = ">= 2.7.0"
      source = "hashicorp/aws"
    }
    gcp = {
      source = "hashicorp/gcp"
    }
    newrelic = "~> 1.19"
  }
}
