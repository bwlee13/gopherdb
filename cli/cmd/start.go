package cmd

import (
	"fmt"
	"strings"

	"github.com/bwlee13/gopherdb/server"

	"github.com/spf13/cobra"
)

var (
	port          string
	policy        string
	validPolicies = []string{"LRU", "LFU", "MRU", "ARC", "TLRU"}
)

func init() {
	policiesDescription := fmt.Sprintf("Cache eviction policies (%s)", strings.Join(validPolicies, ", "))
	startCmd.Flags().StringVarP(&port, "port", "p", "42069", "TCP port for the server to listen on")
	startCmd.Flags().StringVar(&policy, "policy", "LRU", policiesDescription)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the GopherDB Server",
	Long:  "Start the GopherDB TCP server to handle incoming cache requests. This command initializes and runs the server in the background.",
	Run: func(cmd *cobra.Command, args []string) {
		policy = strings.ToUpper(policy)
		isValidPolicy := false
		for _, v := range validPolicies {
			if policy == v {
				isValidPolicy = true
				break
			}
		}
		if !isValidPolicy {
			fmt.Printf("Invalid policy: %s. Valid options are: %v\n", policy, validPolicies)
			return
		}

		server.InitServer(port, policy)
	},
}
