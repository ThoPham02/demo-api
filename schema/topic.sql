CREATE TABLE `topics` (
    `id` bigint,
    `name` varchar(255) NOT NULL,
    `description` varchar(255) NOT NULL,
    `file_url` varchar(255),
    `create_by` varchar(255) NOT NULL,
    `created_at` bigint NOT NULL default now(),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;