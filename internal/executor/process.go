package executor

import (
	"context"
	"os/exec"
	"strings"
)

func runRawProcess(
	ctx context.Context,
	command string,
	args []string,
	stdin string,
	workDir string,
) ([]byte, error) {

	cmd := exec.CommandContext(ctx, command, args...)

	if workDir != "" {
		cmd.Dir = workDir
	}

	cmd.Stdin = strings.NewReader(stdin)

	return cmd.CombinedOutput()
}

func runProcess(
	ctx context.Context,
	command string,
	args []string,
	stdin string,
	workDir string,
) ([]byte, error) {

	nsjailArgs := []string{
		"--really_quiet",

		"--chroot",
		"/",

		"--user",
		"65534",

		"--group",
		"65534",

		"--disable_proc",

		"--iface_no_lo",

		"--cwd",
		workDir,

		"--",
		command,
	}

	nsjailArgs = append(nsjailArgs, args...)

	cmd := exec.CommandContext(
		ctx,
		"nsjail",
		nsjailArgs...,
	)

	cmd.Stdin = strings.NewReader(stdin)

	return cmd.CombinedOutput()
}
