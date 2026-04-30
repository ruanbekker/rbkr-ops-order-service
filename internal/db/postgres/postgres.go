package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func New(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	return db
}
