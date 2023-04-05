DROP TABLE IF EXISTS `t_mechanism`;
CREATE TABLE `t_mechanism`
(
    `f_id`                      bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `f_name`                    char(66)  NOT NULL DEFAULT '' COMMENT '机构名称',
    `f_key`                     char(42)  NOT NULL DEFAULT '' COMMENT '机构key',
    `f_secret`                  char(42)  NOT NULL DEFAULT '' COMMENT '机构密钥',
    PRIMARY KEY (`f_id`) /*T![clustered_index] CLUSTERED */
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='机构数据表';


DROP TABLE IF EXISTS `t_action`;
CREATE TABLE `t_action`
(
    `f_id`                      bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `f_op`                      text COMMENT 'uid',
    `f_params`                  text COMMENT 'appid',
    `f_mechanism_name`          text COMMENT 'mechanism_name',
    `f_state`                   int DEFAULT NULL,     # 0-正在进行 1-成功 -1-失败
    `f_error`                   text COMMENT 'error', # 对应失败的具体原因
    `f_created_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'time',
    `f_updated_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON update CURRENT_TIMESTAMP COMMENT 'time',
    PRIMARY KEY (`f_id`) /*T![clustered_index] CLUSTERED */
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='动作执行表';

