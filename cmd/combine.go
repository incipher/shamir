package cmd

import (
	"fmt"
	"io"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"incipher.io/shamir/shamir"
	"incipher.io/shamir/utils"
)

// Generates the combine command.
func generateCombineCommand(
	inputSource io.Reader,
	outputDestination io.Writer,
	errorDestination io.Writer,
) *cobra.Command {
	// Declare command flag values
	var thresholdCount int

	// Define command
	combineCommand := &cobra.Command{
		Use:   "combine",
		Short: "Reconstruct a secret from shares",
		Long:  "Reconstructs a secret from shares.",
		Args:  cobra.NoArgs,
		Run: runCombineCommand(
			inputSource,
			outputDestination,
			errorDestination,
			&thresholdCount,
		),
	}

	// Define command flags
	combineCommand.Flags().IntVarP(
		&thresholdCount,
		"threshold",
		"t",
		0,
		"number of shares necessary to reconstruct the secret",
	)

	combineCommand.MarkFlagRequired("threshold")

	return combineCommand
}

// Runs the combine command.
func runCombineCommand(
	inputSource io.Reader,
	outputDestination io.Writer,
	errorDestination io.Writer,
	thresholdCount *int,
) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Validate flag values
		if *thresholdCount < 2 || *thresholdCount > 255 {
			utils.ExitWithError(errorDestination, fmt.Errorf("threshold must be between 2 and 255"))
		}

		// Prompt user for shares
		shares := make([]string, *thresholdCount)

		for i := 0; i < *thresholdCount; i++ {
			prompt := promptui.Prompt{
				Stdin:  utils.NopReadCloser(inputSource),
				Stdout: utils.NopWriteCloser(errorDestination),
				Label:  fmt.Sprintf("Share #%d", i+1),
				Validate: func(input string) error {
					if len(input) == 0 {
						return fmt.Errorf("share must not be empty")
					}

					if len(input)%2 != 0 {
						return fmt.Errorf("share must not be of odd length")
					}

					if _, err := utils.HexToByteArray(input); err != nil {
						return fmt.Errorf("share must be in hexadecimal encoding")
					}

					return nil
				},
			}

			share, err := prompt.Run()
			if err != nil {
				utils.ExitWithError(errorDestination, err)
			}

			shares[i] = share
		}

		// Reconstruct secret from shares
		secret, err := shamir.Combine(shares)
		if err != nil {
			utils.ExitWithError(errorDestination, err)
		}

		// Print secret
		_, err = fmt.Fprintln(outputDestination, secret)
		if err != nil {
			utils.ExitWithError(errorDestination, err)
		}
	}
}
