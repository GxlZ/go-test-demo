// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 7dc4d5d85c
// Version Date: 2018年 5月28日 星期一 22时12分59秒 UTC

package main

import (
	"flag"

	// This Service
	"go-test-demo/go-kit-truss/svc/server"
	"go-test-demo/go-kit-truss/svc/server/cli"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	server.Run(cli.Config)
}
