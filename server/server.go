package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Log struct {
	LogType string `json:"logType"`
	Info    string `json:"info"`
}

func handleLog(w http.ResponseWriter, r *http.Request) {

	var log Log
	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Println(log)
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/log", handleLog).Methods("POST")
	fmt.Println("Server running on port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Error starting server", err)
	}

}
