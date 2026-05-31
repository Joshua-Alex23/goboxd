package validate

import (
	"errors"
	"path/filepath"
	"strings"
)

func Filename(name string) error {

	if name == "" {
		return errors.New("filename required")
	}

	if filepath.Base(name) != name {
		return errors.New("invalid filename")
	}

	if strings.Contains(name, "..") {
		return errors.New("invalid filename")
	}

	if strings.HasPrefix(name, ".") {
		return errors.New("invalid filename")
	}

	return nil
}
