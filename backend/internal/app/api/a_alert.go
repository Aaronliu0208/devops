package api

import (
	"casicloud.com/ylops/backend/internal/app/bll"
	"casicloud.com/ylops/backend/internal/app/ginplus"
	"casicloud.com/ylops/backend/internal/app/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// AlertSet 注入Alert
var AlertSet = wire.NewSet(wire.Struct(new(Alert), "*"))

// Alert 示例程序
type Alert struct {
	AlertBll bll.IAlert
}

// Query 查询数据
func (a *Alert) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.AlertQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.AlertBll.Query(ctx, params)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *Alert) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.AlertBll.Get(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}

// Create 创建数据
func (a *Alert) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Alert
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	result, err := a.AlertBll.Create(ctx, item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, result)
}

// Update 更新数据
func (a *Alert) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Alert
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	err := a.AlertBll.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Delete 删除数据
func (a *Alert) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.AlertBll.Delete(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
