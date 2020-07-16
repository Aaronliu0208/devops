package logger

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	SetLevel(int(logrus.DebugLevel))
	var b bytes.Buffer
	SetOutput(io.Writer(&b))

	logger := Get("test", nil)
	logger.Debug("test log")

	assert.True(t, strings.Contains(b.String(), "test log"))
	assert.True(t, strings.Contains(b.String(), "name=test"))
	assert.True(t, strings.Contains(b.String(), "hostname="))
}
