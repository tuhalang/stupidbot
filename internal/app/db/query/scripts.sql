-- name: GetScriptByCode :one
SELECT  *
FROM scripts
WHERE code = $1
LIMIT 1;

-- name: GetScriptByParent :many
SELECT *
FROM scripts
WHERE parent_code = $1;
