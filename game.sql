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
INSERT INTO `games` (id, name) VALUES (1001, '一分三快');
INSERT INTO `games` (id, name) VALUES (1002, '賽車');
INSERT INTO `games` (id, name) VALUES (1006, '百家樂');
INSERT INTO `games` (id, name) VALUES (1007, '龍虎');
INSERT INTO `games` (id, name) VALUES (1008, '牛牛');
INSERT INTO `games` (id, name) VALUES (1009, '輪盤');

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