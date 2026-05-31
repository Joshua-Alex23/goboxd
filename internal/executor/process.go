package executor

import (
	"context"
	"os/exec"
	"strings"
)

func runProcess(
	ctx context.Context,
	command string,
	args []string,
	stdin string,
) ([]byte, error) {

	nsjailArgs := []string{
		"--quiet",
		"--disable_proc",
		"--iface_no_lo",
		"--user",
		"65534",
		"--group",
		"65534",
		"--",
		command,
	}

	nsjailArgs = append(
		nsjailArgs,
		args...,
	)

	cmd := exec.CommandContext(
		ctx,
		"nsjail",
		nsjailArgs...,
	)

	cmd.Stdin = strings.NewReader(stdin)

	return cmd.CombinedOutput()
	// out, err := cmd.CombinedOutput()

	// if err != nil {
	// 	return out, err
	// }

	// return out, nil
}
