package database

import (
	"database/sql"
	"log"
)

func DbConn() (db *sql.DB) {
	connStr := "user=suwayouta dbname=todo sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
