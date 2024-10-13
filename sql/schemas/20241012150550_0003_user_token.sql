-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user_token` (
    `token_id` INT AUTO_INCREMENT PRIMARY KEY,
    `refresh_token` VARCHAR(500) NOT NULL,
    `refresh_token_used` VARCHAR(500) DEFAULT NULL,
    `public_key` VARCHAR(500) NOT NULL,
    `user_id` INT NOT NULL,
    `token_created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `token_updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_token`;
-- +goose StatementEnd
