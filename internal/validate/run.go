package validate

import (
	"errors"

	"github.com/thesouldev/goboxd/internal/models"
)

func RunRequest(req models.RunRequest) error {
	if req.Language == "" {
		return errors.New("language is requrired")
	}

	if req.Source == "" {
		return errors.New("source is required")
	}
	return nil
}
