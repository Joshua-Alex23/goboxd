package api


import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/thesouldev/goboxd/internal/models"
	"github.com/thesouldev/goboxd/internal/validate"
	
)

func Run(w http.ResponseWriter, r *http.Request) {
	var req models.RunRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := validate.RunRequest(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf(
		"language=%s tests=%d",
		req.Language,
		len(req.Tests),
	)

	resp := models.RunResponse{
		Status: "accepted",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}