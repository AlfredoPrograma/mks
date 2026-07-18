resource "aws_secretsmanager_secret" "config" {
  name = var.api_config_secret_name
}

resource "aws_secretsmanager_secret_version" "config" {
  secret_id = aws_secretsmanager_secret.config.id
  secret_string = jsonencode(var.api_config)
}

data "aws_iam_policy_document" "secrets_reader" {
  statement {
    effect    = "Allow"
    actions   = ["secretsmanager:GetSecretValue", "secretsmanager:DescribeSecret"]
    resources = [aws_secretsmanager_secret.config.arn]
  }
}

resource "aws_iam_user" "secrets_reader" {
  name = var.secrets_manager_user_name
}

resource "aws_iam_access_key" "secrets_reader" {
  user = aws_iam_user.secrets_reader.name
}

resource "aws_iam_user_policy" "secrets_reader" {
  name   = var.secrets_manager_user_name
  user   = aws_iam_user.secrets_reader.name
  policy = data.aws_iam_policy_document.secrets_reader.json
}
