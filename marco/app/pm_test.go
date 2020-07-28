package app

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"casicloud.com/ylops/marco/config"
	"casicloud.com/ylops/marco/pkg/models"
	"casicloud.com/ylops/marco/pkg/nginx"
)

func TestPMStart(t *testing.T) {
	conf := &config.Config{
		Workspace: "/home/shanyou/src/ops/ylops/marco/bin",
	}
	conf.EnsureDirectoryExists()
	defer os.RemoveAll(conf.GetTempDir())
	pm := NewPM(conf)
	ngxConf := nginx.NewDefaultRestyConfig(conf.GetPrefix(), conf.GetLogDir(), "")
	s := models.Site{
		Domain:    "www.baidu.com",
		EnableSSL: false,
		Port:      8080,
		Routes: []models.Route{
			{
				Pattern: "",
				Path:    "/",
				Extras: []nginx.Pair{
					{Key: "proxy_pass", Value: "http://www.baidu.com"},
				},
			},
		},
		Root:      "html/site1",
		AccessLog: conf.GetLogDir() + "/www.baidu.com-access.log",
	}
	cluster := &models.Cluster{
		Config: ngxConf,
		Sites: []models.Site{
			s,
		},
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := pm.Start(cluster)
		if err != nil {
			fmt.Print(err)
		}
	}()
	wg.Wait()
}
