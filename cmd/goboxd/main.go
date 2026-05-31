package main

import (
	"log"
	"net/http"

	"github.com/thesouldev/goboxd/internal/api"
	"github.com/thesouldev/goboxd/internal/config"
)

func main() {
	http.HandleFunc("/healthz", api.Healthz)
	http.HandleFunc("/run", api.Run)
	http.HandleFunc("/readyz", api.Readyz)
	http.HandleFunc("/info", api.Info)

	cfg, err := config.Load("/configs/languages.yaml")
	if err != nil {
		log.Fatal(err)
	}

	for _, lang := range cfg.Languages {
		config.Registry[lang.ID] = lang
	}

	log.Printf("loaded %d languages", len(config.Registry))

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
