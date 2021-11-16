package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./devices"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func main() {
	// Setup DB Connection
	connectionString := "host=0.0.0.0 port=5432 user=user password=password dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS devices (name VARCHAR);`)
	if err != nil {
		panic(err)
	}

	// Setup API Routes
	my_devices := devices.New()
	log.Println("Starting up on http://localhost:8080")

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World!"))
	})

	r.Get("/devices", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(my_devices.GetAll())
	})

	r.Post("/devices", func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		my_devices.Add(devices.Device{
			Name: request["Name"],
			Type: request["Type"],
			Ip:   request["Ip"],
		})

	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func write(db sql.DB) {
	insertStatement := `
	INSERT INTO Devices (Name)
	VALUES ('hehehhe');
	`

	_, err := db.Exec(insertStatement)
	if err != nil {
		panic(err)
	}

	fmt.Println("Added Device to Table")
}

func read(db sql.DB) {
	selectQuery := "SELECT Name FROM Devices;"

	rows, err := db.Query(selectQuery)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var Name string

		err = rows.Scan(&Name)
		if err != nil {
			panic(err)
		}
		fmt.Println(Name)
	}
}
