/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50739 (5.7.39)
 Source Host           : 127.0.0.1:3306
 Source Schema         : nft_platform

 Target Server Type    : MySQL
 Target Server Version : 50739 (5.7.39)
 File Encoding         : 65001

 Date: 24/04/2023 20:33:49
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for nft_banner
-- ----------------------------
DROP TABLE IF EXISTS `nft_banner`;
CREATE TABLE `nft_banner`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title`     varchar(50)  DEFAULT NULL COMMENT '标题',
    `image`     varchar(255) DEFAULT NULL COMMENT 'banner图片',
    `link_addr` varchar(255) DEFAULT NULL COMMENT '链接地址',
    `status`    tinyint(4) unsigned DEFAULT '1' COMMENT '是否启用 1 启用 2 下架',
    `mark`      varchar(255) DEFAULT NULL COMMENT '备注',
    `type`      tinyint(4) DEFAULT NULL COMMENT '类型 1、自营平台 2、转售平台',
    `dt_create` int(11) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update` int(11) unsigned DEFAULT '0' COMMENT '跟新时间',
    PRIMARY KEY (`id`),
    KEY         `type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Banner list';

-- ----------------------------
-- Table structure for nft_complex_conf
-- ----------------------------
DROP TABLE IF EXISTS `nft_complex_conf`;
CREATE TABLE `nft_complex_conf`
(
    `id`         int(11) NOT NULL,
    `title`      varchar(255) DEFAULT NULL COMMENT '合成标题',
    `complex_id` int(11) unsigned DEFAULT '0' COMMENT '合成藏品',
    `goods_id`   int(11) unsigned DEFAULT '0' COMMENT '碎片藏品',
    `status`     tinyint(3) unsigned DEFAULT '1' COMMENT '是否可用',
    `dt_create`  int(11) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`  int(11) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY          `complex_id` (`complex_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合成碎片表';

-- ----------------------------
-- Table structure for nft_goods
-- ----------------------------
DROP TABLE IF EXISTS `nft_goods`;
CREATE TABLE `nft_goods`
(
    `id`                    int(10) unsigned NOT NULL AUTO_INCREMENT,
    `no`                    varchar(20)  DEFAULT NULL COMMENT '批次',
    `name`                  varchar(20)  DEFAULT NULL COMMENT '藏品名称',
    `image`                 varchar(255) DEFAULT NULL COMMENT '藏品首图 400*400',
    `type`                  tinyint(4) DEFAULT NULL COMMENT '藏品类型 1普通藏品 2盲盒',
    `detail`                varchar(255) DEFAULT NULL COMMENT '藏品详情图',
    `label`                 varchar(255) DEFAULT NULL COMMENT '标签 ， |分割',
    `price`                 int(11) unsigned DEFAULT '0' COMMENT '价格，单位 分',
    `apple_product_id`      varchar(255) DEFAULT NULL COMMENT '苹果IAP 商品ID',
    `is_show`               tinyint(4) unsigned DEFAULT '1' COMMENT '是否显示',
    `has_saled`             tinyint(4) unsigned DEFAULT '0' COMMENT '是否售罄',
    `sale_starttime`        int(10) unsigned DEFAULT '0' COMMENT '销售时间',
    `sale_endtime`          int(11) unsigned DEFAULT '0' COMMENT '销售结束时间',
    `appointment_starttime` int(10) unsigned DEFAULT '0' COMMENT '预约时间',
    `appointment_endtime`   int(11) unsigned DEFAULT '0' COMMENT '预约结束时间',
    `stock`                 int(11) unsigned DEFAULT '0' COMMENT '库存',
    `sales_volume`          int(11) unsigned DEFAULT '0' COMMENT '销量',
    `publisher_id`          int(11) unsigned DEFAULT '0' COMMENT '发行商',
    `source`                varchar(255) DEFAULT NULL COMMENT '资源地址',
    `source_type`           tinyint(4) unsigned DEFAULT '2' COMMENT '资源类型 1 图片 2 3D',
    `has_cast`              tinyint(4) unsigned DEFAULT '0' COMMENT '是否已铸造',
    `dt_create`             int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`             int(10) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='自营平台藏品表';

-- ----------------------------
-- Table structure for nft_goods_stocks
-- ----------------------------
DROP TABLE IF EXISTS `nft_goods_stocks`;
CREATE TABLE `nft_goods_stocks`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
    `goods_id`   int(11) unsigned DEFAULT '0' COMMENT '藏品ID',
    `goods_no`   varchar(255) DEFAULT NULL COMMENT '藏品编号',
    `goods_hash` varchar(64)  DEFAULT NULL COMMENT '藏品HASH地址',
    `status`     tinyint(4) unsigned DEFAULT '1' COMMENT '是否可用 0不可用 1 可用',
    `user_id`    varchar(20)  DEFAULT NULL COMMENT '用户编号',
    `dt_create`  int(11) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`  int(10) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY          `goods_id` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='自营平台库存表';

-- ----------------------------
-- Table structure for nft_mall_goods
-- ----------------------------
DROP TABLE IF EXISTS `nft_mall_goods`;
CREATE TABLE `nft_mall_goods`
(
    `id`             int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`           varchar(50)  DEFAULT NULL COMMENT '藏品名称',
    `image`          varchar(255) DEFAULT NULL COMMENT '藏品首图',
    `detal`          varchar(255) DEFAULT NULL COMMENT '藏品介绍',
    `price`          int(11) unsigned DEFAULT '0' COMMENT '藏品出售价格',
    `original_price` int(11) unsigned DEFAULT '0' COMMENT '购买价格',
    `platform_id`    int(11) unsigned DEFAULT '0' COMMENT '所属平台',
    `no`             varchar(255) DEFAULT NULL COMMENT '编号',
    `source`         varchar(255) DEFAULT NULL COMMENT '藏品资源',
    `source_type`    tinyint(4) unsigned DEFAULT '1' COMMENT '资源类型 1 图片 2 3D',
    `hash`           varchar(255) DEFAULT NULL COMMENT 'HASH地址',
    `status`         tinyint(4) unsigned DEFAULT '1' COMMENT '是否可售 1是 0下架 2 已售',
    `dt_create`      int(11) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`      int(11) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='寄售平台藏品表';

-- ----------------------------
-- Table structure for nft_mall_orders
-- ----------------------------
DROP TABLE IF EXISTS `nft_mall_orders`;
CREATE TABLE `nft_mall_orders`
(
    `id`               int(10) unsigned NOT NULL AUTO_INCREMENT,
    `order_id`         varchar(30) DEFAULT NULL COMMENT '订单号',
    `goods_id`         int(10) unsigned DEFAULT '0' COMMENT '藏品ID',
    `mall_platform_id` int(11) unsigned DEFAULT '0' COMMENT '平台ID',
    `from_user_id`     varchar(20) DEFAULT '0' COMMENT '出售用户编号',
    `user_id`          varchar(20) DEFAULT NULL COMMENT '用户ID',
    `num`              int(10) unsigned DEFAULT '1' COMMENT '购买数量',
    `amount`           int(10) unsigned DEFAULT '0' COMMENT '订单金额 单位 分',
    `pay_channel_id`   int(10) unsigned DEFAULT '0' COMMENT '支付渠道',
    `pay_endtime`      int(10) unsigned DEFAULT '0' COMMENT '支付超时时间',
    `status`           tinyint(3) unsigned DEFAULT '1' COMMENT '订单状态 1待支付 2 已支付 3 已取消 4 超时 5 已完成',
    `dt_create`        int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`        int(11) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY                `order_id` (`order_id`),
    KEY                `user_id` (`user_id`),
    KEY                `goods_id` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='寄售平台藏品';

-- ----------------------------
-- Table structure for nft_mall_platform
-- ----------------------------
DROP TABLE IF EXISTS `nft_mall_platform`;
CREATE TABLE `nft_mall_platform`
(
    `id`             int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`           varchar(50)  DEFAULT NULL COMMENT '平台名称',
    `logo`           varchar(255) DEFAULT NULL COMMENT '平台logo',
    `site`           varchar(255) DEFAULT NULL COMMENT '平台web地址',
    `api_user`       varchar(255) DEFAULT NULL COMMENT '登记api',
    `api_goods_list` varchar(255) DEFAULT NULL COMMENT '藏品列表API',
    `api_transfer`   varchar(255) DEFAULT NULL COMMENT '藏品转移API',
    `api_sales`      varchar(255) DEFAULT NULL COMMENT '藏品上架销售查询API',
    `status`         tinyint(3) unsigned DEFAULT '1' COMMENT '是否可用 1可用 0 不可用',
    `dt_create`      int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`      int(11) unsigned DEFAULT '0' COMMENT '跟新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='寄售平台类型表';

-- ----------------------------
-- Table structure for nft_mall_user_goods
-- ----------------------------
DROP TABLE IF EXISTS `nft_mall_user_goods`;
CREATE TABLE `nft_mall_user_goods`
(
    `id`               int(10) unsigned NOT NULL AUTO_INCREMENT,
    `goods_id`         int(11) unsigned DEFAULT '0' COMMENT '藏品ID',
    `goods_name`       varchar(20) DEFAULT NULL COMMENT '藏品名称',
    `goods_no`         varchar(10) DEFAULT NULL COMMENT '藏品编号',
    `from_user_id`     varchar(20) DEFAULT NULL COMMENT '出售用户ID',
    `mall_platform_id` varchar(20) DEFAULT NULL COMMENT '平台ID',
    `user_id`          varchar(20) DEFAULT NULL COMMENT '用户编号',
    `goods_hash`       varchar(64) DEFAULT NULL COMMENT '藏品HASh',
    `status`           tinyint(3) unsigned DEFAULT '1' COMMENT '是否可用',
    `dt_create`        int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`        int(10) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY                `goods_id` (`goods_id`),
    KEY                `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='寄售平台用户藏品表';

-- ----------------------------
-- Table structure for nft_manghe_conf
-- ----------------------------
DROP TABLE IF EXISTS `nft_manghe_conf`;
CREATE TABLE `nft_manghe_conf`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT,
    `manghe_id` int(11) unsigned DEFAULT '0' COMMENT '盲盒ID',
    `goods_id`  int(11) unsigned DEFAULT '0' COMMENT '藏品ID\n',
    `rate`      int(11) unsigned DEFAULT '0' COMMENT '概率',
    `level`     varchar(255) DEFAULT NULL COMMENT '等级 SSSR SSR SR R',
    `status`    tinyint(4) unsigned DEFAULT '1' COMMENT '是否可用',
    `dt_create` int(11) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update` int(11) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY         `manghe_id` (`manghe_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='自营平台盲盒表';

-- ----------------------------
-- Table structure for nft_note
-- ----------------------------
DROP TABLE IF EXISTS `nft_note`;
CREATE TABLE `nft_note`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title`     varchar(255) DEFAULT NULL COMMENT '标题',
    `intro`     varchar(500) DEFAULT NULL COMMENT '简介',
    `content`   text COMMENT '内容',
    `status`    tinyint(4) unsigned DEFAULT '1' COMMENT '是否展示 0 不展示 1 展示',
    `type`      tinyint(3) unsigned DEFAULT '1' COMMENT '类型 1 自营平台2 寄售平台',
    `dt_create` int(10) unsigned DEFAULT NULL COMMENT '创建时间',
    `dt_update` int(10) unsigned DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='公告表';

-- ----------------------------
-- Table structure for nft_orders
-- ----------------------------
DROP TABLE IF EXISTS `nft_orders`;
CREATE TABLE `nft_orders`
(
    `id`             int(10) unsigned NOT NULL AUTO_INCREMENT,
    `order_id`       varchar(30) DEFAULT NULL COMMENT '订单号',
    `goods_id`       int(10) unsigned DEFAULT '0' COMMENT '藏品ID',
    `user_id`        int(10) unsigned DEFAULT '0' COMMENT '用户ID',
    `num`            int(10) unsigned DEFAULT '0' COMMENT '购买数量',
    `amount`         int(10) unsigned DEFAULT '0' COMMENT '订单金额 单位 分',
    `pay_channel_id` int(10) unsigned DEFAULT '0' COMMENT '支付渠道',
    `pay_endtime`    int(10) unsigned DEFAULT '0' COMMENT '支付超时时间',
    `status`         tinyint(3) unsigned DEFAULT '0' COMMENT '订单状态 1待支付 2 已支付 3 已取消 4 超时',
    `dt_create`      int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`      int(11) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY              `order_id` (`order_id`),
    KEY              `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='自营平台订单表';

-- ----------------------------
-- Table structure for nft_publisher
-- ----------------------------
DROP TABLE IF EXISTS `nft_publisher`;
CREATE TABLE `nft_publisher`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`      varchar(255) DEFAULT NULL COMMENT '名称',
    `image`     varchar(255) DEFAULT NULL COMMENT '头像',
    `status`    tinyint(4) unsigned DEFAULT '1' COMMENT '状态 0 禁用 1 启用',
    `dt_create` int(11) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update` int(11) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='发行商表';

-- ----------------------------
-- Table structure for nft_user_balance_log
-- ----------------------------
DROP TABLE IF EXISTS `nft_user_balance_log`;
CREATE TABLE `nft_user_balance_log`
(
    `id`          int(11) NOT NULL,
    `user_id`     varchar(20)  DEFAULT NULL COMMENT '用户编号',
    `amount`      int(11) unsigned DEFAULT '0' COMMENT '变动金额 + 出售 - 提现',
    `balance`     int(10) unsigned DEFAULT '0' COMMENT '账号余额',
    `goods_id`    int(10) unsigned DEFAULT '0' COMMENT '藏品ID',
    `goods_name`  varchar(20)  DEFAULT NULL COMMENT '藏品名称',
    `goods_image` varchar(255) DEFAULT NULL COMMENT '藏品图片',
    `info`        varchar(255) DEFAULT NULL COMMENT '说明',
    `type`        tinyint(3) unsigned DEFAULT '1' COMMENT '类型 1 藏品转售 2 提现 3 支付',
    `dt_create`   int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`   int(10) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户余额流水表';

-- ----------------------------
-- Table structure for nft_user_bank
-- ----------------------------
DROP TABLE IF EXISTS `nft_user_bank`;
CREATE TABLE `nft_user_bank`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`      varchar(20)  DEFAULT NULL COMMENT '用户Id',
    `bank`         varchar(10)  DEFAULT NULL COMMENT '开户行',
    `bank_name`    varchar(10)  DEFAULT NULL COMMENT '银行户名',
    `bank_account` varchar(30)  DEFAULT NULL COMMENT '银行账号',
    `bank_address` varchar(255) DEFAULT NULL COMMENT '开户地址',
    `status`       tinyint(4) unsigned DEFAULT '1' COMMENT '是否可用 1 可用 0 不可用',
    `dt_create`    int(11) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`    int(11) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户银行账号表';

-- ----------------------------
-- Table structure for nft_user_goods
-- ----------------------------
DROP TABLE IF EXISTS `nft_user_goods`;
CREATE TABLE `nft_user_goods`
(
    `id`               int(10) unsigned NOT NULL AUTO_INCREMENT,
    `goods_id`         int(11) unsigned DEFAULT '0' COMMENT '藏品ID',
    `goods_name`       varchar(255) DEFAULT NULL COMMENT '藏品名称',
    `goods_no`         varchar(10)  DEFAULT NULL COMMENT '藏品编号',
    `user_id`          varchar(20)  DEFAULT NULL COMMENT '用户编号',
    `goods_type`       tinyint(4) DEFAULT NULL COMMENT '藏品类型 1 普通藏品 2 盲盒',
    `goods_hash`       varchar(64)  DEFAULT NULL COMMENT '藏品HASh',
    `from_wallet_addr` varchar(64)  DEFAULT NULL COMMENT '来源用户钱包地址，空来源系统发行',
    `status`           tinyint(3) unsigned DEFAULT '1' COMMENT '是否可用',
    `dt_create`        int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`        int(10) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY                `goods_id` (`goods_id`),
    KEY                `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='自营平台用户藏品表';

-- ----------------------------
-- Table structure for nft_user_msg
-- ----------------------------
DROP TABLE IF EXISTS `nft_user_msg`;
CREATE TABLE `nft_user_msg`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title`        varchar(255) DEFAULT NULL COMMENT '标题',
    `content`      varchar(500) DEFAULT NULL COMMENT '消息内容',
    `user_id`      varchar(20)  DEFAULT NULL COMMENT '用户编号',
    `from_user_id` varchar(20)  DEFAULT NULL COMMENT '来自',
    `has_read`     tinyint(4) DEFAULT '0' COMMENT '是否一度',
    `dt_create`    int(11) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`    int(11) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY            `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户消息表';

-- ----------------------------
-- Table structure for nft_user_withdraw
-- ----------------------------
DROP TABLE IF EXISTS `nft_user_withdraw`;
CREATE TABLE `nft_user_withdraw`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`      varchar(20)  DEFAULT NULL COMMENT '用户编号',
    `amount`       int(11) unsigned DEFAULT '0' COMMENT '提现金额',
    `bank`         varchar(10)  DEFAULT NULL COMMENT '开户行',
    `bank_name`    varchar(10)  DEFAULT NULL COMMENT '户名',
    `bank_account` varchar(20)  DEFAULT NULL COMMENT '账号',
    `bank_addr`    varchar(255) DEFAULT NULL COMMENT '开户地址',
    `status`       tinyint(4) unsigned DEFAULT '1' COMMENT '状态，1 审核中 2 待打款 3 已打款 4 撤销',
    `dt_create`    int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `dt_update`    int(10) unsigned DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY            `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户提现表';

-- ----------------------------
-- Table structure for nft_users
-- ----------------------------
DROP TABLE IF EXISTS `nft_users`;
CREATE TABLE `nft_users`
(
    `id`                 int(10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`            varchar(20)  DEFAULT NULL COMMENT '用户ID',
    `username`           varchar(20)  DEFAULT NULL COMMENT '账号',
    `password`           varchar(32)  DEFAULT NULL COMMENT '密码',
    `nickname`           varchar(20)  DEFAULT NULL COMMENT '昵称',
    `mobile`             varchar(11)  DEFAULT NULL COMMENT '手机号',
    `avatar`             varchar(255) DEFAULT NULL COMMENT '头像',
    `wallet_addr`        varchar(64)  DEFAULT NULL COMMENT '钱包地址',
    `wallet_private_key` varchar(255) DEFAULT NULL COMMENT '钱包私钥',
    `status`             tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，1 正常 0 冻结',
    `has_real_auth`      tinyint(4) DEFAULT '0' COMMENT '是否实名认证',
    `real_name`          varchar(10)  DEFAULT NULL COMMENT '真实姓名',
    `card_no`            varchar(18)  DEFAULT NULL COMMENT '身份证',
    `amount`             int(10) unsigned DEFAULT '0' COMMENT '账号余额',
    `freeze_amount`      int(11) unsigned DEFAULT '0' COMMENT '冻结金额',
    `dt_create`          int(11) DEFAULT '0' COMMENT '创建时间',
    `dt_update`          int(11) DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_id` (`user_id`),
    UNIQUE KEY `card_no` (`card_no`),
    UNIQUE KEY `mobile` (`mobile`),
    UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';


CREATE TABLE `nft_mall_platform_account`
(
    `id`          int UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     varchar(50) NULL COMMENT '用户编号',
    `wallet_hash` varchar(255) NULL COMMENT '钱包地址',
    `platform_id` int NULL COMMENT '平台ID',
    `status`      tinyint NULL DEFAULT 1 COMMENT '是否可用',
    `dt_create`   int UNSIGNED NULL COMMENT '创建时间',
    `dt_update`   int UNSIGNED NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
);
SET
FOREIGN_KEY_CHECKS = 1;
