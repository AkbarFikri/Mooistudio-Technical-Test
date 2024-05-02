package database

import (
	"github.com/jmoiron/sqlx"
	"log"
)

const userTable = `
CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(255) NOT NULL PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255)
);`

func MigrateDB(db *sqlx.DB) error {
	tables := []string{userTable}
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			log.Printf("Error creating table: %v\n", err)
			return err
		}
	}
	return nil
}
