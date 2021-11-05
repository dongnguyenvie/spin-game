CREATE TABLE `users` (
    `id` varchar(36) NOT NULL,
    `fullname` varchar(255) DEFAULT NULL,
    `wallet_address` varchar(36) NOT NULL,
    `password` varchar(255) NOT NULL,
    `status` int(11) NOT NULL DEFAULT '1',
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    KEY `users_address_idx` (`wallet_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `wallet` (
    `id` varchar(36) NOT NULL,
    `balance` int(11) NOT NULL DEFAULT '1',
    `status` int(11) NOT NULL DEFAULT '1',
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    KEY `wallet_status_idx` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `transactions` (
    `id` varchar(36) NOT NULL,
    `user_id` varchar(255) NOT NULL,
    `wallet_id` varchar(255) NOT NULL,
    `type` int(11) NOT NULL,
    `description` text,
    `credit` decimal(10,0) NOT NULL DEFAULT '0',
    `debit` decimal(10,0) NOT NULL DEFAULT '0',
    `create_by` varchar(36) DEFAULT NULL,
    `status` int(11) NOT NULL DEFAULT '1',
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    KEY `transactions_status_idx` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `game_spin` (
    `id` varchar(36) NOT NULL,
    `quantity` varchar(255) NOT NULL,
    `status` int(11) NOT NULL DEFAULT '1',
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `game_spin_histories` (
    `id` varchar(36) NOT NULL,
    `quantity` varchar(255) NOT NULL,
    `result` int(11) NOT NULL,
    `status` int(11) NOT NULL DEFAULT '1',
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    KEY `game_spin_histories_result_idx` (`result`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;