package main

import (
	"fmt"
	"os"

	"://github.com"
	svccmd "://github.com/cmd"
	"mycustoml1/app"
)

func main() {
	rootCmd, _ := server.NewRootCmd("mycustoml1d", "l1", "/.mycustoml1", "custom-chain-id", app.ModuleBasics)
	if err := svccmd.Execute(rootCmd, "MYCUSTOML1", app.DefaultNodeHome); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing blockchain daemon: %v\n", err)
		os.Exit(1)
	}
}
