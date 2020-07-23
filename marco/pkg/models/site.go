package models

import (
	"time"

	"casicloud.com/ylops/marco/pkg/nginx"
)

//Site represent nginx server config
type Site struct {
	ID        string
	CreatedAt time.Time `json:"created_at"  kv:"-"`
	UpdatedAt time.Time `json:"updated_at" kv:"-"`
	Port      int
	Domain    string
	EnableSSL bool
	Cert      Certificate
	Routes    []Route
	Root      string
	AccessLog string
	ErrorLog  string
	Extras    nginx.Options
}

//Marshal implements directive Marshaler
func (s Site) Marshal() ([]nginx.Directive, error) {
	return nil, nil
}
