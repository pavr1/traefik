package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/endpoint1", endpoint1).Methods("GET")
	router.HandleFunc("/endpoint2", endpoint2).Methods("GET")
	router.HandleFunc("/endpoint3", endpoint3).Methods("GET")

	err := http.ListenAndServe(":9002", router)
	if err != nil {
		fmt.Println(err)
	}
}

func endpoint1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server2.endpoint1")
}

func endpoint2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server2.endpoint2")
}

func endpoint3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server2.endpoint3")
}
