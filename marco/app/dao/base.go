package dao

import (
	"strings"
	"time"

	"casicloud.com/ylops/marco/pkg/log"
	"github.com/jinzhu/gorm"
)

type UpdateFunc func(*gorm.DB) error
type Model interface{}

type DBMigrator struct {
	model Model
	f     UpdateFunc
}

var migrators []*DBMigrator

// Config 配置参数
type Config struct {
	Debug        bool
	DBType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TablePrefix  string
}

// NewDB 创建DB实例
func NewDB(c *Config) (*gorm.DB, func(), error) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return c.TablePrefix + defaultTableName
	}
	db, err := gorm.Open(c.DBType, c.DSN)
	if err != nil {
		return nil, nil, err
	}

	if strings.ToLower(c.DBType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB  CHARSET=utf8mb4")
	}

	if strings.ToLower(c.DBType) == "mysql" || strings.ToLower(c.DBType) == "postgres" {
		db = db.Set("gorm:query_option", "FOR UPDATE")
	}
	if c.Debug {
		db = db.Debug()
	}

	cleanFunc := func() {
		err := db.Close()
		if err != nil {
			log.Errorf("Gorm db close error: %s", err.Error())
		}
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, cleanFunc, err
	}

	db.SingularTable(true)
	if c.MaxIdleConns > 0 {
		db.DB().SetMaxIdleConns(c.MaxIdleConns)
	}

	if c.MaxOpenConns > 0 {
		db.DB().SetMaxOpenConns(c.MaxOpenConns)
	}
	if c.MaxLifetime > 0 {
		db.DB().SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)
	}

	return db, cleanFunc, nil
}

func RegistryMigrater(m interface{}, f UpdateFunc) {
	migrators = append(migrators, &DBMigrator{
		model: m,
		f:     f,
	})
}

func MigerateDB(db *gorm.DB) {
	for _, m := range migrators {
		tx := db.Begin()
		db.AutoMigrate(m.model)
		err := m.f(db)
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}
}
