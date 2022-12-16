CREATE TABLE `d_welfare_classify` (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `platform_pk` int(11) DEFAULT NULL COMMENT '平台PK',
  `name` varchar(20) DEFAULT NULL COMMENT '福利类目名称',
  `description` varchar(200) DEFAULT '' COMMENT '描述',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态; 1:启用 2：禁用',
  `created_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
  PRIMARY KEY (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_福利类目表'