package validate

import "testing"

func TestFilenameValid(t *testing.T) {

	if err := Filename("solution.py"); err != nil {
		t.Fatal(err)
	}
}

func TestFilenameTraversal(t *testing.T) {

	if Filename("../../etc/passwd") == nil {
		t.Fatal("expected error")
	}
}

func TestFilenameHidden(t *testing.T) {

	if Filename(".secret") == nil {
		t.Fatal("expected error")
	}
}
