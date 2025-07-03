variable "instance_type" {
  default = "t2.micro"
}

variable "ami_id" {
  default = "ami-0c55b159cbfafe1f0" # Amazon Linux 2 (you can update it)
}

locals {
  services = {
    "api-gateway"  = 80,
    "auth-service" = 50051,
    "pg-db"        = 5432,
    "mongo-db"     = 27017,
    "jwt-service"  = 6000
  }
}

resource "aws_instance" "authx_instances" {
  for_each = local.services

  ami                         = var.ami_id
  instance_type               = var.instance_type
  subnet_id                   = aws_subnet.authx_subnet.id
  security_groups             = [aws_security_group.authx_sg.id]
  associate_public_ip_address = true
  tags = {
    Name = each.key
  }

  user_data = <<-EOF
              #!/bin/bash
              sudo yum update -y
              echo "Launching ${each.key} on port ${each.value}" > /home/ec2-user/${each.key}.log
              # Install specific services per role later
              EOF
}

output "instance_ips" {
  value = {
    for name, instance in aws_instance.authx_instances :
    name => instance.public_ip
  }
}
