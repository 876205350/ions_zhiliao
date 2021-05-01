/*
SQLyog Ultimate v11.25 (64 bit)
MySQL - 5.7.23-log : Database - ions_zhiliao
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`ions_zhiliao` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `ions_zhiliao`;

/*Table structure for table `sys_auth` */

DROP TABLE IF EXISTS `sys_auth`;

CREATE TABLE `sys_auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_name` varchar(64) NOT NULL DEFAULT '' COMMENT '权限名称',
  `url_for` varchar(255) NOT NULL DEFAULT '' COMMENT 'url反转',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父节点id',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `is_active` int(11) NOT NULL DEFAULT '0' COMMENT '1:启用 0:停用',
  `is_delete` int(11) NOT NULL DEFAULT '0' COMMENT '1:删除 0:未删除',
  `weight` int(11) NOT NULL DEFAULT '0' COMMENT '权重',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;

/*Data for the table `sys_auth` */

insert  into `sys_auth`(`id`,`auth_name`,`url_for`,`pid`,`desc`,`create_time`,`is_active`,`is_delete`,`weight`) values (1,'用户管理','#',0,'用户管理，一级菜单','2021-04-03 06:20:50',1,0,100),(2,'用户列表','UserController.List',1,'用户列表，二级菜单','2021-04-04 05:26:00',1,0,99),(3,'权限管理','#',0,'权限管理，一级菜单','2021-04-04 05:26:39',1,0,90),(4,'权限列表','AuthController.List',3,'权限列表，二级菜单','2021-04-04 16:19:37',1,0,89),(5,'角色管理','RoleController.List',3,'角色管理，二级菜单','2021-04-04 16:24:39',1,0,88),(6,'个人中心','#',0,'个人中心，一级菜单','2021-04-04 16:31:32',1,0,85),(7,'财务中心','#',0,'财务中心，一级菜单','2021-04-04 16:33:53',1,0,80),(8,'财务报表','CaiWuEchartDataController.Get',7,'财务报表，二级菜单','2021-04-07 00:51:40',1,0,79),(9,'个人信息','MyCenterController.Get',6,'个人信息，二级菜单','2021-04-15 23:52:33',1,0,84),(10,'工资条','SalarySlipController.Get',6,'工资条，二级菜单','2021-04-15 23:53:11',1,0,83),(11,'工资条','CaiWuSalarySlipController.Get',7,'工资条，二级菜单','2021-04-20 22:35:24',1,0,78),(12,'内容管理','#',0,'内容管理，一级菜单','2021-04-24 10:13:19',1,0,79),(13,'栏目列表','CategoryController.Get',12,'栏目管理，二级菜单','2021-04-24 10:15:16',1,0,78),(14,'新闻列表','NewsController.Get',12,'新闻列表，二级菜单','2021-04-24 10:16:01',1,0,78),(15,'车辆管理','#',0,'车辆管理，一级菜单','2021-04-25 23:35:42',1,0,88),(16,'车辆品牌','CarBrandController.Get',15,'车辆品牌管理，二级菜单','2021-04-25 23:36:48',1,0,85),(17,'车辆列表','CarsController.Get',15,'车辆列表，二级菜单','2021-04-25 23:37:27',1,0,84),(18,'车辆申请','CarsApplyController.Get',15,'车辆申请，二级菜单','2021-04-25 23:38:09',1,0,83),(19,'车辆审批','CarsApplyController.AuditApply',15,'车辆审批，二级菜单','2021-04-25 23:38:38',1,0,82),(20,'我的申请','CarsApplyController.MyApply',15,'我的申请，二级菜单','2021-04-27 23:07:30',1,0,81),(21,'报表管理','#',0,'报表管理，一级菜单','2021-04-30 01:37:32',1,0,50),(22,'财务报表','EchartsCaiwuController.Get',21,'财务报表，二级菜单','2021-04-30 01:38:26',1,0,49),(23,'业务报表','EchartsBusinessController.Get',21,'业务报表，二级菜单','2021-04-30 01:38:57',1,0,48),(24,'课程报表','EchartsCourseController.Get',21,'课程报表，二级菜单','2021-04-30 01:39:30',1,0,47);

