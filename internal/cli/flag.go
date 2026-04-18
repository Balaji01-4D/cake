package cli

import "github.com/spf13/cobra"

type debugFlag bool

func (f *debugFlag) bind(cmd *cobra.Command) {
	cmd.Flags().BoolVar((*bool)(f), "debug", false, "Enable debug logging.")
}
