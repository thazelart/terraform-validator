/* here would be the provider definition */
provider "google" {
  version = "foo"
}

provider "aws" {
  version = "foo2"
  assume_role {
    role_arn = "role"
  }
}
