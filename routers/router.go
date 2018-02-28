/**
路由配置
**/
package routers

import (
	"beego-app/apis"
	"beego-app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 页面的相应路由配置
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/index", &controllers.UserController{}, "*:Index")
	beego.Router("/user/list", &controllers.UserController{}, "get:List")

	// api的路由配置
	ns := beego.NewNamespace("/api",
		beego.NSRouter("/user/list", &apis.UserController{}, "get:List"),
	)
	beego.AddNamespace(ns)
}
