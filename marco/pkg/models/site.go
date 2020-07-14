package models

//Site represent nginx server config
type Site struct {
	Domain    string      `kv:"server_name"`
	EnableSSL bool        `kv:"-"`
	Cert      Certificate `kv:"-"`
	Routes    []Route
}
