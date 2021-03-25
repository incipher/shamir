package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"incipher.io/shamir/shamir"
	"incipher.io/shamir/utils"
)

// Generates the split command
func generateSplitCommand() *cobra.Command {
	// Declare command flag values
	var sharesCount int
	var thresholdCount int

	// Define command
	splitCommand := &cobra.Command{
		Use:   "split",
		Short: "Split a secret into shares",
		Long: `Splits a secret into shares (of length n), which a subset 
thereof (of length t) is necessary to reconstruct the 
original secret.`,
		Args: cobra.NoArgs,
		Run:  runSplitCommand(&sharesCount, &thresholdCount),
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
		"t",
		0,
		"number of shares necessary to reconstruct the secret",
	)

	splitCommand.MarkFlagRequired("shares")
	splitCommand.MarkFlagRequired("threshold")

	return splitCommand
}

// Runs the split command
func runSplitCommand(
	sharesCount *int,
	thresholdCount *int,
) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Define secret prompt
		prompt := promptui.Prompt{
			Label: "Secret",
			Mask:  '*',
			Validate: func(input string) error {
				if len(input) == 0 {
					return fmt.Errorf("secret must not be empty")
				}

				return nil
			},
		}

		// Prompt user for secret
		secret, err := prompt.Run()

		if err != nil {
			utils.ExitWithError(err.Error())
		}

		// Split secret into shares
		shares, err := shamir.Split(
			secret,
			*sharesCount,
			*thresholdCount,
		)

		if err != nil {
			utils.ExitWithError(err.Error())
		}

		// Print shares
		for _, share := range shares {
			fmt.Println(share)
		}
	}
}
