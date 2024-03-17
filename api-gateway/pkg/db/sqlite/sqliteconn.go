package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
}

func (sqlite SQLiteDB) GetConnection() (*sql.DB, error) {
	dbPath := "data.db"

	// Open the database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("error opening database: ", err)
		return nil, err
	}

	defer db.Close()

	// Ping the database to ensure it's accessible
	if err := db.Ping(); err != nil {
		log.Fatal("error connecting to database: ", err)
		return nil, err
	}

	log.Println("Connected to SQLite database successfully")
	return db, nil
}

