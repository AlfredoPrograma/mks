output "container_registry_name" {
  value = aws_ecr_repository.this.name
  description = "Name assigned to the Container Registry"
}

output "container_registry_url" {
  value = aws_ecr_repository.this.repository_url
  description = "URL assigned to the Container Registry"
}
