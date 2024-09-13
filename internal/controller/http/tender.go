package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/entity"
)

func (c *Controller) createTender(ctx *gin.Context) {
	var tender entity.Tender

	if err := ctx.ShouldBindJSON(&tender); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	ctx.Status(http.StatusCreated)
}
