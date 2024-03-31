package main

import (
	"fmt"

	dbconfig "github.com/Sangram-Barge/Problems-Repository/dbconfig"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := dbconfig.Init()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	results, err := dbconfig.FindAll(db)
	if err != nil {
		panic(err.Error())
	}
	for i, result := range results {
		fmt.Println(i, " ", result)
	}

	p := dbconfig.Problem {
		Problem: "water storage",
		Platform: "leetcode, scaler",
		Description: "store water on buildings",
		Intiution: "lmax, rmax arrays",
		Link: "leetcode.com",
	}
	dbconfig.Insert(p, db)
	results, err = dbconfig.FindAll(db)
	if err != nil {
		panic(err.Error())
	}
	for i, result := range results {
		fmt.Println(i, " ", result)
	}
}
