package main

import (
	_ "github.com/friday182/ttm-go/app/service/setting"
	_ "github.com/friday182/ttm-go/boot"
	_ "github.com/friday182/ttm-go/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
