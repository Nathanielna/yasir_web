package main

import (
	_ "yasir_web/routers"
	_ "yasir_web/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

