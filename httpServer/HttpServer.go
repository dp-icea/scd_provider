package httpServer

import "log"

type HttpServer struct{}

func (h HttpServer) Serve() {
	log.Println("Hello World!")
}
