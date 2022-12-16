CREATE TABLE `d_platform` (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `created_time` timestamp NULL DEFAULT NULL,
  `updated_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_平台表'