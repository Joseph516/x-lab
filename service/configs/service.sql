CREATE DATABASE IF NOT EXISTS x_lab
    DEFAULT CHARACTER
        SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

USE x_lab;

# 用户表
CREATE TABLE IF NOT EXISTS users
(
    `id`          int(4) unsigned AUTO_INCREMENT NOT NULL,
    `username`    varchar(64) DEFAULT '',
    `password`    varchar(64) DEFAULT '',
    `status`      tinyint(4)  DEFAULT '0',
    `create_time` int(10)     DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';