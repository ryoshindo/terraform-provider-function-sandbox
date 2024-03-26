terraform {
  required_providers {
    fs = {
      source = "hashicorp.com/edu/function-sandbox"
    }
  }
}

provider "fs" {}

data "fs_sample" "example" {}
