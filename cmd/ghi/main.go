package main

import (
	"fmt"
	"os"
)

var (
	BuildVersion = "0.1" // TODO: Need to make it part of build
)

func main() {
	Execute()
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
