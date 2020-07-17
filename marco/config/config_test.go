package config

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigLoadString(t *testing.T) {

	data, err := ioutil.ReadFile("config_example.yaml")
	if err != nil {
		t.Fatal(err)
	}
	config, err := LoadConfig(data)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, config)
	assert.True(t, strings.Compare(config.Mode, "debug") == 0)
}

func TestConfigLoadFile(t *testing.T) {
	configFile := "config_example.yaml"

	config, err := LoadConfigFile(configFile)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, config)
	assert.True(t, strings.Compare(config.Mode, "debug") == 0)
}
