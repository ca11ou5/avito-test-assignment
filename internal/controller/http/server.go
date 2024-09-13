package http

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/config"
)

func (c *Controller) StartHTTPListening(cfg *config.Config) error {
	r := gin.Default()

	c.initRoutes(r)

	return r.Run(cfg.PostgresPort)
}

func (c *Controller) initRoutes(r *gin.Engine) {
	api := r.Group("api")

	tenderGroup := api.Group("tenders")
	{
		tenderGroup.POST("new", c.newTender)
	}

	api.GET("/ping", ping)

}

func ping(c *gin.Context) {
	c.Status(http.StatusOK)
}
