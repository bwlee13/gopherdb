/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "output version information",
	Long:  "output current version information for GopherDB",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GopherDB Version 1.0.0")
	},
}

func init() {
	// this is added in root for tracking purposes since its not a rollup
	// rootCmd.AddCommand(versionCmd)
}
