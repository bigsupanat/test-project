package server

import (
	"net/http"

	"github.com/bigsupanat/test-project/handler"
	"github.com/bigsupanat/test-project/service"
	"github.com/gin-gonic/gin"
)

func getCovidSummaryFn(svc service.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		res := handler.CovidSummary(svc)
		c.JSON(http.StatusOK, res)
	}
}
