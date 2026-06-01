package executor

import (
	"strings"
	"testing"
)

func TestTruncateOutput(t *testing.T) {
	input := strings.Repeat("a", MaxOutputBytes+100)

	output := truncateOutput(input)

	if !strings.Contains(output, "[OUTPUT TRUNCATED]") {
		t.Fatal("expected truncation marker")
	}
}
