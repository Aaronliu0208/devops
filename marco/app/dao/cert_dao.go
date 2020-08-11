package dao

import (
	"context"

	"casicloud.com/ylops/marco/app/entity"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type CertificateDao struct {
	db *gorm.DB
}

// NewCertificateDao 工厂函数
func NewCertificateDao(DB *gorm.DB) *CertificateDao {
	return &CertificateDao{db: DB}
}

// GetCertDB 获取特定模型的数据库连接
func (c *CertificateDao) GetCertDB(ctx context.Context) *gorm.DB {
	return c.db
}

// Create 插入一条数据记录
func (c *CertificateDao) Create(ctx context.Context, obj *entity.Certificate) error {
	if result := c.db.Create(obj); result.Error != nil {
		return errors.WithMessage(result.Error, "插入数据失败")
	}
	return nil
}
