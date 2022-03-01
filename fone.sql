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

 Date: 22/09/2021 12:36:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `mobile` varchar(16) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '手机号',
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '密码',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '启用',
  `created` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` bigint NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1：已删除，0：未删除',
  `is_admin` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0 超级管理员 1 普通管理员',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=22224 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='超管表';

-- ----------------------------
-- Table structure for admin_bak
-- ----------------------------
DROP TABLE IF EXISTS `admin_bak`;
CREATE TABLE `admin_bak` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `email` varchar(16) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '邮箱',
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '密码',
  `is_admin` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0 超级管理员 1普通管理员',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '启用',
  `created` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` bigint NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1：已删除，0：未删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='超级管理员表';

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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3 COMMENT='币种管理';

-- ----------------------------
-- Table structure for collect_transaction
-- ----------------------------
DROP TABLE IF EXISTS `collect_transaction`;
CREATE TABLE `collect_transaction` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `from_address` varchar(128) COLLATE utf8_bin DEFAULT NULL,
  `to_address` varchar(128) COLLATE utf8_bin DEFAULT NULL,
  `amount` double DEFAULT NULL,
  `cid` varchar(256) COLLATE utf8_bin DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for company
-- ----------------------------
DROP TABLE IF EXISTS `company`;
CREATE TABLE `company` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '企业名称',
  `email` varchar(255) DEFAULT '' COMMENT '企业邮箱',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3 COMMENT='企业管理';

-- ----------------------------
-- Table structure for company_member
-- ----------------------------
DROP TABLE IF EXISTS `company_member`;
CREATE TABLE `company_member` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `company_id` int NOT NULL COMMENT '企业表ID',
  `name` varchar(255) DEFAULT '' COMMENT '姓名',
  `password` varchar(32) DEFAULT '' COMMENT '密码',
  `email` varchar(150) DEFAULT '' COMMENT '邮箱',
  `desc` varchar(255) DEFAULT '' COMMENT '备注',
  `is_admin` tinyint unsigned DEFAULT '0' COMMENT '是否管理员：1 管理员 0 普通',
  `enabled` tinyint unsigned DEFAULT '1' COMMENT '状态 1正常、0禁用',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb3 COMMENT='成员管理';

-- ----------------------------
-- Table structure for fil_wallet_meta
-- ----------------------------
DROP TABLE IF EXISTS `fil_wallet_meta`;
CREATE TABLE `fil_wallet_meta` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `wallet_id` int unsigned DEFAULT '0' COMMENT '钱包id',
  `is_multi_wallet` tinyint unsigned DEFAULT '2' COMMENT '是否多签：1-是，2-否',
  `required_signer` int unsigned DEFAULT '1' COMMENT '需要签名人数',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='FIL钱包详情表';

