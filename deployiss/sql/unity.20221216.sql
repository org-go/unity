/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50734
 Source Host           : 127.0.0.1:3306
 Source Schema         : unity

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 16/12/2022 13:35:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for d_platform
-- ----------------------------
DROP TABLE IF EXISTS `d_platform`;
CREATE TABLE `d_platform`  (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` tinyint(4) NULL DEFAULT NULL,
  `created_time` timestamp(0) NULL DEFAULT NULL,
  `updated_time` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_平台表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_platform
-- ----------------------------

-- ----------------------------
-- Table structure for d_rank
-- ----------------------------
DROP TABLE IF EXISTS `d_rank`;
CREATE TABLE `d_rank`  (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '头衔名称',
  `experience` bigint(20) NULL DEFAULT NULL COMMENT '晋级所需经验',
  `description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '头衔描述',
  `rights_pks` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '权益PK, 逗号(,)分隔',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_头衔表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_rank
-- ----------------------------

-- ----------------------------
-- Table structure for d_rank_rights
-- ----------------------------
DROP TABLE IF EXISTS `d_rank_rights`;
CREATE TABLE `d_rank_rights`  (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `rights_pk` int(11) NULL DEFAULT NULL COMMENT '权益PK',
  `rank_pk` int(11) NULL DEFAULT NULL COMMENT '头衔PK',
  PRIMARY KEY (`pk`) USING BTREE,
  INDEX `rank_rights_pks`(`rank_pk`, `rights_pk`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_头衔权益关联表1N' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_rank_rights
-- ----------------------------

-- ----------------------------
-- Table structure for d_rights
-- ----------------------------
DROP TABLE IF EXISTS `d_rights`;
CREATE TABLE `d_rights`  (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权益名称',
  `rights_classify_pk` int(11) NULL DEFAULT NULL COMMENT '权益类目PK',
  `description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权益描述',
  `status` tinyint(4) NULL DEFAULT NULL COMMENT '权益状态 1:启用 2：禁用',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_权益表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_rights
-- ----------------------------

-- ----------------------------
-- Table structure for d_rights_classify
-- ----------------------------
DROP TABLE IF EXISTS `d_rights_classify`;
CREATE TABLE `d_rights_classify`  (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `supperior_pk` int(11) NULL DEFAULT NULL COMMENT '上级PK',
  `description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_权益类目表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_rights_classify
-- ----------------------------

-- ----------------------------
-- Table structure for d_task
-- ----------------------------
DROP TABLE IF EXISTS `d_task`;
CREATE TABLE `d_task`  (
  `pk` bigint(20) NOT NULL AUTO_INCREMENT,
  `id` bigint(20) NULL DEFAULT NULL COMMENT '任务ID',
  `task_classify_pk` int(11) NULL DEFAULT NULL COMMENT '任务类目PK',
  `owner_user_id` bigint(20) NULL DEFAULT NULL COMMENT '任务负责人',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '名称',
  `title` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '标题',
  `description` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '描述',
  `claim` json NULL COMMENT '要求',
  `task_score` int(11) NULL DEFAULT NULL COMMENT '任务积分/单',
  `task_number` int(11) NULL DEFAULT NULL COMMENT '单数',
  `task_start_time` datetime(0) NULL DEFAULT NULL COMMENT '任务开始时间',
  `task_end_time` datetime(0) NULL DEFAULT NULL COMMENT '任务结束时间',
  `status` tinyint(4) NULL DEFAULT NULL COMMENT '状态：1:草稿， 3：审核中， 5：进行中， 7：下架， 9：结束',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `auditor` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '审核人',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_任务表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_task
-- ----------------------------

-- ----------------------------
-- Table structure for d_task_classify
-- ----------------------------
DROP TABLE IF EXISTS `d_task_classify`;
CREATE TABLE `d_task_classify`  (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `platform_pk` int(11) NULL DEFAULT NULL COMMENT '平台PK',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '任务类目名称',
  `description` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '描述',
  `status` tinyint(4) NULL DEFAULT NULL COMMENT '状态：1：启用， 2：禁用',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `created_user` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_任务类目表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_task_classify
-- ----------------------------

-- ----------------------------
-- Table structure for d_task_order
-- ----------------------------
DROP TABLE IF EXISTS `d_task_order`;
CREATE TABLE `d_task_order`  (
  `pk` bigint(20) NOT NULL AUTO_INCREMENT,
  `worker_id` int(11) NULL DEFAULT NULL COMMENT '接单人ID',
  `task_id` bigint(20) NULL DEFAULT NULL COMMENT '任务ID',
  `order_number` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务接单号（任务ID+自定义）',
  `score` int(11) NULL DEFAULT NULL COMMENT '任务积分',
  `step` json NULL COMMENT '完成步骤',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建接单时间',
  `updated_time` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新接单时间',
  `evaluate` tinyint(4) NULL DEFAULT NULL COMMENT '评价 1：未完成 2：未达标 10：已达标',
  `auditor` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '评审人',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_任务订单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_task_order
-- ----------------------------

-- ----------------------------
-- Table structure for d_task_order_snapshoot
-- ----------------------------
DROP TABLE IF EXISTS `d_task_order_snapshoot`;
CREATE TABLE `d_task_order_snapshoot`  (
  `pk` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_number` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务订单号',
  `task_owner_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务负责人',
  `task_id` bigint(20) NULL DEFAULT NULL COMMENT '任务ID',
  `task_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务名称',
  `start_time` timestamp(0) NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` timestamp(0) NULL DEFAULT NULL COMMENT '结束时间',
  `worker_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '接单人',
  `claim` json NULL COMMENT '要求',
  `step` json NULL COMMENT '步骤',
  `score` int(11) NULL DEFAULT NULL COMMENT '积分',
  `evaluate` tinyint(255) NULL DEFAULT NULL COMMENT '评价',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_任务订单快照' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_task_order_snapshoot
-- ----------------------------

-- ----------------------------
-- Table structure for d_user
-- ----------------------------
DROP TABLE IF EXISTS `d_user`;
CREATE TABLE `d_user`  (
  `pk` bigint(20) NOT NULL AUTO_INCREMENT,
  `id` bigint(20) NULL DEFAULT NULL COMMENT '对外ID',
  `supperior_pk` bigint(20) NULL DEFAULT 0 COMMENT '上级PK',
  `path_pk` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '级联PK',
  `phone` bigint(20) NULL DEFAULT NULL COMMENT '手机号',
  `account` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '账号',
  `password` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '名称',
  `rank` int(11) NULL DEFAULT NULL COMMENT '权益级别',
  `score` bigint(20) NULL DEFAULT NULL COMMENT '当前积分',
  `scope` set('x','xx','x21','x22') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'x' COMMENT 'x： 正常，\r\nxx： 禁用，\r\nx21： 交易冻结，\r\nx22： 任务冻结',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `last_login_time` timestamp(0) NULL DEFAULT NULL COMMENT '最后登陆时间',
  `last_login_ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '最后登陆IP',
  PRIMARY KEY (`pk`) USING BTREE,
  INDEX `id`(`id`) USING BTREE,
  INDEX `mix_up_path_pk`(`supperior_pk`, `path_pk`) USING BTREE,
  INDEX `mix_account_phone`(`account`, `phone`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_user
-- ----------------------------

-- ----------------------------
-- Table structure for d_user_address
-- ----------------------------
DROP TABLE IF EXISTS `d_user_address`;
CREATE TABLE `d_user_address`  (
  `pk` bigint(22) NOT NULL AUTO_INCREMENT,
  `user_pk` bigint(22) NOT NULL COMMENT '用户PK',
  `state_code` int(11) NULL DEFAULT NULL COMMENT '省',
  `city_code` int(11) NULL DEFAULT NULL COMMENT '市',
  `street_code` int(11) NULL DEFAULT NULL COMMENT '街道',
  `address` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '详细地址',
  `postal_code` int(11) NULL DEFAULT NULL COMMENT '邮编',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `is_default` tinyint(4) NULL DEFAULT NULL COMMENT '1: 默认',
  PRIMARY KEY (`pk`) USING BTREE,
  INDEX `user_pk_default`(`user_pk`, `is_default`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_用户地址表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_user_address
-- ----------------------------

-- ----------------------------
-- Table structure for d_user_oauth_account
-- ----------------------------
DROP TABLE IF EXISTS `d_user_oauth_account`;
CREATE TABLE `d_user_oauth_account`  (
  `user_pk` bigint(20) NOT NULL,
  `wechat` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `qq` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `dingtalk` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `github` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_用户三方账户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_user_oauth_account
-- ----------------------------

-- ----------------------------
-- Table structure for d_user_portrait
-- ----------------------------
DROP TABLE IF EXISTS `d_user_portrait`;
CREATE TABLE `d_user_portrait`  (
  `user_pk` bigint(20) NOT NULL,
  `unit_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '身份证',
  `borth` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '生日',
  `hobbies` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '兴趣爱好',
  `email` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  PRIMARY KEY (`user_pk`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_用户基础画像表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_user_portrait
-- ----------------------------

-- ----------------------------
-- Table structure for d_welfare
-- ----------------------------
DROP TABLE IF EXISTS `d_welfare`;
CREATE TABLE `d_welfare`  (
  `pk` bigint(20) NOT NULL AUTO_INCREMENT,
  `id` bigint(20) NULL DEFAULT NULL COMMENT '对外ID',
  `platform_pk` int(11) NULL DEFAULT NULL COMMENT '平台PK',
  `welfare_classify_pk` int(11) NULL DEFAULT NULL COMMENT '福利类型PK',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '福利名称',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '福利标题',
  `description` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '福利描述',
  `link` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '福利链接',
  `status` tinyint(4) NULL DEFAULT NULL COMMENT '状态; 1:启用； 2：禁用',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `created_user` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  PRIMARY KEY (`pk`) USING BTREE,
  INDEX `classify_pk`(`welfare_classify_pk`) USING BTREE,
  INDEX `mix_name_title_status`(`name`, `title`, `status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_福利信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_welfare
-- ----------------------------

-- ----------------------------
-- Table structure for d_welfare_classify
-- ----------------------------
DROP TABLE IF EXISTS `d_welfare_classify`;
CREATE TABLE `d_welfare_classify`  (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `platform_pk` int(11) NULL DEFAULT NULL COMMENT '平台PK',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '福利类目名称',
  `description` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '描述',
  `status` tinyint(4) NULL DEFAULT NULL COMMENT '状态; 1:启用 2：禁用',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '删除时间',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'data_福利类目表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of d_welfare_classify
-- ----------------------------

-- ----------------------------
-- Table structure for meta_asset_suffix
-- ----------------------------
DROP TABLE IF EXISTS `meta_asset_suffix`;
CREATE TABLE `meta_asset_suffix`  (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `scene` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '场景',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of meta_asset_suffix
-- ----------------------------
INSERT INTO `meta_asset_suffix` VALUES (1, 'BAL', '余额，用于标识余额信息', '如：话费余额，信用卡余额');
INSERT INTO `meta_asset_suffix` VALUES (2, 'CNT', '次数，用于统计数量', '如：登陆次数，充值次数');
INSERT INTO `meta_asset_suffix` VALUES (3, 'DATE', '日期，用于标识到天的信息', '如:  生效日期，失效日期');
INSERT INTO `meta_asset_suffix` VALUES (4, 'CODE', '编码， 用于前台业务使用的编码', '如：公司编码');
INSERT INTO `meta_asset_suffix` VALUES (5, 'DESC', NULL, NULL);
INSERT INTO `meta_asset_suffix` VALUES (6, 'DUR', NULL, NULL);
INSERT INTO `meta_asset_suffix` VALUES (7, 'FEE', NULL, NULL);
INSERT INTO `meta_asset_suffix` VALUES (8, 'FLAG', NULL, NULL);
INSERT INTO `meta_asset_suffix` VALUES (9, 'FLUX', NULL, NULL);
INSERT INTO `meta_asset_suffix` VALUES (10, 'MON', NULL, NULL);
INSERT INTO `meta_asset_suffix` VALUES (11, 'NUM', NULL, NULL);
INSERT INTO `meta_asset_suffix` VALUES (12, 'PRC', NULL, NULL);
INSERT INTO `meta_asset_suffix` VALUES (13, 'SCORE', NULL, NULL);
INSERT INTO `meta_asset_suffix` VALUES (14, 'SN', NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
