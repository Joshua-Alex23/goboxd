package executor

import "testing"

func TestRenderTemplate(t *testing.T) {
	out := renderTemplate(
		"{{source}} -> {{artifact}}",
		map[string]string{
			"source":   "a.cpp",
			"artifact": "a.out",
		},
	)

	if out != "a.cpp -> a.out" {
		t.Fatalf("got %q", out)
	}
}
