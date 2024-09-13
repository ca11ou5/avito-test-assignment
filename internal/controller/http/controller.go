package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface{}

type Controller struct {
	Service
}

func NewController(svc Service) *Controller {
	return &Controller{
		Service: svc,
	}
}

func (c *Controller) ping(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
