terraform {
  required_version = ">= 1.5.0"

  backend "s3" {
    bucket         = "bnt-web-hetzner-tf-state"
    key            = "bnt/ops/web/hetzner/terraform.tfstate"
    region         = "us-east-2"
    encrypt        = true
  }

}

