-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id, to_account_id, amount)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTransfer :one
SELECT *  FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListAccountTransfers :many
SELECT * FROM transfers
WHERE to_account_id = $1
OR from_account_id = $1
LIMIT $2
OFFSET $3;

-- name: ListDebitTransfers :many
SELECT * FROM transfers
WHERE from_account_id = $1
LIMIT $2
OFFSET $3;

-- name: ListCreditTransfers :many
SELECT * FROM transfers
WHERE to_account_id = $1
LIMIT $2
OFFSET $3;

-- name: UpdateTransfer :one
UPDATE transfers
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;
