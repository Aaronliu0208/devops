package models

import (
	"strings"
	"testing"

	"casicloud.com/ylops/marco/pkg/nginx"
	"github.com/stretchr/testify/assert"
)

func TestRouteMarshal(t *testing.T) {
	r := &Route{
		Pattern: "",
		Path:    "/",
		Extras: []nginx.Pair{
			{Key: "echo", Value: "hello"},
		},
	}
	emptyBlk := nginx.NewEmptyBlock()
	emptyBlk.AddInterface(r)
	assert.True(t, strings.Contains(emptyBlk.String(), "location"))
	assert.True(t, strings.Contains(emptyBlk.String(), "/"))
	assert.True(t, strings.Contains(emptyBlk.String(), "echo hello;"))
}
