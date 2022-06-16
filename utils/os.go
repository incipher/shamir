package utils

import (
	"fmt"
	"io"
	"os"
)

// Prints to stderr and exits with an error code.
func ExitWithError(errorDestination io.Writer, err error) {
	fmt.Fprintln(errorDestination, err.Error())
	os.Exit(1)
}
