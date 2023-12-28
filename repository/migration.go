package repository

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"os"

	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(postgresConUrl, migrationUrl string) (err error) {
	postgres := &pgx.Postgres{}
	driver, err := postgres.Open(postgresConUrl)

	if err != nil {
		return fmt.Errorf("unable to open connection: %s", err)
	}

	defer func() {
		if closeErr := driver.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Unable to close connection: %v\n", closeErr)
		}
	}()

	migration, err := migrate.NewWithDatabaseInstance(migrationUrl, "pgx", driver)

	if err != nil {
		return fmt.Errorf("unable to open migration: %s", err)
	}

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("unable to run migration: %s", err)
	}

	return nil
}
