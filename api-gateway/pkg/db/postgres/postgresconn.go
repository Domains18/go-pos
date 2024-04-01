package postgres

import (
	"fmt"
	"github.com/Nerds-Catapult/notiwa/api-gateway/pkg/db"
	"log"
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

	if err = db.CreateTables(); err != nil {
		log.Fatal("errr creating table schema v", err)
	}

}

func Disconnect() error {
	return DB.Close()
}
