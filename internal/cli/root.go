package cli

import (
	"context"
	"os"

	"github.com/balaji01-4d/cake/internal/logger"
	"github.com/spf13/cobra"
)

func NewRootCommand(ctx context.Context, cliCtx CLIContext) *cobra.Command {
	var debug debugFlag
	var f *os.File
	arg := "Makefile"

	cmd := &cobra.Command{
		Use:   "cake [file]",
		Short: "Interactive utility to discover and run Makefile targets with fuzzy search and live preview of execution steps.",

		Long: `Cake is a lightweight CLI utility that improves the ergonomics of working with Makefiles.

Instead of manually inspecting or remembering targets, it scans the project's Makefile and presents all available commands in an interactive interface. You can quickly filter targets using fuzzy search, navigate through them, and execute a selected command directly.

It also provides a live preview of the underlying shell commands for the currently selected target, giving clear visibility into what will run before execution.`,

		Args: cobra.MaximumNArgs(1),

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			logger, err := logger.New(bool(debug))
			if err != nil {
				return err
			}

			cliCtx.Logger = logger
			return nil
		},

		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				cliCtx.Logger.Debug("User specified Makefile", "Makefile", args[0])
				arg = args[0]
			}

			var errOpen error
			f, errOpen = os.Open(arg)
			if errOpen != nil {
				cliCtx.Logger.Error("Failed to open Makefile", "Makefile", arg, "error", errOpen)
				return errOpen
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},

		PostRunE: func(cmd *cobra.Command, args []string) error {
			if f == nil {
				return nil
			}

			errClose := f.Close()
			if errClose != nil {
				cliCtx.Logger.Error("Failed to close Makefile", "Makefile", arg, "error", errClose)
				return errClose
			}

			return nil
		},
	}

	debug.bind(cmd)
	return cmd
}
