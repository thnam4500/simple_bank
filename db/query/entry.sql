-- name: CreateEntry :one
INSERT INTO entries (
  account_id, amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListEntriesByAccount :many
SELECT * FROM entries
WHERE account_id = $1
ORDER by created_at DESC
LIMIT $2
OFFSET $3;
