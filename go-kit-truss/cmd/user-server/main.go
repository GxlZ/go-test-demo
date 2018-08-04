package main

import (
	"flag"

	// This Service
	"go-test-demo/go-kit-truss/svc/server"
	"go-test-demo/go-kit-truss/svc/server/cli"

	_ "go-test-demo/go-kit-truss/di"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	server.Run(cli.Config)
}
