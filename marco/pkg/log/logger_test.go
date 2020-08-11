package log

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	SetLevel(2)
	var b bytes.Buffer
	SetOutput(io.Writer(&b))
	Debug("hello")
	assert.True(t, strings.Contains(b.String(), "hello"))
}

func TestNotPrint(t *testing.T) {
	var b bytes.Buffer
	SetOutput(io.Writer(&b))
	Debug("hello")
	assert.False(t, strings.Contains(b.String(), "hello"))
}
