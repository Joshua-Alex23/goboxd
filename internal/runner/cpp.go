package runner

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"time"
)

func RunCpp(source string, stdin string) (Result, error) {
	tmpDir, err := os.MkdirTemp("", "goboxd-*")
	if err != nil {
		return Result{}, err
	}
	defer os.RemoveAll(tmpDir)

	src := tmpDir + "/main.cpp"
	bin := tmpDir + "/main"

	err = os.WriteFile(src, []byte(source), 0644)
	if err != nil {
		return Result{}, err
	}

	compile := exec.Command(
		"g++",
		src,
		"-o",
		bin,
	)

	out, err := compile.CombinedOutput()
	if err != nil {
		return Result{
			Stderr: string(out),
		}, nil
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()
	cmd := exec.CommandContext(
		ctx,
		bin,
	)
	cmd.Stdin = strings.NewReader(stdin)
	out, err = cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return Result{
			Stderr: "execution timed out",
		}, nil
	}
	return Result{
		Stdout: string(out),
	}, err
}
