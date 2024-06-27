-- name: CreateProduct :one
INSERT INTO products(
    name,
    price,
    description
) VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
SET
    name = COALESCE(sqlc.narg(name), name),
    price = COALESCE(sqlc.narg(price), price),
    description = COALESCE(sqlc.narg(description), description)
WHERE
    id = $1
RETURNING *;


-- name: CountProduct :one
SELECT COUNT(id) FROM products;