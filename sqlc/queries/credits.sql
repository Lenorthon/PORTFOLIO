-- name: AddCreditEntry :one
INSERT INTO credits_ledger (org_id, execution_id, amount, reason)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOrgCredits :one
SELECT SUM(amount) as total_credits FROM credits_ledger WHERE org_id=$1;
-- name: GetOrgUsedCredits :one 