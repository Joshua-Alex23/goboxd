package config

type Language struct {
	Extension string   `yaml:"extension"`
	Command   []string `yaml:"command"`
}

type Languages map[string]Language
