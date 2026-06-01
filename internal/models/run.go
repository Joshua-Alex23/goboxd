package models

type TestCase struct {
	Stdin          string `json:"stdin"`
	ExpectedStdout string `json:"expected_stdout"`
}

type Limits struct {
	WallTimeS    int `json:"wall_time_s"`
	MemoryKB     int `json:"memory_kb"`
	MaxProcesses int `json:"max_processes"`
}

type BuildSpec struct {
	Limits Limits   `json:"limits"`
	Flags  []string `json:"flags"`
}

type RunSpec struct {
	Limits Limits   `json:"limits"`
	Flags  []string `json:"flags"`
}

type RunRequest struct {
	Language         string     `json:"language"`
	Source           string     `json:"source"`
	SourceFilename   string     `json:"source_filename"`
	ArtifactFilename string     `json:"artifact_filename"`
	Build            *BuildSpec `json:"build"`
	Run              *RunSpec   `json:"run"`
	Tests            []TestCase `json:"tests"`
}

type TestResult struct {
	Status     string `json:"status"`
	Stdout     string `json:"stdout"`
	Stderr     string `json:"stderr"`
	DurationMS int64  `json:"duration_ms"`
}

type BuildResult struct {
	Status     string `json:"status"`
	Stdout     string `json:"stdout"`
	Stderr     string `json:"stderr"`
	DurationMS int64  `json:"duration_ms"`
}

type RunResponse struct {
	Status string       `json:"status"`
	Build  BuildResult  `json:"build"`
	Tests  []TestResult `json:"tests"`
}
