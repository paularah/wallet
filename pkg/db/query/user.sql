-- name: CreateUser :one
INSERT INTO users (username, email, firstname, lastname, password) 
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUser :one 
SELECT * FROM users
WHERE id = $1 LIMIT 1;
