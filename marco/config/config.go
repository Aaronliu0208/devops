package config

import "fmt"

// Config of marco
type Config struct {
	RunMode   string `yaml:"RunMode:omitempty"`
	Workspace string `yaml:"Workspace:omitempty"`
	HTTP      HTTP   `yaml:"http,omitempty"`
	Log       Log    `yaml:"log,omitempty"`
	Git       Git    `yaml:"git,omitempty"`
}

// HTTP config for http server
type HTTP struct {
}

// Log config for log
type Log struct {
}

// Git config for git
type Git struct {
}

// Postgres postgres配置参数
type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// DSN 数据库连接串
func (a Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		a.Host, a.Port, a.User, a.DBName, a.Password, a.SSLMode)
}
