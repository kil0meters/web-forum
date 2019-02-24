package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initializeDB() {
	fileName := "database.db"

	log.Printf("loading database file %s", fileName)

	_db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal("encountered an error opening database")
	}
	db = _db
}
