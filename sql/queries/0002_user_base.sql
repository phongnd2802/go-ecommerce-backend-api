-- name: CheckUserBaseExists :one
SELECT COUNT(*)
FROM user_base
WHERE user_account = ?;

-- name: InsertUserBase :execresult
INSERT INTO `user_base` (
    user_account,
    user_password,
    user_salt,
    user_created_at,
    user_updated_at
) VALUES (?, ?, ?, NOW(), NOW());

-- name: GetUserBase :one
SELECT user_id, user_account, user_password, user_salt
FROM `user_base`
WHERE user_account = ?;

-- name: UpdateInfoLogin :exec
UPDATE `user_base`
SET user_login_time = NOW(), user_login_ip = ?
WHERE user_id = ?;

