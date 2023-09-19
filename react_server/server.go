package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/header", getHeader).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./msg-app/build")))

	http.Handle("/", r)
	http.ListenAndServe(":19999", nil)
}

func getHeader(w http.ResponseWriter, r *http.Request) {
	header := "Simple Message App"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(header)
}
