-- name: CreateIntegration :one
INSERT INTO integrations (org_id, type, credentials)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetIntegration :one
SELECT * FROM integrations WHERE id=$1;

-- name: ListIntegrationsByOrg :many
SELECT * FROM integrations WHERE org_id=$1;
-- name: UpdateIntegration :one
UPDATE integrations SET type=$2, credentials=$3, updated_at=now()
WHERE id=$1 RETURNING *;