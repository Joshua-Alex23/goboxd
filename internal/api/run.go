package api

import (
	"encoding/json"
	"net/http"

	"github.com/thesouldev/goboxd/internal/config"
	"github.com/thesouldev/goboxd/internal/executor"
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

	lang, ok := config.Registry[req.Language]
	if !ok {
		http.Error(
			w,
			"unsupported language",
			http.StatusBadRequest,
		)
		return
	}

	result, err := executor.Execute(
		lang,
		req.Source,
		"",
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	resp := models.RunResponse{
		Stdout: result.Stdout,
		Stderr: result.Stderr,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(resp)
}
