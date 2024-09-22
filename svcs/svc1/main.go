package main

import (
	"net/http"

	"github.com/pavr1/people/go/pkg/mod/gopkg.in/!data!dog/dd-trace-go.v1@v1.64.0/contrib/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":80", router)
}

func endpoint1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")