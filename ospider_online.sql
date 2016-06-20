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
INSERT INTO `sp_menu_template` VALUES (2,'定时器管理','crontabview','2016-02-29 04:19:04'),(2,'API管理','floodapiview','2016-06-19 21:13:19'),(2,'目标管理','floodtargetview','2016-06-19 21:13:19'),(4,'爬取规则','ruleview','2016-03-01 05:56:56'),(4,'网站管理','siteview','2016-03-03 04:37:12'),(2,'爬虫管理','spiderview','2016-02-29 04:18:43'),(1,'用户管理','usersview','2016-02-29 04:18:43');
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
INSERT INTO `sp_node_privilege` VALUES (2,'首页','/','2016-02-29 04:07:22'),(2,'获取全部可用的定时器列表Api','/api/crontabs','2016-02-29 04:07:22'),(2,'获取当前用户可用的爬虫规则列表','/api/rules','2016-02-29 16:48:56'),(2,'获取全部可用的爬虫列表Api','/api/spiders','2016-02-29 04:07:22'),(2,'消费记录','/consumelog','2016-02-29 04:07:22'),(4,'定时器添加','/crontab/add','2016-02-29 04:07:22'),(4,'定时器编辑','/crontab/edit','2016-02-29 04:07:22'),(4,'获取单个定时器','/crontab/one','2016-02-29 04:07:22'),(4,'定时器列表','/crontabview','2016-02-29 04:07:22'),(4,'api添加','/floodapi/add','2016-06-19 16:55:03'),(4,'api编辑','/floodapi/edit','2016-06-19 16:55:03'),(4,'获取单个api','/floodapi/one','2016-06-19 16:55:03'),(4,'api列表','/floodapiview','2016-06-19 17:07:48'),(4,'目标添加','/floodtarget/add','2016-06-19 21:18:14'),(4,'目标编辑','/floodtarget/edit','2016-06-19 21:18:14'),(4,'获取单个目标','/floodtarget/one','2016-06-19 21:18:14'),(4,'目标列表','/floodtargetview','2016-06-19 21:18:14'),(1,'登录页','/login','2015-06-03 08:35:33'),(1,'登录验证请求','/login/check','2016-02-29 04:07:22'),(1,'退出登录','/logout','2016-02-29 04:07:22'),(4,'充值','/pay','2016-02-29 04:07:22'),(2,'充值记录','/paylog','2016-02-29 04:07:22'),(2,'爬取规则添加','/rule/add','2016-02-29 16:48:56'),(2,'爬取规则编辑','/rule/edit','2016-02-29 16:48:56'),(2,'获取单个爬取规则','/rule/one','2016-02-29 16:48:56'),(2,'爬取规则列表','/ruleview','2016-02-29 16:48:56'),(2,'网站添加','/site/add','2016-03-09 14:49:00'),(2,'网站编辑','/site/edit','2016-03-09 14:49:00'),(2,'获取单个网站设置','/site/one','2016-03-09 14:49:00'),(2,'网站列表','/siteview','2016-03-09 14:49:00'),(4,'爬虫添加','/spider/add','2016-02-29 04:07:22'),(4,'爬虫编辑','/spider/edit','2016-02-29 04:07:22'),(4,'获取单个爬虫记录','/spider/one','2016-02-29 04:07:22'),(4,'爬虫列表','/spiderview','2016-02-29 04:07:22'),(4,'添加用户','/user/add','2016-02-29 04:07:22'),(2,'更新用户资料','/user/edit','2016-02-29 04:07:22'),(2,'查看用户','/user/view','2016-02-29 04:07:22'),(2,'用户管理','/usersview','2016-02-29 04:07:22');
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
-- Table structure for table `spider_flood_api`
--

DROP TABLE IF EXISTS `spider_flood_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `spider_flood_api` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'site id',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'user id',
  `name` varchar(20) NOT NULL DEFAULT '',
  `api` varchar(200) NOT NULL DEFAULT '',
  `powerlevel` int(11) NOT NULL DEFAULT '4500',
  `time` int(11) NOT NULL DEFAULT '9000',
  `status` int(10) unsigned NOT NULL DEFAULT '1',
  `uptime` int(11) NOT NULL DEFAULT '0',
  `logtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spider_flood_api`
--

LOCK TABLES `spider_flood_api` WRITE;
/*!40000 ALTER TABLE `spider_flood_api` DISABLE KEYS */;
INSERT INTO `spider_flood_api` VALUES (1,1,'alphastress','https://alphastress.com/api.php?user=huzorro&key=zs2xmyF89B6n4rC',4200,7200,1,1466385912,'2016-06-20 01:15:12'),(2,1,'s','s1',4500,9000,0,0,'2016-06-19 18:39:05');
/*!40000 ALTER TABLE `spider_flood_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `spider_flood_target`
--

DROP TABLE IF EXISTS `spider_flood_target`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `spider_flood_target` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'site id',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'user id',
  `url` varchar(255) NOT NULL DEFAULT '',
  `host` varchar(20) NOT NULL DEFAULT '',
  `port` varchar(10) NOT NULL DEFAULT '80',
  `method` varchar(10) NOT NULL DEFAULT 'DNS',
  `time` int(11) NOT NULL DEFAULT '600',
  `powerlevel` int(11) NOT NULL DEFAULT '100',
  `cronid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'crontab id',
  `status` int(10) unsigned NOT NULL DEFAULT '1',
  `logtime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `uptime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spider_flood_target`
--

