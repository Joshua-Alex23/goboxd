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

	ready := true

	for _, cmd := range commands {

		_, err := exec.LookPath(cmd)

		ok := err == nil

		checks[cmd] = ok

		if !ok {
			ready = false
		}
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]any{
			"ready":  ready,
			"checks": checks,
		},
	)
}
