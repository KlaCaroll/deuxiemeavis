package main

import (
	"log"
	"net/http"
	"strings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type StringSlice []string

func (ss *StringSlice) Scan(src any) error {
	var chunks []string
	switch data := src.(type) {
		case string:
				chunks = strings.Split(data, ",")
		case []byte:
				chunks = strings.Split(string(data), ",")
		case nil:
			return nil
		default:
				return fmt.Errorf(`invalid scan pair %T => %T`, src, ss)
	}
	*ss = chunks
	return nil
}

func (s Service) v2Doctors(w http.ResponseWriter, r *http.Request) {
	var items []struct{
		ID int64 `db:"id" json:"id"`
		LastName string `db:"last_name" json:"last_name"`
		FirstName string `db:"first_name" json:"first_name"`
		Diseases StringSlice `json:"diseases"`
		Hospital string `json:"hospital"`
	}

	err := s.DB.Select(&items, `
		SELECT d.id, d.last_name, d.first_name, 
			CONCAT(h.name, ' (', h.city, ')') AS hospital, 
			GROUP_CONCAT(dis.name) AS diseases
		FROM doctors d
		JOIN hospitals h ON h.id = d.hospital_id
		LEFT JOIN doctors_diseases ddis ON ddis.doctor_id = d.id
		LEFT JOIN diseases dis ON dis.id = ddis.disease_id
		GROUP BY d.id
		ORDER BY d.last_name;
	`)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, items)
}
