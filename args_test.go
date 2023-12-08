package main

import (
	"os"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestArgs(t *testing.T) {
	oldArgs := os.Args
	os.Args = []string{"", "g", "r", "test.go"}
	action, category, path := parseArgs()
	assert.Equal(t, "g", action)
	assert.Equal(t, "r", category)
	assert.Equal(t, "test.go", path)
	os.Args = oldArgs
}
