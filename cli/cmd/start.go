package cmd

import (
	"github.com/bwlee13/gopherdb/server"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the GopherDB Server",
	Long:  "Start the GopherDB TCP server to handle incoming cache requests. This command initializes and runs the server in the background.",
	Run: func(cmd *cobra.Command, args []string) {
		server.InitServer()
	},
}
