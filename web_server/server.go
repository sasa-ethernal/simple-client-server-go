package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/send", sendHandler).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./react-app/build")))

	http.Handle("/", r)
	http.ListenAndServe(":19999", nil)
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	data := "Success."

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
