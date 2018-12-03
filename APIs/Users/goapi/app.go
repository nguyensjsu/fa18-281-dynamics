package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func PingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Shayona Grocery Store API is alive!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
