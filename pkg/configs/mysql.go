package configs

import (
	"database/sql"
	"log"
)

// Connect - db connection
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Asulahlo31>@tcp(localhost:3306)/book_test")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
