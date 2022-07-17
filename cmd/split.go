package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"incipher.io/shamir/shamir"
	"incipher.io/shamir/utils"
)

// Generates the split command.
func generateSplitCommand(
	inputSource io.Reader,
	outputDestination io.Writer,
	errorDestination io.Writer,
	isTerminal bool,
) *cobra.Command {
	// Declare command flag values
	var sharesCount int
	var thresholdCount int

	// Define command
	splitCommand := &cobra.Command{
		Use:   "split",
		Short: "Split a secret into shares",
		Long: `Splits a secret into shares (of length n), which a subset 
thereof (of length k) is necessary to reconstruct the 
original secret.`,
		Args: cobra.NoArgs,
		Run: runSplitCommand(
			inputSource,
			outputDestination,
			errorDestination,
			isTerminal,
			&sharesCount,
			&thresholdCount,
		),
	}

	// Define command flags
	splitCommand.Flags().IntVarP(
		&sharesCount,
		"shares",
		"n",
		0,
		"number of shares to be generated",
	)

	splitCommand.Flags().IntVarP(
		&thresholdCount,
		"threshold",
		"k",
		0,
		"number of shares necessary to reconstruct the secret",
	)

	splitCommand.MarkFlagRequired("shares")
	splitCommand.MarkFlagRequired("threshold")

	return splitCommand
}

// Runs the split command.
func runSplitCommand(
	inputSource io.Reader,
	outputDestination io.Writer,
	errorDestination io.Writer,
	isTerminal bool,
	sharesCount *int,
	thresholdCount *int,
) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Procure secret
		secret, err := readSecretFromPrompt(
			inputSource,
			outputDestination,
			errorDestination,
			isTerminal,
			sharesCount,
			thresholdCount,
		)

		// Split secret into shares
		shares, err := shamir.Split(
			secret,
			*sharesCount,
			*thresholdCount,
		)
		if err != nil {
			utils.ExitWithError(errorDestination, err)
		}

		// Print shares
		sharesConcatenated := strings.Join(shares, "\n")
		_, err = fmt.Fprintln(outputDestination, sharesConcatenated)
		if err != nil {
			utils.ExitWithError(errorDestination, err)
		}
	}
}

func readSecretFromPrompt(
	inputSource io.Reader,
	outputDestination io.Writer,
	errorDestination io.Writer,
	isTerminal bool,
	sharesCount *int,
	thresholdCount *int,
) (string, error) {
	prompt := promptui.Prompt{
		Stdin:  utils.NopReadCloser(inputSource),
		Stdout: utils.NopWriteCloser(errorDestination),
		Label:  "Secret",
		Mask:   '*',
		Validate: func(input string) error {
			if len(input) == 0 {
				return fmt.Errorf("secret must not be empty")
			}

			return nil
		},
	}

	return prompt.Run()
}
