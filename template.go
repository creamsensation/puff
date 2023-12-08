package main

import "fmt"

func createComponentTemplate(packageName, structName, name string) string {
	return fmt.Sprintf(
		`package %[1]s

import (
	"github.com/creamsensation/cp"
	. "github.com/creamsensation/gox"
)

type %[2]s struct {
	cp.Component
}

func (m *%[2]s) Name() string {
	return "%[3]s"
}

func (m *%[2]s) Model() {}

func (m *%[2]s) Node() Node {
	return Div(Text("%[2]s"))
}
`, packageName, structName, name,
	)
}

func createControllerTemplate(packageName, structName, name string) string {
	return fmt.Sprintf(
		`package %[1]s

import (
	. "github.com/creamsensation/cp"
)

type %[2]s struct {}

func (m *%[2]s) Name() string {
	return "%[3]s"
}

func (m *%[2]s) Routes() Routes {
	return Routes{}
}
`, packageName, structName, name,
	)
}

func createModuleTemplate(packageName, structName, name string) string {
	return fmt.Sprintf(
		`package %[1]s

import (
	. "github.com/creamsensation/cp"
)

type %[2]s struct {}

func (m *%[2]s) Name() string {
	return "%[3]s"
}

func (m *%[2]s) Controllers() Controllers {
	return Controllers{}
}
`, packageName, structName, name,
	)
}
