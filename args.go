package main

import "os"

const (
	actionGenerate         = "generate"
	actionGenerateShortcut = "g"
	
	actionHelp         = "help"
	actionHelpShortcut = "h"
)

const (
	categoryComponent         = "component"
	categoryComponentShortcut = "c"
	
	categoryController         = "controller"
	categoryControllerShortcut = "r"
	
	categoryModule         = "module"
	categoryModuleShortcut = "m"
)

func parseArgs() (string, string, string) {
	var action, category, path string
	for i, arg := range os.Args {
		if i == 1 {
			action = arg
		}
		if i == 2 {
			category = arg
		}
		if i == 3 {
			path = arg
		}
	}
	return action, category, path
}
