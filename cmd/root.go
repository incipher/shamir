package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

// Defines and runs the CLI
func Execute() {
	examples := []string{"  $ shamir split -n 3 -t 2", "  $ shamir combine -t 2"}

	// Define root command
	rootCommand := &cobra.Command{
		Use:     "shamir",
		Short:   "Split and combine secrets using Shamir's Secret Sharing algorithm.",
		Long:    "Split and combine secrets using Shamir's Secret Sharing algorithm.",
		Version: "0.1.1",
		Example: strings.Join(examples, "\n"),
	}

	// Define commands
	rootCommand.AddCommand(generateSplitCommand())
	rootCommand.AddCommand(generateCombineCommand())

	// Run CLI
	cobra.CheckErr(rootCommand.Execute())
}
