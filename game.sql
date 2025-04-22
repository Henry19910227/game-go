SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '用戶id',
    `user_id` INT UNSIGNED UNIQUE NOT NULL DEFAULT 0 COMMENT '用戶id(外部)',
    `password` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '密碼',
    `nickname` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '暱稱',
    `score` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '帳戶餘額',
    `is_deleted` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否已被刪除',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '修改日期',
    INDEX `index_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用戶表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` (user_id, password, nickname, score) VALUES (10001, 'Cv4XkLoA', 'chase', 1000000);
INSERT INTO `users` (user_id, password, nickname, score) VALUES (10002, 'Cv4XkLoA', 'Henry', 1000000);

-- ----------------------------
-- Table structure for games
-- ----------------------------
DROP TABLE IF EXISTS `games`;
CREATE TABLE IF NOT EXISTS `games` (
    `id` INT UNSIGNED PRIMARY KEY COMMENT '遊戲id',
    `name` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '遊戲名稱',
    `type` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '遊戲type',
    `icon` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '遊戲icon',
    `is_deleted` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否已被刪除',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '修改日期'
) ENGINE=InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '遊戲表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `games` (id, name) VALUES (1001, '一分快三');
INSERT INTO `games` (id, name) VALUES (1002, '賽車');
INSERT INTO `games` (id, name) VALUES (1003, '時時彩');
INSERT INTO `games` (id, name) VALUES (1004, '六合彩');
INSERT INTO `games` (id, name) VALUES (1005, 'PC28');
INSERT INTO `games` (id, name) VALUES (1006, '百家樂');
INSERT INTO `games` (id, name) VALUES (1007, '龍虎');
INSERT INTO `games` (id, name) VALUES (1008, '牛牛');
INSERT INTO `games` (id, name) VALUES (1009, '輪盤');
INSERT INTO `games` (id, name) VALUES (1010, '魚蝦蟹');

