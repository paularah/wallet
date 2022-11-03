-- name: CreateFunding :one 
INSERT INTO fundings (user_id, wallet_id, amount, success) 
VALUES ($1, $2, $3, $4)
RETURNING *; 

-- name: GetFunding :one 
SELECT * FROM entries 
WHERE id = $1 LIMIT 1;

-- name: GetWalletFunding :one 
SELECT * FROM entries 
WHERE wallet_id = $1 
ORDER BY id 
LIMIT $2
OFFSET $3;

-- name: ListWalletFunding :one 
SELECT * FROM entries 
WHERE wallet_id = $1 LIMIT 1;