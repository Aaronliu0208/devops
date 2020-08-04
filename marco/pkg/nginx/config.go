package nginx

// Events nginx events directive
type Events struct {
	WorkerConnections int    `kv:"worker_connections" json:"worker_connections"`
	Use               string `kv:"use" json:"use,omitempty"`                   //epoll
	MultiAccept       bool   `kv:"multi_accept" json:"multi_accept,omitempty"` //default on
}

//HTTP nginx http config section
type HTTP struct {
	DefalutType string  `kv:"default_type,omitempty"`
	LogFormat   string  `kv:"log_format"`
	MimeTypes   Options `kv:"types"`
	AccessLog   string  `kv:"access_log,omitempty"`
	ErrorLog    string  `kv:"error_log,omitempty"`
	SendFile    bool    `kv:"sendfile"`
	ExtConfig   Options
}

//Config represent nginx config
// follows https://www.nginx.com/resources/wiki/start/topics/examples/full/ to build nginx base config
type Config struct {
	User            string `kv:"user,omitempty" json:"user,omitempty"`
	WorkerProcesses string `kv:"worker_processes"`
	PId             string `kv:"pid"`
	ErrorLog        string `kv:"error_log,omitempty"`
	LimitNofile     int    `kv:"worker_rlimit_nofile"`
	Events          Events `kv:"events"`
	HTTP            HTTP   `kv:"http"`
	Extras          Options
}
