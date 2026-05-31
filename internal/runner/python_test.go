package runner

import "testing"

func TestRunPython(t *testing.T) {

	result, err := RunPython(`print("hello")`, "")
	if err != nil {
		t.Fatal(err)
	}

	if result.Stdout != "hello\n" {
		t.Fatalf("expected hello, got %q", result.Stdout)
	}
}
