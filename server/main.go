package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwlee13/gopherdb/storage/base"
)

// use below for CLI future work
// "github.com/spf13/cobra"

var (
	listen = flag.String("listen", ":42069", "address to listen to")
)

var store *base.Store

func main() {
	flag.Parse()

	// default LRU because thats all i have right now
	store = base.NewStore("LRU")
	service := NewService(*listen, store)
	go service.Start()
	log.Println("Starting server...")

	t := make(chan os.Signal, 1)
	signal.Notify(t, os.Interrupt, syscall.SIGTERM)
	<-t

	log.Println("Shutting down server...")
}

/*
https://github.com/cockroachdb/cockroach/blob/master/pkg/cli/start.go#L362

func runStart(cmd *cobra.Command, args []string, startSingleNode bool) error {
	const serverType redact.SafeString = "node"

	newServerFn := func(_ context.Context, serverCfg server.Config, stopper *stop.Stopper) (serverctl.ServerStartupInterface, error) {
		// Beware of not writing simply 'return server.NewServer()'. This is
		// because it would cause the serverctl.ServerStartupInterface reference to
		// always be non-nil, even if NewServer returns a nil pointer (and
		// an error). The code below is dependent on the interface
		// reference remaining nil in case of error.
		s, err := server.NewServer(serverCfg, stopper)
		if err != nil {
			return nil, err
		}
		return s, nil
	}

	return runStartInternal(cmd, serverType, serverCfg.InitNode, newServerFn, startSingleNode)
}

*/
