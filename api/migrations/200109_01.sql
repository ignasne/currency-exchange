-- +migrate Up

ALTER TABLE `currency_cache` CHANGE `v` `v` VARBINARY(32)  NULL  DEFAULT NULL;
