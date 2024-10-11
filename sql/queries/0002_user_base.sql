-- name: CheckUserBaseExists :one
SELECT COUNT(*)
FROM user_base
WHERE user_account = ?;