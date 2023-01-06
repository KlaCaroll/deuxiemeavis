package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func (s Service) RequestListStatus(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Status string 
	}

	err := read(r, &input)
	if err!= nil {
		log.Println("parsing input", err)
		writeError(w, "input_error", err)
		return
	}

	var request []struct{
		PatientID int64 `db:"patient_id" json:"patient_id"`
		Patient string `json:"patient"`
		Doctor string `json:"doctor"`
		Disease string `json:"disease"`
		Hospital string `json:"hospital"`
	}

	err = s.DB.Select(&request, `
		SELECT r.patient_id, CONCAT(p.first_name,' ', p.last_name) AS patient, 
			CONCAT(d.first_name, ' ', d.last_name) AS doctor, dis.name AS disease, 
			CONCAT(h.name, ' (', h.city, ')') AS hospital
		FROM requests r
		JOIN patients p ON p.id = r.patient_id
		JOIN doctors d ON d.id = r.doctor_id
		JOIN diseases dis ON dis.id = r.disease_id
		JOIN hospitals h ON h.id = d.hospital_id
		WHERE status = ?
	`, input.Status)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, request)
}

