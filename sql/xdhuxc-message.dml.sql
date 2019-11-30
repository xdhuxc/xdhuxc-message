USE `xdhuxc-message`;

ALTER TABLE `xdhuxc-message_message` ADD COLUMN `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间';


ALTER TABLE `xdhuxc-message_message` ADD COLUMN `user` varchar(100) NOT NULL COMMENT '使用者';

ALTER TABLE `xdhuxc-message_message` ADD FULLTEXT INDEX ft_index 
    (`user`, `sender`, `message_type`, `content`, `description`, `email_type`, `subject`) WITH PARSER ngram;

SHOW VARIABLES LIKE "ngram%";

SELECT * FROM `xdhuxc-message_message` WHERE (Match (`user`, sender, message_type, content, description, email_type, subject) AGAINST ('+"email"' IN BOOLEAN MODE)) ORDER BY update_time desc;  
 
SELECT count(*) FROM `xdhuxc-message_message` WHERE (Match (`user`, sender, message_type, content, description, email_type, subject) AGAINST ('dingTalk' IN BOOLEAN MODE));