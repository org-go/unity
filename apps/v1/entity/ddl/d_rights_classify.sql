CREATE TABLE `d_rights_classify` (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL COMMENT '名称',
  `supperior_pk` int(11) DEFAULT NULL COMMENT '上级PK',
  `description` varchar(100) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`pk`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_权益类目表'