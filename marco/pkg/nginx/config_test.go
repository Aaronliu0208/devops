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
