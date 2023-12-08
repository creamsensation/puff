package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	action, category, path := parseArgs()
	if !strings.HasSuffix(path, ".go") {
		path = path + ".go"
	}
	path = wd + "/" + path
	switch action {
	case actionHelp, actionHelpShortcut:
		showHelp()
	case actionGenerate, actionGenerateShortcut:
		if len(path) == 0 || path == "/" {
			log.Fatalln("path is empty")
		}
		generate(category, path)
	default:
		fmt.Printf("unknown action: %s \n", action)
		fmt.Println("----------")
		showHelp()
	}
}
