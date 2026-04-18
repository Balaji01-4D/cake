package app

import (
	"context"
	"os"
	"os/exec"

	"github.com/google/shlex"
)

type App interface {
	Run(ctx context.Context, commandString string) error
}

type Cake struct{}

func (c *Cake) Run(ctx context.Context, commandString string) error {
	args, err := shlex.Split(commandString)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
