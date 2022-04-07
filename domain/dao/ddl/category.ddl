CREATE TABLE `category`  (
                         `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                         `created_time` datetime(3) NULL DEFAULT NULL,
                         `updated_time` datetime(3) NULL DEFAULT NULL,
                         `state` bigint(20) NULL DEFAULT NULL,
                         `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                         PRIMARY KEY (`id`) USING BTREE,
                         INDEX `idx_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;