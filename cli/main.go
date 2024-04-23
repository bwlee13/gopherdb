/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/bwlee13/gopherdb/cli/cmd"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "help")
	}
	cmd.Execute()
}
