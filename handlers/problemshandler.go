package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Sangram-Barge/Problems-Repository/dbconfig"
	"github.com/Sangram-Barge/Problems-Repository/persistance"
)


func GetAllProblems(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("find all problems")
	db, err := dbconfig.Init()
	if err != nil {
		errorReturn(w, err, 500)
		return
	}
	defer db.Close()
	problems, err := dbconfig.FindAll(db)
	if err != nil {
		errorReturn(w, err, 500)
		return
	}
	output, err := json.Marshal(problems)
	if err != nil {
		errorReturn(w, err, 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func InsertProblem(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("insert problem")
	if r.Method == "POST" {
		db, err := dbconfig.Init()
		if err != nil {
			errorReturn(w, err, 500)
			return
		}
		defer db.Close()
		b, err := io.ReadAll(r.Body)
		if err != nil {
			errorReturn(w, err, 500)
			return
		}
		var problem persistance.Problem
		if err = json.Unmarshal(b, &problem); err != nil {
			errorReturn(w, err, 400)
			return
		}
		p, err := dbconfig.Insert(problem, db)
		if err != nil {
			errorReturn(w, err, 500)
			return
		}

		output, err := json.Marshal(p)
		if err != nil {
			errorReturn(w, err, 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
	}
}

func errorReturn(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}

func Init() {
	http.HandleFunc("/problems", GetAllProblems)
	http.HandleFunc("/problems/add", InsertProblem)
}