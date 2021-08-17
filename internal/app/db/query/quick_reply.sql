-- name: GetByScriptCode :many
SELECT *
FROM quick_reply
WHERE script_code = $1
AND status = 1;