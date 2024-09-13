package http

import (
	"github.com/gin-gonic/gin"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/config"
)

func (c *Controller) StartHTTPListening(cfg *config.Config) error {
	r := gin.Default()

	c.initRoutes(r)

	return r.Run(":8080")
}

func (c *Controller) initRoutes(r *gin.Engine) {
	api := r.Group("api")

	tenderGroup := api.Group("tenders")
	{
		tenderGroup.POST("new", c.createTender)
	}

	api.GET("/ping", c.ping)

}
