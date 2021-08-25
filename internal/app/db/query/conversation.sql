-- name: CreateConversation :one
INSERT INTO conversation(user_id, mentor_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetConversationByUser :one
SELECT *
FROM conversation
WHERE user_id = $1
AND status = 1
AND valid_time > now();

-- name: GetConversationByMentor :one
SELECT *
FROM conversation
WHERE mentor_id = $1
AND status = 1
AND valid_time > now();

-- name: InvalidConversation :exec
UPDATE conversation
SET status = 0
WHERE id = $1;