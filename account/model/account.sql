CREATE TABLE `accounts` (
    `name` varchar(255),
    `hash_password` varchar NOT NULL,
    `email` varchar(255) NOT NULL,
    PRIMARY KEY (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;