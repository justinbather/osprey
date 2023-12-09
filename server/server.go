package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Log struct {
	Account string `json:"account"`
	LogType string `json:"logType"`
	Info    string `json:"info"`
}

var logs = []Log{}

func handleLog(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	apiKey := queryParams.Get("api_key")
	var log Log
	err := json.NewDecoder(r.Body).Decode(&log)
	log.Account = apiKey
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logs = append(logs, log)
	w.WriteHeader(http.StatusCreated)
	fmt.Println(logs)
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/log", handleLog).Methods("POST")
	fmt.Println("Server running on port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Error starting server", err)
	}

}