-- ----------------------------
-- Table structure for bet_areas
-- ----------------------------
DROP TABLE IF EXISTS `bet_areas`;
CREATE TABLE IF NOT EXISTS `bet_areas` (
    `id` INT UNSIGNED COMMENT '注區id',
    `game_id` INT UNSIGNED COMMENT '遊戲id(關聯)',
    `name` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '注區名稱',
    `min_limit` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '最小限額',
    `max_limit` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '最大限額',
    `is_deleted` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否已被刪除',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '修改日期',
    PRIMARY KEY (id, game_id)
) ENGINE=InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '注區表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of bet_areas    正則表達式去除多餘訊息 ^2024-12-05.*\[24554:25574523\]
-- ----------------------------
-- 一分三快 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 1001, '大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 1001, '小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 1001, '单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 1001, '双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 1001, '总和4点', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 1001, '总和5点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 1001, '总和6点', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 1001, '总和7点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (9, 1001, '总和8点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (10, 1001, '总和9点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (11, 1001, '总和10点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (12, 1001, '总和11点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (13, 1001, '总和12点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (14, 1001, '总和13点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (15, 1001, '总和14点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (16, 1001, '总和15点', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (17, 1001, '总和16点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (18, 1001, '总和17点', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (19, 1001, '单骰1', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (20, 1001, '单骰2', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (21, 1001, '单骰3', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (22, 1001, '单骰4', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (23, 1001, '单骰5', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (24, 1001, '单骰6', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (25, 1001, '对子1', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (26, 1001, '对子2', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (27, 1001, '对子3', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (28, 1001, '对子4', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (29, 1001, '对子5', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (30, 1001, '对子6', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (31, 1001, '豹子1', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (32, 1001, '豹子2', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (33, 1001, '豹子3', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (34, 1001, '豹子4', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (35, 1001, '豹子5', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (36, 1001, '豹子6', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (37, 1001, '豹子1-6', 1000, 500000);

-- 賽車 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 1002, '和大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 1002, '和小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 1002, '和单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 1002, '和双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 1002, '3', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 1002, '4', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 1002, '5', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 1002, '6', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (9, 1002, '7', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (10, 1002, '8', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (11, 1002, '9', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (12, 1002, '10', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (13, 1002, '11', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (14, 1002, '12', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (15, 1002, '13', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (16, 1002, '14', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (17, 1002, '15', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (18, 1002, '16', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (19, 1002, '17', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (20, 1002, '18', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (21, 1002, '19', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (22, 1002, '冠军大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (23, 1002, '冠军小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (24, 1002, '冠军单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (25, 1002, '冠军双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (26, 1002, '冠军1', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (27, 1002, '冠军2', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (28, 1002, '冠军3', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (29, 1002, '冠军4', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (30, 1002, '冠军5', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (31, 1002, '冠军6', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (32, 1002, '冠军7', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (33, 1002, '冠军8', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (34, 1002, '冠军9', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (35, 1002, '冠军10', 1000, 1000000);

-- 時時彩 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 1003, '总和-大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 1003, '总和-小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 1003, '总和-单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 1003, '总和-双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 1003, '球1-大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 1003, '球1-小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 1003, '球1-单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 1003, '球1-双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (9, 1003, '球2-大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (10, 1003, '球2-小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (11, 1003, '球2-单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (12, 1003, '球2-双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (13, 1003, '球3-大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (14, 1003, '球3-小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (15, 1003, '球3-单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (16, 1003, '球3-双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (17, 1003, '球4-大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (18, 1003, '球4-小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (19, 1003, '球4-单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (20, 1003, '球4-双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (21, 1003, '球5-大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (22, 1003, '球5-小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (23, 1003, '球5-单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (24, 1003, '球5-双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (25, 1003, '球1-0号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (26, 1003, '球1-1号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (27, 1003, '球1-2号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (28, 1003, '球1-3号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (29, 1003, '球1-4号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (30, 1003, '球1-5号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (31, 1003, '球1-6号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (32, 1003, '球1-7号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (33, 1003, '球1-8号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (34, 1003, '球1-9号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (35, 1003, '球2-0号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (36, 1003, '球2-1号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (37, 1003, '球2-2号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (38, 1003, '球2-3号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (39, 1003, '球2-4号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (40, 1003, '球2-5号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (41, 1003, '球2-6号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (42, 1003, '球2-7号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (43, 1003, '球2-8号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (44, 1003, '球2-9号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (45, 1003, '球3-0号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (46, 1003, '球3-1号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (47, 1003, '球3-2号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (48, 1003, '球3-3号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (49, 1003, '球3-4号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (50, 1003, '球3-5号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (51, 1003, '球3-6号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (52, 1003, '球3-7号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (53, 1003, '球3-8号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (54, 1003, '球3-9号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (55, 1003, '球4-0号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (56, 1003, '球4-1号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (57, 1003, '球4-2号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (58, 1003, '球4-3号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (59, 1003, '球4-4号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (60, 1003, '球4-5号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (61, 1003, '球4-6号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (62, 1003, '球4-7号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (63, 1003, '球4-8号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (64, 1003, '球4-9号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (65, 1003, '球5-0号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (66, 1003, '球5-1号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (67, 1003, '球5-2号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (68, 1003, '球5-3号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (69, 1003, '球5-4号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (70, 1003, '球5-5号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (71, 1003, '球5-6号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (72, 1003, '球5-7号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (73, 1003, '球5-8号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (74, 1003, '球5-9号', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (75, 1003, '前三-豹子', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (76, 1003, '前三-顺子', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (77, 1003, '前三-对子', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (78, 1003, '前三-半顺', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (79, 1003, '前三-杂六', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (80, 1003, '中三-豹子', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (81, 1003, '中三-顺子', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (82, 1003, '中三-对子', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (83, 1003, '中三-半顺', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (84, 1003, '中三-杂六', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (85, 1003, '后三-豹子', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (86, 1003, '后三-顺子', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (87, 1003, '后三-对子', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (88, 1003, '后三-半顺', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (89, 1003, '后三-杂六', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (90, 1003, '龙', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (91, 1003, '和', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (92, 1003, '虎', 1000, 5000000);

-- 六合彩 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 1004, '大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 1004, '小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 1004, '单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 1004, '双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 1004, '正1', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 1004, '正2', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 1004, '正3', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 1004, '正4', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (9, 1004, '正5', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (10, 1004, '正6', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (11, 1004, '正7', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (12, 1004, '正8', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (13, 1004, '正9', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (14, 1004, '正10', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (15, 1004, '正11', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (16, 1004, '正12', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (17, 1004, '正13', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (18, 1004, '正14', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (19, 1004, '正15', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (20, 1004, '正16', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (21, 1004, '正17', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (22, 1004, '正18', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (23, 1004, '正19', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (24, 1004, '正20', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (25, 1004, '正21', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (26, 1004, '正22', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (27, 1004, '正23', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (28, 1004, '正24', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (29, 1004, '正25', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (30, 1004, '正26', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (31, 1004, '正27', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (32, 1004, '正28', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (33, 1004, '正29', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (34, 1004, '正30', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (35, 1004, '正31', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (36, 1004, '正32', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (37, 1004, '正33', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (38, 1004, '正34', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (39, 1004, '正35', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (40, 1004, '正36', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (41, 1004, '正37', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (42, 1004, '正38', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (43, 1004, '正39', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (44, 1004, '正40', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (45, 1004, '正41', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (46, 1004, '正42', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (47, 1004, '正43', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (48, 1004, '正44', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (49, 1004, '正45', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (50, 1004, '正46', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (51, 1004, '正47', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (52, 1004, '正48', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (53, 1004, '正49', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (54, 1004, '正特', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (55, 1004, '正特', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (56, 1004, '正特', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (57, 1004, '正特', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (58, 1004, '正特', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (59, 1004, '特6', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (60, 1004, '特7', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (61, 1004, '特8', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (62, 1004, '特9', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (63, 1004, '特10', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (64, 1004, '特11', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (65, 1004, '特12', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (66, 1004, '特13', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (67, 1004, '特14', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (68, 1004, '特15', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (69, 1004, '特16', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (70, 1004, '特17', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (71, 1004, '特18', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (72, 1004, '特19', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (73, 1004, '特20', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (74, 1004, '特21', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (75, 1004, '特22', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (76, 1004, '特23', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (77, 1004, '特24', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (78, 1004, '特25', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (79, 1004, '特26', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (80, 1004, '特27', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (81, 1004, '特28', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (82, 1004, '特29', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (83, 1004, '特30', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (84, 1004, '特31', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (85, 1004, '特32', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (86, 1004, '特33', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (87, 1004, '特34', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (88, 1004, '特35', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (89, 1004, '特36', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (90, 1004, '特37', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (91, 1004, '特38', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (92, 1004, '特39', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (93, 1004, '特40', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (94, 1004, '特41', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (95, 1004, '特42', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (96, 1004, '特43', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (97, 1004, '特44', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (98, 1004, '特45', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (99, 1004, '特46', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (100, 1004, '特47', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (101, 1004, '特48', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (102, 1004, '特49', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (103, 1004, '鼠', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (104, 1004, '牛', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (105, 1004, '虎', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (106, 1004, '兔', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (107, 1004, '龙', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (108, 1004, '蛇', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (109, 1004, '马', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (110, 1004, '羊', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (111, 1004, '猴', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (112, 1004, '鸡', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (113, 1004, '狗', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (114, 1004, '猪', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (115, 1004, '红波', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (116, 1004, '蓝波', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (117, 1004, '黄波', 1000, 3000000);

-- PC28 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 1005, '大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 1005, '小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 1005, '单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 1005, '双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 1005, '0点', 1000, 50000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 1005, '1点', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 1005, '2点', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 1005, '3点', 1000, 200000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (9, 1005, '4点', 1000, 200000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (10, 1005, '5点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (11, 1005, '6点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (12, 1005, '7点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (13, 1005, '8点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (14, 1005, '9点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (15, 1005, '10点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (16, 1005, '11点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (17, 1005, '12点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (18, 1005, '13点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (19, 1005, '14点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (20, 1005, '15点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (21, 1005, '16点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (22, 1005, '17点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (23, 1005, '18点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (24, 1005, '19点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (25, 1005, '20点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (26, 1005, '21点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (27, 1005, '22点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (28, 1005, '23点', 1000, 200000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (29, 1005, '24点', 1000, 200000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (30, 1005, '25点', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (31, 1005, '26点', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (32, 1005, '27点', 1000, 50000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (33, 1005, '豹子', 1000, 200000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (34, 1005, '顺子', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (35, 1005, '对子', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (36, 1005, '半顺', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (37, 1005, '杂六', 1000, 3000000);

-- 百家樂 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 1006, '庄', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 1006, '闲', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 1006, '和', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 1006, '庄对', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 1006, '闲对', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 1006, '大', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 1006, '小', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 1006, '幸运6', 1000, 500000);

-- 龍虎 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 1007, '龙', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 1007, '虎', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 1007, '和', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 1007, '龙红', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 1007, '龙黑', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 1007, '虎红', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 1007, '虎黑', 1000, 1000000);

-- 牛牛 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 1008, '蓝胜', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 1008, '红胜', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 1008, '蓝无牛', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 1008, '蓝牛一', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 1008, '蓝牛二', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 1008, '蓝牛三', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 1008, '蓝牛四', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 1008, '蓝牛五', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (9, 1008, '蓝牛六', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (10, 1008, '蓝牛七', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (11, 1008, '蓝牛八', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (12, 1008, '蓝牛九', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (13, 1008, '蓝牛牛', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (14, 1008, '蓝特殊牌型', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (15, 1008, '红无牛', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (16, 1008, '红牛一', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (17, 1008, '红牛二', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (18, 1008, '红牛三', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (19, 1008, '红牛四', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (20, 1008, '红牛五', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (21, 1008, '红牛六', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (22, 1008, '红牛七', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (23, 1008, '红牛八', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (24, 1008, '红牛九', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (25, 1008, '红牛牛', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (26, 1008, '红特殊牌型', 1000, 100000);

-- 輪盤 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 1009, '小', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 1009, '大', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 1009, '红', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 1009, '黑', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 1009, '单', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 1009, '双', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 1009, '打注1-12', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 1009, '打注13-24', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (9, 1009, '打注25-36', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (10, 1009, '直注0', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (11, 1009, '直注1', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (12, 1009, '直注2', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (13, 1009, '直注3', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (14, 1009, '直注4', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (15, 1009, '直注5', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (16, 1009, '直注6', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (17, 1009, '直注7', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (18, 1009, '直注8', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (19, 1009, '直注9', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (20, 1009, '直注10', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (21, 1009, '直注11', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (22, 1009, '直注12', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (23, 1009, '直注13', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (24, 1009, '直注14', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (25, 1009, '直注15', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (26, 1009, '直注16', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (27, 1009, '直注17', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (28, 1009, '直注18', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (29, 1009, '直注19', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (30, 1009, '直注20', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (31, 1009, '直注21', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (32, 1009, '直注22', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (33, 1009, '直注23', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (34, 1009, '直注24', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (35, 1009, '直注25', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (36, 1009, '直注26', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (37, 1009, '直注27', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (38, 1009, '直注28', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (39, 1009, '直注29', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (40, 1009, '直注30', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (41, 1009, '直注31', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (42, 1009, '直注32', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (43, 1009, '直注33', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (44, 1009, '直注34', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (45, 1009, '直注35', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (46, 1009, '直注36', 1000, 100000);

-- 魚蝦蟹 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 1010, '大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 1010, '小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 1010, '单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 1010, '双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 1010, '豹子', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 1010, '4点', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 1010, '5点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 1010, '6点', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (9, 1010, '7点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (10, 1010, '8点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (11, 1010, '9点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (12, 1010, '10点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (13, 1010, '11点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (14, 1010, '12点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (15, 1010, '13点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (16, 1010, '14点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (17, 1010, '15点', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (18, 1010, '16点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (19, 1010, '17点', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (20, 1010, '鱼', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (21, 1010, '虾', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (22, 1010, '葫芦', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (23, 1010, '铜币', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (24, 1010, '蟹', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (25, 1010, '鸡', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (26, 1010, '红', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (27, 1010, '绿', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (28, 1010, '蓝', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (29, 1010, '红对', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (30, 1010, '绿对', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (31, 1010, '蓝对', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (32, 1010, '红豹', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (33, 1010, '绿豹', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (34, 1010, '蓝豹', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (35, 1010, '同色全豹', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (36, 1010, '豹子鱼', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (37, 1010, '豹子虾', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (38, 1010, '豹子葫芦', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (39, 1010, '豹子铜币', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (40, 1010, '豹子蟹', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (41, 1010, '豹子鸡', 1000, 100000);

-- ----------------------------
-- Table structure for odds
-- ----------------------------
DROP TABLE IF EXISTS `odds`;
CREATE TABLE IF NOT EXISTS `odds` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '賠率id',
    `game_id` INT UNSIGNED COMMENT '遊戲id(關聯)',
    `bet_area_id` INT UNSIGNED COMMENT '注區id(關聯)',
    `odd` FLOAT NOT NULL DEFAULT 0 COMMENT '賠率',
    `is_deleted` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否已被刪除',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '修改日期'
) ENGINE=InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '賠率表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of odds
-- ----------------------------
-- 一分三快 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 1, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 2, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 3, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 4, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 5, 63.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 6, 32.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 7, 19.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 8, 13.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 9, 9.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 10, 8.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 11, 7.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 12, 7.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 13, 8.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 14, 9.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 15, 13.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 16, 19.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 17, 32.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 18, 63.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 19, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 19, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 19, 4.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 20, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 20, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 20, 4.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 21, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 21, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 21, 4.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 22, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 22, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 22, 4.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 23, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 23, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 23, 4.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 24, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 24, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 24, 4.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 25, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 26, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 27, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 28, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 29, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 30, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 31, 180.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 32, 180.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 33, 180.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 34, 180.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 35, 180.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 36, 180.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1001, 37, 32.00);

-- 賽車 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 1, 2.18);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 2, 1.75);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 3, 1.75);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 4, 2.18);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 5, 40.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 6, 40.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 7, 20.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 8, 20.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 9, 13.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 10, 13.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 11, 10.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 12, 10.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 13, 8.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 14, 10.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 15, 10.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 16, 13.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 17, 13.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 18, 20.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 19, 20.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 20, 40.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 21, 40.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 22, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 23, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 24, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 25, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 26, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 27, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 28, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 29, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 30, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 31, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 32, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 33, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 34, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1002, 35, 9.60);

-- 時時彩 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 1, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 2, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 3, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 4, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 5, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 6, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 7, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 8, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 9, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 10, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 11, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 12, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 13, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 14, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 15, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 16, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 17, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 18, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 19, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 20, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 21, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 22, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 23, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 24, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 25, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 26, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 27, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 28, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 29, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 30, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 31, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 32, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 33, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 34, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 35, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 36, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 37, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 38, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 39, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 40, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 41, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 42, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 43, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 44, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 45, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 46, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 47, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 48, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 49, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 50, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 51, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 52, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 53, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 54, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 55, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 56, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 57, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 58, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 59, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 60, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 61, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 62, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 63, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 64, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 65, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 66, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 67, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 68, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 69, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 70, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 71, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 72, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 73, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 74, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 75, 88.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 76, 15.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 77, 3.30);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 78, 2.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 79, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 80, 88.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 81, 15.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 82, 3.30);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 83, 2.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 84, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 85, 88.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 86, 15.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 87, 3.30);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 88, 2.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 89, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 90, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 90, 0.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 91, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 92, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1003, 92, 0.50);

-- 六合彩
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 1, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 1, 1.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 2, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 2, 1.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 3, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 3, 1.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 4, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 4, 1.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 5, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 6, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 7, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 8, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 9, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 10, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 11, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 12, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 13, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 14, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 15, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 16, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 17, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 18, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 19, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 20, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 21, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 22, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 23, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 24, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 25, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 26, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 27, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 28, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 29, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 30, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 31, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 32, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 33, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 34, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 35, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 36, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 37, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 38, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 39, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 40, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 41, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 42, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 43, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 44, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 45, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 46, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 47, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 48, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 49, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 50, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 51, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 52, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 53, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 54, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 55, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 56, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 57, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 58, 7.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 59, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 60, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 61, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 62, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 63, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 64, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 65, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 66, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 67, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 68, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 69, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 70, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 71, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 72, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 73, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 74, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 75, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 76, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 77, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 78, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 79, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 80, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 81, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 82, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 83, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 84, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 85, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 86, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 87, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 88, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 89, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 90, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 91, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 92, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 93, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 94, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 95, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 96, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 97, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 98, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 99, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 100, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 101, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 102, 46.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 103, 11.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 104, 11.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 105, 11.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 106, 11.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 107, 9.40);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 108, 11.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 109, 11.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 110, 11.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 111, 11.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 112, 11.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 113, 11.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 114, 11.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 115, 2.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 116, 2.97);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1004, 117, 2.97);

-- PC28 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 1, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 2, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 3, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 4, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 5, 850.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 6, 285.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 7, 145.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 8, 88.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 9, 60.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 10, 43.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 11, 33.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 12, 26.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 13, 21.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 14, 17.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 15, 15.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 16, 13.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 17, 13.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 18, 12.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 19, 12.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 20, 13.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 21, 13.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 22, 15.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 23, 17.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 24, 21.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 25, 26.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 26, 33.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 27, 43.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 28, 60.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 29, 88.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 30, 145.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 31, 285.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 32, 850.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 33, 88.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 34, 15.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 35, 3.30);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 36, 2.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1005, 37, 3.00);

-- 百家樂 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1006, 1, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1006, 1, 1.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1006, 2, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1006, 2, 1.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1006, 3, 9.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1006, 4, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1006, 5, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1006, 6, 1.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1006, 7, 2.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1006, 8, 13.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1006, 8, 21.00);

-- 龍虎 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1007, 1, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1007, 1, 0.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1007, 2, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1007, 2, 0.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1007, 3, 9.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1007, 4, 1.90);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1007, 5, 1.90);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1007, 6, 1.90);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1007, 7, 1.90);

-- 牛牛 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 1, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 2, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 3, 2.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 4, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 5, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 6, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 7, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 8, 13.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 9, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 10, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 11, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 12, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 13, 13.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 14, 500.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 15, 2.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 16, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 17, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 18, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 19, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 20, 13.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 21, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 22, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 23, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 24, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 25, 13.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1008, 26, 500.00);

-- 輪盤 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 1, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 2, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 3, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 4, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 5, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 6, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 7, 2.98);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 8, 2.98);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 9, 2.98);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 10, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 11, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 12, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 13, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 14, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 15, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 16, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 17, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 18, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 19, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 20, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 21, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 22, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 23, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 24, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 25, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 26, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 27, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 28, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 29, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 30, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 31, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 32, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 33, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 34, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 35, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 36, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 37, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 38, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 39, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 40, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 41, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 42, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 43, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 44, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 45, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1009, 46, 35.00);

-- 魚蝦蟹 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 1, 1.9900);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 2, 1.9900);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 3, 1.9900);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 4, 1.9900);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 5, 32.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 6, 63.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 7, 32.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 8, 19.5000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 9, 13.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 10, 9.5000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 11, 8.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 12, 7.5000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 13, 7.5000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 14, 8.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 15, 9.5000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 16, 13.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 17, 19.5000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 18, 32.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 19, 63.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 20, 2.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 20, 3.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 20, 4.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 21, 2.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 21, 3.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 21, 4.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 22, 2.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 22, 3.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 22, 4.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 23, 2.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 23, 3.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 23, 4.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 24, 2.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 24, 3.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 24, 4.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 25, 2.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 25, 3.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 25, 4.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 26, 2.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 27, 2.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 28, 2.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 29, 4.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 30, 4.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 31, 4.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 32, 24.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 33, 24.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 34, 24.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 35, 8.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 36, 180.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 37, 180.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 38, 180.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 39, 180.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 40, 180.0000);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1010, 41, 180.0000);

SET FOREIGN_KEY_CHECKS = 1;



-- ----------------------------
-- Table structure for odds
-- ----------------------------
DROP TABLE IF EXISTS `bets`;
CREATE TABLE IF NOT EXISTS `bets` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '投注 id',
    `user_id` INT UNSIGNED UNIQUE NOT NULL DEFAULT 0 COMMENT '用戶id',
    `round_info_id` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '期號',
    `game_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '遊戲 id',
    `bet_area_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '注區 id',
    `score` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '投注金額',
    `is_deleted` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否已被刪除',
    `create_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `update_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '修改日期'
) ENGINE=InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '投注表' ROW_FORMAT = DYNAMIC;

INSERT INTO `bets` (user_id, round_info_id, game_id, bet_area_id, score) VALUES (1, 202412171135, 1009, 1, 10000);