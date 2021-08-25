-- name: ListQuestion :many
SELECT *
FROM question
WHERE topic_code = $1
and subject_code = $2
and status = 1
ORDER BY difficult
LIMIT $3 OFFSET $4;

-- name: GetRandomQuestion :one
SELECT *
FROM question
WHERE topic_code = $1 
AND subject_code = $2
and status = 1
ORDER BY random()
LIMIT 1;

-- name: GetQuestionById :one
SELECT *
FROM question
WHERE id = $1
and status = 1
LIMIT 1;