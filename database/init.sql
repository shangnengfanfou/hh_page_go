CREATE TABLE IF NOT EXISTS `visits_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `access_time` int(11) DEFAULT NULL,
  `closed_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tag_id` int(11) DEFAULT NULL,
  `unique_id` varchar(32) DEFAULT NULL,
  `title` varchar(50) DEFAULT NULL,
  `summary` varchar(255) DEFAULT NULL,
  `banner_url` varchar(255) DEFAULT NULL,
  `time` int(11) DEFAULT NULL,
  `views_count` int(11) DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uk_unique_id`(`unique_id`) USING BTREE,
  INDEX `idx_tag_id`(`tag_id`) USING BTREE,
  INDEX `idx_views_count`(`views_count`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
