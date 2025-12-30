package cmd

import (
	"fmt"

	"github.com/damon/gogofly/conf"
	"github.com/damon/gogofly/router"
)

func Start() {
	conf.InitConfig()
	router.InitRouter()
}

func Clean() {
	fmt.Println("======================Clean======================")
}
