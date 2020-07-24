package models

import (
	"fmt"
	"strings"
	"testing"

	"casicloud.com/ylops/marco/pkg/nginx"
	"github.com/stretchr/testify/assert"
)

func TestSiteMarshal(t *testing.T) {
	s := &Site{
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
	emptyBlk := nginx.NewEmptyBlock()
	emptyBlk.AddInterface(s)
	fmt.Print(emptyBlk.String())
	assert.True(t, strings.Contains(emptyBlk.String(), "location"))
	assert.True(t, strings.Contains(emptyBlk.String(), "/"))
	assert.True(t, strings.Contains(emptyBlk.String(), "proxy_pass http://www.baidu.com;"))
}
