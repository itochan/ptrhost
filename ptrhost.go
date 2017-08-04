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
	args, err := parser.Parse()

	if err != nil {
		if ferr, ok := err.(*flags.Error); ok && ferr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

	if len(args) == 0 {
		if opts.Version {
			fmt.Println("ptrhost version", version)
			os.Exit(0)
		}

		parser.WriteHelp(os.Stderr)
		os.Exit(0)
	}

	hostname := args[0]
	lookupHosts, _ := net.LookupHost(hostname)
	for _, host := range lookupHosts {
		lookupAddrs, err := net.LookupAddr(host)
		if err == nil {
			fmt.Println(host, "->", lookupAddrs[0])
		} else {
			fmt.Println(host, "->", err.(*net.DNSError).Err)
		}
	}
}