/*Table structure for table `sys_caiwu_data` */

DROP TABLE IF EXISTS `sys_caiwu_data`;

CREATE TABLE `sys_caiwu_data` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cai_wu_date` varchar(32) NOT NULL DEFAULT '' COMMENT '财务月份',
  `sales_volume` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '本月销售额',
  `student_increase` int(11) NOT NULL DEFAULT '0' COMMENT '学员增加数',
  `django` int(11) NOT NULL DEFAULT '0' COMMENT 'django课程卖出数量',
  `vue_django` int(11) NOT NULL DEFAULT '0' COMMENT 'vue+django课程卖出数量',
  `celery` int(11) NOT NULL DEFAULT '0' COMMENT 'celery课程卖出数量',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

/*Data for the table `sys_caiwu_data` */

insert  into `sys_caiwu_data`(`id`,`cai_wu_date`,`sales_volume`,`student_increase`,`django`,`vue_django`,`celery`,`create_time`) values (1,'2021-01','35.21',10,50,20,10,'2021-04-23 00:12:40'),(2,'2021-02','50.21',10,50,20,10,'2021-04-23 00:12:40'),(3,'2021-03','25.56',260,60,30,10,'2021-04-24 02:02:15'),(4,'2021-04','45.25',260,60,30,10,'2021-04-24 02:02:15');

/*Table structure for table `sys_cars` */

DROP TABLE IF EXISTS `sys_cars`;

CREATE TABLE `sys_cars` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '车辆名称',
  `car_brand_id` int(11) NOT NULL COMMENT '车辆品牌外键',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '1:可借,2:不可借',
  `is_active` int(11) NOT NULL DEFAULT '1' COMMENT '启用:1,停用:0',
  `is_delete` int(11) NOT NULL DEFAULT '0' COMMENT '删除:1,未删除:0',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

/*Data for the table `sys_cars` */

insert  into `sys_cars`(`id`,`name`,`car_brand_id`,`status`,`is_active`,`is_delete`,`create_time`) values (1,'宝马X5',1,1,1,0,'2021-04-26 23:32:11'),(2,'奔驰300',2,1,1,0,'2021-04-26 23:32:11'),(3,'大众200',3,1,1,0,'2021-04-27 00:27:05'),(4,'宝马X3',1,1,1,0,'2021-04-27 00:50:12'),(5,'奔驰500',2,1,1,0,'2021-04-27 00:59:45'),(6,'桑塔纳2000',3,1,1,0,'2021-04-27 01:05:22');

/*Table structure for table `sys_cars_apply` */

DROP TABLE IF EXISTS `sys_cars_apply`;

CREATE TABLE `sys_cars_apply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `cars_id` int(11) NOT NULL,
  `reason` varchar(255) NOT NULL DEFAULT '' COMMENT '申请理由',
  `destination` varchar(64) NOT NULL DEFAULT '' COMMENT '目的地',
  `return_date` date NOT NULL COMMENT '归还日期',
  `return_status` int(11) NOT NULL DEFAULT '0' COMMENT '1:已归还，0：未归还',
  `audit_status` int(11) NOT NULL DEFAULT '3' COMMENT '1:同意，2:未同意，3:未审批',
  `audit_option` varchar(255) NOT NULL DEFAULT '' COMMENT '审批意见',
  `is_active` int(11) NOT NULL DEFAULT '1' COMMENT '启用:1,停用:0',
  `is_delete` int(11) NOT NULL DEFAULT '0' COMMENT '删除:1,未删除:0',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `notify_tag` int(11) NOT NULL DEFAULT '0' COMMENT '1:已发送通知，0：未发送通知',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

/*Data for the table `sys_cars_apply` */

