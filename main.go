package main

import (
	"flag"

	"github.com/deane/go-distance/api"
)

func main() {
	var host string

	flag.StringVar(
		&host,
		"host",
		"127.0.0.1:8000",
		"host and port the service should run on. example: '-host=127.0.0.1:8000'",
	)
	flag.Parse()

	api.RunServer(host)
}
