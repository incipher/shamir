package utils

import (
	"fmt"
	"os"
)

// Prints to stderr and exits with an error code.
func ExitWithError(err string) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
