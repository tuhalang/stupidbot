-- name: CreateUser :one
INSERT INTO users (id, user_name, email, status)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
    AND status = 1
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
WHERE status = 1
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: ListAllUsers :many
SELECT *
FROM users
where status = 1
ORDER BY id;

-- name: UpdateUserStatus :one 
UPDATE users
SET STATUS = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserEmail :one 
UPDATE users
SET email = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserPhone :one 
UPDATE users
SET phone = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec 
DELETE FROM users
WHERE id = $1;