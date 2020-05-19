package model

import (
	"context"

	"casicloud.com/ylops/backend/internal/app/model"
	"casicloud.com/ylops/backend/internal/app/model/impl/gorm/entity"
	"casicloud.com/ylops/backend/internal/app/schema"
	"casicloud.com/ylops/backend/pkg/errors"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var _ model.IAlert = (*Alert)(nil)

// AlertSet 注入Alert
var AlertSet = wire.NewSet(wire.Struct(new(Alert), "*"), wire.Bind(new(model.IAlert), new(*Alert)))

// Alert 示例存储
type Alert struct {
	DB *gorm.DB
}

func (a *Alert) getQueryOption(opts ...schema.AlertQueryOptions) schema.AlertQueryOptions {
	var opt schema.AlertQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Alert) Query(ctx context.Context, params schema.AlertQueryParam, opts ...schema.AlertQueryOptions) (*schema.AlertQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetAlertDB(ctx, a.DB)
	if v := params.RawID; v != "" {
		db = db.Where("code=?", v)
	}

	if v := params.Source; v != "" {
		db = db.Where("source=?", v)
	}

	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("raw_id LIKE ? OR name LIKE ? OR source LIKE ?", v, v, v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.Alerts
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.AlertQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaAlerts(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Alert) Get(ctx context.Context, id string, opts ...schema.AlertQueryOptions) (*schema.Alert, error) {
	db := entity.GetAlertDB(ctx, a.DB).Where("id=?", id)
	var item entity.Alert
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaAlert(), nil
}

// Create 创建数据
func (a *Alert) Create(ctx context.Context, item schema.Alert) error {
	eitem := entity.SchemaAlert(item).ToAlert()
	result := entity.GetAlertDB(ctx, a.DB).Create(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更新数据
func (a *Alert) Update(ctx context.Context, id string, item schema.Alert) error {
	eitem := entity.SchemaAlert(item).ToAlert()
	result := entity.GetAlertDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete 删除数据
func (a *Alert) Delete(ctx context.Context, id string) error {
	result := entity.GetAlertDB(ctx, a.DB).Where("id=?", id).Delete(entity.Alert{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
