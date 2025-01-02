package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:12345678@localhost/openapi-go-crud?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}
	log.Println("Database connection established")
}
