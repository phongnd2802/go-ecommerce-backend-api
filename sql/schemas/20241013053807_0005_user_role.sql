-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user_role` (
    `user_id` INT NOT NULL,
    `role_id` INT NOT NULL,
    `assigned_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`, `role_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_role`;
-- +goose StatementEnd
