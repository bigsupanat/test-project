package server

import (
	"github.com/bigsupanat/test-project/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	service service.Service
}

func NewServer(svc service.Service) *Server {
	r := newGin(svc)

	return &Server{
		router:  r,
		service: svc,
	}
}

func newGin(svc service.Service) *gin.Engine {
	r := gin.Default()

	r.GET("/covid/summary", getCovidSummaryFn(svc))

	return r
}

func (s *Server) Start() {
	s.router.Run()
}

//TODO: server Stop with gracefully shutdown
