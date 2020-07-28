package nginx

import (
	"fmt"
	"testing"
)

func TestDefaultRestyConfig(t *testing.T) {
	prefix := "/home/shanyou/macro/app"
	logpath := "/home/shanyou/macro/logs"
	libPath := "/home/shanyou/macro/app/lib"
	conf := NewDefaultRestyConfig(prefix, logpath, libPath)
	emptyBlk := NewEmptyBlock()
	emptyBlk.AddInterface(conf)
	fmt.Println(emptyBlk)
}
