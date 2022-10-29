-- name: CreateWalletEntry :one 
INSERT INTO entries (wallet_id, amount) 
VALUES ($1, $2)
RETURNING *; 

-- name: GetWalletEntry :one 
SELECT * FROM entries 
WHERE id = $1 LIMIT 1;

-- name: ListWalletEntries :one 
SELECT * FROM entries 
WHERE wallet_id = $1 
ORDER BY id 
LIMIT $2 
OFFSET $3;