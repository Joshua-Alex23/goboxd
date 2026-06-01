package executor

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/thesouldev/goboxd/internal/config"
	"github.com/thesouldev/goboxd/internal/validate"
)

const MaxOutputBytes = 1024 * 1024

type Result struct {
	Stdout string
	Stderr string
}

func Execute(
	lang config.Language,
	source string,
	stdin string,
	flags []string,
) (Result, error) {

	tmpDir, err := os.MkdirTemp("", "goboxd-*")
	if err != nil {
		return Result{}, err
	}
	defer os.RemoveAll(tmpDir)

	if err := validate.Filename(lang.SourceFilename); err != nil {
		return Result{}, err
	}

	if lang.Artifact != "" {
		if err := validate.Filename(lang.Artifact); err != nil {
			return Result{}, err
		}
	}

	sourceFile := filepath.Join(
		tmpDir,
		lang.SourceFilename,
	)

	if err := os.WriteFile(
		sourceFile,
		[]byte(source),
		0644,
	); err != nil {
		return Result{}, err
	}

	artifactPath := filepath.Join(
		tmpDir,
		"a.out",
	)

	if lang.Artifact != "" {
		artifactPath = filepath.Join(
			tmpDir,
			lang.Artifact,
		)
	}

	replacements := map[string]string{
		"source":   sourceFile,
		"artifact": artifactPath,
		"flags":    "",
	}

	if len(flags) > 0 {
		replacements["flags"] = strings.Join(
			flags,
			" ",
		)
	}

	// ================= BUILD =================

	if lang.Build != nil {

		buildCmd := renderTemplate(
			lang.Build.Cmd,
			replacements,
		)

		buildArgs := renderArgs(
			lang.Build.Args,
			replacements,
		)

		out, err := runRawProcess(
			context.Background(),
			buildCmd,
			buildArgs,
			"",
			tmpDir,
		)

		if err != nil {
			return Result{
				Stderr: truncateOutput(string(out)),
			}, nil
		}
	}

	// ================= RUN =================

	runCmd := renderTemplate(
		lang.Run.Cmd,
		replacements,
	)

	runArgs := renderArgs(
		lang.Run.Args,
		replacements,
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(lang.Run.Limits.WallTimeS)*time.Second,
	)
	defer cancel()

	out, err := runProcess(
		ctx,
		runCmd,
		runArgs,
		stdin,
		tmpDir,
	)

	if ctx.Err() == context.DeadlineExceeded {
		return Result{
			Stderr: "execution timed out",
		}, nil
	}

	if err != nil {
		return Result{
			Stderr: truncateOutput(string(out)),
		}, nil
	}

	return Result{
		Stdout: truncateOutput(string(out)),
	}, nil
}
