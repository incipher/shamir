package cmd

import (
	"io"
	"strings"

	"github.com/spf13/cobra"
)

// Generates the root command.
func GenerateRootCommand(
	inputSource io.Reader,
	outputDestination io.Writer,
	errorDestination io.Writer,
	isTerminal bool,
) *cobra.Command {
	examples := []string{"  $ shamir split -n 3 -t 2", "  $ shamir combine -t 2"}

	// Define root command
	rootCommand := &cobra.Command{
		Use:     "shamir",
		Short:   "Split and combine secrets using Shamir's Secret Sharing algorithm.",
		Long:    "Split and combine secrets using Shamir's Secret Sharing algorithm.",
		Version: "0.3.0",
		Example: strings.Join(examples, "\n"),
	}

	// Set inputs & outputs
	rootCommand.SetIn(inputSource)
	rootCommand.SetOut(outputDestination)
	rootCommand.SetErr(errorDestination)

	// Define commands
	rootCommand.AddCommand(
		generateSplitCommand(
			inputSource,
			outputDestination,
			errorDestination,
			isTerminal,
		),
	)
	rootCommand.AddCommand(
		generateCombineCommand(
			inputSource,
			outputDestination,
			errorDestination,
			isTerminal,
		),
	)

	return rootCommand
}
