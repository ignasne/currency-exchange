-- +migrate Up

CREATE TABLE IF NOT EXISTS `currency_cache`
(
    `key`   varbinary(6) NOT NULL DEFAULT '',
    `value` int(10)               DEFAULT NULL,
    `ttl`   timestamp    NULL     DEFAULT NULL,
    PRIMARY KEY (`key`),
    KEY `ttl` (`ttl`)
) ENGINE = InnoDB
  DEFAULT CHARSET = latin1;