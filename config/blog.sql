/*
Navicat MySQL Data Transfer

Source Server         : mysql
Source Server Version : 50018
Source Host           : 127.0.0.1:3306
Source Database       : blog

Target Server Type    : MYSQL
Target Server Version : 50018
File Encoding         : 65001

Date: 2017-12-28 20:59:28
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for system_menu
-- ----------------------------
DROP TABLE IF EXISTS `system_menu`;
CREATE TABLE `system_menu` (
  `ID` int(8) NOT NULL auto_increment COMMENT '主键标识',
  `PID` int(8) NOT NULL COMMENT '父级',
  `NAME` varchar(64) default NULL COMMENT '名字',
  `PATH` varchar(64) default '' COMMENT '访问路径',
  `ICON` varchar(64) default '' COMMENT '图标',
  `STATUS` char(1) default '0' COMMENT '状态 0:正常 1:删除',
  `CREATOR` varchar(64) default NULL COMMENT '创建人',
  `UPDATOR` varchar(64) default NULL COMMENT '修改人',
  `CREATETIME` datetime default NULL COMMENT '创建时间',
  `UPDATETIME` datetime default NULL COMMENT '更新时间',
  `DESCRIPTION` varchar(512) default NULL COMMENT '描述',
  PRIMARY KEY  (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='菜单表';

-- ----------------------------
-- Records of system_menu
-- ----------------------------
INSERT INTO `system_menu` VALUES ('1', '0', '博客管理', '', 'el-icon-menu', '0', 'admin1', 'admin1', '2017-12-27 21:12:38', '2017-12-27 21:12:42', '博客管理');
INSERT INTO `system_menu` VALUES ('2', '0', '系统管理', '', 'el-icon-menu', '0', 'admin1', 'admin1', '2017-12-27 21:12:38', '2017-12-27 21:12:38', '系统管理');
INSERT INTO `system_menu` VALUES ('3', '0', '数据统计', '', 'el-icon-menu', '0', 'admin1', 'admin1', '2017-12-27 21:12:38', '2017-12-27 21:12:38', '数据统计');
INSERT INTO `system_menu` VALUES ('4', '0', '日志查询', '', 'el-icon-menu', '0', 'admin1', 'admin1', '2017-12-27 21:12:38', '2017-12-27 21:12:38', '日志查询');
INSERT INTO `system_menu` VALUES ('5', '1', '文章管理', '/article', 'el-icon-menu', '0', 'admin1', 'admin1', '2017-12-27 21:12:38', '2017-12-27 21:12:38', '文章管理');
INSERT INTO `system_menu` VALUES ('6', '2', '用户管理', '/user', 'el-icon-menu', '0', 'admin1', 'admin1', '2017-12-27 21:12:38', '2017-12-27 21:12:38', '用户管理');
INSERT INTO `system_menu` VALUES ('7', '3', '访问统计', '/access', 'el-icon-menu', '0', 'admin1', 'admin1', '2017-12-27 21:12:38', '2017-12-27 21:12:38', '访问统计');
INSERT INTO `system_menu` VALUES ('8', '3', '点击统计', '/click', 'el-icon-menu', '0', 'admin1', 'admin1', '2017-12-27 21:12:38', '2017-12-27 21:12:38', '点击统计');
INSERT INTO `system_menu` VALUES ('9', '4', '博客日志', '/blogLog', 'el-icon-menu', '0', 'admin1', 'admin1', '2017-12-27 21:12:38', '2017-12-27 21:12:38', '博客日志');
INSERT INTO `system_menu` VALUES ('10', '4', '后台日志', '/managerLog', 'el-icon-menu', '0', 'admin1', 'admin1', '2017-12-27 21:12:38', '2017-12-27 21:12:38', '后台日志');

-- ----------------------------
-- Table structure for system_privilege
-- ----------------------------
DROP TABLE IF EXISTS `system_privilege`;
CREATE TABLE `system_privilege` (
  `ID` int(8) NOT NULL auto_increment COMMENT '主键标识',
  `CODE` varchar(32) NOT NULL COMMENT '权限标识',
  `PCODE` varchar(32) NOT NULL COMMENT '父权限标识',
  `NAME` varchar(64) NOT NULL COMMENT '权限名称',
  `CREATETIME` datetime default NULL COMMENT '创建时间',
  `UPDATETIME` datetime default NULL COMMENT '更新时间',
  PRIMARY KEY  (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of system_privilege
-- ----------------------------

-- ----------------------------
-- Table structure for system_role
-- ----------------------------
DROP TABLE IF EXISTS `system_role`;
CREATE TABLE `system_role` (
  `ID` int(8) NOT NULL auto_increment COMMENT '主键标识',
  `NAME` varchar(64) NOT NULL COMMENT '角色标识',
  `STATUS` varchar(64) NOT NULL COMMENT '状态',
  `CREATOR` varchar(64) default NULL COMMENT '创建人',
  `UPDATOR` varchar(64) default NULL COMMENT '修改人',
  `CREATETIME` datetime default NULL COMMENT '创建时间',
  `UPDATETIME` datetime default NULL COMMENT '更新时间',
  `DESCRIPTION` varchar(512) default NULL COMMENT '角色描述',
  `ISSYSTEM` tinyint(1) NOT NULL COMMENT '1,系统管理员 2,不是系统管理员',
  PRIMARY KEY  (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of system_role
-- ----------------------------

-- ----------------------------
-- Table structure for system_role_privilege
-- ----------------------------
DROP TABLE IF EXISTS `system_role_privilege`;
CREATE TABLE `system_role_privilege` (
  `ID` int(11) NOT NULL auto_increment COMMENT '主键标识',
  `ROLEID` int(8) default NULL COMMENT '角色编号',
  `CREATOR` varchar(64) NOT NULL COMMENT '创建人',
  `CREATETIME` datetime NOT NULL COMMENT '创建时间',
  `PRIVILEGECODE` varchar(32) default NULL COMMENT '权限编号',
  PRIMARY KEY  (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of system_role_privilege
-- ----------------------------

-- ----------------------------
-- Table structure for system_user
-- ----------------------------
DROP TABLE IF EXISTS `system_user`;
CREATE TABLE `system_user` (
  `ID` int(8) NOT NULL auto_increment COMMENT '主键标识',
  `LOGINNAME` varchar(64) default NULL COMMENT '登录名称',
  `NAME` varchar(64) default NULL COMMENT '用户名称',
  `ISADMIN` int(2) default NULL COMMENT '是否管理员 1.是 2.否',
  `PASSWORD` varchar(128) default NULL COMMENT '密码',
  `STATUS` int(2) default NULL COMMENT '1.正常 2.锁定',
  `LASTLOGINTIME` datetime default NULL COMMENT '最后登录时间',
  `CREATOR` varchar(64) default NULL COMMENT '创建人',
  `UPDATOR` varchar(64) default NULL COMMENT '修改人',
  `CREATETIME` datetime default NULL COMMENT '创建时间',
  `UPDATETIME` datetime default NULL COMMENT '更新时间',
  `PHONE` varchar(12) default NULL COMMENT '手机号',
  `SEX` int(1) default NULL COMMENT '性别 1.男 2.女',
  `EMAIL` varchar(32) default NULL COMMENT '邮箱',
  PRIMARY KEY  (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of system_user
-- ----------------------------
INSERT INTO `system_user` VALUES ('1', 'admin1', 'admin', '1', '123456', '1', '2017-12-22 20:43:59', 'admin', 'admin', '2017-12-22 20:43:31', '2017-12-22 20:43:34', '12345678911', '1', '11@qq.com');

-- ----------------------------
-- Table structure for system_user_role
-- ----------------------------
DROP TABLE IF EXISTS `system_user_role`;
CREATE TABLE `system_user_role` (
  `ID` int(11) NOT NULL auto_increment COMMENT '主键标识',
  `USERID` int(8) default NULL COMMENT '用户编号',
  `ROLEID` int(8) default NULL COMMENT '角色编号',
  `CREATOR` varchar(64) default NULL COMMENT '创建人',
  `CREATETIME` datetime default NULL COMMENT '创建时间',
  PRIMARY KEY  (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of system_user_role
-- ----------------------------
