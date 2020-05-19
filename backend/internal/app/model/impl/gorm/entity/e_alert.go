package entity

import (
	"context"
	"fmt"
	"time"

	"casicloud.com/ylops/backend/internal/app/config"
	"casicloud.com/ylops/backend/internal/app/schema"
	"casicloud.com/ylops/backend/pkg/util"
	"github.com/jinzhu/gorm"
)

// GetAlertDB 获取Alert存储
func GetAlertDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Alert))
}

// SchemaAlert Alert对象
type SchemaAlert schema.Alert

// ToAlert 转换为Alert实体
func (a SchemaAlert) ToAlert() *Alert {
	item := new(Alert)
	util.StructMapToStruct(a, item)
	return item
}

// Alert Alert实体
type Alert struct {
	ID          string    `gorm:"column:id;primary_key;size:36;"`
	StartsAt    time.Time `gorm:"column:start_at;index;"`
	EndsAt      time.Time `gorm:"column:end_at;index;"`
	Source      string    `gorm:"column:source;size:50;index;default:'';not null;"` // 编号
	Name        string    `gorm:"column:name;size:100;index;default:'';not null;"`  // 名称
	Description string    `gorm:"column:desc;size:200;"`                            // 备注
	Raw         string    `gorm:"column:raw;size:200;"`                             // 原始信息
	State       string    `gorm:"column:state;size:200;"`                           //报警状态
	RawID       string    `gorm:"column:raw_id;size:36;"`
}

// TableName 表名
func (a Alert) TableName() string {
	return fmt.Sprintf("%s%s", config.C.Gorm.TablePrefix, "Alert")
}

// ToSchemaAlert 转换为Alert对象
func (a Alert) ToSchemaAlert() *schema.Alert {
	item := new(schema.Alert)
	util.StructMapToStruct(a, item)
	return item
}

// Alerts Alert列表
type Alerts []*Alert

// ToSchemaAlerts 转换为Alert对象列表
func (a Alerts) ToSchemaAlerts() []*schema.Alert {
	list := make([]*schema.Alert, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaAlert()
	}
	return list
}
