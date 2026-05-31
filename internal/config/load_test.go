package config

import "testing"

func TestLoad(t *testing.T) {
	cfg, err := Load("../../configs/languages.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if len(cfg.Languages) != 2 {
		t.Fatalf("expected 2 languages got %d", len(cfg.Languages))
	}

	if cfg.Languages[0].ID != "py3" {
		t.Fatal("py3 not loaded")
	}
}
