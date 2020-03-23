package setting

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

var (
	//server string
	//database string
)

func init() {
	g.Cfg().SetPath("../../../config")
	g.Cfg().SetFileName("config.toml")
	ser := g.Cfg().Get("server")
	dat := g.Cfg().Get("database")

	glog.Println("Current Settings:")

	g.Dump(ser)
	g.Dump(dat)
}
