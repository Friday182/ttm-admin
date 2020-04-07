package main

import (
	_ "github.com/friday182/ttm-admin/app/service/setting"
	_ "github.com/friday182/ttm-admin/boot"
	_ "github.com/friday182/ttm-admin/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
