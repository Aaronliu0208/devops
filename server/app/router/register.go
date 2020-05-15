package router

import (
	"casicloud.com/ylops/backend/app/apis"
	"github.com/gin-gonic/gin"
)

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")
	v1 := g.Group("/v1")
	{
		v1.GET("ping", apis.Ping)
	}
}
