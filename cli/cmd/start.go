package cmd

import (
	"bufio"
	"fmt"
	"os"
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
		if !contains(validPolicies, policy) {
			fmt.Printf("Invalid policy: %s. Valid options are: %v\n", policy, validPolicies)
			return
		}
		server.InitServer(port, policy)
	},
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func interact(port string, policy string) {
	server.InitServer(port, policy)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("gopherdb:%s> ", port)
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if strings.TrimSpace(line) == "exit" {
			break
		}
		fmt.Println("Executed: ", line)
	}

}
