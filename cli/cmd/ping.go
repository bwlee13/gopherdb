package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"

	"github.com/bwlee13/gopherdb/storage/response"
)

var cacheResp response.CacheResponse

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping test",
	Long:  "Ping test",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := net.Dial("tcp", "localhost:42069")
		if err != nil {
			log.Fatalf("Could not connect to server: %v", err)
		}
		defer conn.Close()

		conn.Write([]byte("ping\n"))
		decoder := json.NewDecoder(conn)
		err = decoder.Decode(&cacheResp)

		if err != nil {
			log.Fatalf("Failed to read response: %v", err)
		}

		if cacheResp.Error != "" {
			log.Fatalf("Error from server: %v", cacheResp.Error)
		}

		fmt.Printf("gopherdb:%s> %s\n", port, cacheResp.Message)
	},
}
