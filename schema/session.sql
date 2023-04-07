CREATE TABLE `sessions` (
    `id` bigint,
    `user_name` varchar(255) NOT NULL,
    `refesh_token` varchar(255) NOT NULL,
    `user_agent` varchar(255) NOT NULL,
    `client_ip` varchar(255) NOT NULL,
    `is_blocked` boolean NOT NULL default false,
    `expires_at` bigint NOT NULL,
    `created_at` bigint NOT NULL default now(),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;