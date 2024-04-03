package dbconfig

import (
	"database/sql"
	"log"
	"github.com/Sangram-Barge/Problems-Repository/persistance"
)

func FindAll(db *sql.DB) ([]persistance.Problem, error) {
	problems, err := db.Query("select * from problems")
	if err != nil {
		return nil, err
	}
	return parseProblems(problems)
}

func Insert(problem persistance.Problem, db *sql.DB) (persistance.Problem, error) {
	statementInsert := `insert into problems 
	(problem, platform, description, intiution, link) values
	(?,?,?,?,?)`
	prep, err := db.Prepare(statementInsert)
	if err != nil {
		return persistance.NewProblem(), err
	}
	prep.Exec(problem.Problem, problem.Platform, problem.Description,
		problem.Intiution, problem.Link)
	log.Printf("inserting %v\n", problem)
	return problem, nil
}

func Delete(id string, db *sql.DB) (error) {
	prep, err := db.Prepare("delete from problems where id = ?")
	if err != nil {
		return err
	}
	prep.Exec(id)
	return nil
}

func parseProblems(problems *sql.Rows) ([]persistance.Problem, error) {
	result := make([]persistance.Problem, 0)
	for problems.Next() {
		problem := persistance.NewProblem()
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
