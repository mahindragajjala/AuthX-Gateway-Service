package createdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	postgresUser     = "mahindra"
	postgresPassword = "1234"
	postgresHost     = "localhost"
	postgresPort     = 5432
)

func CreateDatabase_Manual() {
	// 1. Connect to default "postgres" DB to create "auth_db"
	defaultConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword)
	defaultDB, err := sql.Open("postgres", defaultConnStr)
	if err != nil {
		log.Fatal("Failed to connect to default DB:", err)
	}
	defer defaultDB.Close()

	// 2. Create "auth_db" if not exists
	_, err = defaultDB.Exec("CREATE DATABASE auth_db")
	if err != nil {
		// Ignore error if DB already exists
		fmt.Println("Skipping DB creation (may already exist):", err)
	} else {
		fmt.Println("Database 'auth_db' created.")
	}

	// 3. Connect to the "auth_db" DB
	authConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=auth_db sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword)
	authDB, err := sql.Open("postgres", authConnStr)
	if err != nil {
		log.Fatal("Failed to connect to auth_db:", err)
	}
	defer authDB.Close()

	// 4. Create "users" table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = authDB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	fmt.Println("Users table created successfully in 'auth_db'.")
}
