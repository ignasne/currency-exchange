-- +migrate Up

CREATE TABLE IF NOT EXISTS `currency_cache`
(
    `k`   varbinary(6) NOT NULL DEFAULT '',
    `v`   int(10)               DEFAULT NULL,
    `ttl` timestamp    NULL     DEFAULT NULL,
    PRIMARY KEY (`k`),
    KEY `ttl` (`ttl`)
) ENGINE = InnoDB
  DEFAULT CHARSET = latin1;