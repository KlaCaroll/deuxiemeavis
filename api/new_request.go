package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func (s Service) NewRequest(w http.ResponseWriter, r *http.Request) {
	var input struct {
		PatientID int64 `db:"patient_id" json:"patient_id"`
		DoctorID int64 `db:"doctor_id" json:"doctor_id"`
		DiseaseID int64 `db:"disease_id" json:"disease_id"`
		Diagnosis string `json:"Lorem ipsum"`
	}

	err := read(r, &input)
	if err!= nil {
		log.Println("parsing input", err)
		writeError(w,"input_error", err)
		return
	}

	res, err := s.DB.Exec(`
		INSERT INTO requests (patient_id, doctor_id, disease_id, status, diagnosis)
		VALUES (?, ? ,? ,'approval-awaiting', ?)
	`, input.PatientID, input.DoctorID, input.DiseaseID, input.Diagnosis)

	var request struct{
		ID int64 `db:"id" json:"id"`
	}

	request.ID, err = res.LastInsertId()
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, request)
}

