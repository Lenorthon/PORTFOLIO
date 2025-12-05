-- name: CreateUser :one
INSERT INTO users (org_id, email, name, password_hash, role)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: ListUsersByOrg :many
SELECT * FROM users WHERE org_id = $1 ORDER BY created_at ASC;
-- name: UpdateUser :one
UPDATE users SET name=$2, password_hash=$3, role=$4, updated_at=now()
WHERE id=$1 RETURNING *;   