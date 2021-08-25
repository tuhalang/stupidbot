-- name: CreateUser :one
INSERT INTO users (id, is_mentor, status)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetMentorAvailable :one
SELECT *
FROM users u
WHERE u.is_mentor = 1
AND u.status = 1
AND not exists (SELECT 1 FROM conversation c WHERE c.mentor_id = u.id and c.status = 1 and c.valid_time > now())
LIMIT 1;

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