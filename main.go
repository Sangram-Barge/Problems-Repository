package main

import (
	"log"
	"net/http"

	"github.com/Sangram-Barge/Problems-Repository/handlers"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	handlers.Init()
	log.Fatal(http.ListenAndServe(":9090", nil))
}


