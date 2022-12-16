CREATE TABLE `d_user_portrait` (
  `user_pk` bigint(20) NOT NULL,
  `unit_id` varchar(20) DEFAULT NULL COMMENT '身份证',
  `borth` varchar(20) DEFAULT NULL COMMENT '生日',
  `hobbies` varchar(100) DEFAULT NULL COMMENT '兴趣爱好',
  `email` varchar(20) DEFAULT NULL COMMENT '邮箱',
  PRIMARY KEY (`user_pk`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_用户基础画像表'