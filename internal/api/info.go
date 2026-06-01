package api

import (
	"encoding/json"
	"net/http"

	"github.com/thesouldev/goboxd/internal/config"
	"github.com/thesouldev/goboxd/internal/queue"
)

func Info(w http.ResponseWriter, r *http.Request) {

	languages := []string{}

	for id := range config.Registry {
		languages = append(languages, id)
	}

	resp := map[string]any{
		"languages": languages,
		"workers":   cap(queue.Workers),
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
