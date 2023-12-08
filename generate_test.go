package main

import (
	"fmt"
	"os"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	dir := t.TempDir()
	t.Run(
		"generate component", func(t *testing.T) {
			componentDir := fmt.Sprintf("%s/test/test_component", dir)
			generate(categoryComponent, componentDir)
			fb, err := os.ReadFile(componentDir)
			assert.Nil(t, err)
			assert.Equal(t, createComponentTemplate("test", "TestComponent", "test"), string(fb))
		},
	)
	t.Run(
		"generate controller", func(t *testing.T) {
			controllerDir := fmt.Sprintf("%s/test/test_controller", dir)
			generate(categoryController, controllerDir)
			fb, err := os.ReadFile(controllerDir)
			assert.Nil(t, err)
			assert.Equal(t, createControllerTemplate("test", "TestController", "test"), string(fb))
		},
	)
	t.Run(
		"generate module", func(t *testing.T) {
			moduleDir := fmt.Sprintf("%s/test/test_module", dir)
			generate(categoryModule, moduleDir)
			fb, err := os.ReadFile(moduleDir)
			assert.Nil(t, err)
			assert.Equal(t, createModuleTemplate("test", "TestModule", "test"), string(fb))
		},
	)
}
