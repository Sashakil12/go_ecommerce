package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, dir string) error {
	// Create a new migrate instance
	migrations := &migrate.FileMigrationSource{
		Dir: dir,
	}
	_, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Printf("Error migrating database: %v\n", err)
		return err
	}
	fmt.Println("Database migrated successfully")
	return nil

}
