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
INSERT INTO `games` (id, name) VALUES (1, '一分三快');
INSERT INTO `games` (id, name) VALUES (2, '賽車');
INSERT INTO `games` (id, name) VALUES (6, '百家樂');
INSERT INTO `games` (id, name) VALUES (7, '龍虎');
INSERT INTO `games` (id, name) VALUES (8, '牛牛');
INSERT INTO `games` (id, name) VALUES (9, '輪盤');

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
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 1, '大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 1, '小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 1, '单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 1, '双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 1, '总和4点', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 1, '总和5点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 1, '总和6点', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 1, '总和7点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (9, 1, '总和8点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (10, 1, '总和9点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (11, 1, '总和10点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (12, 1, '总和11点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (13, 1, '总和12点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (14, 1, '总和13点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (15, 1, '总和14点', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (16, 1, '总和15点', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (17, 1, '总和16点', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (18, 1, '总和17点', 1000, 300000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (19, 1, '单骰1', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (20, 1, '单骰2', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (21, 1, '单骰3', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (22, 1, '单骰4', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (23, 1, '单骰5', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (24, 1, '单骰6', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (25, 1, '对子1', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (26, 1, '对子2', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (27, 1, '对子3', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (28, 1, '对子4', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (29, 1, '对子5', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (30, 1, '对子6', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (31, 1, '豹子1', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (32, 1, '豹子2', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (33, 1, '豹子3', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (34, 1, '豹子4', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (35, 1, '豹子5', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (36, 1, '豹子6', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (37, 1, '豹子1-6', 1000, 500000);

-- 賽車 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 2, '和大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 2, '和小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 2, '和单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 2, '和双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 2, '3', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 2, '4', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 2, '5', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 2, '6', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (9, 2, '7', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (10, 2, '8', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (11, 2, '9', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (12, 2, '10', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (13, 2, '11', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (14, 2, '12', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (15, 2, '13', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (16, 2, '14', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (17, 2, '15', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (18, 2, '16', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (19, 2, '17', 1000, 800000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (20, 2, '18', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (21, 2, '19', 1000, 500000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (22, 2, '冠军大', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (23, 2, '冠军小', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (24, 2, '冠军单', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (25, 2, '冠军双', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (26, 2, '冠军1', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (27, 2, '冠军2', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (28, 2, '冠军3', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (29, 2, '冠军4', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (30, 2, '冠军5', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (31, 2, '冠军6', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (32, 2, '冠军7', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (33, 2, '冠军8', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (34, 2, '冠军9', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (35, 2, '冠军10', 1000, 1000000);

-- 百家樂 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 6, '庄', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 6, '闲', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 6, '和', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 6, '庄对', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 6, '闲对', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 6, '大', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 6, '小', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 6, '幸运6', 1000, 500000);

-- 龍虎 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 7, '龙', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 7, '虎', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 7, '和', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 7, '龙红', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 7, '龙黑', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 7, '虎红', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 7, '虎黑', 1000, 1000000);

-- 牛牛 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 8, '蓝胜', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 8, '红胜', 1000, 5000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 8, '蓝无牛', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 8, '蓝牛一', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 8, '蓝牛二', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 8, '蓝牛三', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 8, '蓝牛四', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 8, '蓝牛五', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (9, 8, '蓝牛六', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (10, 8, '蓝牛七', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (11, 8, '蓝牛八', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (12, 8, '蓝牛九', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (13, 8, '蓝牛牛', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (14, 8, '蓝特殊牌型', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (15, 8, '红无牛', 1000, 3000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (16, 8, '红牛一', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (17, 8, '红牛二', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (18, 8, '红牛三', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (19, 8, '红牛四', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (20, 8, '红牛五', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (21, 8, '红牛六', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (22, 8, '红牛七', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (23, 8, '红牛八', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (24, 8, '红牛九', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (25, 8, '红牛牛', 1000, 1000000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (26, 8, '红特殊牌型', 1000, 100000);

-- 輪盤 --
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (1, 9, '小', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (2, 9, '大', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (3, 9, '红', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (4, 9, '黑', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (5, 9, '单', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (6, 9, '双', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (7, 9, '打注1-12', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (8, 9, '打注13-24', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (9, 9, '打注25-36', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (10, 9, '直注0', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (11, 9, '直注1', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (12, 9, '直注2', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (13, 9, '直注3', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (14, 9, '直注4', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (15, 9, '直注5', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (16, 9, '直注6', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (17, 9, '直注7', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (18, 9, '直注8', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (19, 9, '直注9', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (20, 9, '直注10', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (21, 9, '直注11', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (22, 9, '直注12', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (23, 9, '直注13', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (24, 9, '直注14', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (25, 9, '直注15', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (26, 9, '直注16', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (27, 9, '直注17', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (28, 9, '直注18', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (29, 9, '直注19', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (30, 9, '直注20', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (31, 9, '直注21', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (32, 9, '直注22', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (33, 9, '直注23', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (34, 9, '直注24', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (35, 9, '直注25', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (36, 9, '直注26', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (37, 9, '直注27', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (38, 9, '直注28', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (39, 9, '直注29', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (40, 9, '直注30', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (41, 9, '直注31', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (42, 9, '直注32', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (43, 9, '直注33', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (44, 9, '直注34', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (45, 9, '直注35', 1000, 100000);
INSERT INTO `bet_areas` (id, game_id, name, min_limit, max_limit) VALUES (46, 9, '直注36', 1000, 100000);

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
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 1, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 2, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 3, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 4, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 5, 63.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 6, 32.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 7, 19.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 8, 13.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 9, 9.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 10, 8.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 11, 7.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 12, 7.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 13, 8.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 14, 9.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 15, 13.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 16, 19.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 17, 32.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 18, 63.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 19, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 19, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 19, 4.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 20, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 20, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 20, 4.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 21, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 21, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 21, 4.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 22, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 22, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 22, 4.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 23, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 23, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 23, 4.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 24, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 24, 3.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 24, 4.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 25, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 26, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 27, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 28, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 29, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 30, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 31, 180.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 32, 180.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 33, 180.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 34, 180.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 35, 180.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 36, 180.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (1, 37, 32.00);

-- 賽車 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 1, 2.18);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 2, 1.75);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 3, 1.75);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 4, 2.18);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 5, 40.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 6, 40.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 7, 20.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 8, 20.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 9, 13.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 10, 13.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 11, 10.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 12, 10.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 13, 8.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 14, 10.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 15, 10.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 16, 13.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 17, 13.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 18, 20.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 19, 20.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 20, 40.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 21, 40.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 22, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 23, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 24, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 25, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 26, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 27, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 28, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 29, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 30, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 31, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 32, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 33, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 34, 9.60);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (2, 35, 9.60);

-- 百家樂 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (6, 1, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (6, 1, 1.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (6, 2, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (6, 2, 1.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (6, 3, 9.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (6, 4, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (6, 5, 12.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (6, 6, 1.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (6, 7, 2.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (6, 8, 13.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (6, 8, 21.00);

-- 龍虎 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (7, 1, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (7, 1, 0.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (7, 2, 2.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (7, 2, 0.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (7, 3, 9.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (7, 4, 1.90);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (7, 5, 1.90);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (7, 6, 1.90);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (7, 7, 1.90);

-- 牛牛 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 1, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 2, 1.95);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 3, 2.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 4, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 5, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 6, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 7, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 8, 13.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 9, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 10, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 11, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 12, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 13, 13.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 14, 500.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 15, 2.80);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 16, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 17, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 18, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 19, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 20, 13.70);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 21, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 22, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 23, 14.50);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 24, 14.20);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 25, 13.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (8, 26, 500.00);

-- 輪盤 --
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 1, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 2, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 3, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 4, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 5, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 6, 1.99);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 7, 2.98);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 8, 2.98);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 9, 2.98);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 10, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 11, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 12, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 13, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 14, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 15, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 16, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 17, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 18, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 19, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 20, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 21, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 22, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 23, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 24, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 25, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 26, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 27, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 28, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 29, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 30, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 31, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 32, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 33, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 34, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 35, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 36, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 37, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 38, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 39, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 40, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 41, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 42, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 43, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 44, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 45, 35.00);
INSERT INTO `odds` (game_id, bet_area_id, odd) VALUES (9, 46, 35.00);

SET FOREIGN_KEY_CHECKS = 1;