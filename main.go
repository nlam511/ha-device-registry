package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./devices"

	"github.com/go-chi/chi"
)

func main() {
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
