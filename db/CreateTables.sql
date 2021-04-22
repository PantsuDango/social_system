CREATE DATABASE `social_db`;

USE `social_db`;

CREATE TABLE `user` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `nick` varchar(32) NOT NULL DEFAULT '初始昵称' COMMENT '用户昵称',
   `username` varchar(32) NOT NULL COMMENT '登录账号',
   `password` varchar(32) NOT NULL COMMENT '登录密码, 存md5',
   `salt` varchar(32) NOT NULL COMMENT '盐',
   `sex` tinyint(4) NOT NULL DEFAULT 0 COMMENT '性别: 0-男, 1-女',
   `head_image` varchar(4096) DEFAULT NULL COMMENT '头像图片url',
   `email` varchar(32) DEFAULT NULL COMMENT '邮箱',
   `phone` varchar(32) DEFAULT NULL COMMENT '手机号',
   `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT  '账号状态: 0-存续, 1-废除',
   `createtime` datetime NOT NULL COMMENT '创建时间',
   `lastupdate` datetime NOT NULL COMMENT '更新时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

INSERT INTO 
  `user` (`id`, `nick`, `username`, `password`, `salt`, `sex`, `head_image`, `email`, `phone`, `status`, `createtime`, `lastupdate`)
VALUES
    (1, '珍珠哥', 'admin', 'e160fa11757ed2883e890e483f1c6208', 'h6du1cxo', 0, '', '394883561@qq.com', '13266871263', 0, NOW(), NOW());


CREATE TABLE `user_attention_map` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `user_id` int(11) NOT NULL COMMENT '被关注用户的id',
   `follower_id` int(11) NOT NULL COMMENT '点关注的用户id',
   `createtime` datetime NOT NULL COMMENT '创建时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY (`user_id`, `follower_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户关注映射表';


CREATE TABLE `post` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `title` varchar(32) NOT NULL COMMENT '标题',
   `content` text DEFAULT NULL COMMENT '内容',
   `user_id` int(11) NOT NULL COMMENT '用户id',
   `type` tinyint(4) NOT NULL DEFAULT 0 COMMENT '帖子类型: 0-原帖, 1-转发贴',
   `from_id` int(11) NOT NULL DEFAULT 0 COMMENT '转发的作者id',
   `createtime` datetime NOT NULL COMMENT '创建时间',
   `lastupdate` datetime NOT NULL COMMENT '更新时间',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='帖子信息表';


CREATE TABLE `post_picture_map` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `post_id` int(11) NOT NULL COMMENT '帖子id',
   `picture_url` varchar(4096) NOT NULL COMMENT '图片地址',
   `createtime` datetime NOT NULL COMMENT '创建时间',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='帖子图片映射表';


CREATE TABLE `post_star_map` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `post_id` int(11) NOT NULL COMMENT '帖子id',
   `user_id` int(11) NOT NULL COMMENT '点赞人',
   `createtime` datetime NOT NULL COMMENT '创建时间',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='帖子点赞映射表';


CREATE TABLE `post_comment_map` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `post_id` int(11) NOT NULL COMMENT '帖子id',
   `content` text DEFAULT NULL COMMENT '评论内容',
   `user_id` int(11) NOT NULL COMMENT '评论人',
   `createtime` datetime NOT NULL COMMENT '创建时间',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='帖子评论映射表';


CREATE TABLE `post_quoted_map` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `post_id` int(11) NOT NULL COMMENT '帖子id',
   `user_id` int(11) NOT NULL COMMENT '转发人',
   `createtime` datetime NOT NULL COMMENT '创建时间',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='帖子转发映射表';