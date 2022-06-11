# initialize terraform using 'terraform init'
# apply this terraform file using 'terraform apply -auto-approve'

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.5.0"
    }
  }
}

provider "aws" {
  shared_config_files      = [".aws/config.txt"]
  shared_credentials_files = [".aws/credentials.txt"]
}

## Create AMI for a Catalyst 9800 in AWS
resource "aws_instance" "app_server" {
  ami           = "ami-0b9a457b6bcf79b71" 
  instance_type = "c5.xlarge"

  tags = {
    Name = "9800-CL-Created-from-TF"
  }
}