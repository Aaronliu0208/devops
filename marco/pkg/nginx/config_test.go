package nginx

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	emptyBlk := NewEmptyBlock()
	emptyBlk.AddInterface(defaultConfig)
	fmt.Println(emptyBlk)
	assert.True(t, strings.Contains(emptyBlk.String(), "proxy_set_header Host $host;"))
	assert.True(t, strings.Contains(emptyBlk.String(), "proxy_set_header X-Real-IP $remote_addr;"))
}

func TestDefaultRestyConfig(t *testing.T) {
	prefix := "/home/shanyou/macro/app"
	logpath := "/home/shanyou/macro/logs"
	libPath := "/home/shanyou/macro/app/lib"
	conf := NewDefaultRestyConfig(prefix, logpath, libPath)
	emptyBlk := NewEmptyBlock()
	emptyBlk.AddInterface(conf)
	fmt.Println(emptyBlk)
}
