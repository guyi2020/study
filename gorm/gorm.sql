/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50529
Source Host           : localhost:3306
Source Database       : gorm

Target Server Type    : MYSQL
Target Server Version : 50529
File Encoding         : 65001

Date: 2020-11-25 18:54:54
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `age` int(11) unsigned NOT NULL,
  `email` varchar(255) DEFAULT '',
  `birthday` datetime DEFAULT NULL,
  `member_number` varchar(255) DEFAULT '',
  `actived_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', '张三', '18', '', '2020-11-25 18:51:36', '', '0000-00-00 00:00:00', '2020-11-25 18:51:36', '2020-11-25 18:51:36');
INSERT INTO `users` VALUES ('2', '张三', '18', '', '2020-11-25 18:52:10', '', '0000-00-00 00:00:00', '2020-11-25 18:52:10', '2020-11-25 18:52:10');
INSERT INTO `users` VALUES ('3', '张三', '18', '', '2020-11-25 18:52:26', '', '0000-00-00 00:00:00', '2020-11-25 18:52:26', '2020-11-25 18:52:26');
INSERT INTO `users` VALUES ('4', '张三', '18', '', '2020-11-25 18:54:25', null, null, '2020-11-25 18:54:25', '2020-11-25 18:54:25');
