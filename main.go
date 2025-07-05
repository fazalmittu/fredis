package main

import (
	"fredis/service"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// CORE
	mux.HandleFunc("GET /cache/{key}", service.GetItem)
	mux.HandleFunc("PUT /cache/{key}", service.SetItem)
	mux.HandleFunc("DELETE /cache/{key}", service.DelItem)
	mux.HandleFunc("GET /cache/{key}/exists", service.Exists)

	port := ":8080"
	log.Printf("Listening on http://localhost%s\n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
