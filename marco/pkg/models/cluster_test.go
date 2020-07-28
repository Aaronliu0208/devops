package models

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"casicloud.com/ylops/marco/pkg/nginx"
	"github.com/stretchr/testify/assert"
)

func TestClusterGenerator(t *testing.T) {
	prefix, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	logdir := filepath.Join(prefix, "logs")

	config := nginx.NewDefaultRestyConfig(prefix, logdir, "")

	s := Site{
		Domain:    "www.baidu.com",
		EnableSSL: false,
		Port:      80,
		Routes: []Route{
			{
				Pattern: "",
				Path:    "/",
				Extras: []nginx.Pair{
					{Key: "proxy_pass", Value: "http://www.baidu.com"},
				},
			},
		},
		Root:      "html/site1",
		AccessLog: "logs/access.log",
	}
	cluster := &Cluster{
		Config: config,
		Sites: []Site{
			s,
		},
	}

	nginxConf, err := cluster.GenerateConfig()
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, strings.Contains(nginxConf, "http {"))
	assert.True(t, strings.Contains(nginxConf, "types {"))
	assert.True(t, strings.Contains(nginxConf, "events {"))
	assert.True(t, strings.Contains(nginxConf, "worker_processes"))
	assert.True(t, strings.Contains(nginxConf, "proxy_pass http://www.baidu.com;"))
}
