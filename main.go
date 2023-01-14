package main

import (
	"os"
	"syscall"

	"golang.org/x/term"
	"incipher.io/shamir/cmd"
	"incipher.io/shamir/utils"
)

func main() {
	inputSource := os.Stdin
	outputDestination := os.Stdout
	errorDestination := os.Stderr
	isTerminal := term.IsTerminal(int(syscall.Stdin))

	rootCommand := cmd.GenerateRootCommand(
		inputSource,
		outputDestination,
		errorDestination,
		isTerminal,
	)

	err := rootCommand.Execute()
	if err != nil {
		utils.ExitWithError(errorDestination, err)
	}
}
