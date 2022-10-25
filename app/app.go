package app

import (
	"fmt"
	//gorilla mux is a router that simplifies route definitions
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	////multiplexer = routing
	//mux := http.NewServeMux()
	//mux var name collides with package name, using router now.
	router := mux.NewRouter()

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	//regex [0-9]+ all numerical values matched, strings return 404
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	//Vars returns the route variables for the current request, if any.
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
