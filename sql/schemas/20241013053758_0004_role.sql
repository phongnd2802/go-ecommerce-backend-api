-- +goose Up
-- +goose StatementBegin
CREATE TABLE `role` (
    `role_id` int AUTO_INCREMENT PRIMARY KEY,
    `role_name` VARCHAR(50) NOT NULL,
    `role_slug` VARCHAR(255) NOT NULL,
    `role_description` TEXT DEFAULT NULL,
    `role_created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `role_updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `unique_role_name` (`role_name`),
    UNIQUE KEY `unique_role_slug` (`role_slug`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `role`;
-- +goose StatementEnd
