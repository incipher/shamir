package cmd

import (
	"github.com/spf13/cobra"
)

// Defines and runs the CLI
func Execute() {
	// Define root command
	rootCommand := &cobra.Command{
		Use:     "shamir",
		Short:   "Split and combine secrets using Shamir's Secret Sharing algorithm.",
		Long:    "Split and combine secrets using Shamir's Secret Sharing algorithm.",
		Version: "0.1.0",
	}

	// Define commands
	rootCommand.AddCommand(generateSplitCommand())
	rootCommand.AddCommand(generateCombineCommand())

	// Run CLI
	cobra.CheckErr(rootCommand.Execute())
}
