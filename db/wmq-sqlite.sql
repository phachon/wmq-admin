-- -------------------------------------------
-- type: sqlite
-- database: wmq-admin
-- author: phachon@163.com
-- -------------------------------------------

-- -------------------------------------------
-- wmq user table 用户表
-- -------------------------------------------
DROP TABLE IF EXISTS `wmq_user`;
CREATE TABLE `wmq_user` (
  `user_id` integer PRIMARY KEY AUTOINCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '',
  `email` varchar(50) NOT NULL DEFAULT '',
  `password` char(32) NOT NULL DEFAULT '',
  `mobile` char(18) NOT NULL DEFAULT '',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0',
  `create_time` integer unsigned NOT NULL DEFAULT '0',
  `update_time` integer unsigned NOT NULL DEFAULT '0'
);
INSERT INTO `wmq_user` (`name`, `email`, `password`,  `mobile`, `is_delete`, `create_time`, `update_time`)
VALUES ('root', 'root@123456.com', 'e10adc3949ba59abbe56e057f20f883e', '1102222', '0', '1460557068', '1461407549');

-- -------------------------------------------
-- wmq node table 节点表
-- -------------------------------------------
DROP TABLE IF EXISTS `wmq_node`;
CREATE TABLE `wmq_node` (
  `node_id` integer PRIMARY KEY AUTOINCREMENT,
  `ip` varchar(20) NOT NULL DEFAULT '',
  `manager_port` int(8) NOT NULL DEFAULT '0',
  `message_port` int(8) NOT NULL DEFAULT '0',
  `token` VARCHAR(32) NOT NULL DEFAULT '',
  `comment` VARCHAR(200) NOT NULL DEFAULT '',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0',
  `create_time` integer unsigned NOT NULL DEFAULT '0',
  `update_time` integer unsigned NOT NULL DEFAULT '0'
);

