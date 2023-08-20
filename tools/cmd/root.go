package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "tools",
		Short: "A tool kit for myself",
		Long:  `tools is a tool kit`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
