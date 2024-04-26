module github.com/bwlee13/gopherdb/cli

go 1.22.0

require github.com/spf13/cobra v1.8.0

require github.com/pkg/errors v0.9.1 // indirect

require (
	github.com/bwlee13/gopherdb v0.0.0-20240424144325-b4c2ae129d64
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
replace github.com/bwlee13/gopherdb => ../
