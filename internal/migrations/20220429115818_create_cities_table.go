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
		"uuid 				varchar(36) PRIMARY KEY," +
		"name            	varchar(255)," +
		"region			 	varchar(255)," +
		"federal_district 	varchar(255)," +
		"latitude        	float," +
		"longitude			float);"

	_, err := tx.Exec(query)
	return err
}

func downCreateCitiesTable(tx *sql.Tx) error {
	query := "DROP TABLE cities;"

	_, err := tx.Exec(query)
	return err
}
