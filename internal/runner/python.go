package runner

import (
	"context"
	"os"
	"os/exec"
	"time"
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

	// out, err := exec.Command("python3", file).CombinedOutput()
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"python3",
		file,
	)
	out, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return Result{
			Stdout: "",
			Stderr: "execution timed out",
		}, nil
	}

	return Result{
		Stdout: string(out),
	}, err
}
