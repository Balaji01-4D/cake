package main

import (
	"context"

	"github.com/balaji01-4d/cake/internal/cli"
)

func main() {
	ctx := context.Background()
	cliCtx := cli.CLIContext{}

	cmd := cli.NewRootCommand(ctx, cliCtx)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
