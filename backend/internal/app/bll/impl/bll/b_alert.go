package bll

import (
	"context"

	"casicloud.com/ylops/backend/internal/app/bll"
	"casicloud.com/ylops/backend/internal/app/iutil"
	"casicloud.com/ylops/backend/internal/app/model"
	"casicloud.com/ylops/backend/internal/app/schema"
	"casicloud.com/ylops/backend/pkg/errors"
	"github.com/google/wire"
)

var _ bll.IAlert = (*Alert)(nil)

// AlertSet 注入Alert
var AlertSet = wire.NewSet(wire.Struct(new(Alert), "*"), wire.Bind(new(bll.IAlert), new(*Alert)))

// Alert 示例程序
type Alert struct {
	AlertModel model.IAlert
}

// Query 查询数据
func (a *Alert) Query(ctx context.Context, params schema.AlertQueryParam, opts ...schema.AlertQueryOptions) (*schema.AlertQueryResult, error) {
	return a.AlertModel.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *Alert) Get(ctx context.Context, id string, opts ...schema.AlertQueryOptions) (*schema.Alert, error) {
	item, err := a.AlertModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *Alert) checkRaw(ctx context.Context, rawid string, source string) error {
	result, err := a.AlertModel.Query(ctx, schema.AlertQueryParam{
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		},
		RawID:  rawid,
		Source: source,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("编号已经存在")
	}

	return nil
}

// Create 创建数据
func (a *Alert) Create(ctx context.Context, item schema.Alert) (*schema.IDResult, error) {
	err := a.checkRaw(ctx, item.RawID, item.Source)
	if err != nil {
		return nil, err
	}

	item.ID = iutil.NewID()
	err = a.AlertModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

// Update 更新数据
func (a *Alert) Update(ctx context.Context, id string, item schema.Alert) error {
	oldItem, err := a.AlertModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.RawID != item.RawID && oldItem.Source != item.Source {
		if err := a.checkRaw(ctx, item.RawID, item.Source); err != nil {
			return err
		}
	}
	item.ID = oldItem.ID

	return a.AlertModel.Update(ctx, id, item)
}

// Delete 删除数据
func (a *Alert) Delete(ctx context.Context, id string) error {
	oldItem, err := a.AlertModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.AlertModel.Delete(ctx, id)
}
