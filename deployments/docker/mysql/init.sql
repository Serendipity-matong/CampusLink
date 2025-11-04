-- 创建数据库
CREATE DATABASE IF NOT EXISTS campuslink DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE campuslink;

-- 用户表
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `student_id` varchar(50) DEFAULT NULL COMMENT '学号',
  `real_name` varchar(50) DEFAULT NULL COMMENT '真实姓名',
  `school` varchar(100) DEFAULT NULL COMMENT '学校',
  `nickname` varchar(50) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `bio` varchar(500) DEFAULT NULL COMMENT '个人简介',
  `role` varchar(20) DEFAULT 'student' COMMENT '角色',
  `verified` tinyint(1) DEFAULT 0 COMMENT '是否认证',
  `status` varchar(20) DEFAULT 'active' COMMENT '状态',
  `credit_score` int DEFAULT 100 COMMENT '信用分',
  `balance` decimal(10,2) DEFAULT 0.00 COMMENT '余额',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`),
  UNIQUE KEY `idx_phone` (`phone`),
  KEY `idx_student_id` (`student_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 商品表
CREATE TABLE IF NOT EXISTS `product` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL COMMENT '商品名称',
  `description` text COMMENT '商品描述',
  `category` varchar(50) DEFAULT NULL COMMENT '分类',
  `product_type` varchar(20) DEFAULT 'sale' COMMENT '类型: sale/rent',
  `price` decimal(10,2) NOT NULL COMMENT '价格',
  `rent_price_daily` decimal(10,2) DEFAULT NULL COMMENT '日租金',
  `stock` int DEFAULT 0 COMMENT '库存',
  `images` json DEFAULT NULL COMMENT '图片列表',
  `seller_id` bigint unsigned NOT NULL COMMENT '卖家ID',
  `condition` varchar(20) DEFAULT NULL COMMENT '成色',
  `location` varchar(100) DEFAULT NULL COMMENT '位置',
  `status` varchar(20) DEFAULT 'offline' COMMENT '状态',
  `view_count` bigint DEFAULT 0 COMMENT '浏览量',
  `favorite_count` bigint DEFAULT 0 COMMENT '收藏量',
  `extra` json DEFAULT NULL COMMENT '额外信息',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_seller_id` (`seller_id`),
  KEY `idx_category` (`category`),
  KEY `idx_product_type` (`product_type`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品表';

-- 任务表
CREATE TABLE IF NOT EXISTS `task` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL COMMENT '任务标题',
  `description` text COMMENT '任务描述',
  `task_type` varchar(20) NOT NULL COMMENT '任务类型',
  `reward` decimal(10,2) NOT NULL COMMENT '报酬',
  `pickup_location` varchar(200) DEFAULT NULL COMMENT '取件地点',
  `delivery_location` varchar(200) DEFAULT NULL COMMENT '送件地点',
  `deadline` timestamp NULL DEFAULT NULL COMMENT '截止时间',
  `publisher_id` bigint unsigned NOT NULL COMMENT '发布者ID',
  `executor_id` bigint unsigned DEFAULT NULL COMMENT '接单者ID',
  `contact_phone` varchar(20) DEFAULT NULL COMMENT '联系电话',
  `images` json DEFAULT NULL COMMENT '图片',
  `status` varchar(20) DEFAULT 'pending' COMMENT '状态',
  `view_count` bigint DEFAULT 0 COMMENT '浏览量',
  `extra` json DEFAULT NULL COMMENT '额外信息',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `accepted_at` timestamp NULL DEFAULT NULL COMMENT '接单时间',
  `completed_at` timestamp NULL DEFAULT NULL COMMENT '完成时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_publisher_id` (`publisher_id`),
  KEY `idx_executor_id` (`executor_id`),
  KEY `idx_task_type` (`task_type`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='任务表';

-- 订单表
CREATE TABLE IF NOT EXISTS `order` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `order_no` varchar(50) NOT NULL COMMENT '订单号',
  `buyer_id` bigint unsigned NOT NULL COMMENT '买家ID',
  `seller_id` bigint unsigned NOT NULL COMMENT '卖家ID',
  `order_type` varchar(20) NOT NULL COMMENT '订单类型',
  `total_amount` decimal(10,2) NOT NULL COMMENT '总金额',
  `paid_amount` decimal(10,2) DEFAULT 0.00 COMMENT '实付金额',
  `delivery_address` varchar(500) DEFAULT NULL COMMENT '配送地址',
  `contact_phone` varchar(20) DEFAULT NULL COMMENT '联系电话',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `status` varchar(20) DEFAULT 'pending' COMMENT '状态',
  `payment_method` varchar(20) DEFAULT NULL COMMENT '支付方式',
  `payment_time` timestamp NULL DEFAULT NULL COMMENT '支付时间',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `completed_at` timestamp NULL DEFAULT NULL COMMENT '完成时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_order_no` (`order_no`),
  KEY `idx_buyer_id` (`buyer_id`),
  KEY `idx_seller_id` (`seller_id`),
  KEY `idx_order_type` (`order_type`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单表';

-- 订单项表
CREATE TABLE IF NOT EXISTS `order_item` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `order_id` bigint unsigned NOT NULL COMMENT '订单ID',
  `product_id` bigint unsigned NOT NULL COMMENT '商品ID',
  `product_name` varchar(200) NOT NULL COMMENT '商品名称',
  `quantity` int NOT NULL COMMENT '数量',
  `price` decimal(10,2) NOT NULL COMMENT '单价',
  `rent_days` int DEFAULT NULL COMMENT '租借天数',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单项表';

-- 支付表
CREATE TABLE IF NOT EXISTS `payment` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `payment_no` varchar(50) NOT NULL COMMENT '支付单号',
  `order_id` bigint unsigned NOT NULL COMMENT '订单ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `amount` decimal(10,2) NOT NULL COMMENT '支付金额',
  `payment_method` varchar(20) NOT NULL COMMENT '支付方式',
  `status` varchar(20) DEFAULT 'pending' COMMENT '状态',
  `transaction_id` varchar(100) DEFAULT NULL COMMENT '第三方交易号',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `paid_at` timestamp NULL DEFAULT NULL COMMENT '支付时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_payment_no` (`payment_no`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='支付表';

-- 通知表
CREATE TABLE IF NOT EXISTS `notification` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `title` varchar(200) NOT NULL COMMENT '标题',
  `content` text COMMENT '内容',
  `notification_type` varchar(20) NOT NULL COMMENT '类型',
  `is_read` tinyint(1) DEFAULT 0 COMMENT '是否已读',
  `extra` json DEFAULT NULL COMMENT '额外信息',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_notification_type` (`notification_type`),
  KEY `idx_is_read` (`is_read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='通知表';

-- 举报表
CREATE TABLE IF NOT EXISTS `report` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `reporter_id` bigint unsigned NOT NULL COMMENT '举报人ID',
  `report_type` varchar(20) NOT NULL COMMENT '举报类型',
  `target_id` bigint unsigned NOT NULL COMMENT '目标ID',
  `reason` varchar(100) NOT NULL COMMENT '原因',
  `description` text COMMENT '详细描述',
  `images` json DEFAULT NULL COMMENT '图片',
  `status` varchar(20) DEFAULT 'pending' COMMENT '状态',
  `handler_id` bigint unsigned DEFAULT NULL COMMENT '处理人ID',
  `handle_result` varchar(500) DEFAULT NULL COMMENT '处理结果',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `handled_at` timestamp NULL DEFAULT NULL COMMENT '处理时间',
  PRIMARY KEY (`id`),
  KEY `idx_reporter_id` (`reporter_id`),
  KEY `idx_report_type` (`report_type`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='举报表';

-- 插入测试数据
INSERT INTO `user` (`username`, `password`, `phone`, `student_id`, `real_name`, `school`, `role`, `verified`)
VALUES 
('admin', '$2a$10$rPPZa7zqLsLbqELqJVvO9.hZOEHd5JGJqJVvO9.hZOEHd5JGJqJV', '13800138000', 'S2024001', '管理员', '某某大学', 'admin', 1),
('test_user', '$2a$10$rPPZa7zqLsLbqELqJVvO9.hZOEHd5JGJqJVvO9.hZOEHd5JGJqJV', '13800138001', 'S2024002', '测试用户', '某某大学', 'student', 1);


