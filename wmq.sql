-- -------------------------------------------
-- wmq-admin database wmq_admin
-- author: phachon@163.com
-- -------------------------------------------

-- -------------------------------------------
-- wmq user table
-- -------------------------------------------
DROP TABLE IF EXISTS `wmq_user`;
CREATE TABLE `wmq_user` (
  `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` char(18) NOT NULL DEFAULT '' COMMENT '手机',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 0 正常 -1 删除',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `wmq_user` (`name`, `email`, `password`,  `mobile`, `status`, `create_time`, `update_time`)
VALUES ('root', 'root@123456.com', 'e10adc3949ba59abbe56e057f20f883e', '1102222', '0', '1460557068', '1461407549');