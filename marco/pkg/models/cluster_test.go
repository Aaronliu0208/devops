package models

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"casicloud.com/ylops/marco/pkg/nginx"
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
		Extras: nginx.Options{
			{"shanyou", "great"},
		},
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

	fmt.Print(nginxConf)
}
