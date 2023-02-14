package db

import (
	"database/sql"
)

// Connect to the database.
func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "bertneng11"
	dbName := "portfolio"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db
}
