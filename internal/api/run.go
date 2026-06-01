package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/thesouldev/goboxd/internal/config"
	"github.com/thesouldev/goboxd/internal/executor"
	"github.com/thesouldev/goboxd/internal/models"
)

func Run(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var req models.RunRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			`{"error":{"code":"invalid_json","message":"bad request"}}`,
			http.StatusBadRequest,
		)
		return
	}

	lang, ok := config.Registry[req.Language]
	if !ok {
		http.Error(
			w,
			`{"error":{"code":"unknown_language","message":"language not registered"}}`,
			http.StatusBadRequest,
		)
		return
	}

	if req.ArtifactFilename != "" {
		lang.Artifact = req.ArtifactFilename
	}

	if len(req.Tests) == 0 {
		http.Error(
			w,
			`{"error":{"code":"invalid_request","message":"at least one test required"}}`,
			http.StatusBadRequest,
		)
		return
	}

	results := make(
		[]models.TestResult,
		0,
		len(req.Tests),
	)

	finalStatus := "accepted"

	build := models.BuildResult{
		Status: "ok",
	}

	for _, tc := range req.Tests {
		out, _ := executor.Execute(
			lang,
			req.Source,
			tc.Stdin,
			nil,
		)

		status := "accepted"

		if out.Stderr != "" {
			status = "runtime_error"
		} else if strings.TrimSpace(out.Stdout) != strings.TrimSpace(tc.ExpectedStdout) {
			status = "wrong_output"
		}

		if status != "accepted" &&
			finalStatus == "accepted" {
			finalStatus = status
		}

		results = append(
			results,
			models.TestResult{
				Status: status,
				Stdout: out.Stdout,
				Stderr: out.Stderr,
			},
		)
	}

	resp := models.RunResponse{
		Status: finalStatus,
		Build:  build,
		Tests:  results,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}
