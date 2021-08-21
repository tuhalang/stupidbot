-- name: CreateSessionContext :one
INSERT INTO session_context (user_id, script_code)
VALUES ($1, $2)
RETURNING *;

-- name: GetLatestSession :many
SELECT *
FROM session_context
WHERE user_id = $1
AND valid_time > now()
ORDER BY valid_time desc;