insert  into `sys_cars_apply`(`id`,`user_id`,`cars_id`,`reason`,`destination`,`return_date`,`return_status`,`audit_status`,`audit_option`,`is_active`,`is_delete`,`create_time`,`notify_tag`) values (1,1,6,'外出谈合作','北京海淀区百度科技园','2021-04-30',1,1,'同意前往',1,0,'2021-04-30 01:15:24',1),(2,1,5,'公车私用','顺义区','2021-04-30',0,2,'禁止公车私用',1,0,'2021-04-27 23:53:12',0),(3,1,6,'谈生意','大兴区','2021-04-30',1,1,'同意前往',1,0,'2021-04-30 01:15:24',1);

/*Table structure for table `sys_cars_brand` */

DROP TABLE IF EXISTS `sys_cars_brand`;

CREATE TABLE `sys_cars_brand` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '品牌名称',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '品牌描述',
  `is_active` int(11) NOT NULL DEFAULT '1' COMMENT '启用:1,停用:0',
  `is_delete` int(11) NOT NULL DEFAULT '0' COMMENT '删除:1,未删除:0',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

/*Data for the table `sys_cars_brand` */

insert  into `sys_cars_brand`(`id`,`name`,`desc`,`is_active`,`is_delete`,`create_time`) values (1,'宝马','宝马',1,0,'2021-04-26 00:52:44'),(2,'奔驰','奔驰',1,0,'2021-04-26 00:53:11'),(3,'大众','一汽大众',1,0,'2021-04-26 01:12:25');

/*Table structure for table `sys_category` */

DROP TABLE IF EXISTS `sys_category`;

