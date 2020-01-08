package datasource

import (
	"database/sql"
	"fmt"
	"github.com/ignasne/currency-exchange/api/config"
	"github.com/ignasne/currency-exchange/api/logger"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type ConnectionPool struct {
	*sql.DB
}

type Reader interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type Writer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type ReaderWriter interface {
	Reader
	Writer
	Begin() (*sql.Tx, error)
}

func Connect(cfg *config.DB) *ConnectionPool {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.User,
			cfg.Pass,
			cfg.Host,
			cfg.Port,
			cfg.Name,
		))

	if err != nil {
		logger.Get().WithError(err).Fatal("Could not open database")
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		logger.Get().WithError(err).Fatal("Could not ping database")
		os.Exit(1)
	}

	return &ConnectionPool{db}
}
