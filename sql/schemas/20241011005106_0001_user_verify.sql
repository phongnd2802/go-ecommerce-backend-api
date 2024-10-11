-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user_verify` (
    `verify_id` int AUTO_INCREMENT PRIMARY KEY,
    `verify_otp` varchar(6) NOT NULL,
    `verify_key` varchar(255) NOT NULL,
    `verify_key_hash` varchar(255) NOT NULL,
    `verify_type` int DEFAULT '1',
    `is_verified` int DEFAULT '0',
    `is_deleted` int DEFAULT '0',
    `verify_created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `verify_updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `unique_verify_key` (`verify_key`),
    KEY `idx_verify_otp` (`verify_otp`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_verify`;
-- +goose StatementEnd
