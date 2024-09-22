package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/svc1/endpoint1", endpoint1).Methods("GET")
	router.HandleFunc("/svc1/endpoint2", endpoint2).Methods("GET")
	router.HandleFunc("/svc1/endpoint3", endpoint3).Methods("GET")

	err := http.ListenAndServe(":9000", router)
	if err != nil {
		fmt.Println(err)
	}
}

func endpoint1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "svc1.endpoint1")
}

func endpoint2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "svc1.endpoint2")
}

func endpoint3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "svc1.endpoint3")
}
