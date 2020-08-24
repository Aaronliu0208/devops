package entity

type ICertStore interface {
	Create(cert *Certificate) (*IDResult, error)
}
type Certificate struct {
	Base
	// 用户用来建立全局唯一ID
	GID     string `gorm:"column:gid;type:varchar(255);index;not null"`
	SNI     string `gorm:"column:sni;type:varchar(255)"`
	Content []byte `gorm:"column:content;type:blob"`
}

// TableName 返回数据库表名称
func (Certificate) TableName() string {
	return "certificate"
}

// CertQueryParam 查询条件
type CertQueryParam struct {
	PaginationParam
	SNI        string `form:"sni"`        //
	QueryValue string `form:"queryValue"` // 模糊查询
	Status     int    `form:"status"`     // 用户状态(1:启用 2:停用)
}

// CertQueryOptions 查询可选参数项
type CertQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// CertQueryResult 查询结果
type CertQueryResult struct {
	Data       []*Certificate
	PageResult *PaginationResult
}
