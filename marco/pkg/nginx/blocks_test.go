package nginx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyBlock(t *testing.T) {
	emptyBlk := NewEmptyBlock()
	emptyBlk.AddKVOption("proxy_set_header", "Host $host")
	emptyBlk.AddKVOption("proxy_set_header", "X-Real-IP $remote_addr")
	blkStr := emptyBlk.String()
	assert.True(t, strings.Contains(blkStr, "proxy_set_header"))
	assert.True(t, strings.Contains(blkStr, "X-Real-IP $remote_addr"))
}
