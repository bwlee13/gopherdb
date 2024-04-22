package server

// use below for CLI future work
// "github.com/spf13/cobra"

func main() {

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
