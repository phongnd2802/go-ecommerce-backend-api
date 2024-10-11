-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user_base` (
    `user_id` int AUTO_INCREMENT PRIMARY KEY,
    `user_account` varchar(255) NOT NULL,
    `user_password` varchar(255) NOT NULL,
    `user_salt` varchar(255) NOT NULL,
    `user_login_time` timestamp NULL DEFAULT NULL,
    `user_logout_time` timestamp NULL DEFAULT NULL,
    `user_login_ip` varchar(45) DEFAULT NULL,
    `user_created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `user_updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `unique_user_account` (`user_account`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_base`;
-- +goose StatementEnd
