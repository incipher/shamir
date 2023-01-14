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
	isTerminal bool,
) *cobra.Command {
	var thresholdCount int

	combineCommand := &cobra.Command{
		Use:   "combine",
		Short: "Reconstruct a secret from shares",
		Long:  "Reconstructs a secret from shares.",
		Args:  cobra.NoArgs,
		Run: runCombineCommand(
			inputSource,
			outputDestination,
			errorDestination,
			isTerminal,
			&thresholdCount,
		),
	}

	combineCommand.Flags().IntVarP(
		&thresholdCount,
		"threshold",
		"k",
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
	isTerminal bool,
	thresholdCount *int,
) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if *thresholdCount < 2 || *thresholdCount > 255 {
			utils.ExitWithError(errorDestination, fmt.Errorf("threshold must be between 2 and 255"))
		}

		var shares []string
		var err error

		if isTerminal {
			shares, err = readSharesFromPrompts(
				inputSource,
				outputDestination,
				errorDestination,
				isTerminal,
				thresholdCount,
			)
		} else {
			shares, err = readSharesFromInputSource(
				inputSource,
				outputDestination,
				errorDestination,
				isTerminal,
				thresholdCount,
			)
		}
		if err != nil {
			utils.ExitWithError(errorDestination, err)
		}

		secret, err := shamir.Combine(shares)
		if err != nil {
			utils.ExitWithError(errorDestination, err)
		}

		_, err = fmt.Fprintln(outputDestination, secret)
		if err != nil {
			utils.ExitWithError(errorDestination, err)
		}
	}
}

func readSharesFromPrompts(
	inputSource io.Reader,
	outputDestination io.Writer,
	errorDestination io.Writer,
	isTerminal bool,
	thresholdCount *int,
) ([]string, error) {
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
			return nil, err
		}

		shares[i] = share
	}

	return shares, nil
}

func readSharesFromInputSource(
	inputSource io.Reader,
	outputDestination io.Writer,
	errorDestination io.Writer,
	isTerminal bool,
	thresholdCount *int,
) ([]string, error) {
	shares := utils.ReadLines(inputSource)

	if len(shares) != *thresholdCount {
		return nil, fmt.Errorf("number of shares must be equal to threshold")
	}

	for i, share := range shares {
		if len(share) == 0 {
			return nil, fmt.Errorf("share #%d must not be empty", i+1)
		}

		if len(share)%2 != 0 {
			return nil, fmt.Errorf("share #%d must not be of odd length", i+1)
		}

		if _, err := utils.HexToByteArray(share); err != nil {
			return nil, fmt.Errorf("share #%d must be in hexadecimal encoding", i+1)
		}
	}

	return shares, nil
}
