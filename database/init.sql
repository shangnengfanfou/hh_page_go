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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

INSERT INTO hh_page.article
(id, tag_id, unique_id, title, summary, banner_url, `time`, views_count)
VALUES(1, 1, '7752743d31fd69e3', '测试文档11', '随着node版本的不断更新，node的最大堆内存也不断的更新，但是网上还是充斥着大量32位机node只能用0.7GB的内存，64位机只能用1.4G内存的资料', 'http://10.226.11.52:8090/views/img/1658236642_30a1fcd295679369.png', 1658236646, 8);
INSERT INTO hh_page.article
(id, tag_id, unique_id, title, summary, banner_url, `time`, views_count)
VALUES(2, 1, '22dc97b03d156e29', 'nodejs事件驱动', '对比运行浏览器上的js，nodejs除了v8引擎之外，V8引擎负责js文件解释，运行，内存，垃圾回收等；还存在一个一个开源库libuv，libuv是一个强大的异步I/O库，提供网络编程（跨平台），文件读写等能力，这也是nodejs能进行后端开发的原因', 'http://10.226.11.52:8090/views/img/1658298479_14a9ec6678ae82b7.png', 1658298486, 12);
INSERT INTO hh_page.article
(id, tag_id, unique_id, title, summary, banner_url, `time`, views_count)
VALUES(3, 1, '9565f395af34ecf5', 'pc应用录屏并推流到rtmp流媒体服务器', '由FFMEPG实现截频录屏、推流的整个行为。通过c#的process类调用ffmpeg进程，使用相关命令行参数实现推流。主要讲解ffmepg的能力在untiy中的使用，以及一些参数的使用和说明', 'http://10.226.11.52:8090/views/img/1658302524_2a55b6be29d7793e.png', 1658302524, 22);

