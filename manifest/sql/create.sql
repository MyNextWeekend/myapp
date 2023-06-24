-- 用户表
CREATE TABLE `user`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    `passport`  varchar(45) NOT NULL COMMENT '账号',
    `password`  varchar(45) NOT NULL COMMENT '密码',
    `nickname`  varchar(45) NOT NULL COMMENT '昵称',
    `create_at` datetime DEFAULT NULL COMMENT '创建时间',
    `update_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- 用例表
CREATE TABLE `api`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT,
    `name`        varchar(20)  NOT NULL COMMENT '名称',
    `describe`    varchar(200) NOT NULL COMMENT '描述',
    `type`        varchar(10)  NOT NULL COMMENT '请求类型',
    `uri`         varchar(30)  NOT NULL COMMENT 'ip地址',
    `path`        varchar(100) NOT NULL COMMENT '路径',
    `body`        varchar(200) NOT NULL COMMENT '请求体',
    `state`       tinyint(1) NOT NULL COMMENT '状态',
    `create_at`   datetime DEFAULT NULL COMMENT '创建时间',
    `create_user` varchar(30)  NOT NULL COMMENT '创建人',
    `update_at`   datetime DEFAULT NULL COMMENT '更新时间',
    `update_user` varchar(30)  NOT NULL COMMENT '更新人',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 用例表
CREATE TABLE `case`
(
    `id`             int(11) NOT NULL AUTO_INCREMENT,
    `name`           varchar(20)  NOT NULL COMMENT '名称',
    `describe`       varchar(200) NOT NULL COMMENT '描述',
    `state`          tinyint(1) NOT NULL COMMENT '状态',
    `create_at`      datetime DEFAULT NULL COMMENT '创建时间',
    `create_user`    varchar(30)  NOT NULL COMMENT '创建人',
    `update_at`      datetime DEFAULT NULL COMMENT '更新时间',
    `update_user`    varchar(30)  NOT NULL COMMENT '更新人',
    `database_id`    int(11) NOT NULL COMMENT '数据库信息',
    `environment_id` int(11) NOT NULL COMMENT '环境信息',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 数据库配置表
CREATE TABLE `database`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT,
    `name`        varchar(20)  NOT NULL COMMENT '名称',
    `describe`    varchar(200) NOT NULL COMMENT '描述',
    `host`        varchar(20)  NOT NULL COMMENT '地址',
    `user`        varchar(20)  NOT NULL COMMENT '用户',
    `password`    varchar(20)  NOT NULL COMMENT '密码',
    `database`    varchar(20)  NOT NULL COMMENT '数据库名',
    `state`       int(11) NOT NULL COMMENT '状态',
    `remarks`     varchar(100) DEFAULT NULL COMMENT '备注',
    `create_at`   datetime     DEFAULT NULL COMMENT '创建时间',
    `create_user` varchar(30)  NOT NULL COMMENT '创建人',
    `update_at`   datetime     DEFAULT NULL COMMENT '更新时间',
    `update_user` varchar(30)  NOT NULL COMMENT '更新人',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 环境信息表
CREATE TABLE `environment`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT,
    `name`        varchar(20)  NOT NULL COMMENT '名称',
    `describe`    varchar(200) NOT NULL COMMENT '描述',
    `uri`         varchar(30)  NOT NULL COMMENT '环境地址',
    `state`       int(11) NOT NULL COMMENT '状态',
    `remarks`     varchar(100) NOT NULL COMMENT '备注',
    `create_at`   datetime DEFAULT NULL COMMENT '创建时间',
    `create_user` varchar(30)  NOT NULL COMMENT '创建人',
    `update_at`   datetime DEFAULT NULL COMMENT '更新时间',
    `update_user` varchar(30)  NOT NULL COMMENT '更新人',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;