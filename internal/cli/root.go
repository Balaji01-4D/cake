package cli

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cake [file]",
		Short: "Interactive utility to discover and run Makefile targets with fuzzy search and live preview of execution steps.",

		Long: `Cake is a lightweight CLI utility that improves the ergonomics of working with Makefiles.

Instead of manually inspecting or remembering targets, it scans the project's Makefile and presents all available commands in an interactive interface. You can quickly filter targets using fuzzy search, navigate through them, and execute a selected command directly.

It also provides a live preview of the underlying shell commands for the currently selected target, giving clear visibility into what will run before execution.`,
	}
	return cmd
}
