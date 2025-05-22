package main

import (
	"scd_provider/httpServer"
	"scd_provider/scd"
)

type Server interface {
	Serve()
}

func main() {
	var server Server = httpServer.HttpServer{
		Deconflictor: scd.InterussDeconflictor{},
	}
	server.Serve()
}
