package main

import (
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDummy(t *testing.T) {
	assert.Equal(t, config.Config.ConfigVar, "Dummy Value")
}
