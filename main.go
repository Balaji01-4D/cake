package main

import (
	"context"
	"errors"
	"os"
	"os/exec"

	"github.com/balaji01-4d/cake/internal/cli"
)

func main() {
	ctx := context.Background()
	cliCtx := cli.CLIContext{}

	cmd := cli.NewRootCommand(ctx, cliCtx)
	if err := cmd.Execute(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.ExitCode())
		}
		os.Exit(1)
	}
}
