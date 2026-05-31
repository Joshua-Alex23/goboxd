package models

type TestCase struct {
	Stdin          string `json:"stdin"`
	ExpectedStdout string `json:"expected_stdout"`
}

type RunRequest struct {
	Language string     `json:"language"`
	Source   string     `json:"source"`
	Flags    []string   `json:"flags,omitempty"`
	Tests    []TestCase `json:"tests"`
}

type TestResult struct {
	Passed   bool   `json:"passed"`
	Actual   string `json:"actual"`
	Expected string `json:"expected"`
}

type RunResponse struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`

	Passed bool `json:"passed,omitempty"`

	Results []TestResult `json:"results,omitempty"`
}
