package injector

import (
	"casicloud.com/ylops/backend/app/router"
	"casicloud.com/ylops/backend/config"
	"github.com/gin-gonic/gin"
)

// InitGinEngine 初始化gin引擎
func InitGinEngine(r router.IRouter) *gin.Engine {
	gin.SetMode(config.C.RunMode)
	app := gin.New()
	// 注册路由
	r.Register(app)
	return app
}
