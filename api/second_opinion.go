package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func (s Service) SecondOpinion(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID int64 `db:"request_id" json:"request_id"` 
		SecondOpinion string `db:"second_opinion" json:"second_opinion"`
	}

	err := read(r, &input)
	if err!= nil {
		log.Println("parsing input", err)
		writeError(w, "input_error", err)
		return
	}

	res, err := s.DB.Exec(`
	UPDATE requests
	SET second_opinion = ?, status = 'done', second_opinion_date = current_timestamp
	WHERE id = ?;	
	`, input.SecondOpinion, input.ID,)

	var output struct {
		Status string `db:"status" json:"result"`
	}
	
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	log.Println(res)
	write(w, output)
}

