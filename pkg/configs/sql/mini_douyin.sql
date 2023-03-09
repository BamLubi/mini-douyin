/*
 Navicat Premium Data Transfer

 Source Server         : 本地虚拟机
 Source Server Type    : MySQL
 Source Server Version : 50740
 Source Host           : 192.168.45.129:3306
 Source Schema         : mini_douyin

 Target Server Type    : MySQL
 Target Server Version : 50740
 File Encoding         : 65001

 Date: 09/03/2023 23:15:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  `video_id` bigint(20) NULL DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_date` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `active` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for favorites_0
-- ----------------------------
DROP TABLE IF EXISTS `favorites_0`;
CREATE TABLE `favorites_0`  (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  `video_id` bigint(20) NULL DEFAULT NULL,
  `action_type` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for favorites_1
-- ----------------------------
DROP TABLE IF EXISTS `favorites_1`;
CREATE TABLE `favorites_1`  (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  `video_id` bigint(20) NULL DEFAULT NULL,
  `action_type` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for favorites_2
-- ----------------------------
DROP TABLE IF EXISTS `favorites_2`;
CREATE TABLE `favorites_2`  (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  `video_id` bigint(20) NULL DEFAULT NULL,
  `action_type` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(32) NOT NULL COMMENT '用户id',
  `username` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL COMMENT '用户名',
  `salt` varchar(10) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL COMMENT '盐值',
  `hash` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL COMMENT '密码加盐后hash',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for userinfo
-- ----------------------------
DROP TABLE IF EXISTS `userinfo`;
CREATE TABLE `userinfo`  (
  `id` bigint(32) NOT NULL COMMENT '用户id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名称',
  `follow_count` int(32) NULL DEFAULT 0 COMMENT '关注总数',
  `follower_count` int(32) NULL DEFAULT 0 COMMENT '粉丝总数',
  `is_follow` tinyint(1) NULL DEFAULT NULL COMMENT '已关注、未关注',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户头像',
  `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户个人页顶部大图',
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '个人简介',
  `total_favorited` int(32) NULL DEFAULT 0 COMMENT '获赞数量',
  `work_count` int(32) NULL DEFAULT 0 COMMENT '作品数量',
  `favorite_count` int(32) NULL DEFAULT 0 COMMENT '点赞数量',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for videoinfo
-- ----------------------------
DROP TABLE IF EXISTS `videoinfo`;
CREATE TABLE `videoinfo`  (
  `id` bigint(32) NOT NULL COMMENT 'id',
  `user_id` bigint(32) NULL DEFAULT NULL COMMENT '用户id',
  `play_url` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL COMMENT '播放地址',
  `cover_url` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL COMMENT '封面地址',
  `favorite_count` int(32) NULL DEFAULT 0 COMMENT '点赞总数',
  `comment_count` int(32) NULL DEFAULT 0 COMMENT '评论总数',
  `is_favorite` tinyint(1) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `create_time` bigint(32) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
