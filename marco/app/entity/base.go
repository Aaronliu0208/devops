package entity

type Base struct {
	ID        int64 `gorm:"column:id;type:bigint(20) auto_increment;primary_key"`
	CreatedAt XTime `gorm:"column:created_at;index;"`
	UpdatedAt XTime `gorm:"column:updated_at;index;"`
	DeletedAt XTime `gorm:"column:deleted_at;index;"`
	Deleted   bool  `gorm:"column:deleted;index;type:tinyint(1);not null;default:0"`
}
