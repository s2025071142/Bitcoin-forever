package main

import (
	"fmt"
	"os"

	"://github.com"
	svccmd "://github.com/cmd"
	"://github.com"
)

func main() {
	rootCmd, _ := server.NewRootCmd("mycustoml1d", "l1", "/.mycustoml1", "custom-chain-id", main.ModuleBasics)
	if err := svccmd.Execute(rootCmd, "MYCUSTOML1", main.DefaultNodeHome); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing blockchain daemon: %v\n", err)
		os.Exit(1)
	}
}
