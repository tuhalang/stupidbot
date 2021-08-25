-- name: GetSubjects :many
SELECT *
FROM subject
WHERE status = 1
ORDER BY order_number;