LOCK TABLES `spider_flood_target` WRITE;
/*!40000 ALTER TABLE `spider_flood_target` DISABLE KEYS */;
INSERT INTO `spider_flood_target` VALUES (1,1,'sss','222','80','DNS',600,100,3,0,'0000-00-00 00:00:00','2016-06-20 01:12:19'),(2,1,'http://www.bjbhyl.com','115.47.100.197','80','DNS',600,100,4,1,'2016-06-19 21:27:54','2016-06-20 01:12:46');
/*!40000 ALTER TABLE `spider_flood_target` ENABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `spider_history`
--

LOCK TABLES `spider_history` WRITE;
/*!40000 ALTER TABLE `spider_history` DISABLE KEYS */;
INSERT INTO `spider_history` VALUES (1,6,15,'http://www.oschina.net/news/71458/apache-tomcat-native-1-2-5','{\"id\":15,\"uid\":6,\"name\":\"测试\",\"url\":\"http://localhost\",\"document_set\":{\"position\":7,\"display\":1,\"uid\":1,\"nickname\":\"Administrator\",\"check\":1,\"category_id\":42,\"title\":\"搜索优化\",\"group_id\":0,\"model_id\":2},\"documentSetStr\":\"{\\\"position\\\":7,\\\"display\\\":1,\\\"uid\\\":1,\\\"nickname\\\":\\\"Administrator\\\",\\\"check\\\":1,\\\"category_id\\\":42,\\\"title\\\":\\\"搜索优化\\\",\\\"group_id\\\":0,\\\"model_id\\\":2}\",\"Rule\":{\"id\":22,\"uid\":0,\"name\":\"开源中国软件更新资讯\",\"spiderid\":1,\"manager\":{\"id\":1,\"uid\":0,\"name\":\"网站更新爬虫\",\"task\":\"spiderTaskQueueRoot1\",\"result\":\"spiderResultQueueRoot1\",\"queueName\":\"{\\\"task\\\":\\\"spiderTaskQueueRoot1\\\",\\\"result\\\":\\\"spiderResultQueueRoot1\\\"}\",\"status\":0,\"logtime\":\"\"},\"url\":\"http://www.oschina.net/news/71458/apache-tomcat-native-1-2-5\",\"selector\":{\"title\":\"Apache Tomcat Native 1.2.5 发布\",\"content\":\"Apache Tomcat Native 1.2.5 发布了，Tomcat Native 这个项目可以让 Tomcat 使用 Apache 的 apr 包来处理包括文件和网络IO操作，以提升性能。\\u003cbr/\\u003e该版本主要特点如下：\\u003cbr/\\u003e- Report OpenSSL runtime version in use rather than compile\\u003cbr/\\u003e  time version used.\\u003cbr/\\u003e- Windows binaries built with APR 1.5.1 and OpenSSL 1.0.2g.\\u003cbr/\\u003e改进记录：\\u003cbr/\\u003e下载地址：\\u003ca href=\\\"http://tomcat.apache.org/download-native.cgi\\\" target=\\\"_blank\\\"\\u003ehttp://tomcat.apache.org/download-native.cgi\\u003c/a\\u003e\",\"section\":\"#ProjectNews \\u003e ul.p1\",\"href\":\"\",\"filter\":\"p[style]\"},\"selectorStr\":\"{\\\"title\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e h1\\\",\\\"content\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e div.Body.NewsContent.TextContent\\\\u003ep\\\",\\\"section\\\":\\\"#ProjectNews \\\\u003e ul.p1\\\",\\\"href\\\":\\\"\\\",\\\"filter\\\":\\\"p[style]\\\"}\",\"status\":0,\"logtime\":\"\"},\"Cron\":{\"id\":4,\"uid\":0,\"name\":\"每10秒执行一次\",\"cron\":\"10 * * * * *\",\"status\":0,\"logtime\":\"\"},\"status\":0,\"logtime\":\"\"}','2016-03-12 04:52:34'),(2,6,15,'http://www.oschina.net/news/71456/y-ppa-manager','{\"id\":15,\"uid\":6,\"name\":\"测试\",\"url\":\"http://localhost\",\"document_set\":{\"position\":7,\"display\":1,\"uid\":1,\"nickname\":\"Administrator\",\"check\":1,\"category_id\":42,\"title\":\"搜索优化\",\"group_id\":0,\"model_id\":2},\"documentSetStr\":\"{\\\"position\\\":7,\\\"display\\\":1,\\\"uid\\\":1,\\\"nickname\\\":\\\"Administrator\\\",\\\"check\\\":1,\\\"category_id\\\":42,\\\"title\\\":\\\"搜索优化\\\",\\\"group_id\\\":0,\\\"model_id\\\":2}\",\"Rule\":{\"id\":22,\"uid\":0,\"name\":\"开源中国软件更新资讯\",\"spiderid\":1,\"manager\":{\"id\":1,\"uid\":0,\"name\":\"网站更新爬虫\",\"task\":\"spiderTaskQueueRoot1\",\"result\":\"spiderResultQueueRoot1\",\"queueName\":\"{\\\"task\\\":\\\"spiderTaskQueueRoot1\\\",\\\"result\\\":\\\"spiderResultQueueRoot1\\\"}\",\"status\":0,\"logtime\":\"\"},\"url\":\"http://www.oschina.net/news/71456/y-ppa-manager\",\"selector\":{\"title\":\"Y PPA Manager 更新，修复PPA搜索功能\",\"content\":\"\\u003cspan id=\\\"result_box\\\" class=\\\"\\\"\\u003e由于一些\\u003cspan class=\\\"\\\"\\u003eLaunchpad的\\u003c/span\\u003e\\u003cspan class=\\\"\\\"\\u003e变化\\u003c/span\\u003e\\u003cspan class=\\\"\\\"\\u003e，\\u003c/span\\u003eY PPA Manager中的PPA搜索功能出现了一些问题，\\u003cspan class=\\\"\\\"\\u003e只返回\\u003c/span\\u003e\\u003cspan class=\\\"\\\"\\u003e搜索结果\\u003c/span\\u003e\\u003cspan class=\\\"\\\"\\u003e的一部分。\\u003c/span\\u003e2016年3月11日，\\u003cspan id=\\\"result_box\\\" class=\\\"\\\"\\u003e发布的最新的Y PPA Manager中\\u003c/span\\u003e，修复了\\u003cspan class=\\\"\\\"\\u003e这个问题。\\u003c/span\\u003e\\u003c/span\\u003e\\u003cbr/\\u003e\\u003cspan id=\\\"result_box\\\" class=\\\"\\\"\\u003e\\u003cspan class=\\\"\\\"\\u003e功能介绍：\\u003c/span\\u003e\\u003c/span\\u003e\\u003cbr/\\u003e\\u003cspan id=\\\"result_box\\\" class=\\\"\\\"\\u003e\\u003cspan class=\\\"\\\"\\u003e\\u003c/span\\u003e\\u003c/span\\u003e\\u003cbr/\\u003e详情请看：\\u003ca target=\\\"_blank\\\" href=\\\"http://www.webupd8.org/2016/03/y-ppa-manager-updated-with-important.html\\\"\\u003ehttp://www.webupd8.org/2016/03/y-ppa-manager-updated-with-important.html\\u003c/a\\u003e\\u003cbr/\\u003e\",\"section\":\"#ProjectNews \\u003e ul.p1\",\"href\":\"\",\"filter\":\"p[style]\"},\"selectorStr\":\"{\\\"title\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e h1\\\",\\\"content\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e div.Body.NewsContent.TextContent\\\\u003ep\\\",\\\"section\\\":\\\"#ProjectNews \\\\u003e ul.p1\\\",\\\"href\\\":\\\"\\\",\\\"filter\\\":\\\"p[style]\\\"}\",\"status\":0,\"logtime\":\"\"},\"Cron\":{\"id\":4,\"uid\":0,\"name\":\"每10秒执行一次\",\"cron\":\"10 * * * * *\",\"status\":0,\"logtime\":\"\"},\"status\":0,\"logtime\":\"\"}','2016-03-12 04:52:34'),(3,6,15,'http://www.oschina.net/news/71452/apache-phoenix-4-7','{\"id\":15,\"uid\":6,\"name\":\"测试\",\"url\":\"http://localhost\",\"document_set\":{\"position\":7,\"display\":1,\"uid\":1,\"nickname\":\"Administrator\",\"check\":1,\"category_id\":42,\"title\":\"搜索优化\",\"group_id\":0,\"model_id\":2},\"documentSetStr\":\"{\\\"position\\\":7,\\\"display\\\":1,\\\"uid\\\":1,\\\"nickname\\\":\\\"Administrator\\\",\\\"check\\\":1,\\\"category_id\\\":42,\\\"title\\\":\\\"搜索优化\\\",\\\"group_id\\\":0,\\\"model_id\\\":2}\",\"Rule\":{\"id\":22,\"uid\":0,\"name\":\"开源中国软件更新资讯\",\"spiderid\":1,\"manager\":{\"id\":1,\"uid\":0,\"name\":\"网站更新爬虫\",\"task\":\"spiderTaskQueueRoot1\",\"result\":\"spiderResultQueueRoot1\",\"queueName\":\"{\\\"task\\\":\\\"spiderTaskQueueRoot1\\\",\\\"result\\\":\\\"spiderResultQueueRoot1\\\"}\",\"status\":0,\"logtime\":\"\"},\"url\":\"http://www.oschina.net/news/71452/apache-phoenix-4-7\",\"selector\":{\"title\":\"Apache Phoenix 4.7 发布\",\"content\":\"Apache Phoenix 4.7 发布了。Apache Phoenix 是 \\u003ca target=\\\"_blank\\\" href=\\\"http://www.oschina.net/p/hbase\\\"\\u003eHBase\\u003c/a\\u003e 的 SQL 驱动。Phoenix 使得 HBase 支持通过 JDBC 的方式进行访问，并将你的 SQL 查询转成 HBase 的扫描和相应的动作。\\u003cbr/\\u003eApache Phoenix 4.7 发布主要包括：\\u003cbr/\\u003e下载地址：\\u003ca target=\\\"_blank\\\" href=\\\"https://phoenix.apache.org/download.html\\\"\\u003ehttps://phoenix.apache.org/download.html\\u003c/a\\u003e\\u003cbr/\\u003e\",\"section\":\"#ProjectNews \\u003e ul.p1\",\"href\":\"\",\"filter\":\"p[style]\"},\"selectorStr\":\"{\\\"title\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e h1\\\",\\\"content\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e div.Body.NewsContent.TextContent\\\\u003ep\\\",\\\"section\\\":\\\"#ProjectNews \\\\u003e ul.p1\\\",\\\"href\\\":\\\"\\\",\\\"filter\\\":\\\"p[style]\\\"}\",\"status\":0,\"logtime\":\"\"},\"Cron\":{\"id\":4,\"uid\":0,\"name\":\"每10秒执行一次\",\"cron\":\"10 * * * * *\",\"status\":0,\"logtime\":\"\"},\"status\":0,\"logtime\":\"\"}','2016-03-12 04:52:34'),(4,6,15,'http://www.oschina.net/news/71451/piwigo-2-8-0','{\"id\":15,\"uid\":6,\"name\":\"测试\",\"url\":\"http://localhost\",\"document_set\":{\"position\":7,\"display\":1,\"uid\":1,\"nickname\":\"Administrator\",\"check\":1,\"category_id\":42,\"title\":\"搜索优化\",\"group_id\":0,\"model_id\":2},\"documentSetStr\":\"{\\\"position\\\":7,\\\"display\\\":1,\\\"uid\\\":1,\\\"nickname\\\":\\\"Administrator\\\",\\\"check\\\":1,\\\"category_id\\\":42,\\\"title\\\":\\\"搜索优化\\\",\\\"group_id\\\":0,\\\"model_id\\\":2}\",\"Rule\":{\"id\":22,\"uid\":0,\"name\":\"开源中国软件更新资讯\",\"spiderid\":1,\"manager\":{\"id\":1,\"uid\":0,\"name\":\"网站更新爬虫\",\"task\":\"spiderTaskQueueRoot1\",\"result\":\"spiderResultQueueRoot1\",\"queueName\":\"{\\\"task\\\":\\\"spiderTaskQueueRoot1\\\",\\\"result\\\":\\\"spiderResultQueueRoot1\\\"}\",\"status\":0,\"logtime\":\"\"},\"url\":\"http://www.oschina.net/news/71451/piwigo-2-8-0\",\"selector\":{\"title\":\"Piwigo 2.8.0 发布\",\"content\":\"Piwigo 2.8.0 发布！截至今日，已有一年未推出大型版本更新了，我们也把服务器搬到了Github， \\n现在Piwigo为您呈现最新的2.8版本。在通知邮件中我们新增了认证密钥功能，为用户使用提供了便捷。支持更多的格式将会提升专业性，为企业及专业摄\\n影师提供了便利。对PHP 7的兼容，让我们向前迈进了一步。希望您能喜欢本次更新！\\u003cbr/\\u003e\\u003cspan style=\\\"font-size: 16px;\\\"\\u003e更新日志：\\u003c/span\\u003e\\u003cbr/\\u003e\\u003cspan style=\\\"font-size: 14px;\\\"\\u003e用户部分:\\u003c/span\\u003e\\u003cbr/\\u003e1、通知\\u003cbr/\\u003e2、多格式\\u003cbr/\\u003e3、新搜索功能\\u003cbr/\\u003e4、遗弃图片\\u003cbr/\\u003e5、水平重复水印\\u003cbr/\\u003e6、用户自定义\\u003cbr/\\u003e7、显示上传进度的favicon\\u003cbr/\\u003e\\u003cspan style=\\\"font-size: 14px;\\\"\\u003e技术部分:\\u003c/span\\u003e\\u003cbr/\\u003e1、PHP 7\\u003cbr/\\u003e2、Logger\\u003cbr/\\u003e3、上传文件处理\\u003cbr/\\u003e4、上传的Chunk尺寸\\u003cbr/\\u003e5、IPTC 关键词分割器\\u003cbr/\\u003e6、库更新\\u003cbr/\\u003e7、支持 Proxy\\u003cbr/\\u003e8、API 提升\\u003cbr/\\u003e\\u003cspan style=\\\"font-size: 16px;\\\"\\u003e更新截图：\\u003c/span\\u003e\\u003cbr/\\u003e\\u003cimg src=\\\"http://piwigo.org/screenshots/piwigo-2.8-auth-key.png\\\" class=\\\"screenshot\\\"/\\u003e\\u003cbr/\\u003e\\u003cimg src=\\\"http://piwigo.org/screenshots/piwigo-2.8-album-notify-users.png\\\" class=\\\"screenshot\\\"/\\u003e\\u003cimg src=\\\"http://piwigo.org/screenshots/piwigo-2.8-multiple-format.png\\\" class=\\\"screenshot\\\"/\\u003e\\u003cimg src=\\\"http://piwigo.org/screenshots/piwigo-2.8-search-tags.png\\\" class=\\\"screenshot\\\"/\\u003e\\u003cimg src=\\\"http://piwigo.org/screenshots/piwigo-2.8-orphan-photos.png\\\" class=\\\"screenshot\\\"/\\u003e\\u003cimg src=\\\"http://piwigo.org/screenshots/piwigo-2.8-watermark-yrepeat.jpg\\\" class=\\\"screenshot\\\"/\\u003e\\u003cimg src=\\\"http://piwigo.org/screenshots/piwigo-2.8-edit-user-popin.png\\\" class=\\\"screenshot\\\"/\\u003e\\u003cimg src=\\\"http://piwigo.org/screenshots/piwigo-2.8-upload-progress-favicon.png\\\" class=\\\"screenshot\\\"/\\u003e\\u003cbr/\\u003e\\u003cspan style=\\\"font-size: 16px;\\\"\\u003e文章来源：\\u003c/span\\u003e\\u003ca target=\\\"_blank\\\" href=\\\"http://cn.piwigo.org/forum/viewtopic.php?id=16059\\\"\\u003ehttp://cn.piwigo.org/forum/viewtopic.php?id=16059\\u003c/a\\u003e\\u003cbr/\\u003e\",\"section\":\"#ProjectNews \\u003e ul.p1\",\"href\":\"\",\"filter\":\"p[style]\"},\"selectorStr\":\"{\\\"title\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e h1\\\",\\\"content\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e div.Body.NewsContent.TextContent\\\\u003ep\\\",\\\"section\\\":\\\"#ProjectNews \\\\u003e ul.p1\\\",\\\"href\\\":\\\"\\\",\\\"filter\\\":\\\"p[style]\\\"}\",\"status\":0,\"logtime\":\"\"},\"Cron\":{\"id\":4,\"uid\":0,\"name\":\"每10秒执行一次\",\"cron\":\"10 * * * * *\",\"status\":0,\"logtime\":\"\"},\"status\":0,\"logtime\":\"\"}','2016-03-12 04:52:34'),(5,6,15,'http://www.oschina.net/news/71450/smartgit-7-2','{\"id\":15,\"uid\":6,\"name\":\"测试\",\"url\":\"http://localhost\",\"document_set\":{\"position\":7,\"display\":1,\"uid\":1,\"nickname\":\"Administrator\",\"check\":1,\"category_id\":42,\"title\":\"搜索优化\",\"group_id\":0,\"model_id\":2},\"documentSetStr\":\"{\\\"position\\\":7,\\\"display\\\":1,\\\"uid\\\":1,\\\"nickname\\\":\\\"Administrator\\\",\\\"check\\\":1,\\\"category_id\\\":42,\\\"title\\\":\\\"搜索优化\\\",\\\"group_id\\\":0,\\\"model_id\\\":2}\",\"Rule\":{\"id\":22,\"uid\":0,\"name\":\"开源中国软件更新资讯\",\"spiderid\":1,\"manager\":{\"id\":1,\"uid\":0,\"name\":\"网站更新爬虫\",\"task\":\"spiderTaskQueueRoot1\",\"result\":\"spiderResultQueueRoot1\",\"queueName\":\"{\\\"task\\\":\\\"spiderTaskQueueRoot1\\\",\\\"result\\\":\\\"spiderResultQueueRoot1\\\"}\",\"status\":0,\"logtime\":\"\"},\"url\":\"http://www.oschina.net/news/71450/smartgit-7-2\",\"selector\":{\"title\":\"SmartGit 7.2 发布，Git 客户端\",\"content\":\"SmartGit 7.2 发布了。SmartGit 是一个 \\u003ca href=\\\"http://www.oschina.net/p/git\\\"\\u003eGit\\u003c/a\\u003e 版本控制系统的图形化客户端程序。\\u003cbr/\\u003e改进：\\u003cbr/\\u003e- Linux: GTK3 is now the default. Please report all bugs and only switch back\\u003cbr/\\u003e  to GTK2 (SWT_GTK3=0) if absolutely necessary.\\u003cbr/\\u003e- updated SWT to version 4.614\\u003cbr/\\u003e新特性：\\u003cbr/\\u003eBug 修复\\u003cbr/\\u003e- SVN:\\u003cbr/\\u003e  - Checkout: line ending correction did not work properly if \\u0026#34;eol\\u0026#34; attribute\\u003cbr/\\u003e    was set, but \\u0026#34;text\\u0026#34; wan\\u0026#39;t\\u003cbr/\\u003e- Linux:\\u003cbr/\\u003e  - GTK3:\\u003cbr/\\u003e    - Preferences, Built-In Text Editors: first tab is not drawn\\u003cbr/\\u003e    - Setup wizard: wrong input field focused by default\\u003cbr/\\u003e下载地址：\\u003ca target=\\\"_blank\\\" href=\\\"http://www.syntevo.com/smartgit/preview\\\"\\u003ehttp://www.syntevo.com/smartgit/preview\\u003c/a\\u003e\\u003cbr/\\u003e\",\"section\":\"#ProjectNews \\u003e ul.p1\",\"href\":\"\",\"filter\":\"p[style]\"},\"selectorStr\":\"{\\\"title\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e h1\\\",\\\"content\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e div.Body.NewsContent.TextContent\\\\u003ep\\\",\\\"section\\\":\\\"#ProjectNews \\\\u003e ul.p1\\\",\\\"href\\\":\\\"\\\",\\\"filter\\\":\\\"p[style]\\\"}\",\"status\":0,\"logtime\":\"\"},\"Cron\":{\"id\":4,\"uid\":0,\"name\":\"每10秒执行一次\",\"cron\":\"10 * * * * *\",\"status\":0,\"logtime\":\"\"},\"status\":0,\"logtime\":\"\"}','2016-03-12 04:52:34'),(6,6,15,'http://www.oschina.net/news/71449/appcode-3-4-eap','{\"id\":15,\"uid\":6,\"name\":\"测试\",\"url\":\"http://localhost\",\"document_set\":{\"position\":7,\"display\":1,\"uid\":1,\"nickname\":\"Administrator\",\"check\":1,\"category_id\":42,\"title\":\"搜索优化\",\"group_id\":0,\"model_id\":2},\"documentSetStr\":\"{\\\"position\\\":7,\\\"display\\\":1,\\\"uid\\\":1,\\\"nickname\\\":\\\"Administrator\\\",\\\"check\\\":1,\\\"category_id\\\":42,\\\"title\\\":\\\"搜索优化\\\",\\\"group_id\\\":0,\\\"model_id\\\":2}\",\"Rule\":{\"id\":22,\"uid\":0,\"name\":\"开源中国软件更新资讯\",\"spiderid\":1,\"manager\":{\"id\":1,\"uid\":0,\"name\":\"网站更新爬虫\",\"task\":\"spiderTaskQueueRoot1\",\"result\":\"spiderResultQueueRoot1\",\"queueName\":\"{\\\"task\\\":\\\"spiderTaskQueueRoot1\\\",\\\"result\\\":\\\"spiderResultQueueRoot1\\\"}\",\"status\":0,\"logtime\":\"\"},\"url\":\"http://www.oschina.net/news/71449/appcode-3-4-eap\",\"selector\":{\"title\":\"AppCode 3.4 EAP 发布：Swift的代码折叠\",\"content\":\"AppCode 3.4 EAP新版本，  \\u003cspan id=\\\"result_box\\\"\\u003e建立145.256\\u003cspan class=\\\"\\\"\\u003e\\u003ca target=\\\"_blank\\\" href=\\\"https://confluence.jetbrains.com/display/OBJC/AppCode+EAP\\\"\\u003e可供下载\\u003c/a\\u003e\\u003c/span\\u003e\\u003cspan class=\\\"\\\"\\u003e。\\u003cspan id=\\\"result_box\\\"\\u003e如果你使用的是以前的3.4 EAP构建\\u003cspan class=\\\"\\\"\\u003e，那么这个补丁更新也能找到。\\u003c/span\\u003e\\u003c/span\\u003e\\u003c/span\\u003e\\u003c/span\\u003e\\u003cbr/\\u003e\\u003cspan style=\\\"font-size: 14px;\\\"\\u003e更新日志：\\u003c/span\\u003e\\u003cbr/\\u003e\\u003cspan style=\\\"font-size: 14px;\\\"\\u003e文章来源：\\u003c/span\\u003e\\u003ca target=\\\"_blank\\\" href=\\\"http://blog.jetbrains.com/objc/2016/03/appcode-3-4-eap-code-folding/\\\"\\u003ehttp://blog.jetbrains.com/objc/2016/03/appcode-3-4-eap-code-folding/\\u003c/a\\u003e\\u003cbr/\\u003e\",\"section\":\"#ProjectNews \\u003e ul.p1\",\"href\":\"\",\"filter\":\"p[style]\"},\"selectorStr\":\"{\\\"title\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e h1\\\",\\\"content\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e div.Body.NewsContent.TextContent\\\\u003ep\\\",\\\"section\\\":\\\"#ProjectNews \\\\u003e ul.p1\\\",\\\"href\\\":\\\"\\\",\\\"filter\\\":\\\"p[style]\\\"}\",\"status\":0,\"logtime\":\"\"},\"Cron\":{\"id\":4,\"uid\":0,\"name\":\"每10秒执行一次\",\"cron\":\"10 * * * * *\",\"status\":0,\"logtime\":\"\"},\"status\":0,\"logtime\":\"\"}','2016-03-12 04:52:34'),(7,6,15,'http://www.oschina.net/news/71448/pycharm-5-1-beta-2','{\"id\":15,\"uid\":6,\"name\":\"测试\",\"url\":\"http://localhost\",\"document_set\":{\"position\":7,\"display\":1,\"uid\":1,\"nickname\":\"Administrator\",\"check\":1,\"category_id\":42,\"title\":\"搜索优化\",\"group_id\":0,\"model_id\":2},\"documentSetStr\":\"{\\\"position\\\":7,\\\"display\\\":1,\\\"uid\\\":1,\\\"nickname\\\":\\\"Administrator\\\",\\\"check\\\":1,\\\"category_id\\\":42,\\\"title\\\":\\\"搜索优化\\\",\\\"group_id\\\":0,\\\"model_id\\\":2}\",\"Rule\":{\"id\":22,\"uid\":0,\"name\":\"开源中国软件更新资讯\",\"spiderid\":1,\"manager\":{\"id\":1,\"uid\":0,\"name\":\"网站更新爬虫\",\"task\":\"spiderTaskQueueRoot1\",\"result\":\"spiderResultQueueRoot1\",\"queueName\":\"{\\\"task\\\":\\\"spiderTaskQueueRoot1\\\",\\\"result\\\":\\\"spiderResultQueueRoot1\\\"}\",\"status\":0,\"logtime\":\"\"},\"url\":\"http://www.oschina.net/news/71448/pycharm-5-1-beta-2\",\"selector\":{\"title\":\"PyCharm 5.1 Beta 2 发布，Python 集成开发环境\",\"content\":\"PyCharm 5.1 Beta 2 发布了。PyCharm是由JetBrains打造的一款Python IDE。我们知道，VS2010的重构插件Resharper就是出自JetBrains之手。那么，PyCharm有什么吸引人的特点呢？\\u003cbr/\\u003e首先，PyCharm用于一般IDE具备的功能，比如， 调试、语法高亮、Project管理、代码跳转、智能提示、自动完成、单元测试、版本控制……\\u003cbr/\\u003e另外，PyCharm还提供了一些很好的功能用于\\u003ca href=\\\"http://www.oschina.net/p/django\\\"\\u003eDjango\\u003c/a\\u003e开发，同时支持Google App Engine，更酷的是，PyCharm支持\\u003ca href=\\\"http://www.oschina.net/p/ironpython\\\"\\u003eIronPython\\u003c/a\\u003e！\\u003cbr/\\u003e改进日志：\\u003cbr/\\u003e下载地址：\\u003cbr/\\u003e\\u003cbr/\\u003e\",\"section\":\"#ProjectNews \\u003e ul.p1\",\"href\":\"\",\"filter\":\"p[style]\"},\"selectorStr\":\"{\\\"title\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e h1\\\",\\\"content\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e div.Body.NewsContent.TextContent\\\\u003ep\\\",\\\"section\\\":\\\"#ProjectNews \\\\u003e ul.p1\\\",\\\"href\\\":\\\"\\\",\\\"filter\\\":\\\"p[style]\\\"}\",\"status\":0,\"logtime\":\"\"},\"Cron\":{\"id\":4,\"uid\":0,\"name\":\"每10秒执行一次\",\"cron\":\"10 * * * * *\",\"status\":0,\"logtime\":\"\"},\"status\":0,\"logtime\":\"\"}','2016-03-12 04:52:34'),(8,6,15,'http://www.oschina.net/news/71447/util-linux-v2-28-rc1','{\"id\":15,\"uid\":6,\"name\":\"测试\",\"url\":\"http://localhost\",\"document_set\":{\"position\":7,\"display\":1,\"uid\":1,\"nickname\":\"Administrator\",\"check\":1,\"category_id\":42,\"title\":\"搜索优化\",\"group_id\":0,\"model_id\":2},\"documentSetStr\":\"{\\\"position\\\":7,\\\"display\\\":1,\\\"uid\\\":1,\\\"nickname\\\":\\\"Administrator\\\",\\\"check\\\":1,\\\"category_id\\\":42,\\\"title\\\":\\\"搜索优化\\\",\\\"group_id\\\":0,\\\"model_id\\\":2}\",\"Rule\":{\"id\":22,\"uid\":0,\"name\":\"开源中国软件更新资讯\",\"spiderid\":1,\"manager\":{\"id\":1,\"uid\":0,\"name\":\"网站更新爬虫\",\"task\":\"spiderTaskQueueRoot1\",\"result\":\"spiderResultQueueRoot1\",\"queueName\":\"{\\\"task\\\":\\\"spiderTaskQueueRoot1\\\",\\\"result\\\":\\\"spiderResultQueueRoot1\\\"}\",\"status\":0,\"logtime\":\"\"},\"url\":\"http://www.oschina.net/news/71447/util-linux-v2-28-rc1\",\"selector\":{\"title\":\"util-linux v2.28-rc1 发布，Linux 实用工具集\",\"content\":\"util-linux v2.28-rc1 发布了。\\u003cbr/\\u003e改进日志：\\u003cbr/\\u003e更多请看：\\u003ca target=\\\"_blank\\\" href=\\\"http://git.kernel.org/cgit/utils/util-linux/util-linux.git/tree/Documentation/releases/v2.28-ReleaseNotes?id=ec9538250cde31be6ac522609624148a4b2d2f71\\\"\\u003ehttp://git.kernel.org/cgit/utils/util-linux/util-linux.git/tree/Documentation/releases/v2.28-ReleaseNotes?id=ec9538250cde31be6ac522609624148a4b2d2f71\\u003c/a\\u003e\\u003cbr/\\u003e\\u003cbr/\\u003e下载地址：\",\"section\":\"#ProjectNews \\u003e ul.p1\",\"href\":\"\",\"filter\":\"p[style]\"},\"selectorStr\":\"{\\\"title\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e h1\\\",\\\"content\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e div.Body.NewsContent.TextContent\\\\u003ep\\\",\\\"section\\\":\\\"#ProjectNews \\\\u003e ul.p1\\\",\\\"href\\\":\\\"\\\",\\\"filter\\\":\\\"p[style]\\\"}\",\"status\":0,\"logtime\":\"\"},\"Cron\":{\"id\":4,\"uid\":0,\"name\":\"每10秒执行一次\",\"cron\":\"10 * * * * *\",\"status\":0,\"logtime\":\"\"},\"status\":0,\"logtime\":\"\"}','2016-03-12 04:52:34'),(9,6,15,'http://www.oschina.net/news/71444/rails-4-1-15','{\"id\":15,\"uid\":6,\"name\":\"测试\",\"url\":\"http://localhost\",\"document_set\":{\"position\":7,\"display\":1,\"uid\":1,\"nickname\":\"Administrator\",\"check\":1,\"category_id\":42,\"title\":\"搜索优化\",\"group_id\":0,\"model_id\":2},\"documentSetStr\":\"{\\\"position\\\":7,\\\"display\\\":1,\\\"uid\\\":1,\\\"nickname\\\":\\\"Administrator\\\",\\\"check\\\":1,\\\"category_id\\\":42,\\\"title\\\":\\\"搜索优化\\\",\\\"group_id\\\":0,\\\"model_id\\\":2}\",\"Rule\":{\"id\":22,\"uid\":0,\"name\":\"开源中国软件更新资讯\",\"spiderid\":1,\"manager\":{\"id\":1,\"uid\":0,\"name\":\"网站更新爬虫\",\"task\":\"spiderTaskQueueRoot1\",\"result\":\"spiderResultQueueRoot1\",\"queueName\":\"{\\\"task\\\":\\\"spiderTaskQueueRoot1\\\",\\\"result\\\":\\\"spiderResultQueueRoot1\\\"}\",\"status\":0,\"logtime\":\"\"},\"url\":\"http://www.oschina.net/news/71444/rails-4-1-15\",\"selector\":{\"title\":\"Rails 4.1.15 发布，开源网络应用框架\",\"content\":\"Rails 4.1.15 发布了，\\u003cstrong\\u003eRails\\u003c/strong\\u003e 是一个用于开发数据库驱动的网络应用程序的完整框架。Rails基于MVC（模型- 视图- \\n控制器）设计模式。从视图中的Ajax应用，到控制器中的访问请求和反馈，到封装数据库的模型，Rails \\n为你提供一个纯Ruby的开发环境。发布网站时，你只需要一个数据库和一个网络服务器即可。\\u003cbr/\\u003eRuby On \\nRails是一个用于编写网络应用程序的软件包.它基于一种计算机软件语言Ruby,给程序开发人员提供了强大的框架支持.你可以用比以前少的多的代码和\\n 短的多的时间编写出一流的网络软件.比较著名的社区网站43things.com, odeo.com和basecamphq.com就是用Ruby \\nOn Rails编写的\\u003cbr/\\u003e暂无相关改进说明，持续关注点击\\u003ca target=\\\"_blank\\\" href=\\\"https://github.com/rails/rails/blob/v4.1.15/actionmailer/CHANGELOG.md\\\"\\u003e这里\\u003c/a\\u003e。\\u003cbr/\\u003e下载地址：\",\"section\":\"#ProjectNews \\u003e ul.p1\",\"href\":\"\",\"filter\":\"p[style]\"},\"selectorStr\":\"{\\\"title\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e h1\\\",\\\"content\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e div.Body.NewsContent.TextContent\\\\u003ep\\\",\\\"section\\\":\\\"#ProjectNews \\\\u003e ul.p1\\\",\\\"href\\\":\\\"\\\",\\\"filter\\\":\\\"p[style]\\\"}\",\"status\":0,\"logtime\":\"\"},\"Cron\":{\"id\":4,\"uid\":0,\"name\":\"每10秒执行一次\",\"cron\":\"10 * * * * *\",\"status\":0,\"logtime\":\"\"},\"status\":0,\"logtime\":\"\"}','2016-03-12 04:52:34'),(10,6,15,'http://www.oschina.net/news/71443/xwiki-6-4-8','{\"id\":15,\"uid\":6,\"name\":\"测试\",\"url\":\"http://localhost\",\"document_set\":{\"position\":7,\"display\":1,\"uid\":1,\"nickname\":\"Administrator\",\"check\":1,\"category_id\":42,\"title\":\"搜索优化\",\"group_id\":0,\"model_id\":2},\"documentSetStr\":\"{\\\"position\\\":7,\\\"display\\\":1,\\\"uid\\\":1,\\\"nickname\\\":\\\"Administrator\\\",\\\"check\\\":1,\\\"category_id\\\":42,\\\"title\\\":\\\"搜索优化\\\",\\\"group_id\\\":0,\\\"model_id\\\":2}\",\"Rule\":{\"id\":22,\"uid\":0,\"name\":\"开源中国软件更新资讯\",\"spiderid\":1,\"manager\":{\"id\":1,\"uid\":0,\"name\":\"网站更新爬虫\",\"task\":\"spiderTaskQueueRoot1\",\"result\":\"spiderResultQueueRoot1\",\"queueName\":\"{\\\"task\\\":\\\"spiderTaskQueueRoot1\\\",\\\"result\\\":\\\"spiderResultQueueRoot1\\\"}\",\"status\":0,\"logtime\":\"\"},\"url\":\"http://www.oschina.net/news/71443/xwiki-6-4-8\",\"selector\":{\"title\":\"XWiki 6.4.8 发布，Java的Wiki 系统 \",\"content\":\"XWiki 6.4.8 发布了，XWiki是一个由Java编写的基于LGPL协议发布的开源wiki和应用平台。它的开发平台特性允许创建协作式Web应用，同时也提供了构建于平台之上的打包应用（第二代wiki）。\\u003cbr/\\u003e改进记录：\\u003cbr/\\u003e详情请看：\\u003ca target=\\\"_blank\\\" href=\\\"http://jira.xwiki.org/secure/Dashboard.jspa?selectPageId=13600\\\"\\u003ehttp://jira.xwiki.org/secure/Dashboard.jspa?selectPageId=13600\\u003c/a\\u003e\\u003cbr/\\u003e下载地址：\\u003ca target=\\\"_blank\\\" href=\\\"http://www.xwiki.org/xwiki/bin/view/Main/Download\\\"\\u003ehttp://www.xwiki.org/xwiki/bin/view/Main/Download\\u003c/a\\u003e\\u003cbr/\\u003e\",\"section\":\"#ProjectNews \\u003e ul.p1\",\"href\":\"\",\"filter\":\"p[style]\"},\"selectorStr\":\"{\\\"title\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e h1\\\",\\\"content\\\":\\\"#NewsChannel \\\\u003e div.NewsBody \\\\u003e div \\\\u003e div.NewsEntity \\\\u003e div.Body.NewsContent.TextContent\\\\u003ep\\\",\\\"section\\\":\\\"#ProjectNews \\\\u003e ul.p1\\\",\\\"href\\\":\\\"\\\",\\\"filter\\\":\\\"p[style]\\\"}\",\"status\":0,\"logtime\":\"\"},\"Cron\":{\"id\":4,\"uid\":0,\"name\":\"每10秒执行一次\",\"cron\":\"10 * * * * *\",\"status\":0,\"logtime\":\"\"},\"status\":0,\"logtime\":\"\"}','2016-03-12 04:52:34');
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
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

-- Dump completed on 2016-06-20  9:53:39
