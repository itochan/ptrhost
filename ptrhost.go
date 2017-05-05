package main

import (
	"flag"
	"net"
)

func main() {
	flag.Parse()
	hostname := flag.Args()[0]

	addr, _ := net.LookupHost(hostname)
	for _, v := range addr {
		ptrAddr, _ := net.LookupAddr(v)
		println(v, "->", ptrAddr[0])
	}
}
