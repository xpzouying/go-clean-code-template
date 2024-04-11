package config

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {

	data := json.RawMessage(`{
    "listen": "127.0.0.1:9090",
    "log_level": "debug"
}`)

	cfg, err := NewConfig(bytes.NewReader(data))
	assert.NoError(t, err)

	want := Config{
		ListenAddr: "127.0.0.1:9090",
		LogLevel:   "debug",
	}
	assert.Equal(t, want, *cfg)
}
