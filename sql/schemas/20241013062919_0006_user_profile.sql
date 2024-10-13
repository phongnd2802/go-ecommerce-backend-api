-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user_profile` (
    `user_id` int PRIMARY KEY,
    `user_email` varchar(255) NOT NULL,
    `user_nickname` varchar(255) DEFAULT NULL,
    `user_avatar` varchar(255) DEFAULT NULL,
    `user_state` tinyint unsigned NOT NULL DEFAULT 2,
    `user_mobile` varchar(20) DEFAULT NULL,
    `user_gender` tinyint unsigned DEFAULT NULL,
    `user_birthday` date DEFAULT NULL,
    `cretead_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `unique_user_email` (`user_email`),
    KEY `idx_user_mobile` (`user_mobile`),
    KEY `idx_user_email` (`user_email`),
    KEY `idx_user_state` (`user_state`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_profile`;
-- +goose StatementEnd
