
#### 创建表都是以tik_作为前缀进行表的设置
#用户表
CREATE TABLE `tik_user` (
                            `user_id` BIGINT not null auto_increment,		#自增创建ID
                            `name` 		VARCHAR(40) DEFAULT '',
                            `follow_count` int DEFAULT '0',
                            `follower_count` int DEFAULT '0',
                            `is_follow` bool DEFAULT '0',
                            `password` VARCHAR(40) DEFAULT '',
                            PRIMARY KEY(`user_id`)
)