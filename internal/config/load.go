package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (Languages, error) {
	var langs Languages

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &langs)
	if err != nil {
		return nil, err
	}

	return langs, nil
}
