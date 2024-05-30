CREATE TABLE `dm_activity` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `act_id` varchar(32) NOT NULL DEFAULT '' COMMENT '活动ID',
  `act_name` varchar(30) NOT NULL DEFAULT '' COMMENT '活动名称',
  `act_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '活动类型',
  `act_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '活动状态',
  `start_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '开始时间',
  `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '结束时间',
  `version` tinyint(4) unsigned DEFAULT '0' COMMENT '版本号',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_act_id` (`act_id`),
  KEY `idx_updated_at` (`updated_at`),
  KEY `idx_act_type_act_status` (`act_type`, `act_status`),
  KEY `idx_start_time_end_time` (`start_time`,`end_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='活动表';