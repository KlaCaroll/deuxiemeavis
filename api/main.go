package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
	_ "github.com/go-sql-driver/mysql"
)

type any = interface{}

type Service struct {
	DB *sqlx.DB
}

func main() {
	var (
		addr string
		dsn  string // Data Source Name
	)
	flag.StringVar(&addr, "addr", "0.0.0.0:8080", "addr to listen on")
	flag.StringVar(&dsn, "dsn", "deuxiemeavis.mariadb", "path to the database to use")
	flag.Parse()

	// INITIALIZE THE DATABASE CONNEXION

	log.Println("opening connection to", dsn)
	db, err := sqlx.Connect("mysql", "deuxiemeavis:L3HmVZ72vMlk@tcp(localhost:3306)/deuxiemeavis")
	if err != nil {
		log.Println("opening connection", err)
		return
	}
	log.Println("opened connection")
	defer db.Close()

	s := Service {
		DB: db,
	}

	var mux = http.NewServeMux()

	mux.HandleFunc("/v2/doctors", s.v2Doctors)
	mux.HandleFunc("/v2/requestslist", s.RequestListStatus)

	// Start the HTTP server.
	handler := cors.Default().Handler(mux)
	// CORS => handler
	var srv = &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	log.Println("listen on addr", addr)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Println("listening", err)
		return
	}
}

func write(w http.ResponseWriter, payload any) {
	w.Header().Set("Content-Type", "application/json")
	raw, _ := json.Marshal(payload)
	w.Write(raw)
}

type apiError struct{
	Code string `json:"code"`
	Err string `json:"err"`
}

func writeError(w http.ResponseWriter, code string, err error) {
	write(w, apiError{
		Code: code,
		Err: err.Error(),
	})
}

func read(r *http.Request, payload any) (err error) {
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, payload)
	if err != nil {
		return err
	}
	return nil		
}