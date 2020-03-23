package boot

import (
	"github.com/friday182/ttm-go/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	//"github.com/gogf/gf/os/gcfg"
)

func init() {
	g.Cfg().SetPath("../config")
	g.Cfg().SetFileName("config.toml")
	c := g.Cfg().Get("database")
	//c := gcfg.GetContent("../config/config.toml")
	if c != "" {
		if err := model.ConnectDb(); err != nil {
			//panic
		}
	}

	path := "temp/glog"
	glog.SetPath(path)
	glog.SetStdoutPrint(true)
	// 使用默认文件名称格式
	glog.Println("标准文件名称格式，使用当前时间时期")
	glog.SetFile("stdout.log")
	glog.Println("设置日志输出文件名称格式为同一个文件")
	glog.Println("---------- glog start ----------")
	
	// 链式操作设置文件名称格式
	//glog.File("stderr.log").Println("支持链式操作")
	//glog.File("error-{Ymd}.log").Println("文件名称支持带gtime日期格式")
	//glog.File("access-{Ymd}.log").Println("文件名称支持带gtime日期格式")
	//glog.File("access-{Ymd}.log").Println("set access log")

	list, err := gfile.ScanDir(path, "*")
	g.Dump(err)
	g.Dump(list)

}
