/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80018
 Source Host           : localhost:3306
 Source Schema         : digital

 Target Server Type    : MySQL
 Target Server Version : 80018
 File Encoding         : 65001

 Date: 22/09/2021 12:38:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `mobile` varchar(16) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '手机号',
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '密码',
  `enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '启用',
  `created` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1：已删除，0：未删除',
  `is_admin` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0 超级管理员 1 普通管理员',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=22223 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='超管表';

-- ----------------------------
-- Records of admin
-- ----------------------------
BEGIN;
INSERT INTO `admin` VALUES (1, '苏木', '15158830795', '21232f297a57a5a743894a0e4a801fc3', 0, 1620286492, 1620389003, 0, 0);
INSERT INTO `admin` VALUES (2, '邓智忠', '18565612837', '35853774658c7e83a3c610332ea9afe4', 1, 1620286492, 1620286492, 0, 0);
INSERT INTO `admin` VALUES (32, '京墨', '13136187803', 'e10adc3949ba59abbe56e057f20f883e', 1, 1620394150, 1620630657, 0, 0);
INSERT INTO `admin` VALUES (33, '椿木', '17310672593', 'e180f68a6504aa0862bbab477fbbab90', 1, 1620628231, 1621845618, 0, 0);
INSERT INTO `admin` VALUES (34, '张三', '15736875508', '53951213192b537af27af9a5719941aa', 1, 1620704133, 1620704133, 0, 0);
INSERT INTO `admin` VALUES (35, '测试账号', '13100009999', '82256468d61cea15e52fd68f367ad56b', 1, 1620812588, 1620812588, 0, 0);
INSERT INTO `admin` VALUES (36, '九胖', '15537636906', '3fcd44ae5d5510dce9dc581d13d3495c', 1, 1627644696, 1627820587, 0, 0);
INSERT INTO `admin` VALUES (37, '一桥', '18158514350', 'bc0af0a6226b15d6c1d79b1388b91e26', 1, 1628492651, 1628492651, 0, 0);
INSERT INTO `admin` VALUES (9999, '外部系统', '1', '1', 0, 0, 0, 0, 0);
INSERT INTO `admin` VALUES (22222, NULL, NULL, NULL, 0, 0, 0, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for chain_height
-- ----------------------------
DROP TABLE IF EXISTS `chain_height`;
CREATE TABLE `chain_height` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `height` bigint(20) DEFAULT NULL,
  `status` tinyint(3) unsigned DEFAULT '0' COMMENT '1:未成功，0:成功',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of chain_height
-- ----------------------------
BEGIN;
INSERT INTO `chain_height` VALUES (1, 12, 0, NULL, NULL, 0);
INSERT INTO `chain_height` VALUES (2, 107547, 0, NULL, NULL, 0);
INSERT INTO `chain_height` VALUES (3, 110380, 0, NULL, NULL, 0);
INSERT INTO `chain_height` VALUES (4, 140168, 0, '2021-08-17 16:30:24', '2021-08-17 17:32:08', 0);
INSERT INTO `chain_height` VALUES (5, 182838, 1, '2021-08-19 17:34:17', '2021-08-19 18:18:40', 0);
COMMIT;

-- ----------------------------
-- Table structure for chain_transaction
-- ----------------------------
DROP TABLE IF EXISTS `chain_transaction`;
CREATE TABLE `chain_transaction` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `block_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `block_hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `from_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `to_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `gas` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `value` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `transaction_index` bigint(20) DEFAULT NULL,
  `timestamp` bigint(20) DEFAULT NULL,
  `nonce` int(11) DEFAULT NULL,
  `block_height` int(11) DEFAULT NULL,
  `method` int(11) DEFAULT '0' COMMENT '0-默认，1-to和from钱包都在系统中 2-单from在系统中（转出） 3-单to在系统中（转入）',
  `status` int(11) DEFAULT NULL,
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=110 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='链上交易数据';

-- ----------------------------
-- Table structure for coin
-- ----------------------------
DROP TABLE IF EXISTS `coin`;
CREATE TABLE `coin` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '币种名称',
  `code` varchar(50) DEFAULT '' COMMENT '币种代号',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='币种管理';

-- ----------------------------
-- Records of coin
-- ----------------------------
BEGIN;
INSERT INTO `coin` VALUES (1, 'filecoin', 'file', 30, 330, 0);
INSERT INTO `coin` VALUES (2, '22', '333', 330, 330, 0);
COMMIT;

-- ----------------------------
-- Table structure for collect_wallet
-- ----------------------------
DROP TABLE IF EXISTS `collect_wallet`;
CREATE TABLE `collect_wallet` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `wallet_id` bigint(20) DEFAULT '0' COMMENT '普通钱包id',
  `address` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '钱包公钥地址',
  `access_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT 'ak',
  `sign` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '签名',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `address` (`address`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='归集钱包表';

-- ----------------------------
-- Table structure for company
-- ----------------------------
DROP TABLE IF EXISTS `company`;
CREATE TABLE `company` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '企业名称',
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '企业密码',
  `email` varchar(255) DEFAULT '' COMMENT '企业邮箱',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  `enabled` tinyint(3) unsigned DEFAULT '0' COMMENT '0 正常 禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='用户管理';

-- ----------------------------
-- Records of company
-- ----------------------------
BEGIN;
INSERT INTO `company` VALUES (11, '22', 'e10adc3949ba59abbe56e057f20f883e', '22', 1631423432, 1631952988, 0, 0);
INSERT INTO `company` VALUES (12, '网易', 'e10adc3949ba59abbe56e057f20f883e', 'wangyi.163.com', 1631518766, 1631518766, 0, 0);
INSERT INTO `company` VALUES (13, '有赞', 'e10adc3949ba59abbe56e057f20f883e', 'youzan.163.com', 1631688576, 1631688576, 0, 0);
INSERT INTO `company` VALUES (14, '有赞', 'e10adc3949ba59abbe56e057f20f883e', 'youzan.163.com', 1631692463, 1631692463, 0, 0);
INSERT INTO `company` VALUES (15, '有赞1', 'fcea920f7412b5da7be0cf42b8c93759', 'youzan.163.com', 1632275891, 1632275891, 0, 0);
INSERT INTO `company` VALUES (16, '公司2', '', '', 1632278115, 1632278115, 0, 0);
INSERT INTO `company` VALUES (17, '公司23', '', '', 1632278575, 1632278575, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for company_member
-- ----------------------------
DROP TABLE IF EXISTS `company_member`;
CREATE TABLE `company_member` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `company_id` int(11) NOT NULL COMMENT '企业表ID',
  `name` varchar(255) DEFAULT '' COMMENT '姓名',
  `password` varchar(32) DEFAULT '' COMMENT '密码',
  `email` varchar(150) DEFAULT '' COMMENT '邮箱',
  `desc` varchar(255) DEFAULT '' COMMENT '备注',
  `is_admin` tinyint(3) unsigned DEFAULT '0' COMMENT '0 普通  1管理员',
  `enabled` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 1正常 0 禁用',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 COMMENT='成员管理';

-- ----------------------------
-- Records of company_member
-- ----------------------------
BEGIN;
INSERT INTO `company_member` VALUES (1, 11, '阿里巴巴1', 'e10adc3949ba59abbe56e057f20f883e', 'alibba1@163.com', '淘宝变更了1', 0, 0, 1631423432, 1631686268, 0);
INSERT INTO `company_member` VALUES (13, 0, '百度', 'e10adc3949ba59abbe56e057f20f883e', 'baidu@qq.com', 'hah111a22a', 1, 0, 1631427701, 1631427701, 0);
INSERT INTO `company_member` VALUES (14, 0, '京东', 'e10adc3949ba59abbe56e057f20f883e', 'jingdong@qq.com', 'hah111a22a', 1, 0, 1631518716, 1631518716, 0);
INSERT INTO `company_member` VALUES (15, 12, '网易', 'e10adc3949ba59abbe56e057f20f883e', 'wangyi.163.com', 'hahhaa', 0, 0, 1631518766, 1631518766, 0);
INSERT INTO `company_member` VALUES (16, 0, '字节跳动', 'e10adc3949ba59abbe56e057f20f883e', 'tiaodong@qq.com', '字节跳动', 1, 0, 1631634048, 1631634048, 0);
INSERT INTO `company_member` VALUES (17, 0, '哈罗', 'e10adc3949ba59abbe56e057f20f883e', 'haluo@qq.com', 'haluo', 0, 0, 1631634221, 1631634221, 0);
INSERT INTO `company_member` VALUES (18, 13, '企业管理员', 'e10adc3949ba59abbe56e057f20f883e', 'youzan.163.com', 'youzan', 0, 0, 1631688576, 1631688576, 0);
INSERT INTO `company_member` VALUES (19, 14, '企业管理员', 'e10adc3949ba59abbe56e057f20f883e', 'youzan.163.com', 'youzan1', 0, 0, 1631692463, 1631692463, 0);
INSERT INTO `company_member` VALUES (20, 15, '企业管理员', 'fcea920f7412b5da7be0cf42b8c93759', 'youzan.163.com', 'youzan1', 0, 0, 1632275891, 1632275891, 0);
INSERT INTO `company_member` VALUES (21, 0, '哈罗2', 'e10adc3949ba59abbe56e057f20f883e', 'haluo2@qq.com', 'haluo1', 1, 0, 1632278850, 1632278850, 0);
INSERT INTO `company_member` VALUES (22, 12, '北京公司', 'e10adc3949ba59abbe56e057f20f883e', 'beijing@qq.com', 'haluo1', 1, 0, 1632279045, 1632279045, 0);
INSERT INTO `company_member` VALUES (23, 0, '北京公司', 'e10adc3949ba59abbe56e057f20f883e', 'beijingq2@qq.com', 'haluo1', 1, 0, 1632281434, 1632281434, 0);
INSERT INTO `company_member` VALUES (24, 0, '北京公司', 'e10adc3949ba59abbe56e057f20f883e', '2302728693@qq.com', 'haluo1', 1, 0, 1632281474, 1632281474, 0);
INSERT INTO `company_member` VALUES (25, 0, '大白兔', 'e10adc3949ba59abbe56e057f20f883e', '23027286932@qq.com', 'haluo1', 1, 0, 1632284576, 1632284576, 0);
INSERT INTO `company_member` VALUES (26, 0, '大白兔2', 'e10adc3949ba59abbe56e057f20f883e', '230272869322@qq.com', 'haluo1', 1, 1, 1632285065, 1632285065, 0);
COMMIT;

-- ----------------------------
-- Table structure for month_bill
-- ----------------------------
DROP TABLE IF EXISTS `month_bill`;
CREATE TABLE `month_bill` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `month` varchar(32) DEFAULT '0' COMMENT '月份',
  `amount` decimal(4,0) unsigned NOT NULL DEFAULT '0' COMMENT '金额',
  `pay_status` tinyint(3) unsigned DEFAULT '0' COMMENT '状态 0未付清、1已付清',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  `wallet_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `name` varchar(32) DEFAULT NULL COMMENT '钱包别名',
  `adress` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='月付账单';

-- ----------------------------
-- Records of month_bill
-- ----------------------------
BEGIN;
INSERT INTO `month_bill` VALUES (1, '202109', 1, 1, 0, 0, 0, 1, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for msw_stat
-- ----------------------------
DROP TABLE IF EXISTS `msw_stat`;
CREATE TABLE `msw_stat` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `wid` int(11) DEFAULT '0' COMMENT '钱包id',
  `date` date DEFAULT NULL COMMENT '日期',
  `trans_num` int(11) DEFAULT '0' COMMENT '交易数',
  `in_num` int(11) DEFAULT '0' COMMENT '收入笔数',
  `out_num` int(11) DEFAULT '0' COMMENT '支出笔数',
  `in_amount` double DEFAULT '0' COMMENT '流入金额',
  `out_amount` double DEFAULT '0' COMMENT '流出金额',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=204 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='多签钱包每日统计';

-- ----------------------------
-- Records of msw_stat
-- ----------------------------
BEGIN;
INSERT INTO `msw_stat` VALUES (50, 16, '2021-08-16', 9, 0, 9, 0, 488, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (51, 17, '2021-08-16', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (52, 18, '2021-08-16', 4, 4, 0, 158, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (53, 19, '2021-08-16', 3, 3, 0, 120, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (54, 20, '2021-08-16', 0, 0, 0, 0, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (55, 21, '2021-08-16', 0, 0, 0, 0, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (56, 22, '2021-08-16', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (57, 16, '2021-08-17', 1, 3, 2, 0, 1.2, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (58, 17, '2021-08-17', 1, 3, 2, 0, 1.2, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (59, 18, '2021-08-17', 1, 3, 2, 0, 1.2, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (60, 19, '2021-08-17', 1, 3, 2, 0, 1.2, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (61, 20, '2021-08-17', 1, 3, 2, 0, 1.2, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (62, 21, '2021-08-17', 1, 3, 2, 0, 1.2, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (63, 22, '2021-08-17', 1, 3, 2, 0, 1.2, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (64, 16, '2021-08-18', 1, 3, 2, 0, 1.2, 20210819173417, 20210819181840, 0);
INSERT INTO `msw_stat` VALUES (65, 17, '2021-08-18', 1, 3, 2, 0, 1.2, 20210819173417, 20210819181840, 0);
INSERT INTO `msw_stat` VALUES (66, 18, '2021-08-18', 1, 3, 2, 0, 1.2, 20210819173417, 20210819181840, 0);
INSERT INTO `msw_stat` VALUES (67, 19, '2021-08-18', 1, 1, 2, 0, 1.2, 20210819173417, 20210819181840, 0);
INSERT INTO `msw_stat` VALUES (68, 20, '2021-08-18', 1, 1, 2, 0, 1.2, 20210819173417, 20210819181840, 0);
INSERT INTO `msw_stat` VALUES (69, 21, '2021-08-18', 1, 1, 2, 0, 1.2, 20210819173417, 20210819181840, 0);
INSERT INTO `msw_stat` VALUES (70, 22, '2021-08-18', 1, 1, 2, 0, 1.2, 20210819173417, 20210819181840, 0);
INSERT INTO `msw_stat` VALUES (71, 16, '2021-06-16', 9, 4, 9, 0, 488, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (72, 17, '2021-06-16', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (73, 18, '2021-06-16', 4, 4, 12, 158, 23.44, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (74, 19, '2021-06-16', 3, 3, 2, 120, 34.53, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (75, 20, '2021-06-16', 2, 2, 32, 3134, 23, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (76, 21, '2021-06-16', 2, 3, 2, 2, 2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (77, 22, '2021-06-16', 1, 1, 23, 40, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (78, 16, '2021-07-17', 1, 16, 3, 23.5, 23.5, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (79, 17, '2021-07-17', 1, 17, 3, 23.5, 23.5, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (80, 18, '2021-07-17', 1, 18, 3, 23.5, 23.5, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (81, 19, '2021-07-17', 1, 19, 3, 23.5, 23.5, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (82, 20, '2021-07-17', 1, 20, 3, 23.5, 23.5, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (83, 21, '2021-07-17', 1, 21, 3, 23.5, 23.5, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (84, 22, '2021-07-17', 1, 22, 3, 23.5, 23.5, 20210818000000, 20210818180500, 0);
INSERT INTO `msw_stat` VALUES (85, 16, '2021-06-18', 16, 16, 16, 16, 16, 20210819173417, 20210819173417, 0);
INSERT INTO `msw_stat` VALUES (86, 17, '2021-06-18', 17, 17, 17, 17, 17, 20210819173417, 20210819173417, 0);
INSERT INTO `msw_stat` VALUES (87, 18, '2021-06-18', 18, 18, 18, 18, 18, 20210819173417, 20210819173417, 0);
INSERT INTO `msw_stat` VALUES (88, 19, '2021-06-18', 19, 19, 19, 19, 19, 20210819173417, 20210819173417, 0);
INSERT INTO `msw_stat` VALUES (89, 20, '2021-06-18', 20, 20, 20, 20, 20, 20210819173417, 20210819173417, 0);
INSERT INTO `msw_stat` VALUES (90, 21, '2021-06-18', 21, 21, 21, 21, 21, 20210819173417, 20210819173417, 0);
INSERT INTO `msw_stat` VALUES (91, 22, '2021-06-18', 22, 22, 22, 22, 22, 20210819173417, 20210819173417, 0);
INSERT INTO `msw_stat` VALUES (92, 16, '2021-06-05', 9, 12, 9, 12, 488, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (93, 17, '2021-06-05', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (94, 18, '2021-06-05', 4, 4, 0, 158, 16, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (95, 19, '2021-06-05', 3, 3, 0, 120, 17, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (96, 20, '2021-06-05', 0, 0, 0, 0, 18, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (97, 21, '2021-06-05', 0, 0, 0, 0, 19, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (98, 22, '2021-06-05', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (99, 16, '2021-06-06', 9, 0, 9, 0, 488, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (100, 17, '2021-06-06', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (101, 18, '2021-06-06', 4, 4, 0, 158, 16, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (102, 19, '2021-06-06', 3, 3, 0, 120, 17, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (103, 20, '2021-06-06', 0, 0, 0, 0, 18, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (104, 21, '2021-06-06', 0, 0, 0, 0, 19, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (105, 22, '2021-06-06', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (106, 16, '2021-06-07', 9, 0, 9, 0, 488, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (107, 17, '2021-06-07', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (108, 18, '2021-06-07', 4, 4, 0, 158, 16, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (109, 19, '2021-06-07', 3, 3, 0, 120, 17, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (110, 20, '2021-06-07', 0, 0, 0, 0, 18, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (111, 21, '2021-06-07', 0, 0, 0, 0, 19, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (112, 22, '2021-06-07', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (113, 16, '2021-06-08', 9, 0, 9, 0, 42, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (114, 17, '2021-06-08', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (115, 18, '2021-06-08', 4, 4, 0, 158, 16, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (116, 19, '2021-06-08', 3, 3, 0, 120, 17, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (117, 20, '2021-06-08', 0, 0, 0, 0, 18, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (118, 21, '2021-06-08', 0, 0, 0, 0, 19, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (119, 22, '2021-06-08', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (120, 16, '2021-06-09', 9, 0, 9, 0, 488, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (121, 17, '2021-06-09', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (122, 18, '2021-06-09', 4, 4, 0, 158, 16, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (123, 19, '2021-06-09', 3, 3, 0, 120, 17, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (124, 20, '2021-06-09', 0, 0, 0, 0, 18, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (125, 21, '2021-06-09', 0, 0, 0, 0, 19, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (126, 22, '2021-06-09', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (127, 16, '2021-06-10', 9, 0, 9, 0, 488, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (128, 17, '2021-06-10', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (129, 18, '2021-06-10', 4, 4, 0, 158, 16, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (130, 19, '2021-06-10', 3, 3, 0, 120, 17, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (131, 20, '2021-06-10', 0, 0, 0, 0, 18, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (132, 21, '2021-06-10', 0, 0, 0, 0, 19, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (133, 22, '2021-06-10', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (134, 16, '2021-08-07', 9, 0, 9, 0, 2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (135, 17, '2021-08-07', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (136, 18, '2021-08-07', 4, 4, 0, 158, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (137, 19, '2021-08-07', 3, 3, 0, 120, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (138, 20, '2021-08-07', 0, 0, 0, 0, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (139, 21, '2021-08-07', 0, 0, 0, 0, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (140, 22, '2021-08-07', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (141, 16, '2021-08-06', 9, 0, 9, 0, 4, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (142, 17, '2021-08-06', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (143, 18, '2021-08-06', 4, 4, 0, 158, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (144, 19, '2021-08-06', 3, 3, 0, 120, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (145, 20, '2021-08-06', 0, 0, 0, 0, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (146, 21, '2021-08-06', 0, 0, 0, 0, 0, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (147, 22, '2021-08-06', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (148, 16, '2021-08-09', 9, 0, 9, 0, 488, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (149, 17, '2021-08-09', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (150, 18, '2021-08-09', 4, 4, 1, 158, 20, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (151, 19, '2021-08-09', 3, 3, 1, 120, 12, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (152, 20, '2021-08-09', 4, 4, 1, 2, 4, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (153, 21, '2021-08-09', 3, 3, 1, 5, 34, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (154, 22, '2021-08-09', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (155, 16, '2021-08-10', 9, 0, 9, 0, 488, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (156, 17, '2021-08-10', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (157, 18, '2021-08-10', 4, 4, 1, 158, 20, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (158, 19, '2021-08-10', 3, 3, 1, 120, 12, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (159, 20, '2021-08-10', 4, 4, 1, 2, 4, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (160, 21, '2021-08-10', 3, 3, 1, 5, 34, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (161, 22, '2021-08-10', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (162, 16, '2021-08-11', 9, 0, 9, 0, 488, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (163, 17, '2021-08-11', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (164, 18, '2021-08-11', 4, 4, 1, 158, 20, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (165, 19, '2021-08-11', 3, 3, 1, 120, 12, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (166, 20, '2021-08-11', 4, 4, 1, 2, 4, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (167, 21, '2021-08-11', 3, 3, 1, 5, 34, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (168, 22, '2021-08-11', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (169, 16, '2021-08-12', 9, 0, 9, 0, 488, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (170, 17, '2021-08-12', 3, 1, 2, 50, 80, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (171, 18, '2021-08-12', 4, 4, 1, 158, 20, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (172, 19, '2021-08-12', 3, 3, 1, 120, 12, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (173, 20, '2021-08-12', 4, 4, 1, 2, 4, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (174, 21, '2021-08-12', 3, 3, 1, 5, 34, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (175, 22, '2021-08-12', 1, 1, 2, 40, 1.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (176, 16, '2021-08-13', 9, 3, 9, 0, 16.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (177, 17, '2021-08-13', 3, 1, 2, 50, 17, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (178, 18, '2021-08-13', 4, 4, 1, 158, 18.1, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (179, 19, '2021-08-13', 3, 3, 1, 120, 19, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (180, 20, '2021-08-13', 4, 4, 1, 2, 20.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (181, 21, '2021-08-13', 3, 3, 1, 5, 21, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (182, 22, '2021-08-13', 1, 1, 2, 40, 22, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (183, 16, '2021-08-14', 9, 3, 9, 0, 16.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (184, 17, '2021-08-14', 3, 1, 2, 50, 17, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (185, 18, '2021-08-14', 4, 4, 1, 158, 18.1, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (186, 19, '2021-08-14', 3, 3, 1, 120, 19, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (187, 20, '2021-08-14', 4, 4, 1, 2, 20.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (188, 21, '2021-08-14', 3, 3, 1, 5, 21, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (189, 22, '2021-08-14', 1, 1, 2, 40, 22, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (197, 16, '2021-08-15', 9, 3, 9, 0, 16.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (198, 17, '2021-08-15', 3, 1, 2, 50, 17, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (199, 18, '2021-08-15', 4, 4, 1, 158, 18.1, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (200, 19, '2021-08-15', 3, 3, 1, 120, 19, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (201, 20, '2021-08-15', 4, 4, 1, 2, 20.2, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (202, 21, '2021-08-15', 3, 3, 1, 5, 21, 20210817201859, 20210817235500, 0);
INSERT INTO `msw_stat` VALUES (203, 22, '2021-08-15', 1, 1, 2, 40, 22, 20210817201859, 20210817235500, 0);
COMMIT;

-- ----------------------------
-- Table structure for multi_sign_wallet_detail
-- ----------------------------
DROP TABLE IF EXISTS `multi_sign_wallet_detail`;
CREATE TABLE `multi_sign_wallet_detail` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `wallet_id` bigint(20) unsigned DEFAULT NULL COMMENT '多签钱包ID',
  `ordinary_wallet_id` bigint(20) unsigned DEFAULT NULL COMMENT '普通钱包ID',
  `status` tinyint(1) DEFAULT '1' COMMENT '签名状态：1-待签，4-驳回，127-完成',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='多签钱包详情/成员表';

-- ----------------------------
-- Table structure for multi_sign_wallet_meta
-- ----------------------------
DROP TABLE IF EXISTS `multi_sign_wallet_meta`;
CREATE TABLE `multi_sign_wallet_meta` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `wallet_id` bigint(20) unsigned DEFAULT '0' COMMENT '多签钱包ID',
  `required_signer` bigint(20) unsigned DEFAULT '0' COMMENT '签名人数',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='多签钱包表';

-- ----------------------------
-- Table structure for ordinary_wallet
-- ----------------------------
DROP TABLE IF EXISTS `ordinary_wallet`;
CREATE TABLE `ordinary_wallet` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '钱包别名',
  `address` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '钱包公钥地址',
  `currency_id` bigint(20) unsigned DEFAULT NULL COMMENT '币种ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `address` (`address`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='元钱包表';

-- ----------------------------
-- Table structure for transaction
-- ----------------------------
DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `chain_transaction_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `type` tinyint(3) unsigned DEFAULT NULL COMMENT '交易类型 1 转出，2 转入，3 其它',
  `sub_type` tinyint(3) unsigned DEFAULT '0' COMMENT '子类型:0-缺省/默认，1-创建多签，2-修改多签签名',
  `comment` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '交易备注',
  `from_wallet_id` bigint(20) unsigned DEFAULT '0' COMMENT '发起方钱包ID',
  `destination_wallet` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '0',
  `amount` double DEFAULT '0' COMMENT '金额',
  `gas_amount` double DEFAULT '0' COMMENT 'gas金额',
  `status` tinyint(3) unsigned DEFAULT '1' COMMENT '交易状态：127-成功，1-已发起，4-失败',
  `serial_number` varchar(36) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '交易流水号',
  `coin_id` tinyint(4) DEFAULT '1' COMMENT '币种id',
  `company_member_id` bigint(20) unsigned DEFAULT '0' COMMENT '发起人id',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=67 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of transaction
-- ----------------------------
BEGIN;
INSERT INTO `transaction` VALUES (10, '1', 1, 1, '123', 1, '1', 1, 1, 1, '1', 1, 1, 0, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for transaction_detail
-- ----------------------------
DROP TABLE IF EXISTS `transaction_detail`;
CREATE TABLE `transaction_detail` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `transaction_id` bigint(20) unsigned DEFAULT '0' COMMENT '交易ID',
  `ordinary_wallet_id` bigint(20) unsigned DEFAULT '0' COMMENT '普通钱包ID',
  `gas_amount` double DEFAULT '0' COMMENT 'gas金额',
  `status` tinyint(3) unsigned DEFAULT '0' COMMENT '交易状态 127:已签 1:待签 4:驳回',
  `chain_transaction_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '链上cid',
  `serial_number` varchar(36) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '交易流水号',
  `is_sponsor` tinyint(1) DEFAULT '0' COMMENT '是否是发起者',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  `wallet_id` bigint(20) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=131 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of transaction_detail
-- ----------------------------
BEGIN;
INSERT INTO `transaction_detail` VALUES (10, 10, 10, 10, 20, '2', '3', 30, 30, 30, 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for wallet
-- ----------------------------
DROP TABLE IF EXISTS `wallet`;
CREATE TABLE `wallet` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '钱包别名',
  `address` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '钱包地址',
  `company_id` bigint(20) unsigned DEFAULT '0',
  `coin_id` bigint(20) unsigned DEFAULT '1' COMMENT '币种id',
  `balance` double DEFAULT '0' COMMENT '余额',
  `status` tinyint(3) unsigned DEFAULT '1' COMMENT '签名状态：1-待签，4-失败，127-成功',
  `type_id` tinyint(1) DEFAULT '2' COMMENT '钱包种类id:1-多签，2-普通',
  `required_signer` int(10) unsigned DEFAULT '0' COMMENT '签名人数',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `address` (`address`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='钱包表';

-- ----------------------------
-- Records of wallet
-- ----------------------------
BEGIN;
INSERT INTO `wallet` VALUES (1, '2', '123233', 1, 1, 0, 1, 2, 0, 20210903150214, 20210903150214, 0);
INSERT INTO `wallet` VALUES (2, 'haha', '123', 1, 1, 0, 1, 2, 0, 0, 0, 0);
INSERT INTO `wallet` VALUES (3, '123123', '1233', 2, 1, 0, 1, 2, 0, 0, 0, 0);
INSERT INTO `wallet` VALUES (4, '333', '22', 2, 1, 0, 1, 2, 0, 0, 0, 0);
INSERT INTO `wallet` VALUES (5, '222', '33', 0, 1, 0, 1, 2, 0, 0, 0, 0);
INSERT INTO `wallet` VALUES (10, '666', '66', 0, 1, 0, 1, 2, 0, 0, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for wallet_authorize
-- ----------------------------
DROP TABLE IF EXISTS `wallet_authorize`;
CREATE TABLE `wallet_authorize` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `wallet_id` bigint(20) unsigned DEFAULT '0' COMMENT '多签钱包ID',
  `company_member_id` bigint(20) unsigned DEFAULT '0' COMMENT '用户ID',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='钱包权限表';

-- ----------------------------
-- Records of wallet_authorize
-- ----------------------------
BEGIN;
INSERT INTO `wallet_authorize` VALUES (19, 1, 12, 1631428754, 1631428754, 0);
INSERT INTO `wallet_authorize` VALUES (20, 2, 12, 1631428754, 1631428754, 0);
INSERT INTO `wallet_authorize` VALUES (21, 3, 12, 1631428754, 1631428754, 0);
INSERT INTO `wallet_authorize` VALUES (37, 4, 13, 1631697504, 1631697504, 0);
INSERT INTO `wallet_authorize` VALUES (38, 5, 13, 1631697504, 1631697504, 0);
INSERT INTO `wallet_authorize` VALUES (39, 6, 13, 1631697504, 1631697504, 0);
INSERT INTO `wallet_authorize` VALUES (51, 1, 17, 1631697798, 1631697798, 0);
INSERT INTO `wallet_authorize` VALUES (52, 2, 17, 1631697798, 1631697798, 0);
INSERT INTO `wallet_authorize` VALUES (53, 3, 17, 1631697798, 1631697798, 0);
INSERT INTO `wallet_authorize` VALUES (56, 1, 14, 1631843355, 1631843355, 0);
INSERT INTO `wallet_authorize` VALUES (57, 2, 14, 1631843355, 1631843355, 0);
COMMIT;

-- ----------------------------
-- Table structure for wallet_type
-- ----------------------------
DROP TABLE IF EXISTS `wallet_type`;
CREATE TABLE `wallet_type` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `coin_id` int(10) unsigned NOT NULL DEFAULT '0',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '钱包名称',
  `code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '代号',
  `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0: 未删除 1: 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='币种管理';

-- ----------------------------
-- Records of wallet_type
-- ----------------------------
BEGIN;
INSERT INTO `wallet_type` VALUES (1, 1, 'fil:多签', 'fil:msig', 0, 0, 0);
INSERT INTO `wallet_type` VALUES (2, 1, 'fil:普通', 'fil:nomal', 0, 0, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
