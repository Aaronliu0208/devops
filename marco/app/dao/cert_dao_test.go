package dao

import (
	"testing"

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
}
