package main

import (
	"github.com/spf13/cobra"
)

var stopCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop server",
		Long:  "Stop server",
		Args:  stopArgs,
		Run:   stopRun,
	}
	return cmd
}()

func stopArgs(cmd *cobra.Command, args []string) error {
	verbosePrint("verbose: %t", verbose)
	verbosePrint("configPath: %s", configPath)

	return nil
}

func stopRun(cmd *cobra.Command, args []string) {
}
