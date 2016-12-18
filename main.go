package main

import (
	"github.com/alok87/github-cli/cmd"
)

var (
	BuildVersion = "0.1" // TODO: Need to make it part of build
)

func main() {
	cmd.Execute()
}
