package executor

func truncateOutput(s string) string {

	if len(s) <= MaxOutputBytes {
		return s
	}

	return s[:MaxOutputBytes] +
		"\n[OUTPUT TRUNCATED]"
}
