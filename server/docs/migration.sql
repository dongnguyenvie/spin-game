CREATE TABLE `users` (
    `id` int NOT NULL AUTO_INCREMENT,
    `email` varchar(255) DEFAULT NULL UNIQUE,
    `wallet_address` varchar(50) NOT NULL UNIQUE,
    `password` varchar(255) NOT NULL,
    `status` int(11) NOT NULL DEFAULT '1',
    `nonce` int(11) NOT NULL,
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    KEY `users_address_idx` (`wallet_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `wallets` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL UNIQUE,
    `balance` decimal(10,0) NOT NULL DEFAULT '0',
    `status` int(11) NOT NULL DEFAULT '1',
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    KEY `wallet_status_idx` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `transactions` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` varchar(255) NOT NULL,
    `wallet_id` varchar(255) NOT NULL,
    `type` int(11) NOT NULL,
    `package` int(11) NOT NULL,
    `credit` decimal(10,0) NOT NULL DEFAULT '0',
    `create_by` varchar(36) DEFAULT NULL,
    `status` int(11) NOT NULL DEFAULT '1',
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    KEY `transactions_status_idx` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `game_spin` (
    `id` int NOT NULL AUTO_INCREMENT,
    `package` varchar(255) NOT NULL,
    `status` int(11) NOT NULL DEFAULT '1',
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `game_spin_histories` (
    `id` int NOT NULL AUTO_INCREMENT,
    `result` int(11) NOT NULL,
    `status` int(11) NOT NULL DEFAULT '1',
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    KEY `game_spin_histories_result_idx` (`result`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;