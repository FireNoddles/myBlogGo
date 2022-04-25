CREATE TABLE `article`  (
                             `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                             `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                             `cid` bigint(20) UNSIGNED NOT NULL,
                             `state` bigint(20) NULL DEFAULT NULL,
                             `desc` varchar(200) NULL DEFAULT NULL,
                             `content` longtext NULL DEFAULT NULL,
                             `Img` varchar(200) NULL DEFAULT NULL,
                             `cmt_count` bigint(20) UNSIGNED NOT NULL,
                             `read_count` bigint(20) UNSIGNED NOT NULL,
                             `created_time` datetime(3) NULL DEFAULT NULL,
                             `updated_time` datetime(3) NULL DEFAULT NULL,
                            PRIMARY KEY (`id`) USING BTREE,
                             INDEX `idx_name`(`name`) USING BTREE
)ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;