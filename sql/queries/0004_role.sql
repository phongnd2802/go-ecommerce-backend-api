-- name: InsertRole :exec
INSERT INTO `role` (
    role_name,
    role_slug,
    role_description
) VALUES (?, ?, ?);

