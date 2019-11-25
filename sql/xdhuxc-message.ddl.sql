CREATE DATABASE IF NOT EXISTS `xdhuxc-message`;

USE `xdhuxc-message`;

# DDL
CREATE TABLE IF NOT EXISTS `xdhuxc-message_message`
(
    `id`           int(11)      NOT NULL AUTO_INCREMENT COMMENT '自增 id',
    `user`       varchar(100) NOT NULL COMMENT '发送者',
    `sender`       varchar(100) NOT NULL COMMENT '发送者',
    `message_type`         varchar(100) NOT NULL COMMENT '消息中介类型',
    `is_sent`      tinyint(1)  NOT NULL DEFAULT '0' COMMENT '是否已发送',
    `content`   text NOT NULL COMMENT '消息内容',
    `description`   varchar(1024) DEFAULT '' COMMENT '消息描述',
    `receivers`   json NOT NULL COMMENT '消息接收者数组',
    `cc`  json          DEFAULT NULL COMMENT '邮件抄送者数组',
    `email_type` varchar(50) DEFAULT '' COMMENT '邮件内容类型',
    `subject`      varchar(512) DEFAULT '' COMMENT '邮件消息主题',
    `create_time`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
    `update_time`       datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `xdhuxc-message_audit`
(
    `id`           int(11)      NOT NULL AUTO_INCREMENT COMMENT '自增 id',
    `user` varchar(100) NOT NULL COMMENT '操作者',
    `operate`         varchar(100) NOT NULL COMMENT '操作',
    `object`     varchar(50)  NOT NULL COMMENT '操作对象',
    `create_time`   datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
