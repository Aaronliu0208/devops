package models

import "time"

//Site represent nginx server config
type Site struct {
	ID        string
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Domain    string      `kv:"server_name"`
	EnableSSL bool        `kv:"-"`
	Cert      Certificate `kv:"-"`
	Routes    []Route
}
