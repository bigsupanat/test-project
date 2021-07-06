package main

import "github.com/bigsupanat/test-project/server"

func main() {
	serv := server.NewServer()
	serv.Start()
}
