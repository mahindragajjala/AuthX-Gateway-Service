package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectPostgres() {
	connStr := "user=postgres password=1234 dbname=postgres sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB Ping failed:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")
}
