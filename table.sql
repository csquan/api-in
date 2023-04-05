DROP TABLE IF EXISTS `t_mechanism`;
CREATE TABLE `t_mechanism`
(
    `f_id`                      bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `f_name`                    char(66)  NOT NULL DEFAULT '' COMMENT '机构名称',
    `f_key`                     char(42)  NOT NULL DEFAULT '' COMMENT '机构key',
    `f_secret`                  char(42)  NOT NULL DEFAULT '' COMMENT '机构密钥',
    PRIMARY KEY (`f_id`) /*T![clustered_index] CLUSTERED */
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='机构数据表';


DROP TABLE IF EXISTS `t_transfer`;
CREATE TABLE `t_transfer`
(
    `f_id`                      bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `f_fromAccount`             text COMMENT 'fromAccount',
    `f_toAccount`               text COMMENT 'toAmount',
    `f_thirdId`                 text COMMENT 'thirdId',
    `f_token`                   text COMMENT 'token',
    `f_amount`                  text COMMENT 'amount',
    `f_callback`                text COMMENT 'callback',
    `f_ext`                     text COMMENT 'ext',
    `f_isSync`                  text COMMENT 'isSync',
    `f_isTransaction`           text COMMENT 'isTransaction',
    `f_state`                   int DEFAULT NULL,     # 0-正在进行 1-成功 -1-失败
    `f_error`                   text COMMENT 'error', # 对应失败的具体原因
    `f_created_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'time',
    `f_updated_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON update CURRENT_TIMESTAMP COMMENT 'time',
    PRIMARY KEY (`f_id`) /*T![clustered_index] CLUSTERED */
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='转账表';

DROP TABLE IF EXISTS `t_withdraw`;
CREATE TABLE `t_withdraw`
(
    `f_id`                      bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `f_Account`                 text COMMENT 'account',
    `f_thirdId`                 text COMMENT 'thirdId',
    `f_symbol`                  text COMMENT 'symbol',
    `f_amount`                  text COMMENT 'amount',
    `f_chain`                   text COMMENT 'chain',
    `f_addr`                    text COMMENT 'addr',
    `f_isSync`                  text COMMENT 'isSync',
    `f_state`                   int DEFAULT NULL,     # 0-正在进行 1-成功 -1-失败
    `f_error`                   text COMMENT 'error', # 对应失败的具体原因
    `f_created_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'time',
    `f_updated_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON update CURRENT_TIMESTAMP COMMENT 'time',
    PRIMARY KEY (`f_id`) /*T![clustered_index] CLUSTERED */
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='转账表';
