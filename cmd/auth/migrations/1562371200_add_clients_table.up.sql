START TRANSACTION;
CREATE TABLE IF NOT EXISTS `clients`
(
    `distinct_id` INT          NOT NULL AUTO_INCREMENT,
    `id`          CHAR(36)     NOT NULL,
    `user_id`     CHAR(36)     NOT NULL,
    `secret`      VARCHAR(255) NOT NULL,
    `domain`      VARCHAR(255) NOT NULL,
    `data`        JSON         NOT NULL,
    PRIMARY KEY (`distinct_id`),
    UNIQUE KEY `id` (`id`),
    INDEX `i_user_id` (`user_id`)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8
    COLLATE = utf8_bin;
COMMIT;
