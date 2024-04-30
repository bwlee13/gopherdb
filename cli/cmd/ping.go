package cmd

import (
	"bufio"
	"log"
	"net"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping test",
	Long:  "Ping test",
	Run: func(cmd *cobra.Command, args []string) {
		// wonder if instead of net Dial if this should have some handlers
		conn, err := net.Dial("tcp", "localhost:42069")
		if err != nil {
			log.Fatalf("Could not connect to server: %v", err)
		}
		defer conn.Close()

		conn.Write([]byte("ping"))
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read response: %v", err)
		}
		log.Printf("Response from server: %s", response)
	},
}
