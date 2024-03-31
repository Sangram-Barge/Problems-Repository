package dbconfig

import (
	"database/sql"
	"log"
)

func FindAll(db *sql.DB) ([]Problem, error) {
	problems, err := db.Query("select * from problems")
	if err != nil {
		return nil, err
	}
	result, err := parseProblems(problems)
	return result, err
}

func Insert(problem Problem, db *sql.DB) (Problem, error) {
	statementInsert := `insert into problems 
	(problem, platform, description, intiution, link) values
	(?,?,?,?,?)`
	prep, err := db.Prepare(statementInsert)
	if err != nil {
		return NewProblem(), err
	}
	prep.Exec(problem.Problem, problem.Platform, problem.Description,
		problem.Intiution, problem.Link)
	log.Printf("inserting %v\n", problem)
	return problem, nil
}

func parseProblems(problems *sql.Rows) ([]Problem, error) {
	result := make([]Problem, 0)
	for problems.Next() {
		problem := NewProblem()
		err := problems.Scan(
			&problem.Id,
			&problem.Problem,
			&problem.Platform,
			&problem.Description,
			&problem.Intiution,
			&problem.Link)
		if err != nil {
			return nil, err
		}
		result = append(result, problem)
	}
	return result, nil
}
