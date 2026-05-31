package executor

import (
	"context"
	"os"
	"path/filepath"
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

	err = os.WriteFile(
		sourceFile,
		[]byte(source),
		0644,
	)
	if err != nil {
		return Result{}, err
	}

	artifactPath := ""
	if lang.Artifact != "" {
		artifactPath = filepath.Join(
			tmpDir,
			lang.Artifact,
		)
	}

	replacements := map[string]string{
		"source":   sourceFile,
		"artifact": artifactPath,
	}

	if lang.Build != nil {

		buildArgs := renderArgs(
			lang.Build.Args,
			replacements,
		)

		out, err := runProcess(
			context.Background(),
			lang.Build.Cmd,
			buildArgs,
			"",
		)
		if err != nil {
			return Result{
				Stderr: truncateOutput(string(out)),
			}, nil
		}
	}

	runCmdStr := renderTemplate(
		lang.Run.Cmd,
		replacements,
	)

	runArgs := renderArgs(
		lang.Run.Args,
		replacements,
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	out, err := runProcess(
		ctx,
		runCmdStr,
		runArgs,
		stdin,
	)

	if ctx.Err() == context.DeadlineExceeded {
		return Result{
			Stderr: "execution timed out",
		}, nil
	}

	return Result{
		Stdout: truncateOutput(string(out)),
	}, err
	// 	if err != nil {
	// 		return Result{
	// 			Stderr: truncateOutput(string(out)),
	// 		}, nil
	// 	}

	//	return Result{
	//		Stdout: truncateOutput(string(out)),
	//	}, nil
}
