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
func (s Site) MarshalD() ([]nginx.Directive, error) {
	serverBlk := nginx.NewBlock("server")
	if len(s.Root) > 0 {
		serverBlk.AddKVOption("root", s.Root)
	}

	if s.EnableSSL {

	} else {
		serverBlk.AddKVOption("listen", s.Port)
	}

	if len(s.Domain) > 0 {
		serverBlk.AddKVOption("server_name", s.Domain)
	}

	if len(s.AccessLog) > 0 {
		serverBlk.AddKVOption("access_log", []string{s.AccessLog, "main"})
	}

	if len(s.ErrorLog) > 0 {
		serverBlk.AddKVOption("error_log", []string{s.ErrorLog, "warn"})
	}

	for _, r := range s.Routes {
		serverBlk.AddInterface(r)
	}
	serverBlk.AddInterface(s.Extras)

	return []nginx.Directive{serverBlk}, nil
}
