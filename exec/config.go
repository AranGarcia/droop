package main

import "flag"

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "the server address (host and port)")
	flag.Parse()
}