CREATE TABLE `sys_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '分类名称',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `is_active` int(11) NOT NULL DEFAULT '1' COMMENT '1启用，0停用',
  `is_delete` int(11) NOT NULL DEFAULT '0' COMMENT '1删除，0未删除',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

/*Data for the table `sys_category` */

insert  into `sys_category`(`id`,`name`,`desc`,`is_active`,`is_delete`,`create_time`) values (1,'高管','高管相关',1,0,'2021-04-24 11:04:00'),(2,'财务','财务专栏',1,0,'2021-04-24 12:45:24'),(3,'XX','XX',1,0,'2021-04-24 12:49:09');

/*Table structure for table `sys_message_notify` */

DROP TABLE IF EXISTS `sys_message_notify`;

CREATE TABLE `sys_message_notify` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `flag` int(11) NOT NULL DEFAULT '1' COMMENT '1:车辆逾期，2:所有通知',
  `title` varchar(64) NOT NULL DEFAULT '' COMMENT '消息标题',
  `content` longtext NOT NULL COMMENT '消息内容',
  `user_id` int(11) NOT NULL,
  `read_tag` int(11) NOT NULL DEFAULT '0' COMMENT '1:已读，0:未读',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

/*Data for the table `sys_message_notify` */

insert  into `sys_message_notify`(`id`,`flag`,`title`,`content`,`user_id`,`read_tag`) values (3,1,'车辆归还逾期','admin0用户，你借的车辆桑塔纳2000，归还时间为2021-04-29，已经逾期，请尽快归还',1,0),(4,1,'车辆归还逾期','admin0用户，你借的车辆桑塔纳2000，归还时间为2021-04-30，已经逾期，请尽快归还',1,1);

/*Table structure for table `sys_news` */

DROP TABLE IF EXISTS `sys_news`;

CREATE TABLE `sys_news` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(64) NOT NULL DEFAULT '' COMMENT '新闻标题',
  `content` longtext NOT NULL COMMENT '新闻内容',
  `category_id` int(11) NOT NULL,
  `is_active` int(11) NOT NULL DEFAULT '1' COMMENT '1启用，0停用',
  `is_delete` int(11) NOT NULL DEFAULT '0' COMMENT '1删除，0未删除',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

/*Data for the table `sys_news` */

insert  into `sys_news`(`id`,`title`,`content`,`category_id`,`is_active`,`is_delete`,`create_time`) values (1,'XX','XX',1,1,0,'2021-04-24 13:56:38'),(2,'招聘','<p>招聘员工</p>',2,1,0,'2021-04-24 17:51:09'),(3,'招聘AD001','<p>会计专业</p>',2,1,0,'2021-04-24 17:53:17'),(4,'测试','<p><img src=\"http://127.0.0.1:8080/upload/news_img/1619261951-95eef01f3a292df5f58a1837bd315c6034a87342.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dii\" data-code=\"200\" data-msg=\"文件保存成功\"></p>',2,1,0,'2021-04-24 19:00:10'),(5,'测试AD001','<p>招聘<img src=\"http://127.0.0.1:8080/upload/news_img/1619264831-c21cb045a36580161a897e6e3a164e2a243038d7.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\" data-code=\"200\" data-msg=\"文件保存成功\"></p>',2,1,0,'2021-04-24 19:47:17'),(6,'测试AD002','<p><img src=\"http://127.0.0.1:8080/upload/news_img/1619265041-timg (1).jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\" data-code=\"200\" data-msg=\"文件保存成功\"></p>',1,1,0,'2021-04-24 19:50:47'),(7,'非的故事','<p>范德萨</p>',2,1,0,'2021-04-24 19:53:43'),(8,'都是','<p>撒旦</p>',2,1,0,'2021-04-24 19:58:06'),(9,'第三方','<p>&nbsp;撒旦</p><p><img src=\"http://127.0.0.1:8080/upload/news_img/1619265501-timg.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\" data-code=\"200\" data-msg=\"文件保存成功\"></p>',2,1,0,'2021-04-24 19:58:28'),(10,'富商大贾','<p>大 士大夫</p>',3,1,0,'2021-04-24 20:09:19'),(11,'第三方','<p>撒旦</p>',2,1,0,'2021-04-24 20:22:05'),(12,'撒旦法','<p>啥打法上</p>',2,1,0,'2021-04-24 20:26:36'),(13,'大是大非','<p>手打</p>',2,1,0,'2021-04-24 20:33:28'),(14,'撒旦法','<p>撒旦法</p>',2,1,0,'2021-04-24 20:35:09'),(15,'实打实','<p>撒旦法</p>',2,1,0,'2021-04-24 20:37:01');

/*Table structure for table `sys_role` */

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(64) NOT NULL DEFAULT '',
  `desc` varchar(255) NOT NULL DEFAULT '',
  `is_active` int(11) NOT NULL DEFAULT '0' COMMENT '启用状态',
  `is_delete` int(11) NOT NULL DEFAULT '0' COMMENT '删除状态',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

/*Data for the table `sys_role` */

insert  into `sys_role`(`id`,`role_name`,`desc`,`is_active`,`is_delete`,`create_time`) values (1,'超级管理员','最大权限',1,0,'2021-04-06 15:23:58'),(2,'管理员','二级权限',1,0,'2021-04-06 16:02:34'),(3,'主任','主任',1,0,'2021-04-07 00:41:03'),(4,'副主任','副主任',1,0,'2021-04-07 00:42:27'),(5,'组长','组长',1,0,'2021-04-07 00:45:09'),(6,'副组长','副组长',1,0,'2021-04-07 01:10:08'),(7,'A小组','A小组',1,0,'2021-04-07 18:14:31');

/*Table structure for table `sys_role_sys_auths` */

DROP TABLE IF EXISTS `sys_role_sys_auths`;

CREATE TABLE `sys_role_sys_auths` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `sys_role_id` int(11) NOT NULL,
  `sys_auth_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=122 DEFAULT CHARSET=utf8;

/*Data for the table `sys_role_sys_auths` */

insert  into `sys_role_sys_auths`(`id`,`sys_role_id`,`sys_auth_id`) values (19,2,1),(20,2,2),(21,2,6),(22,2,7),(23,2,8),(98,1,1),(99,1,2),(100,1,3),(101,1,4),(102,1,5),(103,1,6),(104,1,9),(105,1,10),(106,1,7),(107,1,8),(108,1,11),(109,1,12),(110,1,13),(111,1,14),(112,1,15),(113,1,16),(114,1,17),(115,1,18),(116,1,19),(117,1,20),(118,1,21),(119,1,22),(120,1,23),(121,1,24);

