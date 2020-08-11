package entity

type ICertStore interface {
	Create(cert *Certificate) error
}
type Certificate struct {
	Base
	SNI     string `gorm:"column:sni;type:varchar(255)"`
	Content []byte `gorm:"column:content;type:blob"`
}
