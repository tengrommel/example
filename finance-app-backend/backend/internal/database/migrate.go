package database

import (
	"awesomeProject/finance-app-backend/backend/internal/config"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func migrateDb(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "connecting to database")
	}

	migrationSource := fmt.Sprintf("file//%sinternal/database/migrations/", *config.DataDirectory)
	migrator, err := migrate.NewWithDatabaseInstance(migrationSource, "postgress", driver)
	if err != nil {
		return errors.Wrap(err, "creating migrator")
	}
	if err := migrator.Up(); err != nil {
		return errors.Wrap(err, "executing migration")
	}
	version, dirty, err := migrator.Version()
	if err != nil {
		return errors.Wrap(err, "getting migration")
	}
	logrus.WithFields(logrus.Fields{
		"version": version,
		"dirty":   dirty,
	}).Debug("Database migrated")
	return nil
}
