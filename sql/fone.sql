/*
 Navicat Premium Data Transfer

 Source Server         : 数字资产
 Source Server Type    : MySQL
 Source Server Version : 80024
 Source Host           : 192.168.2.131:3306
 Source Schema         : fone

 Target Server Type    : MySQL
 Target Server Version : 80024
 File Encoding         : 65001

 Date: 07/09/2021 15:18:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for chain_height
-- ----------------------------
DROP TABLE IF EXISTS `chain_height`;
CREATE TABLE `chain_height` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `height` bigint DEFAULT NULL,
  `status` tinyint unsigned DEFAULT '0' COMMENT '1:未成功，0:成功',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for chain_transaction
-- ----------------------------
DROP TABLE IF EXISTS `chain_transaction`;
CREATE TABLE `chain_transaction` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `block_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `block_hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `from_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `to_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `gas` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `transaction_index` bigint DEFAULT NULL,
  `timestamp` bigint DEFAULT NULL,
  `nonce` int DEFAULT NULL,
  `block_height` int DEFAULT NULL,
  `method` int DEFAULT '0' COMMENT '0-默认，1-to和from钱包都在系统中 2-单from在系统中（转出） 3-单to在系统中（转入）',
  `status` int DEFAULT NULL,
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=110 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='链上交易数据';

-- ----------------------------
-- Table structure for coin
-- ----------------------------
DROP TABLE IF EXISTS `coin`;
CREATE TABLE `coin` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '币种名称',
  `code` varchar(50) DEFAULT '' COMMENT '币种代号',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='币种管理';

-- ----------------------------
-- Table structure for collect_wallet
-- ----------------------------
DROP TABLE IF EXISTS `collect_wallet`;
CREATE TABLE `collect_wallet` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `wallet_id` bigint DEFAULT '0' COMMENT '普通钱包id',
  `address` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '钱包公钥地址',
  `access_key` varchar(255) COLLATE utf8_bin DEFAULT '' COMMENT 'ak',
  `sign` varchar(255) COLLATE utf8_bin DEFAULT '' COMMENT '签名',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `address` (`address`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='归集钱包表';

-- ----------------------------
-- Table structure for company
-- ----------------------------
DROP TABLE IF EXISTS `company`;
CREATE TABLE `company` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '企业名称',
  `email` varchar(255) DEFAULT '' COMMENT '企业邮箱',
  `password` varchar(50) DEFAULT '' COMMENT '企业密码',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='企业管理';

-- ----------------------------
-- Table structure for member
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `company_id` int NOT NULL COMMENT '企业表ID',
  `name` varchar(255) DEFAULT '' COMMENT '姓名',
  `password` varchar(32) DEFAULT '' COMMENT '密码',
  `email` varchar(150) DEFAULT '' COMMENT '邮箱',
  `desc` varchar(255) DEFAULT '' COMMENT '备注',
  `is_admin` tinyint unsigned DEFAULT '0' COMMENT '0 管理员 1 普通',
  `wallet_auth` varchar(150) DEFAULT '' COMMENT '钱包权限',
  `enabled` tinyint unsigned DEFAULT '0' COMMENT '状态 0正常、1禁用',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='成员管理';

-- ----------------------------
-- Table structure for month_bill
-- ----------------------------
DROP TABLE IF EXISTS `month_bill`;
CREATE TABLE `month_bill` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(150) DEFAULT '' COMMENT '钱包名称',
  `address` varchar(255) DEFAULT '' COMMENT '钱包地址',
  `desc` varchar(255) DEFAULT '' COMMENT '归集钱包备注',
  `date` int unsigned DEFAULT '0' COMMENT '月份',
  `amount` decimal(4,0) unsigned NOT NULL DEFAULT '0',
  `state` tinyint unsigned DEFAULT '0' COMMENT '状态 0未付清、1已付清',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='月付账单';

-- ----------------------------
-- Table structure for msw_stat
-- ----------------------------
DROP TABLE IF EXISTS `msw_stat`;
CREATE TABLE `msw_stat` (
  `id` int NOT NULL AUTO_INCREMENT,
  `tid` int DEFAULT '0' COMMENT '多签id',
  `date` date DEFAULT NULL COMMENT '日期',
  `trans_num` int DEFAULT '0' COMMENT '交易数',
  `in_num` int DEFAULT '0' COMMENT '收入笔数',
  `out_num` int DEFAULT '0' COMMENT '支出笔数',
  `in_amount` double DEFAULT '0' COMMENT '流入金额',
  `out_amount` double DEFAULT '0' COMMENT '流出金额',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=204 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='多签钱包每日统计';

-- ----------------------------
-- Table structure for multi_sign_wallet_detail
-- ----------------------------
DROP TABLE IF EXISTS `multi_sign_wallet_detail`;
CREATE TABLE `multi_sign_wallet_detail` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `wallet_id` bigint unsigned DEFAULT NULL COMMENT '多签钱包ID',
  `ordinary_wallet_id` bigint unsigned DEFAULT NULL COMMENT '普通钱包ID',
  `status` tinyint(1) DEFAULT '1' COMMENT '签名状态：1-待签，4-驳回，127-完成',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='多签钱包详情/成员表';

-- ----------------------------
-- Table structure for ordinary_wallet
-- ----------------------------
DROP TABLE IF EXISTS `ordinary_wallet`;
CREATE TABLE `ordinary_wallet` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '钱包别名',
  `address` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '钱包公钥地址',
  `currency_id` bigint unsigned DEFAULT NULL COMMENT '币种ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `address` (`address`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='元钱包表';

-- ----------------------------
-- Table structure for sys_transaction
-- ----------------------------
DROP TABLE IF EXISTS `sys_transaction`;
CREATE TABLE `sys_transaction` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `block_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `block_hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `from_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `to_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `gas` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `transaction_index` bigint DEFAULT NULL,
  `timestamp` bigint DEFAULT NULL,
  `nonce` int DEFAULT NULL,
  `block_height` int DEFAULT NULL,
  `method` int DEFAULT '0' COMMENT '0-默认，1-to和from钱包都在系统中 2-单from在系统中（转出） 3-单to在系统中（转入）',
  `status` int DEFAULT NULL,
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=110 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='平台交易数据';

-- ----------------------------
-- Table structure for transaction
-- ----------------------------
DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `chain_transaction_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `type` tinyint unsigned DEFAULT NULL COMMENT '交易类型 1 转出，2 转入，3 其它',
  `sub_type` tinyint unsigned DEFAULT '0' COMMENT '子类型:0-缺省/默认，1-创建多签，2-修改多签签名',
  `comment` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '交易备注',
  `multi_sign_wallet_id` bigint unsigned DEFAULT '0' COMMENT '发起方钱包ID',
  `destination_wallet` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '0',
  `amount` double DEFAULT '0' COMMENT '金额',
  `gas_amount` double DEFAULT '0' COMMENT 'gas金额',
  `status` tinyint unsigned DEFAULT '1' COMMENT '交易状态：127-成功，1-已发起，4-失败',
  `serial_number` varchar(36) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '交易流水号',
  `coin_id` tinyint DEFAULT '1' COMMENT '币种id',
  `user_id` bigint unsigned DEFAULT '0' COMMENT '发起人id',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=67 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for transaction_detail
-- ----------------------------
DROP TABLE IF EXISTS `transaction_detail`;
CREATE TABLE `transaction_detail` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `transaction_id` bigint unsigned DEFAULT '0' COMMENT '交易ID',
  `ordinary_wallet_id` bigint unsigned DEFAULT '0' COMMENT '普通钱包ID',
  `gas_amount` double DEFAULT '0' COMMENT 'gas金额',
  `status` tinyint unsigned DEFAULT '0' COMMENT '交易状态 127:已签 1:待签 4:驳回',
  `chain_transaction_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '链上cid',
  `serial_number` varchar(36) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '交易流水号',
  `is_sponsor` tinyint(1) DEFAULT '0' COMMENT '是否是发起者',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=131 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for wallet
-- ----------------------------
DROP TABLE IF EXISTS `wallet`;
CREATE TABLE `wallet` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '钱包别名',
  `address` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '钱包地址',
  `currency_id` bigint unsigned DEFAULT '0' COMMENT '币种id',
  `is_multi_sign` tinyint(1) DEFAULT '2' COMMENT '是否多签:1-是，2-否',
  `is_collect` tinyint(1) DEFAULT '2' COMMENT '是否归集：1-归集，2-普通',
  `required_signer` int unsigned DEFAULT '1' COMMENT '签名个数',
  `balance` double DEFAULT '0' COMMENT '余额',
  `status` tinyint unsigned DEFAULT '1' COMMENT '签名状态：1-待签，4-失败，127-成功',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `address` (`address`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='钱包表';

-- ----------------------------
-- Table structure for wallet_authorize
-- ----------------------------
DROP TABLE IF EXISTS `wallet_authorize`;
CREATE TABLE `wallet_authorize` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `wallet_id` bigint unsigned DEFAULT NULL COMMENT '多签钱包ID',
  `member_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='钱包权限表';

SET FOREIGN_KEY_CHECKS = 1;
