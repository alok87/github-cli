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
	// viper.SetConfigName(".ghi")
	// viper.AddConfigPath("$HOME")
	// viper.SetConfigType("json")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("No configuration file loaded - using defaults")
	// }
	// fmt.Println(viper.Get("msg")) // this would be "steve"
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "\n%v\n", err)
	os.Exit(1)
}
