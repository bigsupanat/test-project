package main

import (
	"github.com/bigsupanat/test-project/internal/client"
	"github.com/bigsupanat/test-project/server"
	"github.com/bigsupanat/test-project/service"
)

func main() {
	clientInternal := client.Init()

	svc := service.Service{
		RequestFn: clientInternal.Request,
	}

	serv := server.NewServer(svc)
	serv.Start()
}
