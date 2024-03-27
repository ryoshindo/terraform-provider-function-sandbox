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

output "arn_build_example" {
  value = provider::fs::arn_build("aws", "ec2", "us-east-1", "123456789012", "vpc/vpc-0e9801d129EXAMPLE")
}
