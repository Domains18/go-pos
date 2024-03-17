package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	var err error

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbConnectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err = sqlx.Connect("postgres", dbConnectionString)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to the database: %v", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(4)

	CreateTables()
}

func CreateTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		email VARCHAR(255) UNIQUE NOT NULL,
		user_id UUID PRIMARY KEY
	);
	`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create the users table: %v", err))
	}
}
