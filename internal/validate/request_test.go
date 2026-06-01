package validate

import (
	"testing"

	"github.com/thesouldev/goboxd/internal/models"
)

func TestRunRequestRequiresLanguage(t *testing.T) {
	err := RunRequest(models.RunRequest{
		Source: "print(1)",
		Tests: []models.TestCase{
			{},
		},
	})

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestRunRequestRequiresTests(t *testing.T) {
	err := RunRequest(models.RunRequest{
		Language: "py3",
		Source:   "print(1)",
	})

	if err == nil {
		t.Fatal("expected error")
	}
}
