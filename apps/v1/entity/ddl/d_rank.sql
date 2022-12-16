CREATE TABLE `d_rank` (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT '' COMMENT '头衔名称',
  `experience` bigint(20) DEFAULT NULL COMMENT '晋级所需经验',
  `description` varchar(100) DEFAULT '' COMMENT '头衔描述',
  `rights_pks` varchar(100) DEFAULT '' COMMENT '权益PK, 逗号(,)分隔',
  `created_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_头衔表'