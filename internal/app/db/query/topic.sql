-- name: ListTopics :many
SELECT *
FROM topic
WHERE subject_code = $1
ORDER BY order_number;