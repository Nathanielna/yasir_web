package main

import (
	_ "yasir_web/routers"
	//_ "yasir_web/sysinit"  //使用数据库请取消注释
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
)

func main() {
	// 日志输出到log目录下
	logs.Async()
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"log/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	beego.Run()
}

