package model

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB = nil
)

func OpenDb() {
	// TODO: Use global variable here.
	db, err := sql.Open("sqlite3", "./go_rest.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := "create table foo (id integer not null primary key, name text);"
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
}

func GetAll() {
	if db != nil {
		rows, err := db.Query("select id, name from foo")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		// log.Printf("%v", rows)
	} else {
		log.Println("Failed to get all data, db is nil")
	}
}

func CloseDb() {
	if db != nil {
		log.Println("Close db")
		db.Close()
	} else {
		log.Println("Failed to close data, db is nil")
	}
}
