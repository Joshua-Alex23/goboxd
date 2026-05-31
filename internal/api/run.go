package api

import (
	"encoding/json"
	"net/http"

	"github.com/thesouldev/goboxd/internal/config"
	"github.com/thesouldev/goboxd/internal/models"
	"github.com/thesouldev/goboxd/internal/runner"
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

	// log.Printf(
	// 	"language=%s tests=%d",
	// 	req.Language,
	// 	len(req.Tests),
	// )

	// resp := models.RunResponse{
	// 	Status: "accepted",
	// }

	// if req.Language != "py3" {
	// 	http.Error(w, "unsupported language", http.StatusBadRequest)
	// 	return
	// }

	// lang, ok := languages[req.Language]
	// if !ok {
	// 	http.Error(w, "unsupported language", http.StatusBadRequest)
	// 	return
	// }

	// result, err := runner.RunPython(req.Source)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	_, ok := config.Registry[req.Language]
	if !ok {
		http.Error(w, "unsupported language", http.StatusBadRequest)
		return
	}
	var (
		result runner.Result
		err    error
	)

	switch req.Language {
	case "py3":
		result, err = runner.RunPython(req.Source, "")

	case "cpp":
		result, err = runner.RunCpp(req.Source, "")

	default:
		http.Error(w, "unsupported language", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.RunResponse{
		Stdout: result.Stdout,
		Stderr: result.Stderr,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
