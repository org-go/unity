CREATE TABLE `d_task_order` (
  `pk` bigint(20) NOT NULL AUTO_INCREMENT,
  `worker_id` int(11) DEFAULT NULL COMMENT '接单人ID',
  `task_id` bigint(20) DEFAULT NULL COMMENT '任务ID',
  `order_number` varchar(64) DEFAULT NULL COMMENT '任务接单号（任务ID+自定义）',
  `score` int(11) DEFAULT NULL COMMENT '任务积分',
  `step` json DEFAULT NULL COMMENT '完成步骤',
  `created_time` timestamp NULL DEFAULT NULL COMMENT '创建接单时间',
  `updated_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新接单时间',
  `evaluate` tinyint(4) DEFAULT NULL COMMENT '评价 1：未完成 2：未达标 10：已达标',
  `auditor` varchar(20) DEFAULT NULL COMMENT '评审人',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_任务订单表'