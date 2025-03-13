package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./projects.db")
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	return db, nil
}

func CreateTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS projects (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"category" TEXT,
		"url" TEXT,
		"technologies" TEXT
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
