package datasource

import (
	"database/sql"
	"github.com/ignasne/currency-exchange/api/logger"
	"time"
)

type CacheDB struct {
	DB ReaderWriter
}

func (c *CacheDB) Get(key string) *string {
	c.clearExpired()

	var value *string

	queryString := `SELECT v FROM currency_cache WHERE k = ? AND (ttl IS NULL OR ttl > ?)`

	err := c.DB.QueryRow(queryString, key, Timestamp()).Scan(&value)

	if err != nil && err != sql.ErrNoRows {
		// silently log error and return empty value that cache could be ignore and real data can be used in api
		logger.Get().WithError(err).Error("failed to get currency cache item")
		return nil
	}

	return value
}

func (c *CacheDB) Set(key string, value string, ttl int) bool {
	err := c.clearExpired()

	if err != nil {
		logger.Get().WithError(err).Warn("failed to clear currency cache")
		return false
	}

	res, err := c.DB.Exec(`
		INSERT IGNORE INTO currency_cache (k, v, ttl)
		VALUES (?, ?, ?)`, key, value, TimeToString(c.getExpire(ttl)))

	if err != nil {
		logger.Get().WithError(err).Warn("failed to insert currency cache")
		return false
	}

	rowsInserted, err := res.RowsAffected()

	if err != nil {
		logger.Get().WithError(err).Warn("failed to get affected rows after currency cache item insert")
		return false
	}

	return rowsInserted == 1
}

func (c *CacheDB) getExpire(expireInSeconds int) time.Time {
	return time.Now().UTC().Add(time.Duration(expireInSeconds)*time.Second)
}

func (c *CacheDB) clearExpired() error {
	_, err := c.DB.Exec(`DELETE FROM currency_cache WHERE ttl < ?`, Timestamp())

	if err != nil {
		logger.Get().WithError(err).Warn("failed to remove expired currency cache")
		return err
	}

	return nil
}