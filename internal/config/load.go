package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	// // debug (keep for now)
	// for i, lang := range cfg.Languages {
	// 	if lang.Build != nil {
	// 		// fmt.Printf("\n[LOAD DEBUG %d]\n", i)
	// 		// fmt.Printf("ARGS PTR: %p\n", lang.Build.Args)
	// 		// fmt.Printf("ARGS: %#v\n", lang.Build.Args)
	// 	}
	// }

	return &cfg, nil
}
