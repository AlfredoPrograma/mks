terraform {
  required_version = "1.15.8"
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "6.55"
    }
  }
}

provider "aws" {
  profile = "default"
  region = "us-east-1"

  default_tags {
    tags = {
      ManagedBy = "Terraform"
      App = "mks"
    }
  }
}
