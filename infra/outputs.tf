output "secrets_reader_access_key_id" {
  value = aws_iam_access_key.secrets_reader.id
  description = "Access key ID for secrets reader user"
  sensitive = true
}

output "secrets_reader_secret_access_key" {
  value = aws_iam_access_key.secrets_reader.secret
  description = "Secret access key for secrets reader user"
  sensitive = true
}

output "api_config_secret_name" {
  value = aws_secretsmanager_secret.config.name
  description = "Name assigned to the API config secret"
}

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