/*Table structure for table `sys_role_sys_users` */

DROP TABLE IF EXISTS `sys_role_sys_users`;

CREATE TABLE `sys_role_sys_users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `sys_role_id` int(11) NOT NULL,
  `sys_user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;

/*Data for the table `sys_role_sys_users` */

insert  into `sys_role_sys_users`(`id`,`sys_role_id`,`sys_user_id`) values (17,1,1),(18,1,12),(19,2,8),(20,2,1);

/*Table structure for table `sys_salary_slip` */

DROP TABLE IF EXISTS `sys_salary_slip`;

CREATE TABLE `sys_salary_slip` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `card_id` varchar(64) NOT NULL DEFAULT '' COMMENT '员工工号',
  `base_pay` decimal(12,2) NOT NULL DEFAULT '0.00' COMMENT '基本工资',
  `working_day` decimal(4,2) NOT NULL DEFAULT '0.00' COMMENT '工作天数',
  `days_off` decimal(4,2) NOT NULL DEFAULT '0.00' COMMENT '请假天数',
  `days_off_no` decimal(4,2) NOT NULL DEFAULT '0.00' COMMENT '调休天数',
  `reward` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '奖金',
  `rent_subsidy` decimal(6,2) NOT NULL DEFAULT '0.00' COMMENT '住房补贴',
  `trans_subsidy` decimal(6,2) NOT NULL DEFAULT '0.00' COMMENT '交通补贴',
  `social_security` decimal(6,2) NOT NULL DEFAULT '0.00' COMMENT '社保',
  `house_provident_fund` decimal(6,2) NOT NULL DEFAULT '0.00' COMMENT '住房公积金',
  `personal_income_tax` decimal(6,2) NOT NULL DEFAULT '0.00' COMMENT '个税',
  `fine` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '罚金',
  `net_salary` decimal(12,2) NOT NULL DEFAULT '0.00' COMMENT '实发工资',
  `pay_date` varchar(64) NOT NULL DEFAULT '' COMMENT '工资月份',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `CARDPAY` (`card_id`,`pay_date`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;

/*Data for the table `sys_salary_slip` */

insert  into `sys_salary_slip`(`id`,`card_id`,`base_pay`,`working_day`,`days_off`,`days_off_no`,`reward`,`rent_subsidy`,`trans_subsidy`,`social_security`,`house_provident_fund`,`personal_income_tax`,`fine`,`net_salary`,`pay_date`,`create_time`) values (1,'zl-001','5000.00','22.00','0.00','0.00','100.00','200.00','100.00','150.00','147.00','212.00','0.00','6033.00','2021-01','2021-04-21 23:34:09'),(2,'zl-001','5000.00','21.00','1.00','0.00','50.00','200.00','100.00','150.00','147.00','212.00','0.00','6033.00','2021-03','2021-04-21 23:34:44'),(4,'zl-001','5000.00','22.00','0.00','0.00','10.00','0.00','100.00','150.00','147.00','212.00','0.00','6033.00','2021-02','2021-04-21 23:34:44'),(5,'zl-001','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:11'),(6,'zl-002','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(7,'zl-003','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(8,'zl-004','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(9,'zl-005','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(10,'zl-006','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(11,'zl-007','5000.00','22.00','0.00','0.00','800.00','200.00','100.00','152.00','147.00','212.00','0.00','6633.00','2021-04','2021-04-22 00:01:12'),(12,'zl-008','5000.00','17.00','5.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(13,'zl-009','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(14,'zl-010','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(15,'zl-011','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(16,'zl-012','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(17,'zl-013','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(18,'zl-014','12000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','342.00','0.00','13163.00','2021-04','2021-04-22 00:01:12'),(19,'zl-015','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(20,'zl-016','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(21,'zl-017','5000.00','22.00','0.00','0.00','500.00','200.00','100.00','152.00','147.00','212.00','50.00','6383.00','2021-04','2021-04-22 00:01:12'),(22,'zl-018','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(23,'zl-019','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(24,'zl-020','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(25,'zl-021','5000.00','22.00','0.00','0.00','200.00','200.00','100.00','152.00','147.00','212.00','0.00','6033.00','2021-04','2021-04-22 00:01:12'),(26,'zl-022','5000.00','22.00','0.00','0.00','300.00','200.00','100.00','152.00','147.00','212.00','0.00','6133.00','2021-04','2021-04-22 00:01:12'),(27,'zl-023','8000.00','19.00','2.00','1.00','200.00','200.00','100.00','152.00','147.00','289.00','0.00','9110.00','2021-04','2021-04-22 00:01:12');

/*Table structure for table `sys_user` */

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `card_id` varchar(64) DEFAULT NULL COMMENT '员工工号',
  `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `age` int(11) DEFAULT NULL COMMENT '年龄',
  `gender` int(11) DEFAULT NULL COMMENT '性别',
  `phone` bigint(20) DEFAULT NULL COMMENT '手机号',
  `addr` varchar(255) DEFAULT NULL COMMENT '地址',
  `is_active` int(11) DEFAULT '1' COMMENT '1启用，0停用',
  `is_delete` int(11) DEFAULT '0' COMMENT '1删除，0未删除',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

/*Data for the table `sys_user` */

insert  into `sys_user`(`id`,`card_id`,`user_name`,`password`,`age`,`gender`,`phone`,`addr`,`is_active`,`is_delete`,`create_time`) values (1,'zl-001','admin0','e10adc3949ba59abbe56e057f20f883e',24,1,18738519352,'北京顺义区',1,0,'2021-03-22 00:50:54'),(2,'zl-002','admin1','25d55ad283aa400af464c76d713c07ad',25,1,18738519352,'北京昌平区',0,0,'2021-03-23 00:50:54'),(3,'zl-003','admin2','25d55ad283aa400af464c76d713c07ad',25,2,18738519352,'北京昌平区',0,0,'2021-03-23 00:50:54'),(4,'zl-004','admin3','25d55ad283aa400af464c76d713c07ad',24,1,18738519352,'北京顺义区',0,0,'2021-03-22 00:50:54'),(5,'zl-005','admin4','25d55ad283aa400af464c76d713c07ad',25,2,18738519352,'北京昌平区',0,0,'2021-03-23 00:50:54'),(6,'zl-006','admin5','25d55ad283aa400af464c76d713c07ad',25,2,18738519352,'北京昌平区',0,0,'2021-03-23 00:50:54'),(7,'zl-007','admin6','25d55ad283aa400af464c76d713c07ad',25,2,18738519352,'北京昌平区',0,0,'2021-03-23 00:50:54'),(8,'zl-008','admin7','25d55ad283aa400af464c76d713c07ad',18,2,13585648856,'北京昌平区',1,0,'2021-03-27 18:27:46'),(9,'zl-009','admin8','25d55ad283aa400af464c76d713c07ad',25,1,13845996652,'北京门头沟',1,0,'2021-03-27 18:31:34'),(10,'zl-010','admin9','25d55ad283aa400af464c76d713c07ad',26,1,13569852215,'北京房山区',1,0,'2021-03-27 18:33:08'),(11,'zl-011','admin10','25d55ad283aa400af464c76d713c07ad',26,1,13578546632,'北京顺义区',1,0,'2021-03-27 18:40:05'),(12,'zl-012','admin','25d55ad283aa400af464c76d713c07ad',26,1,13596586623,'XXX',1,0,'2021-03-27 18:43:34'),(13,'zl-013','admin12','25f9e794323b453885f5181f1b624d0b',26,1,13595688852,'XXX',1,0,'2021-03-27 18:50:58'),(15,'zl-015','admin13','25d55ad283aa400af464c76d713c07ad',25,1,13568529962,'顺义区',1,0,'2021-04-07 01:13:29');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
