-- +migrate Up
CREATE TABLE
    IF NOT EXISTS `user` (
        `id` BIGINT auto_increment NOT NULL,
        `username` VARCHAR(50) NOT NULL,
        `email` VARCHAR(100) NOT NULL,
        `phone` VARCHAR(25) NOT NULL,
        `division_id` BIGINT NOT NULL,
        `password` VARCHAR(100) NOT NULL,
        `is_deleted` BOOLEAN NOT NULL DEFAULT FALSE,
        `created_at` TIMESTAMP(6) NOT NULL,
        `updated_at` TIMESTAMP(6) DEFAULT NULL,
        `deleted_at` TIMESTAMP(6) NULL,
        PRIMARY KEY (`id`),
        UNIQUE KEY `username_unique` (`username`) USING BTREE,
        UNIQUE KEY `email_unique` (`email`) USING BTREE,
        FOREIGN KEY `user_division_fk` (`division_id`) REFERENCES `division` (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS `user`;