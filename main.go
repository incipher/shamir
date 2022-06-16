package main

import (
	"os"

	"incipher.io/shamir/cmd"
	"incipher.io/shamir/utils"
)

func main() {
	inputSource := os.Stdin
	outputDestination := os.Stdout
	errorDestination := os.Stderr

	// Generate root command
	rootCommand := cmd.GenerateRootCommand(
		inputSource,
		outputDestination,
		errorDestination,
	)

	// Run root command
	err := rootCommand.Execute()
	if err != nil {
		utils.ExitWithError(errorDestination, err)
	}
}
