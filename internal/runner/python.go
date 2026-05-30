package runner

import (
	"os"
	"os/exec"
)

type Result struct {
	Stdout string
	Stderr string
}




func RunPython(source string) (Result, error) {

	tmpDir, err := os.MkdirTemp("", "goboxd-*")
	if err != nil {
		return Result{}, err
	}
	defer os.RemoveAll(tmpDir)

	file := tmpDir + "/main.py"

	err = os.WriteFile(file, []byte(source), 0644)
	if err != nil {
		return Result{}, err
	}

	out, err := exec.Command("python3", file).CombinedOutput()

	return Result{
		Stdout: string(out),
	}, err
}