package api

import (
	"encoding/json"
	"net/http"

	"github.com/thesouldev/goboxd/internal/config"
	"github.com/thesouldev/goboxd/internal/executor"
	"github.com/thesouldev/goboxd/internal/models"
	"github.com/thesouldev/goboxd/internal/queue"
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
	if lang.Build != nil {
		if err := validate.Flags(
			req.Flags,
			lang.Build.FlagAllowlist,
		); err != nil {

			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}
	}

	queue.Acquire()
	defer queue.Release()
	results := make([]models.TestResult, 0, len(req.Tests))
	allPassed := true

	for _, test := range req.Tests {

		result, err := executor.Execute(
			lang,
			req.Source,
			test.Stdin,
		)

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		passed := result.Stdout == test.ExpectedStdout

		if !passed {
			allPassed = false
		}

		results = append(results, models.TestResult{
			Passed:   passed,
			Actual:   result.Stdout,
			Expected: test.ExpectedStdout,
		})
	}

	resp := models.RunResponse{
		Passed:  allPassed,
		Results: results,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(resp)
}
