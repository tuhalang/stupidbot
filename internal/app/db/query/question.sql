-- name: ListQuestion :many
SELECT *
FROM question
WHERE topic_code = $1
and subject_code = $2
ORDER BY difficult
LIMIT $3 OFFSET $4;