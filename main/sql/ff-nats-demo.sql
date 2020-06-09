 CREATE TABLE `ff-nats-demo` (
          `id` int(11) NOT NULL COMMENT '自己生成随机id',
          `this_time` int(20) DEFAULT NULL COMMENT '当前时间',
          `device_id` varchar(20) DEFAULT NULL COMMENT '设备Id号',
          `tag` varchar(50) DEFAULT NULL COMMENT '标签信息',
          `data_id` int(20) DEFAULT NULL COMMENT 'dataId',
          `tag_id` int(20) DEFAULT NULL COMMENT 'tagId',
          `value` int(20) DEFAULT NULL COMMENT '该条数据上传的值',
          PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;