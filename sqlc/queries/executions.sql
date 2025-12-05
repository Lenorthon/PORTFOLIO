-- name: CreateExecution :one
INSERT INTO executions (workflow_id, org_id, input, status)
VALUES ($1, $2, $3, 'queued')
RETURNING *;

-- name: GetExecution :one
SELECT * FROM executions WHERE id = $1;

-- name: ListExecutionsByWorkflow :many
SELECT * FROM executions WHERE workflow_id = $1 ORDER BY created_at DESC;

-- name: UpdateExecutionStatus :one
UPDATE executions
SET status=$2, output=$3, updated_at=now()
WHERE id=$1
RETURNING *;

-- name: CreateNodeLog :one
INSERT INTO nodes_logs (execution_id, node_id, status, input)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateNodeLog :one
UPDATE nodes_logs
SET status=$2, output=$3, error=$4, updated_at=now()
WHERE id=$1
RETURNING *;
-- name: ListNodeLogsByExecution :many
SELECT * FROM nodes_logs WHERE execution_id = $1 ORDER BY created_at ASC; 