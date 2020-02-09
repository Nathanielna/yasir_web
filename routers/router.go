package routers

import (
	"yasir_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// ERROR 页面  404/403/401/500/503
	beego.ErrorController(&controllers.ErrorController{})
	//自动匹配路由
	//beego.AutoRouter(&controllers.ObjectController{})
	//object/login   调用 ObjectController 中的 Login 方法
    //object/logout  调用 ObjectController 中的 Logout 方法
	beego.AutoRouter(&controllers.MainController{})
}
