CREATE TABLE IF NOT EXISTS `visits_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `access_time` int(11) DEFAULT NULL,
  `closed_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
