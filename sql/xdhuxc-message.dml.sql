USE `xdhuxc-message`;

ALTER TABLE `xdhuxc-message_message` ADD COLUMN `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间';


ALTER TABLE `xdhuxc-message_message` ADD COLUMN `user` varchar(100) NOT NULL COMMENT '使用者';
