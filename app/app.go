package app

import (
	"log"
	"net/http"
)

func Start() {
	//define routes

	mux := http.NewServeMux()

	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)
	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
