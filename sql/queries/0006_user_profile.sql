-- name: InsertUserProfileRegister :execresult
INSERT INTO `user_profile` (
    user_id,
    user_email
) VALUES (?, ?);

