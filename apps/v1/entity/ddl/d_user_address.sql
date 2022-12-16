CREATE TABLE `d_user_address` (
  `pk` bigint(22) NOT NULL AUTO_INCREMENT,
  `user_pk` bigint(22) NOT NULL COMMENT '用户PK',
  `state_code` int(11) DEFAULT NULL COMMENT '省',
  `city_code` int(11) DEFAULT NULL COMMENT '市',
  `street_code` int(11) DEFAULT NULL COMMENT '街道',
  `address` varchar(60) DEFAULT '' COMMENT '详细地址',
  `postal_code` int(11) DEFAULT NULL COMMENT '邮编',
  `created_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_default` tinyint(4) DEFAULT NULL COMMENT '1: 默认',
  PRIMARY KEY (`pk`) USING BTREE,
  KEY `user_pk_default` (`user_pk`,`is_default`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_用户地址表'