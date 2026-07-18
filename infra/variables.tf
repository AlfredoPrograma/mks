variable "container_repository_name" {
  type        = string
  description = "Name of the container repository"
  default     = "mks"
}

variable "secrets_manager_user_name" {
  type        = string
  description = "Name of the user used for manage Secrets Manager for external secrets"
  default     = "mks_secrets_reader"
}

variable "api_config_secret_name" {
  type = string
  description = "Name of the configuration secret"
  default = "mks_api_config"
}

variable "api_config" {
  type = map(string)
  description = "Key/pair values that are going to be injected to MKS API"
}
