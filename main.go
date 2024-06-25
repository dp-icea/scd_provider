package main

import (
	"icea_uss/httpServer"
	"icea_uss/scd"
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
