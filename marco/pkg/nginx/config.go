package nginx

// KVDirect represent ngin key value directive
type KVDirect = map[string]interface{}

// Events nginx events directive
type Events struct {
	WorkerConnections int    `kv:"worker_connections"`
	Use               string `kv:"use"`          //epoll
	MultiAccept       bool   `kv:"multi_accept"` //default on
}

//MimeTypes for nginx config
type MimeTypes struct {
	Types KVDirect
}

var (
	defaultProxyConfig = KVDirect{
		"proxy_redirect":          "off",
		"proxy_set_header":        "Host $host",
		"proxy_set_header":        "X-Real-IP $remote_addr",
		"proxy_set_header":        "X-Forwarded-For $proxy_add_x_forwarded_for",
		"client_max_body_size":    "10m",
		"client_body_buffer_size": "128k",
		"proxy_connect_timeout":   "90",
		"proxy_send_timeout":      "90",
		"proxy_read_timeout":      "90",
		"proxy_buffers":           "32 4k",
	}
)

//HTTP nginx http config section
type HTTP struct {
}

//Config represent nginx config
// follows https://www.nginx.com/resources/wiki/start/topics/examples/full/ to build nginx base config
type Config struct {
	User            string `kv:"user"`
	WorkerProcesses string `kv:"worker_processes"`
	PId             string `kv:"pid"`
	ErrorLog        string `kv:"error_log,omitempty"`
	LimitNofile     int    `kv:"worker_rlimit_nofile"`
	ExtConfig       map[string]interface{}
}
