CREATE TABLE `d_task` (
  `pk` bigint(20) NOT NULL AUTO_INCREMENT,
  `id` bigint(20) DEFAULT NULL COMMENT '任务ID',
  `task_classify_pk` int(11) DEFAULT NULL COMMENT '任务类目PK',
  `owner_user_id` bigint(20) DEFAULT NULL COMMENT '任务负责人',
  `name` varchar(20) DEFAULT '' COMMENT '名称',
  `title` varchar(60) DEFAULT '' COMMENT '标题',
  `description` varchar(200) DEFAULT '' COMMENT '描述',
  `claim` json DEFAULT NULL COMMENT '要求',
  `task_score` int(11) DEFAULT NULL COMMENT '任务积分/单',
  `task_number` int(11) DEFAULT NULL COMMENT '单数',
  `task_start_time` datetime DEFAULT NULL COMMENT '任务开始时间',
  `task_end_time` datetime DEFAULT NULL COMMENT '任务结束时间',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态：1:草稿， 3：审核中， 5：进行中， 7：下架， 9：结束',
  `created_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `auditor` varchar(20) DEFAULT '' COMMENT '审核人',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_任务表'