package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "help")
	}

	// ignore errors so cobra can handle
	cmd, _, _ := gopherCmd.Find(os.Args[1:])

	cmdName := commandName(cmd)

	err := doMain(cmd, cmdName)
	// errCode := exit.Success()
	if err != nil {
		fmt.Printf("Failed running %q\n", cmdName)
		fmt.Printf("Error: %s\n", err)
	}
	// 	// Display the error and its details/hints.
	// 	clierror.OutputError(stderr, err, true /*showSeverity*/, false /*verbose*/)

	// 	// Remind the user of which command was being run.

	// 	// Finally, extract the error code, as optionally specified
	// 	// by the sub-command.
	// 	errCode = getExitCode(err)
	// }
	// // Finally, gracefully shutdown logging facilities.
	// cliCtx.logShutdownFn()

	// exit.WithCode(errCode)

}

func doMain(cmd *cobra.Command, cmdName string) error {
	fmt.Println("Commands: ", cmd)
	fmt.Println("CommandName: ", cmdName)
	return Run(os.Args[1:])
}

func Run(args []string) error {
	gopherCmd.SetArgs(args)
	return gopherCmd.Execute()
}

func commandName(cmd *cobra.Command) string {
	rootName := gopherCmd.CommandPath()
	if cmd != nil {
		return strings.TrimPrefix(cmd.CommandPath(), rootName+" ")
	}
	return rootName
}

var gopherCmd = &cobra.Command{
	Use:   "gopherdb [command] (flags)",
	Short: "GopherDB command-line interface and server",
	// TODO(cdo): Add a pointer to the docs in Long.
	Long: `GopherDB command-line interface and server.`,
	// Disable automatic printing of usage information whenever an error
	// occurs. Many errors are not the result of a bad command invocation,
	// e.g. attempting to start a node on an in-use port, and printing the
	// usage information in these cases obscures the cause of the error.
	// Commands should manually print usage information when the error is,
	// in fact, a result of a bad invocation, e.g. too many arguments.
	SilenceUsage: true,
	// Disable automatic printing of the error. We want to also print
	// details and hints, which cobra does not do for us. Instead
	// we do the printing in Main().
	SilenceErrors: true,
	// Version causes cobra to automatically support a --version flag
	// that reports this string.
	Version: "GopherDB Version 1.0.0",
	// Prevent cobra from auto-generating a completions command,
	// since we provide our own.
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "output version information",
	Long:  "output current version information for GopherDB",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GopherDB Version 1.0.0")
	},
}

// GopherCmd returns the root gopherdb Command object.
func GopherCmd() *cobra.Command {
	return gopherCmd
}

func init() {
	cobra.EnableCommandSorting = false

	// Set an error function for flag parsing which prints the usage message.
	gopherCmd.SetFlagErrorFunc(
		func(c *cobra.Command, err error) error {
			if err := c.Usage(); err != nil {
				return err
			}
			fmt.Fprintln(c.OutOrStderr()) // just provides line break between usage and err
			return err
		})
	gopherCmd.AddCommand(
		// startCmd,
		// initCmd,
		versionCmd,
	)
}

// use this in other cmd/***.go commands to print errors to the terminal
func UsageAndErr(cmd *cobra.Command, args []string) error {
	if err := cmd.Usage(); err != nil {
		return err
	}
	return fmt.Errorf("unknown sub-command: %q", strings.Join(args, " "))
}
