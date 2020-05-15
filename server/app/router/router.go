package router

import "github.com/gin-gonic/gin"

var _ IRouter = (*Router)(nil) //告送编译器静态检查Router是否实现接口IRouter, 下划线告诉编译器要有效地放弃RHS值

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

// Router 路由管理器
type Router struct {
}

// Register 注册路由
func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

// Prefixes 路由前缀列表
func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}
