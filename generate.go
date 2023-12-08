package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	
	"github.com/iancoleman/strcase"
)

const (
	componentCamelSuffix  = "Component"
	controllerCamelSuffix = "Controller"
	moduleCamelSuffix     = "Module"
	
	componentKebabSuffix  = "-component"
	controllerKebabSuffix = "-controller"
	moduleKebabSuffix     = "-module"
)

func generate(category, path string) {
	var packageName, structName, name string
	parts := strings.Split(path, "/")
	partsN := len(parts)
	if partsN == 0 {
		log.Fatalln("invalid generate path")
	}
	for i, part := range parts {
		if partsN > 1 && i == partsN-2 {
			packageName = part
		}
		if partsN > 0 && i == partsN-1 {
			name = strcase.ToKebab(strings.TrimSuffix(part, ".go"))
			structName = strcase.ToCamel(name)
		}
	}
	switch category {
	case categoryComponent, categoryComponentShortcut:
		if !strings.HasSuffix(structName, componentCamelSuffix) {
			structName += componentCamelSuffix
		}
		if strings.HasSuffix(name, componentKebabSuffix) {
			name = strings.TrimSuffix(name, componentKebabSuffix)
		}
		generateComponent(packageName, structName, name, path)
	case categoryController, categoryControllerShortcut:
		if !strings.HasSuffix(structName, controllerCamelSuffix) {
			structName += controllerCamelSuffix
		}
		if strings.HasSuffix(name, controllerKebabSuffix) {
			name = strings.TrimSuffix(name, controllerKebabSuffix)
		}
		generateController(packageName, structName, name, path)
	case categoryModule, categoryModuleShortcut:
		if !strings.HasSuffix(structName, moduleCamelSuffix) {
			structName += moduleCamelSuffix
		}
		if strings.HasSuffix(name, moduleKebabSuffix) {
			name = strings.TrimSuffix(name, moduleKebabSuffix)
		}
		generateModule(packageName, structName, name, path)
	default:
		fmt.Printf("unknown category: %s \n", category)
		fmt.Println("----------")
		showHelp()
	}
}

func generateComponent(packageName, structName, name, path string) {
	content := createComponentTemplate(packageName, structName, name)
	createFile(path, content)
	fmt.Printf("Component [%s] was generated\n", name)
}

func generateController(packageName, structName, name, path string) {
	content := createControllerTemplate(packageName, structName, name)
	createFile(path, content)
	fmt.Printf("Controller [%s] was generated\n", name)
}

func generateModule(packageName, structName, name, path string) {
	content := createModuleTemplate(packageName, structName, name)
	createFile(path, content)
	fmt.Printf("Module [%s] was generated\n", name)
}

func createFile(path, content string) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		log.Fatalln("already exists: ", path)
	}
	if err := os.MkdirAll(path[:strings.LastIndex(path, "/")+1], os.ModePerm); err != nil {
		log.Fatalln(err)
	}
	file, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := file.WriteString(content); err != nil {
		log.Fatalln(err)
	}
}
