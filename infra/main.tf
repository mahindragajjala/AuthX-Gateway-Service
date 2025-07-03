provider "aws" {
  region  = var.aws_region
  profile = "default"  # or your profile name like "authx-dev"
}


resource "aws_vpc" "authx_vpc" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "authx_subnet" {
  vpc_id            = aws_vpc.authx_vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-east-1a"
}

resource "aws_security_group" "authx_sg" {
  name        = "authx-sg"
  vpc_id      = aws_vpc.authx_vpc.id
  description = "Allow HTTP, gRPC, DB traffic"

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # API Gateway
  }

  ingress {
    from_port   = 50051
    to_port     = 50051
    protocol    = "tcp"
    cidr_blocks = ["10.0.1.0/24"] # gRPC
  }

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["10.0.1.0/24"] # PostgreSQL
  }

  ingress {
    from_port   = 27017
    to_port     = 27017
    protocol    = "tcp"
    cidr_blocks = ["10.0.1.0/24"] # MongoDB
  }

  ingress {
    from_port   = 6000
    to_port     = 6000
    protocol    = "tcp"
    cidr_blocks = ["10.0.1.0/24"] # JWT Service
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
