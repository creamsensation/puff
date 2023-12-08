package main

import "fmt"

func showHelp() {
	fmt.Println(`show help: puff help/h`)
	fmt.Println(`----------`)
	fmt.Println(`generate: puff generate/g <category> <path>`)
	fmt.Println(`generate controller: puff generate/g controller/r <path>`)
	fmt.Println(`generate component: puff generate/g component/c <path>`)
	fmt.Println(`generate module: puff generate/g module/m <path>`)
}
