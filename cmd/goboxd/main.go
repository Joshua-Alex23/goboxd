package main

import (
	"log"
	"net/http"

	"github.com/thesouldev/goboxd/internal/api"
	"github.com/thesouldev/goboxd/internal/config"
)

func main() {
	// http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request){
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write([]byte(`{"status":"ok"}`))
	http.HandleFunc("/healthz", api.Healthz)
	http.HandleFunc("/run", api.Run)

	cfg, err := config.Load("/configs/languages.yaml")
	if err != nil {
		log.Fatal(err)
	}

	for _, lang := range cfg.Languages {
		config.Registry[lang.ID] = lang
	}

	log.Printf("loaded %d languages", len(config.Registry))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("loaded languages: %v", langs)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
