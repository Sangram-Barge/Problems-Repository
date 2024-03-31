package dbconfig

import (
	"database/sql"
	"fmt"
)


func Init() (db *sql.DB, err error) {
	dbName := "dsa"
	dbUser := "dsa"
	dbPass := "dsa"
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName))
	return db, err
}
