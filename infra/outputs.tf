output "container_registry_name" {
  value = aws_ecr_repository.this.name
  description = "Name assigned to the Container Registry"
}

output "container_registry_repository_url" {
  value = aws_ecr_repository.this.repository_url
  description = "URL assigned to the Container Registry Repository"
}

output "workflows_role_arn" {
  value = aws_iam_role.workflows.arn
  description = "ARN of the role assumed over workflows"
}
