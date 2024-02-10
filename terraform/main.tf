terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region = "eu-west-1"
}

data "aws_key_pair" "key_pair" {
  key_name = "laptop"
}

data "aws_vpc" "default" {
    default = true
}

resource "aws_key_pair" "key" {
    key_name = "pokemon_keypair"
    public_key = file("./keys.pub")
}

// Allow ssh access to the instance
resource "aws_security_group" "whos_that_pokemon_sg" {
  name = "whos_that_pokemon_sg"
  description = "Allow ssh access to the instance"
  vpc_id = data.aws_vpc.default.id

    ingress {
        description = "SSH from VPC"
        from_port = 22
        to_port = 22
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }

    ingress {
        description = "HTTP from VPC"
        from_port = 80
        to_port = 80
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }

    ingress {
        description = "HTTPS from VPC"
        from_port = 443
        to_port = 443
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }

    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
}

resource "aws_instance" "whos_that_pokemon_ec2" {
    ami = "ami-0fef2f5dd8d0917e8"
    instance_type = "t2.micro"
    key_name = resource.aws_key_pair.key.key_name
    availability_zone = "eu-west-1a"
    security_groups = [aws_security_group.whos_that_pokemon_sg.name]

    tags = {
        Name = "whos_that_pokemon_ec2"
    }
}

output "aws_public_ip" {
  value = aws_instance.whos_that_pokemon_ec2.public_ip
}

// Add an A record to the public DNS zone
# resource "aws_route53_zone" "whos_that_pokemon_zone" {
#   name = "whosthatpokemon.com"
# }