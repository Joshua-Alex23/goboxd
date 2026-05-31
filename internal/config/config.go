package config

type Limits struct {
	WallTimeS    int `yaml:"wall_time_s"`
	MemoryKB     int `yaml:"memory_kb"`
	MaxProcesses int `yaml:"max_processes"`
}

type Phase struct {
	Cmd           string   `yaml:"cmd"`
	Args          []string `yaml:"args"`
	Limits        Limits   `yaml:"limits"`
	FlagAllowlist []string `yaml:"flag_allowlist"`
}

type Language struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`

	SourceFilename string `yaml:"source_filename"`
	Artifact       string `yaml:"artifact"`

	Build *Phase `yaml:"build"`
	Run   Phase  `yaml:"run"`
}

type Config struct {
	Languages []Language `yaml:"languages"`
}
