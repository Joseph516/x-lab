CREATE DATABASE IF NOT EXISTS x_lab
    DEFAULT CHARACTER
        SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

USE x_lab;

# 用户表
CREATE TABLE IF NOT EXISTS users
(
    `id`          int(4) unsigned AUTO_INCREMENT NOT NULL,
    `username`    varchar(64) DEFAULT '' COMMENT '用户名',
    `password`    varchar(64) DEFAULT '' COMMENT '用户密码MD5',
    `create_time` int(10)     DEFAULT '0' COMMENT '创建时间',
    `status`      tinyint(4)  DEFAULT '0' COMMENT '状态0为正常，状态1为删除',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

# 文件表
CREATE TABLE IF NOT EXISTS files
(
    `id`              int(10) unsigned NOT NULL AUTO_INCREMENT,
    `filename`        varchar(100)        DEFAULT '' COMMENT '文件名称',
    `desc`            varchar(255)        DEFAULT '' COMMENT '文件简述',
    `file_access_url` varchar(255)        DEFAULT '' COMMENT '文件地址',
    `type`            tinyint(3)          DEFAULT '1' COMMENT '文件类型',
    ## 公共字段 ##
    `created_time`    int(10) unsigned    DEFAULT '0' COMMENT '创建时间',
    `created_by`      varchar(100)        DEFAULT '' COMMENT '创建人',
    `modified_time`   int(10) unsigned    DEFAULT '0' COMMENT '修改时间',
    `modified_by`     varchar(100)        DEFAULT '' COMMENT '修改人',
    `deleted_time`    int(10) unsigned    DEFAULT '0' COMMENT '删除时间',
    ## 公共字段 ##
    `status`          tinyint(3) unsigned DEFAULT '0' COMMENT '状态0为正常，状态1为删除',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='文件表';