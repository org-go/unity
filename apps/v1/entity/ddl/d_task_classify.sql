CREATE TABLE `d_task_classify` (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `platform_pk` int(11) DEFAULT NULL COMMENT '平台PK',
  `name` varchar(20) DEFAULT '' COMMENT '任务类目名称',
  `description` varchar(200) DEFAULT '' COMMENT '描述',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态：1：启用， 2：禁用',
  `created_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_user` varchar(20) DEFAULT '' COMMENT '创建人',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_任务类目表'