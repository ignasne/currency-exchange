package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ignasne/currency-exchange/api/config"
	"github.com/ignasne/currency-exchange/api/datasource"
	"github.com/ignasne/currency-exchange/api/logger"
	"github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations",
	Run: func(cmd *cobra.Command, args []string) {
		runMigrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func runMigrate() {
	logger.Get().Debug("running migrations")

	migrationCfg := config.GetMigration()
	connPool := datasource.Connect(config.GetDB())

	migrations := &migrate.FileMigrationSource{Dir: migrationCfg.Path}
	migrate.SetTable("migrations")

	num, err := migrate.Exec(connPool.DB, "mysql", migrations, migrate.Up)
	if err != nil {
		logger.Get().WithError(err).Warn("could not apply migrations")
		return
	}

	logger.Get().WithFields(logrus.Fields{"total_migrations": num}).Info("migrations applied")
}