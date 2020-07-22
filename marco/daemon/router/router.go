package router

import "github.com/gin-gonic/gin"

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}
