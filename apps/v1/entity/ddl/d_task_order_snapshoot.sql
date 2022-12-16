CREATE TABLE `d_task_order_snapshoot` (
  `pk` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_number` varchar(64) DEFAULT NULL COMMENT '任务订单号',
  `task_owner_name` varchar(20) DEFAULT NULL COMMENT '任务负责人',
  `task_id` bigint(20) DEFAULT NULL COMMENT '任务ID',
  `task_name` varchar(200) DEFAULT NULL COMMENT '任务名称',
  `start_time` timestamp NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` timestamp NULL DEFAULT NULL COMMENT '结束时间',
  `worker_id` varchar(20) DEFAULT NULL COMMENT '接单人',
  `claim` json DEFAULT NULL COMMENT '要求',
  `step` json DEFAULT NULL COMMENT '步骤',
  `score` int(11) DEFAULT NULL COMMENT '积分',
  `evaluate` tinyint(255) DEFAULT NULL COMMENT '评价',
  `created_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC