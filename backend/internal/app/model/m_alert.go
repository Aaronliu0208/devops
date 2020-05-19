package model

import (
	"context"

	"casicloud.com/ylops/backend/internal/app/schema"
)

// IAlert 告警信息存储接口
type IAlert interface {
	// 查询数据
	Query(ctx context.Context, params schema.AlertQueryParam, opts ...schema.AlertQueryOptions) (*schema.AlertQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, id string, opts ...schema.AlertQueryOptions) (*schema.Alert, error)
	// 创建数据
	Create(ctx context.Context, item schema.Alert) error
	// 更新数据
	Update(ctx context.Context, id string, item schema.Alert) error
	// 删除数据
	Delete(ctx context.Context, id string) error
}
