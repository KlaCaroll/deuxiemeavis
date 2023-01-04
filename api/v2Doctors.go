package main

import (
	"log"
	"net/http"

	//"github.com/jmoiron/sqlx"
	//"github.com/rs/cors"
	_ "github.com/go-sql-driver/mysql"
)

func (s Service) v2Doctors(w http.ResponseWriter, r *http.Request) {
	var items []struct{
		ID int64 `db:"id" json:"id"`
		Name string `db:"last_name" json:"last_name"`
		FirstName string `db:"first_name" json:"first_name"`
		Diseases string `db:"name" json:"diseases"`
		Hospital_name string `db:"name" json:"hospital_name"` 
		Hospital_city string `db:"city" json:"hospital_city"`
	}

	err := s.DB.Select(&items, `
		SELECT d.id, d.last_name, d.first_name, concat(h.name, ' (', h.city, ')')
		FROM doctors d
		JOIN hospitals h ON h.id = d.hospital_id
		ORDER BY d.last_name;
	`)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, items)
}

//JOIN doctors_diseases ddis ON ddis.doctor_id = d.id 
//JOIN diseases dis ON dis.id = ddis.disease_id