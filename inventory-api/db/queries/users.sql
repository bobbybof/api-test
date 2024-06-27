-- name: CreateUser :one
INSERT INTO users (
    name,
    password,
    email
) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * 
FROM USERS
WHERE email = $1;