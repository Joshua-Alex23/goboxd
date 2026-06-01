package api

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

func Readyz(w http.ResponseWriter, r *http.Request) {
	checks := map[string]bool{}

	commands := []string{
		"python3",
		"g++",
		"nsjail",
	}

	status := "ready"

	for _, cmd := range commands {
		_, err := exec.LookPath(cmd)

		ok := err == nil
		checks[cmd] = ok

		if !ok {
			status = "not_ready"
		}
	}

	w.Header().Set("Content-Type", "application/json")

	resp := map[string]any{
		"status": status,
		"checks": checks,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
