package main

import (
	"fmt"
	"net"
	"os"

	flags "github.com/jessevdk/go-flags"
)

const version = "1.0.1"

var opts struct {
	Version bool `short:"v" long:"version" description:"Show version"`
}

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Usage = "HOSTNAME [OPTIONS]"
	args, _ := parser.Parse()

	if len(args) == 0 {
		if opts.Version {
			fmt.Println("ptrhost version", version)
		}
		os.Exit(1)
	}

	hostname := args[0]
	addr, _ := net.LookupHost(hostname)
	for _, v := range addr {
		resolvedHost, err := net.LookupAddr(v)
		if err == nil {
			fmt.Println(v, "->", resolvedHost[0])
		} else {
			fmt.Println(v, "->", err.(*net.DNSError).Err)
		}
	}
}
