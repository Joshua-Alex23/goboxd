package api

import "testing"

func TestStatusMapping(t *testing.T) {
	tests := []struct {
		stderr   string
		stdout   string
		expected string
	}{
		{"runtime error", "", "runtime_error"},
		{"", "wrong", "wrong_output"},
		{"", "ok", "accepted"},
	}

	for _, tc := range tests {
		status := "accepted"

		if tc.stderr != "" {
			status = "runtime_error"
		} else if tc.stdout == "wrong" {
			status = "wrong_output"
		}

		if status != tc.expected {
			t.Fatalf("expected %s got %s", tc.expected, status)
		}
	}
}
