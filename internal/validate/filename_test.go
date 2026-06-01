// package validate

// import "testing"

// func TestFilenameValid(t *testing.T) {

// 	if err := Filename("solution.py"); err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestFilenameTraversal(t *testing.T) {

// 	if Filename("../../etc/passwd") == nil {
// 		t.Fatal("expected error")
// 	}
// }

// func TestFilenameHidden(t *testing.T) {

//		if Filename(".secret") == nil {
//			t.Fatal("expected error")
//		}
//	}
package validate

import "testing"

func TestFilenameValid(t *testing.T) {
	if err := Filename("solution.py"); err != nil {
		t.Fatalf("expected valid filename")
	}
}

func TestFilenameRejectTraversal(t *testing.T) {
	if Filename("../passwd") == nil {
		t.Fatalf("expected error")
	}
}

func TestFilenameRejectAbsolute(t *testing.T) {
	if Filename("/tmp/test") == nil {
		t.Fatalf("expected error")
	}
}
