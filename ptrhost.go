package main

import (
	"net"
	"os"

	flags "github.com/jessevdk/go-flags"
)

func main() {
	parser := flags.NewParser(nil, flags.Default)
	parser.Usage = "HOSTNAME [OPTIONS]"
	args, _ := parser.Parse()

	if len(args) == 0 {
		os.Exit(1)
	}

	hostname := args[0]
	addr, _ := net.LookupHost(hostname)
	for _, v := range addr {
		ptrAddr, _ := net.LookupAddr(v)
		println(v, "->", ptrAddr[0])
	}
}
