-- name: GetUserByID :one
SELECT id, email FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email;
