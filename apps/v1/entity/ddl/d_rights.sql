CREATE TABLE `d_rights` (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL COMMENT '权益名称',
  `rights_classify_pk` int(11) DEFAULT NULL COMMENT '权益类目PK',
  `description` varchar(100) DEFAULT NULL COMMENT '权益描述',
  `status` tinyint(4) DEFAULT NULL COMMENT '权益状态 1:启用 2：禁用',
  `created_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_权益表'