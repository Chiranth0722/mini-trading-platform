package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error

	connStr := "host=localhost port=5432 user=postgres password=123456 dbname=trading sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}

	log.Println("Connected to PostgreSQL!")
}
