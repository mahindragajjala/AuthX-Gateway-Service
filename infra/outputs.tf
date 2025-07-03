output "vpc_id" {
  description = "ID of the created VPC"
  value       = aws_vpc.authx_vpc.id
}

output "subnet_id" {
  description = "ID of the created subnet"
  value       = aws_subnet.authx_subnet.id
}

output "security_group_id" {
  description = "ID of the created security group"
  value       = aws_security_group.authx_sg.id
}

output "instance_public_ips" {
  description = "Public IPs of all EC2 instances by name"
  value = {
    for name, instance in aws_instance.authx_instances :
    name => instance.public_ip
  }
}

output "instance_private_ips" {
  description = "Private IPs of all EC2 instances by name"
  value = {
    for name, instance in aws_instance.authx_instances :
    name => instance.private_ip
  }
}
