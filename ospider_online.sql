-- MySQL dump 10.13  Distrib 5.5.46, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: ospider
-- ------------------------------------------------------
-- Server version	5.5.46-0ubuntu0.14.04.2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `sp_access_privilege`
--

DROP TABLE IF EXISTS `sp_access_privilege`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_access_privilege` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pri_group` varchar(500) NOT NULL DEFAULT '' COMMENT '1;2;3;4;5',
  `pri_rule` int(11) NOT NULL DEFAULT '0' COMMENT '1:all, 2:allow, 4:ban',
  `logtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_access_privilege`
--

LOCK TABLES `sp_access_privilege` WRITE;
/*!40000 ALTER TABLE `sp_access_privilege` DISABLE KEYS */;
INSERT INTO `sp_access_privilege` VALUES (1,'',1,'2015-06-03 08:35:45'),(8,'6',2,'2016-03-09 15:15:33'),(9,'7',2,'2016-03-09 15:15:33'),(10,'8',2,'2016-03-09 15:26:53');
/*!40000 ALTER TABLE `sp_access_privilege` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sp_menu_template`
--

DROP TABLE IF EXISTS `sp_menu_template`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_menu_template` (
  `id` int(11) NOT NULL DEFAULT '0' COMMENT '1 2 4 8',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '关键词管理',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT 'show',
  `logtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_menu_template`
--

LOCK TABLES `sp_menu_template` WRITE;
/*!40000 ALTER TABLE `sp_menu_template` DISABLE KEYS */;
INSERT INTO `sp_menu_template` VALUES (2,'定时器管理','crontabview','2016-02-29 04:19:04'),(4,'爬取规则','ruleview','2016-03-01 05:56:56'),(4,'网站管理','siteview','2016-03-03 04:37:12'),(2,'爬虫管理','spiderview','2016-02-29 04:18:43'),(1,'用户管理','usersview','2016-02-29 04:18:43');
/*!40000 ALTER TABLE `sp_menu_template` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sp_node_privilege`
--

DROP TABLE IF EXISTS `sp_node_privilege`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_node_privilege` (
  `id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(100) NOT NULL DEFAULT '',
  `node` varchar(200) NOT NULL DEFAULT '' COMMENT '1:/login, 2:/login/check, 4:/logout, 8:/key/add, 16:/key/update, 32:/key/show, 64:/key/one',
  `logtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`node`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_node_privilege`
--

LOCK TABLES `sp_node_privilege` WRITE;
/*!40000 ALTER TABLE `sp_node_privilege` DISABLE KEYS */;
INSERT INTO `sp_node_privilege` VALUES (2,'首页','/','2016-02-29 04:07:22'),(2,'获取全部可用的定时器列表Api','/api/crontabs','2016-02-29 04:07:22'),(2,'获取当前用户可用的爬虫规则列表','/api/rules','2016-02-29 16:48:56'),(2,'获取全部可用的爬虫列表Api','/api/spiders','2016-02-29 04:07:22'),(2,'消费记录','/consumelog','2016-02-29 04:07:22'),(4,'定时器添加','/crontab/add','2016-02-29 04:07:22'),(4,'定时器编辑','/crontab/edit','2016-02-29 04:07:22'),(4,'获取单个定时器','/crontab/one','2016-02-29 04:07:22'),(4,'定时器列表','/crontabview','2016-02-29 04:07:22'),(1,'登录页','/login','2015-06-03 08:35:33'),(1,'登录验证请求','/login/check','2016-02-29 04:07:22'),(1,'退出登录','/logout','2016-02-29 04:07:22'),(4,'充值','/pay','2016-02-29 04:07:22'),(2,'充值记录','/paylog','2016-02-29 04:07:22'),(2,'爬取规则添加','/rule/add','2016-02-29 16:48:56'),(2,'爬取规则编辑','/rule/edit','2016-02-29 16:48:56'),(2,'获取单个爬取规则','/rule/one','2016-02-29 16:48:56'),(2,'爬取规则列表','/ruleview','2016-02-29 16:48:56'),(2,'网站添加','/site/add','2016-03-09 14:49:00'),(2,'网站编辑','/site/edit','2016-03-09 14:49:00'),(2,'获取单个网站设置','/site/one','2016-03-09 14:49:00'),(2,'网站列表','/siteview','2016-03-09 14:49:00'),(4,'爬虫添加','/spider/add','2016-02-29 04:07:22'),(4,'爬虫编辑','/spider/edit','2016-02-29 04:07:22'),(4,'获取单个爬虫记录','/spider/one','2016-02-29 04:07:22'),(4,'爬虫列表','/spiderview','2016-02-29 04:07:22'),(4,'添加用户','/user/add','2016-02-29 04:07:22'),(2,'更新用户资料','/user/edit','2016-02-29 04:07:22'),(2,'查看用户','/user/view','2016-02-29 04:07:22'),(2,'用户管理','/usersview','2016-02-29 04:07:22');
/*!40000 ALTER TABLE `sp_node_privilege` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sp_role`
--

DROP TABLE IF EXISTS `sp_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_role` (
  `id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT 'user, services, admin, guess',
  `privilege` int(11) NOT NULL DEFAULT '0',
  `menu` int(11) NOT NULL DEFAULT '0',
  `logtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_role`
--

LOCK TABLES `sp_role` WRITE;
/*!40000 ALTER TABLE `sp_role` DISABLE KEYS */;
INSERT INTO `sp_role` VALUES (1,'匿名用户',1,0,'2016-02-29 04:09:43'),(2,'管理员',7,7,'2016-03-01 05:58:16'),(3,'普通用户',3,5,'2016-03-03 04:58:32');
/*!40000 ALTER TABLE `sp_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sp_user`
--

DROP TABLE IF EXISTS `sp_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL DEFAULT '',
  `password` varchar(100) NOT NULL DEFAULT '',
  `roleid` int(11) NOT NULL DEFAULT '3',
  `accessid` int(11) NOT NULL DEFAULT '0',
  `logtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_user`
--

LOCK TABLES `sp_user` WRITE;
/*!40000 ALTER TABLE `sp_user` DISABLE KEYS */;
INSERT INTO `sp_user` VALUES (1,'root','admin',2,1,'2015-06-03 08:35:04'),(6,'test','test',3,8,'2016-03-09 14:29:38'),(7,'test1','test1',3,9,'2016-03-09 15:05:40'),(8,'test2','test2',3,10,'2016-03-09 15:26:53');
/*!40000 ALTER TABLE `sp_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `spider_crontab`
--

DROP TABLE IF EXISTS `spider_crontab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `spider_crontab` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'crontab id',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'user id',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT 'crontab name',
  `cron` varchar(30) NOT NULL DEFAULT '' COMMENT 'crontab ex',
  `status` int(10) unsigned NOT NULL DEFAULT '1' COMMENT 'status 0:inactive 1:active',
  `logtime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `uptime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spider_crontab`
--

LOCK TABLES `spider_crontab` WRITE;
/*!40000 ALTER TABLE `spider_crontab` DISABLE KEYS */;
INSERT INTO `spider_crontab` VALUES (1,1,'每天8、18点执行','* 8,18 * * *',1,'2016-02-29 04:28:34','2016-02-29 05:30:48'),(2,1,'每隔4小时执行一次','* 4/* * * * *',1,'2016-02-29 05:37:38','2016-03-09 14:16:16'),(3,1,'每天8,12,18点执行','* 8,12,18 * * *',1,'2016-03-01 05:44:05','2016-03-09 14:17:37'),(4,1,'每10秒执行一次','10 * * * * *',1,'2016-03-04 21:09:08','2016-03-05 03:44:01');
/*!40000 ALTER TABLE `spider_crontab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `spider_history`
--

DROP TABLE IF EXISTS `spider_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `spider_history` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'history id',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'user id',
  `siteid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'site id',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT 'site url',
  `result_json` text COMMENT '{title:, content:}',
  `logtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spider_history`
--

LOCK TABLES `spider_history` WRITE;
/*!40000 ALTER TABLE `spider_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `spider_history` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `spider_manager`
--

DROP TABLE IF EXISTS `spider_manager`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `spider_manager` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'spider id',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'user id',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT 'spider name',
  `queue_name_json` varchar(100) NOT NULL COMMENT '{spider:spider task name, result:spider result queue name}',
  `status` int(10) unsigned NOT NULL DEFAULT '1' COMMENT 'status 0:inactive 1:active',
  `logtime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `uptime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `queue_name` (`queue_name_json`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spider_manager`
--

LOCK TABLES `spider_manager` WRITE;
/*!40000 ALTER TABLE `spider_manager` DISABLE KEYS */;
INSERT INTO `spider_manager` VALUES (1,1,'网站更新爬虫','{\"task\":\"spiderTaskQueueRoot1\",\"result\":\"spiderResultQueueRoot1\"}',1,'2016-02-25 04:55:46','2016-03-09 14:22:03');
/*!40000 ALTER TABLE `spider_manager` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `spider_rule`
--

DROP TABLE IF EXISTS `spider_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `spider_rule` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'rule id',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'user_id',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT 'rule name',
  `spiderid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'spider id',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT 'p/?.html, c=mode&p=?',
  `rule_json` varchar(1000) NOT NULL DEFAULT '' COMMENT 'title, content, page',
  `status` int(10) unsigned NOT NULL DEFAULT '1',
  `logtime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `uptime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spider_rule`
--

LOCK TABLES `spider_rule` WRITE;
/*!40000 ALTER TABLE `spider_rule` DISABLE KEYS */;
INSERT INTO `spider_rule` VALUES (20,1,'夜店信息',1,'http://www.nc-yechang.com/','{\"title\":\"#body \\u003e div.inner \\u003e div.right \\u003e div.main \\u003e div \\u003e div.title \\u003e h3\",\"content\":\"#body \\u003e div.inner \\u003e div.right \\u003e div.main \\u003e div \\u003e div.maincontent.clearfix \\u003e p:nth-child(1)\",\"section\":\"#body \\u003e div.MainBlock \\u003e div.left \\u003e div \\u003e ul.tab-bd \\u003e ul\",\"href\":\"\",\"filter\":\"\"}',1,'2016-03-02 01:30:10','2016-03-05 08:08:01'),(22,6,'开源中国软件更新资讯',1,'http://www.oschina.net/','{\"title\":\"#NewsChannel \\u003e div.NewsBody \\u003e div \\u003e div.NewsEntity \\u003e h1\",\"content\":\"#NewsChannel \\u003e div.NewsBody \\u003e div \\u003e div.NewsEntity \\u003e div.Body.NewsContent.TextContent\\u003ep\",\"section\":\"#ProjectNews \\u003e ul.p1\",\"href\":\"\",\"filter\":\"p[style]\"}',1,'2016-03-10 00:52:04','2016-03-10 03:11:55');
/*!40000 ALTER TABLE `spider_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `spider_site`
--

DROP TABLE IF EXISTS `spider_site`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `spider_site` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'site id',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'user id',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT 'site name',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT 'site url',
  `document_set` varchar(500) NOT NULL DEFAULT '' COMMENT '{document_position:,document_display:}',
  `ruleid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'rule id',
  `cronid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'crontab id',
  `status` int(10) unsigned NOT NULL DEFAULT '1',
  `logtime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `uptime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spider_site`
--

LOCK TABLES `spider_site` WRITE;
/*!40000 ALTER TABLE `spider_site` DISABLE KEYS */;
INSERT INTO `spider_site` VALUES (12,1,'测试','http://localhost','{\"position\":7,\"display\":1,\"uid\":1,\"nickname\":\"Administrator\",\"check\":1,\"category_id\":41,\"title\":\"新闻公告\",\"group_id\":0,\"model_id\":2}',20,4,1,'2016-03-03 17:36:31','2016-03-10 03:10:14'),(15,6,'测试','http://localhost','{\"position\":7,\"display\":1,\"uid\":1,\"nickname\":\"Administrator\",\"check\":1,\"category_id\":42,\"title\":\"搜索优化\",\"group_id\":0,\"model_id\":2}',22,4,1,'2016-03-10 00:52:50','2016-03-09 16:52:50');
/*!40000 ALTER TABLE `spider_site` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-03-10  8:05:29
