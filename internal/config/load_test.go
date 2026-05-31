package config

import "testing"

func TestLoad(t *testing.T) {

	langs, err := Load("../../configs/languages.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if langs["py3"].Extension != "py" {
		t.Fatal("py3 not loaded")
	}
}
