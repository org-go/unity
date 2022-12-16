CREATE TABLE `d_welfare` (
  `pk` bigint(20) NOT NULL AUTO_INCREMENT,
  `id` bigint(20) DEFAULT NULL COMMENT '对外ID',
  `platform_pk` int(11) DEFAULT NULL COMMENT '平台PK',
  `welfare_classify_pk` int(11) DEFAULT NULL COMMENT '福利类型PK',
  `name` varchar(20) DEFAULT '' COMMENT '福利名称',
  `title` varchar(100) DEFAULT '' COMMENT '福利标题',
  `description` varchar(200) DEFAULT '' COMMENT '福利描述',
  `link` varchar(100) DEFAULT '' COMMENT '福利链接',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态; 1:启用； 2：禁用',
  `created_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_user` varchar(20) DEFAULT '' COMMENT '创建人',
  PRIMARY KEY (`pk`) USING BTREE,
  KEY `classify_pk` (`welfare_classify_pk`) USING BTREE,
  KEY `mix_name_title_status` (`name`,`title`,`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_福利信息表'