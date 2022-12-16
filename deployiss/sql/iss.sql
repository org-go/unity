/*
 Navicat Premium Data Transfer

 Source Server         : A_DEV
 Source Server Type    : MySQL
 Source Server Version : 50735
 Source Host           : rm-8vbjy34g96075qpoklo.mysql.zhangbei.rds.aliyuncs.com:3408
 Source Schema         : iss

 Target Server Type    : MySQL
 Target Server Version : 50735
 File Encoding         : 65001

 Date: 09/10/2022 14:17:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for iss_biz
-- ----------------------------
DROP TABLE IF EXISTS `iss_biz`;
CREATE TABLE `iss_biz`  (
  `biz_id` int(11) NOT NULL AUTO_INCREMENT,
  `biz_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商户号(唯一)',
  `biz_name_cn` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商户汉字',
  `biz_name_kana` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商户假名',
  `biz_name_en` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商户英文',
  `industry_code` int(11) NULL DEFAULT NULL COMMENT '行业code',
  `post_code` int(11) NULL DEFAULT NULL COMMENT '邮编code',
  `status` tinyint(4) NULL DEFAULT NULL COMMENT '状态: 1:新增2：更新3：删除',
  `start_time` date NULL DEFAULT NULL COMMENT '生效时间',
  `updated_time` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`biz_id`) USING BTREE,
  INDEX `i_biz_code`(`biz_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '公司表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for iss_biz_log
-- ----------------------------
DROP TABLE IF EXISTS `iss_biz_log`;
CREATE TABLE `iss_biz_log`  (
  `log_id` int(11) NOT NULL AUTO_INCREMENT,
  `biz_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商户code',
  `detail` json NULL COMMENT '详情记录',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '公司日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for iss_biz_store
-- ----------------------------
DROP TABLE IF EXISTS `iss_biz_store`;
CREATE TABLE `iss_biz_store`  (
  `store_id` int(11) NOT NULL AUTO_INCREMENT,
  `biz_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商户号(唯一)',
  `biz_store_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '店铺code',
  `biz_store_cn` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '店铺汉字',
  `biz_store_kana` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '店铺假名',
  `biz_store_en` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '店铺英文',
  `industry_code` int(11) NULL DEFAULT NULL COMMENT '行业code',
  `post_code` int(11) NULL DEFAULT NULL COMMENT '邮编code',
  `address` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址 ',
  `status` tinyint(4) NULL DEFAULT NULL COMMENT '状态: 1:新增2：更新3：删除',
  `start_time` date NULL DEFAULT NULL COMMENT '生效时间',
  `updated_time` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`store_id`) USING BTREE,
  INDEX `i_biz_store_code`(`biz_store_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '公司表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for iss_biz_store_log
-- ----------------------------
DROP TABLE IF EXISTS `iss_biz_store_log`;
CREATE TABLE `iss_biz_store_log`  (
  `log_id` int(11) NOT NULL AUTO_INCREMENT,
  `biz_store_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '店铺code',
  `detail` json NULL COMMENT '详情记录',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 63 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '公司日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for iss_cpm_log
-- ----------------------------
DROP TABLE IF EXISTS `iss_cpm_log`;
CREATE TABLE `iss_cpm_log`  (
  `log_id` int(11) NOT NULL AUTO_INCREMENT,
  `qrcode` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'qrcode' COMMENT '二维码',
  `barcode` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '一维码',
  `request_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付ID',
  `cpm_args` json NULL COMMENT '参数',
  `created_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  `use_time` int(11) NULL DEFAULT NULL COMMENT '耗时',
  `updated_time` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4749 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '扫码日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for iss_method
-- ----------------------------
DROP TABLE IF EXISTS `iss_method`;
CREATE TABLE `iss_method`  (
  `method_id` int(11) NOT NULL AUTO_INCREMENT,
  `method_sign` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `method_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `method_system_args` json NULL,
  `is_delete` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`method_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '支付渠道表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for iss_payment
-- ----------------------------
DROP TABLE IF EXISTS `iss_payment`;
CREATE TABLE `iss_payment`  (
  `pay_id` int(11) NOT NULL AUTO_INCREMENT,
  `request_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `origin_request_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求ID',
  `one_time_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '加权时间戳',
  `type` enum('pay','refund') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'pay' COMMENT 'pay支付， refund退款',
  `store_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商户店铺',
  `term_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商户终端号',
  `biz_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商户号',
  `biz_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商户名',
  `sotre_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商户店铺名称',
  `biz_pay_consumer` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '消费者',
  `iss_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'iss订单号',
  `receipt_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商户出示的收据号',
  `core_system_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Core System 流水号',
  `amount` bigint(20) NULL DEFAULT NULL COMMENT '支付金额',
  `fee` int(11) NULL DEFAULT NULL COMMENT '费率',
  `sign` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '签证',
  `req_time` varchar(14) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求时间YYYYMMDDhhmmss',
  `created_time` bigint(14) NULL DEFAULT NULL COMMENT 'iss创建时间YYYYMMDDhhmmss',
  `callback_time` bigint(14) NULL DEFAULT NULL COMMENT 'core回调时间YYYYMMDDhhmmss',
  `date` int(11) NULL DEFAULT NULL COMMENT '日期YYYYMMDD',
  `status` tinyint(4) NULL DEFAULT NULL COMMENT '状态：1:发起订单 2：回调订单 3：支付成功  4：退款成功 5：失败',
  PRIMARY KEY (`pay_id`) USING BTREE,
  INDEX `i_pay_r_id`(`request_id`) USING BTREE,
  INDEX `i_pay_iss_id`(`iss_no`) USING BTREE,
  INDEX `i_pay_c_id`(`biz_pay_consumer`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4786 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '流水表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for iss_payment_log
-- ----------------------------
DROP TABLE IF EXISTS `iss_payment_log`;
CREATE TABLE `iss_payment_log`  (
  `log_id` int(11) NOT NULL AUTO_INCREMENT,
  `type` enum('pay','refund') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'pay' COMMENT 'pay支付， refund退款',
  `pay_id` int(11) NULL DEFAULT NULL COMMENT '支付ID',
  `req_detail` json NULL COMMENT '支付详情',
  `code_with_consumer` json NULL,
  `callback_detail` json NULL COMMENT '回调详情',
  `req_time` timestamp(0) NULL DEFAULT NULL COMMENT '发起时间',
  `callback_time` timestamp(0) NULL DEFAULT NULL COMMENT '回调时间',
  PRIMARY KEY (`log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4756 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '流水日志表' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
