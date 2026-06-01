// package config

// import "testing"

// func TestLoad(t *testing.T) {
// 	cfg, err := Load("../../configs/languages.yaml")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(cfg.Languages) != 2 {
// 		t.Fatalf("expected 2 languages got %d", len(cfg.Languages))
// 	}

//		if cfg.Languages[0].ID != "py3" {
//			t.Fatal("py3 not loaded")
//		}
//	}
package config

import (
	"os"
	"testing"
)

func TestLoadLanguages(t *testing.T) {
	content := `
languages:
  - id: py3
    name: Python 3
    source_filename: solution.py
    run:
      cmd: python3
`

	tmp, err := os.CreateTemp("", "cfg-*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.WriteString(content); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}

	if len(cfg.Languages) != 1 {
		t.Fatalf("expected 1 language, got %d", len(cfg.Languages))
	}

	if cfg.Languages[0].ID != "py3" {
		t.Fatalf("expected py3")
	}
}
