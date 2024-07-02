-- +migrate Up
CREATE TABLE
  IF NOT EXISTS `division` (
    id BIGINT auto_increment NOT NULL,
    public_id varchar(13) NOT NULL,
    name varchar(50) NOT NULL,
    created_at DATETIME (6) NOT NULL,
    updated_at DATETIME (6) NULL,
    CONSTRAINT ID_PK PRIMARY KEY (id),
    CONSTRAINT PUBLIC_ID_UNIQUE UNIQUE KEY (public_id),
    CONSTRAINT NAME_UNIQUE UNIQUE KEY (name)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS `division`;