-- ----------------------------
-- Table structure for month_bill
-- ----------------------------
DROP TABLE IF EXISTS `month_bill`;
CREATE TABLE `month_bill` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '钱包名称',
  `month` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '月份',
  `amount` decimal(4,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '金额',
  `pay_status` tinyint unsigned DEFAULT '0' COMMENT '状态 0未付清、1已付清',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  `wallet_id` bigint unsigned NOT NULL DEFAULT '0',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1 归集',
  `adress` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '钱包地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3 COMMENT='月付账单';

-- ----------------------------
-- Table structure for msw_stat
-- ----------------------------
DROP TABLE IF EXISTS `msw_stat`;
CREATE TABLE `msw_stat` (
  `id` int NOT NULL AUTO_INCREMENT,
  `wid` int DEFAULT '0' COMMENT '钱包id',
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
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='多签钱包详情/成员表';

-- ----------------------------
-- Table structure for multi_sign_wallet_meta
-- ----------------------------
DROP TABLE IF EXISTS `multi_sign_wallet_meta`;
CREATE TABLE `multi_sign_wallet_meta` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `wallet_id` bigint unsigned DEFAULT '0' COMMENT '多签钱包ID',
  `required_signer` bigint unsigned DEFAULT '0' COMMENT '签名人数',
  `is_multi_wallet` tinyint unsigned DEFAULT '2' COMMENT ' 是否多签：1-是，2-否',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='多签钱包表';

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
-- Table structure for transaction
-- ----------------------------
DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `chain_transaction_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `type` tinyint unsigned DEFAULT NULL COMMENT '交易类型 1 转出，2 转入，3 其它',
  `sub_type` tinyint unsigned DEFAULT '0' COMMENT '子类型:1-缺省/默认，2-创建多签，3-修改多签签名',
  `comment` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '交易备注',
  `from_wallet_id` bigint unsigned DEFAULT '0' COMMENT '发起方钱包ID',
  `destination_wallet` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '0',
  `amount` double DEFAULT '0' COMMENT '金额',
  `gas_amount` double DEFAULT '0' COMMENT 'gas金额',
  `status` tinyint unsigned DEFAULT '1' COMMENT '交易状态：127-成功，1-已发起，4-失败',
  `serial_number` varchar(36) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '交易流水号',
  `coin_id` tinyint DEFAULT '1' COMMENT '币种id',
  `company_member_id` bigint unsigned DEFAULT '0' COMMENT '发起人id',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='交易表';

-- ----------------------------
-- Table structure for transaction_detail
-- ----------------------------
DROP TABLE IF EXISTS `transaction_detail`;
CREATE TABLE `transaction_detail` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `transaction_id` bigint unsigned DEFAULT '0' COMMENT '交易ID',
  `wallet_id` bigint unsigned DEFAULT '0' COMMENT '普通钱包ID',
  `gas_amount` double DEFAULT '0' COMMENT 'gas金额',
  `status` tinyint unsigned DEFAULT '1' COMMENT '交易状态 127:已签 1:待签 4:驳回',
  `chain_transaction_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '链上cid',
  `serial_number` varchar(36) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '交易流水号',
  `is_sponsor` tinyint(1) DEFAULT '0' COMMENT '是否是发起者',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=133 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='交易详情';

-- ----------------------------
-- Table structure for wallet
-- ----------------------------
DROP TABLE IF EXISTS `wallet`;
CREATE TABLE `wallet` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '钱包别名',
  `address` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '钱包地址',
  `company_id` bigint unsigned DEFAULT '0',
  `coin_id` bigint unsigned DEFAULT '1' COMMENT '币种id',
  `balance` double DEFAULT '0' COMMENT '余额',
  `status` tinyint unsigned DEFAULT '1' COMMENT '签名状态：1-待签，4-失败，127-成功',
  `type_id` tinyint(1) DEFAULT '2' COMMENT '钱包种类id:1-多签，2-普通',
  `required_signer` int unsigned DEFAULT '1' COMMENT '签名人数',
  `is_collect` tinyint(1) DEFAULT '2' COMMENT '是否归集 1-是  2-否',
  `access_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT 'ak',
  `sign` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '签名',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `address` (`address`) USING BTREE,
  KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='钱包表';

-- ----------------------------
-- Table structure for wallet_authorize
-- ----------------------------
DROP TABLE IF EXISTS `wallet_authorize`;
CREATE TABLE `wallet_authorize` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `wallet_id` bigint unsigned DEFAULT '0' COMMENT '钱包ID',
  `company_member_id` bigint unsigned DEFAULT '0' COMMENT '用户ID',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='钱包权限表';

-- ----------------------------
-- Table structure for wallet_collect
-- ----------------------------
DROP TABLE IF EXISTS `wallet_collect`;
CREATE TABLE `wallet_collect` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `wallet_id` bigint DEFAULT '0' COMMENT '普通钱包id',
  `address` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '子钱包地址',
  `symbol` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '归集标识',
  `private_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '私钥',
  `balance` double unsigned DEFAULT '0' COMMENT '余额',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `address` (`address`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin COMMENT='归集钱包表';

-- ----------------------------
-- Table structure for wallet_type
-- ----------------------------
DROP TABLE IF EXISTS `wallet_type`;
CREATE TABLE `wallet_type` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `coin_id` int unsigned NOT NULL DEFAULT '0',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '钱包名称',
  `code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '代号',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3 COMMENT='币种管理';

SET FOREIGN_KEY_CHECKS = 1;
mysql> show tables;
Empty set (0.01 sec)
mysql> tee /Users/yangyi/Desktop/test.sql
