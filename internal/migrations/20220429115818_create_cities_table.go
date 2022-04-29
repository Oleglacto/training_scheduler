package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateCitiesTable, downCreateCitiesTable)
}

func upCreateCitiesTable(tx *sql.Tx) error {
	query := "CREATE TABLE cities (" +
		"name            varchar(80)," +
		"location        point);"
	return nil
}

func downCreateCitiesTable(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
