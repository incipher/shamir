package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"incipher.io/shamir/shamir"
	"incipher.io/shamir/utils"
)

// Generates the combine command
func generateCombineCommand() *cobra.Command {
	// Declare command flag values
	var thresholdCount int

	// Define command
	combineCommand := &cobra.Command{
		Use:   "combine",
		Short: "Reconstruct a secret from shares",
		Long:  "Reconstructs a secret from shares.",
		Args:  cobra.NoArgs,
		Run:   runCombineCommand(&thresholdCount),
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

// Runs the combine command
func runCombineCommand(
	thresholdCount *int,
) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Validate flag values
		if *thresholdCount < 2 || *thresholdCount > 255 {
			utils.ExitWithError("threshold must be between 2 and 255")
		}

		// Prompt user for shares
		shares := make([]string, *thresholdCount)

		for i := 0; i < *thresholdCount; i++ {
			prompt := promptui.Prompt{
				Label: fmt.Sprintf("Share #%d", i+1),
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
				utils.ExitWithError(err.Error())
			}

			shares[i] = share
		}

		// Reconstruct secret from share
		secret, err := shamir.Combine(shares)

		if err != nil {
			utils.ExitWithError(err.Error())
		}

		// Print secret
		fmt.Println(secret)
	}
}
