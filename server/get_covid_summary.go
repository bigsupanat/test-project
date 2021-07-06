package server

import (
	"net/http"

	"github.com/bigsupanat/test-project/handler"
	"github.com/bigsupanat/test-project/internal/client"
	"github.com/bigsupanat/test-project/service"
	"github.com/gin-gonic/gin"
)

func getCovidSummary(c *gin.Context) {

	cli := client.Init()

	res := handler.CovidSummary(cli)
	c.JSON(http.StatusOK, res)
}

func getCovidSummaryFn(svc service.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		res := handler.CovidSummary(svc)
		c.JSON(http.StatusOK, res)
	}
}
