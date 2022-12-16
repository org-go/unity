CREATE TABLE `d_user_oauth_account` (
  `user_pk` bigint(20) NOT NULL,
  `wechat` varchar(100) DEFAULT NULL,
  `qq` varchar(100) DEFAULT NULL,
  `dingtalk` varchar(100) DEFAULT NULL,
  `github` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_用户三方账户表'