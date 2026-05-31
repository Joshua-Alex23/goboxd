package validate

import "testing"

func TestAllowedFlags(t *testing.T) {

	err := Flags(
		[]string{
			"-O2",
			"-Wall",
		},
		[]string{
			"-O0",
			"-O1",
			"-O2",
			"-O3",
			"-Wall",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestBlockedFlag(t *testing.T) {

	err := Flags(
		[]string{
			"-fplugin=evil.so",
		},
		[]string{
			"-O0",
			"-O1",
			"-O2",
			"-O3",
		},
	)

	if err == nil {
		t.Fatal("expected error")
	}
}
