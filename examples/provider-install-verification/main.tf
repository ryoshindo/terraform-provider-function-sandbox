terraform {
  required_providers {
    fs = {
      source = "hashicorp.com/edu/function-sandbox"
    }
  }
}

provider "fs" {
  name = "example"
}

data "fs_hello_world" "example" {
  name = "Ryo Shindo"
}

output "message" {
  value = data.fs_hello_world.example.message
}
