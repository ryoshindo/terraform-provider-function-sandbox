terraform {
  required_providers {
    hashicups = {
      source = "hashicorp.com/edu/function-sandbox"
    }
  }
}

provider "hashicups" {}

data "hashicups_coffees" "example" {}
