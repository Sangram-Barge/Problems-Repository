package dbconfig

import (
	"database/sql"
	"fmt"
)

type Problem struct {
	Id          int
	Problem     string
	Platform    string
	Description string
	Intiution   string
	Link        string
}

func NewProblem() Problem {
	return Problem{}
}

func Init() (db *sql.DB, err error) {
	dbName := "dsa"
	dbUser := "dsa"
	dbPass := "dsa"
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName))
	return db, err
}
