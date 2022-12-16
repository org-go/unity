CREATE TABLE `d_rank_rights` (
  `pk` int(11) NOT NULL AUTO_INCREMENT,
  `rights_pk` int(11) DEFAULT NULL COMMENT '权益PK',
  `rank_pk` int(11) DEFAULT NULL COMMENT '头衔PK',
  PRIMARY KEY (`pk`) USING BTREE,
  KEY `rank_rights_pks` (`rank_pk`,`rights_pk`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='data_头衔权益关联表1N'