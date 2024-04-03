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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func InsertProblem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Default().Println("insert problem")
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept")
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
}

func DeleteProblem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, PUT, POST, DELETE")
	if r.Method == "DELETE" {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "no id passed", 400)
			return
		}
		log.Default().Println("delete for id ", id)
		db, err := dbconfig.Init()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		dbconfig.Delete(id, db);
	}
}

func errorReturn(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}

func Init() {
	http.HandleFunc("/problem", DeleteProblem)
	http.HandleFunc("/problems", GetAllProblems)
	http.HandleFunc("/problems/add", InsertProblem)
}
