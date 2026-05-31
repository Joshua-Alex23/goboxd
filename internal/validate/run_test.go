package validate

import (
	"strings"
	"testing"

	"github.com/thesouldev/goboxd/internal/models"
)

func TestMissingTests(t *testing.T) {

	req := models.RunRequest{
		Language: "py3",
		Source:   "print(1)",
	}

	err := RunRequest(req)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSourceTooLarge(t *testing.T) {

	req := models.RunRequest{
		Language: "py3",
		Source: strings.Repeat(
			"a",
			MaxSourceBytes+1,
		),
		Tests: []models.TestCase{
			{},
		},
	}

	err := RunRequest(req)

	if err == nil {
		t.Fatal("expected error")
	}
}
