package validate

import (
	"errors"

	"github.com/thesouldev/goboxd/internal/models"
)

const (
	MaxSourceBytes = 256 * 1024
	MaxTests       = 50
)

func RunRequest(req models.RunRequest) error {

	if req.Language == "" {
		return errors.New("language is required")
	}

	if req.Source == "" {
		return errors.New("source is required")
	}

	if len(req.Source) > MaxSourceBytes {
		return errors.New("source exceeds maximum size")
	}

	if len(req.Tests) == 0 {
		return errors.New("at least one test is required")
	}

	if len(req.Tests) > MaxTests {
		return errors.New("too many tests")
	}

	return nil
}
