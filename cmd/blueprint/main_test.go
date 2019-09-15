package main

import (
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Example test to show usage of `make test`
func TestDummy(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestLoadConfig(t *testing.T) {
	if err := config.LoadConfig("/config"); err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "Dummy Value", config.Config.ConfigVar)
}
