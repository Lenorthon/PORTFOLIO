-- name: CreateWorkflow :one
INSERT INTO workflows (org_id, name, description, definition, version)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetWorkflow :one
SELECT * FROM workflows WHERE id = $1;

-- name: ListWorkflowsByOrg :many
SELECT * FROM workflows WHERE org_id = $1 ORDER BY created_at ASC;

-- name: UpdateWorkflow :one
UPDATE workflows
SET name=$2, description=$3, definition=$4, version=$5, updated_at=now()
WHERE id=$1
RETURNING *;

-- name: DeleteWorkflow :exec
DELETE FROM workflows WHERE id=$1; 
