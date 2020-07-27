package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNginxControllerStart(t *testing.T) {
	controller := &NginxController{
		Prefix:     "test",
		ConfigFile: "test/nginx.conf",
		PidFile:    "test/logs/nginx.pid",
	}

	controller.Stop()
	err := controller.Start()
	if err != nil {
		t.Log(err)
	}

	ok, err := controller.Test()
	if err != nil {
		t.Log(err)
	}

	assert.True(t, ok)

	err = controller.Reload()
	if err != nil {
		t.Log(err)
	}
	controller.Stop()
}
