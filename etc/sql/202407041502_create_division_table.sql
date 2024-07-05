-- +migrate Up
CREATE TABLE
    IF NOT EXISTS `division` (
        `id` BIGINT auto_increment NOT NULL,
        `name` VARCHAR(50) NOT NULL,
        `is_deleted` BOOLEAN NOT NULL DEFAULT FALSE,
        `created_at` TIMESTAMP(6) NOT NULL,
        `updated_at` TIMESTAMP(6) NULL,
        `deleted_at` TIMESTAMP(6) NULL,
        PRIMARY KEY (`id`),
        UNIQUE KEY `name_unique` (`name`) USING BTREE
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS `division`;