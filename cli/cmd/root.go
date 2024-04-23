/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gopherdb [command] (flags)",
	Short: "GopherDB command-line interface and server",
	// TODO: Add a pointer to the docs in Long.
	Long: `GopherDB command-line interface and server.
	gopherdb is a CLI library for GopherDB that allows easy interface with GopherDB actions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Inside GopherDB root run func")
	},
	SilenceUsage: true,

	// consider adding this in early stages.
	// however, each func should handle errors appropriately instead
	// SilenceErrors: true,
	Version: "GopherDB Version 1.0.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		versionCmd,
		timezoneCmd,
		// startCmd,
		// initCmd,
	)

	cobra.EnableCommandSorting = false

	// Set an error function for flag parsing which prints the usage message.
	rootCmd.SetFlagErrorFunc(
		func(c *cobra.Command, err error) error {
			if err := c.Usage(); err != nil {
				return err
			}
			fmt.Fprintln(c.OutOrStderr()) // just provides line break between usage and err
			return err
		})
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// use this in other cmd/***.go commands to print errors to the terminal
func UsageAndErr(cmd *cobra.Command, args []string) error {
	if err := cmd.Usage(); err != nil {
		return err
	}
	return fmt.Errorf("unknown sub-command: %q", strings.Join(args, " "))
}
