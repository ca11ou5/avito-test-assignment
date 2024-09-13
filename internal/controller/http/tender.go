package http

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/controller/http/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) newTender(ctx *gin.Context) {
	var tender model.NewTenderRequest

	if err := ctx.ShouldBindJSON(&tender); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	tenderEntity := model.NewTenderEntity(&tender)

	tenderEntity, err := c.TenderService.CreateTender(tenderEntity)
	if err != nil {

	}

	ctx.Status(http.StatusCreated)
}
