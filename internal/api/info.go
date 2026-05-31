package api

import (
	"encoding/json"
	"net/http"

	"github.com/thesouldev/goboxd/internal/config"
)

func Info(w http.ResponseWriter, r *http.Request) {

	languages := []string{}

	for id := range config.Registry {
		languages = append(
			languages,
			id,
		)
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]any{
			"languages": languages,
		},
	)
}
