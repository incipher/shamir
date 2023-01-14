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
	examples := []string{"  $ shamir split -n 5 -k 3", "  $ shamir combine -k 3"}

	rootCommand := &cobra.Command{
		Use:     "shamir",
		Short:   "Split and combine secrets using Shamir's Secret Sharing algorithm.",
		Long:    "Split and combine secrets using Shamir's Secret Sharing algorithm.",
		Version: "0.5.0",
		Example: strings.Join(examples, "\n"),
	}

	rootCommand.SetIn(inputSource)
	rootCommand.SetOut(outputDestination)
	rootCommand.SetErr(errorDestination)

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
