package main

import (
	"log"
	"net/http"
	"github.com/thesouldev/goboxd/internal/api"
	
)

func main() {
	// http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request){
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write([]byte(`{"status":"ok"}`))
	http.HandleFunc("/healthz", api.Healthz)
	http.HandleFunc("/run", api.Run)
	

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}