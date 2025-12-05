-- name: CreateAiRequest :one
INSERT INTO ai_requests (execution_id, node_id, provider, model, input, output, cost)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
-- name: GetAiRequest :one
SELECT * FROM ai_requests WHERE id = $1;