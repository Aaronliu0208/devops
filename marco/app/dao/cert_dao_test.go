package dao

import (
	"testing"

	"casicloud.com/ylops/marco/app/entity"
	"casicloud.com/ylops/marco/pkg/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCertCreateDao(t *testing.T) {
	conf := getTestConfig()
	db, f, err := NewDB(conf)
	defer f()
	if err != nil {
		t.Fatal(err)
	}
	dao := NewCertificateDao(db)
	assert.NotNil(t, dao)

	MigerateDB(db)

	err = dao.Create(&entity.Certificate{
		GID:     utils.NewID(),
		SNI:     "*.casicloud.com",
		Content: []byte("证书"),
	})
	if err != nil {
		t.Fatal(err)
	}
}
