-- name: InsertUserToken :execresult
INSERT INTO `user_token` (
    refresh_token,
    public_key,
    user_id,
    token_created_at,
    token_updated_at
) VALUES (?, ?, ?, NOW(), NOW());

