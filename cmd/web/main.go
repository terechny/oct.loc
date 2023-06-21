package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "status Sreated"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Server Start on http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Panicln(err)
}
