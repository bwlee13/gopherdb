package cmd

import (
	"fmt"
	"log"

	"github.com/bwlee13/gopherdb/server"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a specified GopherDB Server",
	Long:  "Stop a specified GopherDB TCP server.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := server.StopServer(port); err != nil {
			log.Fatalf("Error stopping the server: %v \n", err)
		}
		fmt.Printf("Server listening on port: %v stopped successfully\n", port)
	},
}
