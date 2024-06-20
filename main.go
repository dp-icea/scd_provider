package main

import (
	"icea_uss/httpServer"
)

type Server interface {
	Serve()
}

func main() {
	var server Server = httpServer.HttpServer{}
	server.Serve()
}
